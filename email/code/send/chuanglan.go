package send

import (
	"bytes"
	"coding.net/baoquan2017/candy-backend/src/common/logger"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type SmsClRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	Msg      string `json:"msg"`
	Phone    string `json:"phone"`
	Sendtime string `json:"sendtime"`
	Report   bool   `json:"report"`
	Extend   string `json:"extend"`
	Uid      string `json:"uid"`
}

type SmsClResponse struct {
	Code     string `json:"code"`
	MsgId    string `json:"msgId"`
	ErrorMsg string `json:"errorMsg"`
	Time     string `json:"time"`
}

func ClSend(requ *SmsClRequest) (*SmsClResponse, error) {
	apiUrl := "https://smssh1.253.com/msg/send/json"
	b, _ := json.Marshal(requ)
	body := bytes.NewReader(b)
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // disable verify
	}
	// Create Http Client
	client := &http.Client{Transport: transCfg}
	//client := &http.Client{}
	req, _ := http.NewRequest("POST", apiUrl, body)
	req.Header.Set("Content-Type", "application/json")

	var response SmsClResponse
	resp, err := client.Do(req)
	if err != nil {
		logger.Warn(err)
		return nil, err
	}

	defer resp.Body.Close()
	rb, err := ioutil.ReadAll(resp.Body)
	logger.Debug("sms response :" + string(rb))
	if err != nil {
		logger.Warn(err)
		return nil, err
	}

	if err = json.Unmarshal(rb, &response); err != nil {
		logger.Warn(err)
		return nil, err
	}
	return &response, nil
}

func AsyncClSend(requ *SmsClRequest, handle func(err error)) error {
	go func() {
		res, err := ClSend(requ)
		if err == nil {
			if res.Code != "0" {
				err = errors.New("chuanglan sms service error!")
			}
		}
		handle(err)
	}()
	return nil
}
