package send

import (
	"coding.net/baoquan2017/candy-backend/src/common/code/validate"
	"coding.net/baoquan2017/candy-backend/src/common/constant"
	"fmt"
	"testing"
)

func TestClSend(t *testing.T) {
	address := "15658836559"
	codeV = validate.NewCodeValidate(validate.NewMemoryStore(constant.DefaultGCInterval))
	code, err := codeV.Generate(address)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(code)
	content := fmt.Sprintf("%s", code)
	request := SmsClRequest{
		Account:  "N6247230",
		Password: "GsVUSELt59c336",
		Msg:      content,
		Phone:    address,
		//Sendtime: time.Now().Format("201204101400"),
		Report: false,
		//Extend:   "",
		//Uid:      "",
	}
	resp, err := ClSend(&request)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(resp)
}
