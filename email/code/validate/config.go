package validate

import "time"

// Config 邮箱验证的配置参数
type Config struct {
	Expire  time.Duration // 过期的持续时间
	CodeLen int           // 验证码的长度
}

// DataItem 存储验证信息的数据项
type DataItem struct {
	Id         int64         `json:"id"`          // 唯一标识
	Email      string        `json:"email"`       // 邮箱
	Code       string        `json:"code"`        // 验证码
	CreateTime time.Time     `json:"create_time"` // 存储时间
	Expire     time.Duration `json:"expire"`      // 过期的持续时间
}
