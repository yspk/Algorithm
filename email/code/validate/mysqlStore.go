/* Gorilla Sessions backend for MySQL.

Copyright (c) 2013 Contributors. See the list of contributors in the CONTRIBUTORS file for details.

This software is licensed under a MIT style license available in the LICENSE file.
*/
package validate

import (
	"coding.net/baoquan2017/candy-backend/src/common/constant"
	"coding.net/baoquan2017/candy-backend/src/common/logger"
	"database/sql"
	"encoding/gob"
	"errors"
	"github.com/go-sql-driver/mysql"
	"strings"
	"time"
)

//创建基于mysql内存表的存储实例
type MySQLStore struct {
	db          *sql.DB
	stmtInsert  *sql.Stmt
	stmtDelete  *sql.Stmt
	stmtDelete2 *sql.Stmt
	stmtUpdate  *sql.Stmt
	stmtSelect  *sql.Stmt
	stmtSelect2 *sql.Stmt
	stmtSelect3 *sql.Stmt
	table       string
}

func init() {
	gob.Register(time.Time{})
}

func NewMySQLStore(endpoint string, tableName string) *MySQLStore {
	db, err := sql.Open("mysql", endpoint)
	if err != nil {
		logger.Warn(err)
		return nil
	}
	var store *MySQLStore
	if store, err = NewMySQLStoreFromConnection(db, tableName); err != nil {
		logger.Warn(err)
		return nil
	}
	return store
}

func NewMySQLStoreFromConnection(db *sql.DB, tableName string) (*MySQLStore, error) {
	// Make sure table name is enclosed.
	tableName = "`" + strings.Trim(tableName, "`") + "`"

	cTableQ := "CREATE TABLE IF NOT EXISTS " +
		tableName + " (id bigint(20) unsigned NOT NULL AUTO_INCREMENT, " +
		"email varchar(255) NOT NULL DEFAULT '', " +
		"code varchar(64) NOT NULL DEFAULT '', " +
		"create_time TIMESTAMP DEFAULT NOW(), " +
		"expire bigint(20) NOT NULL," +
		"PRIMARY KEY(`id`), " +
		"UNIQUE KEY `email` (`email`)" +
		") ENGINE=MEMORY DEFAULT CHARSET=utf8"
	if _, err := db.Exec(cTableQ); err != nil {
		switch err.(type) {
		case *mysql.MySQLError:
			// Error 1142 means permission denied for create command
			if err.(*mysql.MySQLError).Number == 1142 {
				break
			} else {
				return nil, err
			}
		default:
			return nil, err
		}
	}

	insQ := "INSERT INTO " + tableName +
		"(email, code, create_time, expire) VALUES (?, ?, ?, ?) " +
		"on duplicate key update " +
		"email=values(email),code=values(code),create_time=values(create_time),expire=values(expire) "
	stmtInsert, stmtErr := db.Prepare(insQ)
	if stmtErr != nil {
		return nil, stmtErr
	}

	delQ := "DELETE FROM " + tableName + " WHERE email = ?"
	stmtDelete, stmtErr := db.Prepare(delQ)
	if stmtErr != nil {
		return nil, stmtErr
	}

	delQ2 := "DELETE FROM " + tableName + " WHERE create_time < ?"
	stmtDelete2, stmtErr := db.Prepare(delQ2)
	if stmtErr != nil {
		return nil, stmtErr
	}

	updQ := "UPDATE " + tableName + " SET code = ?, create_time = ?, expire = ? " +
		"WHERE email = ?"
	stmtUpdate, stmtErr := db.Prepare(updQ)
	if stmtErr != nil {
		return nil, stmtErr
	}

	selQ := "SELECT id ,email, code, create_time, expire from " +
		tableName + " WHERE email = ?"
	stmtSelect, stmtErr := db.Prepare(selQ)
	if stmtErr != nil {
		return nil, stmtErr
	}

	selQ2 := "SELECT id ,email, code, create_time, expire from " +
		tableName + " WHERE id = ?"
	stmtSelect2, stmtErr := db.Prepare(selQ2)
	if stmtErr != nil {
		return nil, stmtErr
	}

	selQ3 := "SELECT id ,email, code, create_time, expire from " +
		tableName + " WHERE email = ? and code=?"
	stmtSelect3, stmtErr := db.Prepare(selQ3)
	if stmtErr != nil {
		return nil, stmtErr
	}

	return &MySQLStore{
		db:          db,
		stmtInsert:  stmtInsert,
		stmtDelete:  stmtDelete,
		stmtDelete2: stmtDelete2,
		stmtUpdate:  stmtUpdate,
		stmtSelect:  stmtSelect,
		stmtSelect2: stmtSelect2,
		stmtSelect3: stmtSelect3,
		table:       tableName,
	}, nil
}

func (m *MySQLStore) Close() {
	m.stmtSelect.Close()
	m.stmtUpdate.Close()
	m.stmtDelete.Close()
	m.stmtDelete2.Close()
	m.stmtInsert.Close()
	m.stmtSelect2.Close()
	m.stmtSelect3.Close()
}

func (m *MySQLStore) Put(item DataItem) (int64, error) {
	var it DataItem
	var exp2 int64
	m.stmtSelect.QueryRow(item.Email).Scan(&it.Id, &it.Email, &it.Code, &it.CreateTime, &exp2)
	nowT := time.Now()
	expireTime := nowT.Add(-item.Expire)
	exp := int64(item.Expire / time.Nanosecond)
	if nowT.Sub(it.CreateTime) < constant.DefaultGCInterval {
		return 0, errors.New("Sending code too frequent!")
	}
	if it.Id == 0 {
		c, err := m.stmtInsert.Exec(&item.Email, &item.Code, &item.CreateTime, &exp)
		if err != nil {
			return 0, err
		}
		//删除过期数据
		m.stmtDelete2.Exec(&expireTime)
		return c.LastInsertId()
	} else {
		//logger.Warn(item)
		u, err := m.stmtUpdate.Exec(&item.Code, &item.CreateTime, &exp, &item.Email)
		if err != nil {
			return 0, err
		}
		//删除过期数据
		m.stmtDelete2.Exec(&expireTime)
		return u.LastInsertId()
	}
}

func (m *MySQLStore) TakeByID(id int64) (*DataItem, error) {
	var item DataItem
	var exp int64
	err := m.stmtSelect2.QueryRow(id).Scan(&item.Id, &item.Email, &item.Code, &item.CreateTime, &exp)
	if err != nil {
		return nil, err
	} else {
		item.Expire = time.Nanosecond * time.Duration(exp)
		return &item, nil
	}
}

func (m *MySQLStore) TakeByEmailAndCode(email, code string) (*DataItem, error) {
	var item DataItem
	var exp int64
	err := m.stmtSelect3.QueryRow(email, code).Scan(&item.Id, &item.Email, &item.Code, &item.CreateTime, &exp)
	if err != nil {
		return nil, err
	} else {
		item.Expire = time.Nanosecond * time.Duration(exp)
		return &item, nil
	}
}
