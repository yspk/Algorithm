package code

import (
	"github.com/yspk/Algorithm/email/code/send"
	"github.com/yspk/Algorithm/email/code/validate"
	"fmt"
	"coding.net/baoquan2017/dataqin-backend/common/logger"

	//"io/ioutil"
	"database/sql"
	"regexp"
	"sync"
)

func SendCode(address string,optype uint32, sell,buy float64) error {
	//chuangLan := config.GetChuangLanConfig()
	//email := send.ValEmail{
	//	"smtp.exmail.qq.com:465",
	//	"CandyTown",
	//	"kefu@aboutico.com",
	//	"Wanta123","wjjnctgzwencchge"
	//}
	email := send.ValEmail{
		"smtp.qq.com:465",
		"SIPC",
		"3550546715@qq.com",
		"wjjnctgzwencchge",
	}
	//validateType := config.GetValidateType()
	//mysqlStore, _ := validate.NewMySQLStoreFromConnection(db, "code")
	//defer mysqlStore.Close()
	//codeV := validate.NewCodeValidate(mysqlStore)
	//code, err := codeV.Generate(address)
	//if err != nil {
	//	logger.Warn(err)
	//	return err
	//}
	reg := fmt.Sprintf("SIPC当前挂单最低价 %0.4f ;接单最高价：%.4f",sell,buy)
	var wg sync.WaitGroup
	wg.Add(1)
	    var err error
		if optype == 1 {
			err = send.AsyncTlsSend(address, "高价播报", reg, &email, func(err error) {
				defer wg.Done()
				if err != nil {
					logger.Warn(err)
				}
			})
			if err != nil {
				logger.Warn(err)
				return err
			}
			wg.Wait()
		} else if optype == 2{
			err = send.AsyncTlsSend(address, "低价播报", reg, &email, func(err error) {
				defer wg.Done()
				if err != nil {
					logger.Warn(err)
				}
			})
			if err != nil {
				logger.Warn(err)
				return err
			}
			wg.Wait()
		} else {
			err = send.AsyncTlsSend(address, "价格播报", reg, &email, func(err error) {
				defer wg.Done()
				if err != nil {
					logger.Warn(err)
				}
			})
			if err != nil {
				logger.Warn(err)
				return err
			}
			wg.Wait()
		}
		return nil
}

func CodeValidate(code, email string, db *sql.DB) bool {
	mysqlStore, _ := validate.NewMySQLStoreFromConnection(db, "code")
	defer mysqlStore.Close()
	codeV := validate.NewCodeValidate(mysqlStore)
	isValid, err := codeV.Validate(email, code)
	if err != nil {
		logger.Warn(err)
	}
	return isValid
}

func ValidatePhone(mobile string) bool {
	//reg := regexp.MustCompile(`^1([38][0-9]|14[57]|5[^4])\d{8}$`)
	//reg := regexp.MustCompile(`^1([38][0-9]|4[57]|5[^4]|7[0-9]|9[0-9])\d{8}$`)
	reg := regexp.MustCompile(`^1\d{10}$`)
	return reg.MatchString(mobile)
}
