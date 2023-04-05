package main

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(EncodeUserSecret("2365-f604-4574-8b78-5136", "jfafifdsjalfjeqjfdajl"))
	account, secret, err := DecodeAuthHeader("MGYwZTZhNTItMmFhNS00NTc4LTk1YjQtNzc0M2E4NWM1NTJhOmI2NWY5YWM1MTIyOTE0N2ExYmQzOTczYjVhYTcyMDlk")
	if err == nil {
		fmt.Println(account, secret)
	}
}

// EncodeUserSecret produces a base64 encoded auth header of a user and secret
func EncodeUserSecret(user string, secret string) string {
	data := fmt.Sprintf("%v:%v", user, secret)
	return base64.StdEncoding.EncodeToString([]byte(data))
}

// DecodeAuthHeader produces the user and secret from an encoded auth header
func DecodeAuthHeader(authHeader string) (user string, secret string, err error) {
	data, err := base64.StdEncoding.DecodeString(authHeader)
	if err != nil {
		return
	}

	splitData := strings.Split(string(data), ":")
	if len(splitData) != 2 {
		err = fmt.Errorf("improperly formatted decoded auth header: %v", string(data))
		return
	}

	user = splitData[0]
	secret = splitData[1]
	return
}
