package main

import (
	"context"
	"fmt"
	hu "github.com/chinaran/httputil"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func main() {
	ch2 := make(chan interface{})
	for i := 0; i < 1000; i++ {
		v4 := uuid.NewV4()

		go SenderEmail(v4.String()+"@qq.com", ch2)
	
	}
	i := 0
	for {
		select {
		case u := <-ch2:
			println(u)
			i = 0
		default:
			println(i)
			if i > 100 {
				goto Loop
			}
			i++
			time.Sleep(time.Second)
		}
	}
Loop:
	print("执行完成")
}
func SenderEmail(email string, channel chan interface{}) {
	urlPost := "https://xf.gl/api/v1/passport/comm/sendEmailVerify"
	req := map[string]string{"email": email}
	var respPost interface{}
	if err := hu.Post(context.TODO(), urlPost, &req, &respPost, hu.WithLogTimeCost()); err != nil {
		log.Printf("Post %s err: %s", urlPost, err)
		return
	}
}
func httpPostForm(email string, chanel chan interface{}) {
	chanel <- email
	resp, err := http.PostForm("https://xf.gl/api/v1/passport/comm/sendEmailVerify",
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
func register(email string, chanel chan interface{}) {
	println(email)
	chanel <- email
	resp, err := http.PostForm("https://www.awslcn.xyz/api/v1/passport/auth/register",
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

	fmt.Println(textQuoted)

	chanel <- textQuoted

}
