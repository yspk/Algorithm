package main

/*
#define PCRE2_CODE_UNIT_WIDTH 8
#cgo pkg-config: libpcre2-8
#include <stdio.h>
#include <stdlib.h>
#include <pcre2.h>
#include <string.h>

int rc;

char** RegexString(char *reString,char *subString){
	pcre2_code *re;
  	PCRE2_SIZE erroffset;
  	int errcode;
  	PCRE2_UCHAR8 buffer[128];

	PCRE2_SIZE* ovector;

	const char *pattern = reString;
	size_t pattern_size = strlen(pattern);

	const char *subject = subString;
	size_t subject_size = strlen(subject);
	uint32_t options = 0;

	pcre2_match_data *match_data;
	uint32_t ovecsize = 128;

	re = pcre2_compile((PCRE2_SPTR)pattern, pattern_size, options, &errcode, &erroffset, NULL);
	if (re == NULL)
	{
		pcre2_get_error_message(errcode, buffer, 120);
		fprintf(stderr,"%d\t%s\n", errcode, buffer);
		return NULL;
	}

	char** paar;
	match_data = pcre2_match_data_create(ovecsize, NULL);
	rc = pcre2_match(re, (PCRE2_SPTR)subject, subject_size, 0, options, match_data, NULL);
	if(rc == 0) {
		fprintf(stderr,"offset vector too small: %d",rc);
	}
	else if(rc > 0)
	{
		ovector = pcre2_get_ovector_pointer(match_data);
		PCRE2_SIZE i;
		paar = malloc(rc*sizeof(char*));
		for(i = 0; i < rc; i++)
		{
			PCRE2_SPTR start = (PCRE2_SPTR)subject + ovector[2*i];
			PCRE2_SIZE slen = ovector[2*i+1] - ovector[2*i];
			char* buf = malloc(slen*sizeof(char));
			// printf( "%.*s\n", (int)slen, (char *)start );
			sprintf( buf, "%.*s\n", (int)slen, (char *)start );
			// printf( "%s\n", buf);
			paar[i] = buf;
		}
	}
	else if (rc < 0)
	{
		printf("No match\n");
	}

	pcre2_match_data_free(match_data);
	pcre2_code_free(re);

	return paar;
}
*/
import "C"
import (
	"fmt"
	"net"
	"reflect"
	"strings"
	"time"
	"unsafe"
)

func main() {
	rep := `\d{4}[^ \t\n\r\d]{3,11}[^ ]{1}` //正则表达式
	input := "a;jhgoqoghqoj0329 u0tyu10hg0h9Y0Y9827342482y(Y0y(G)_)lajf;lqjfgqhgpqjopjqa=)*(^!@#$%^&*())9999999" //匹配字符串
	pattern := C.CString(rep)
	subject := C.CString(input)
	defer C.free(unsafe.Pointer(pattern))
	defer C.free(unsafe.Pointer(subject))
	result := C.RegexString(pattern, subject)
	str := GoStrings(int(C.rc), result)
	for _, v := range str {
		matchString := SplitString(v, 4, len(v)-1)
		if err := SendUdpMessage(matchString); err != nil {
			fmt.Println("SendUdpMessage err", err)
			return
		}
		fmt.Println("Successfully send:", matchString)
		time.Sleep(time.Second * 2)
	}
}

// GoStrings      字符串数组转换
// @description   C语言到Go语言的字符串数组转换
// @auth      yspk51           时间（2022/4/9   22:57 ）
// @param     length           int            "字符串数组长度"
// @param     argv             **C.char       "C语言中字符串数组指针"
// @return            		   []string       "Go语言的字符串数组"
func GoStrings(length int, argv **C.char) []string {
	if argv == nil {
		return nil
	}
	var str []string
	var pbuf []*C.char
	header := (*reflect.SliceHeader)(unsafe.Pointer(&pbuf))
	header.Cap = length
	header.Len = length
	header.Data = uintptr(unsafe.Pointer(argv))
	for _, v := range pbuf {
		str = append(str, strings.Replace(C.GoString(v), "\n", "", -1)) //去掉换行符
	}
	return str
}

// SendUdpMessage 发送Udp消息
// @description   以Udp的方式发送字符串
// @auth      yspk51           时间（2022/4/9   22:57 ）
// @param     message          string         "消息字符串"
// @return            		   error          "错误信息"
func SendUdpMessage(message string) error {
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:16523")
	if err != nil {
		return err
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return err
	}
	defer conn.Close()
	_, err = conn.Write([]byte(message))
	if err != nil {
		return err
	}
	return nil
}

// SplitString    字符串截取
// @description   截取字符串中间部分
// @auth      yspk51           时间（2022/4/9   21:57 ）
// @param     input        	   string         "输入字符串"
// @param     start        	   int            "起始位置"
// @param     end        	   int            "终止位置"
// @return            		   string         "字符串结果"
func SplitString(input string, start, end int) string {
	if start > end || end > len(input) || start < 0 {
		return ""
	}
	return input[start:end]
}
