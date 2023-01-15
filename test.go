package main

import (
	"context"
	"fmt"
	hu "github.com/chinaran/httputil"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	Loop()
	println("master branch")
}

func Loop() {
	ch2 := make(chan interface{})
	rand.Seed(time.Now().UnixNano())
	var q [15]string = [15]string{"qq.com", "gmail.com", "163.com", "126.com", "foxmail.com","outlook.com"}

	for i := 0; i < 10000; i++ {
		var qq string
		for i := 0; i < 9; i++ {
			intn := rand.Intn(9)
			qq = qq + strconv.Itoa(intn)

		}
		print(qq+"@"+q[rand.Intn(7)])
		go SenderEmail(qq+"@"+q[rand.Intn(3)], ch2)
		go Login(qq+"@"+q[rand.Intn(3)], ch2)
		//go Register(qq+"@qq.com", ch2)
		//go Check(ch2)
		//go SenderEmail(v4.String()+"@gmail.com", ch2)
		//go SenderEmail(v4.String()+"@163.com", ch2)
	}
	i := 0
	for {
		select {
		case u := <-ch2:
			println(u)
			i = 0
		default:
			println("ok")
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

//这个只是一个简单的版本只是获取QQ邮箱并且没有进行封装操作，另外爬出来的数据也没有进行去重操作
var (
	// \d是数字
	reQQEmail = `(\d+)@qq.com`
)

// 处理异常
func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)

	}
}
func Check(channel chan interface{}) {
	var resp interface{}
	if err := hu.Get(context.TODO(), "http://flying1008.top/", &resp); err != nil {
		log.Printf("Post %s err: %s", "http://flying1008.top/", err)
		return
	}
	channel <- resp

}
func Register(email string, channel chan interface{}) {
	urlPost := "https://meoso.net/api/v1/passport/auth/register"
	intn := rand.Intn(9)

	req := map[string]string{"email": email, "password": strconv.Itoa(intn) + "DuD_9ZvpFikujfn" + strconv.Itoa(intn)}
	var respPost interface{}
	if err := hu.Post(context.TODO(), urlPost, &req, &respPost, hu.WithLogTimeCost()); err != nil {
		log.Printf("Post %s err: %s", urlPost, err)
		return
	}
	channel <- respPost

}
func GetEmail() {
	// 1.去网站拿数据
	resp, err := http.Get("https://tieba.baidu.com/p/6051076813?red_tag=1573533731")
	HandleError(err, "http.Get url")
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	// 2.读取页面内容
	pageBytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "ioutil.ReadAll")

	// 字节转字符串
	pageStr := string(pageBytes)
	//fmt.Println(pageStr)
	// 3.过滤数据，过滤qq邮箱
	re := regexp.MustCompile(reQQEmail)
	// -1代表取全部
	results := re.FindAllStringSubmatch(pageStr, -1)
	//fmt.Println(results)

	// 遍历结果
	for _, v := range results {
		fmt.Println("email:", v[0])
	}
}
func SenderEmail(email string, channel chan interface{}) {
	urlPost := "https://72vpn.xyz/api/v1/passport/comm/sendEmailVerify"
	req := map[string]string{"email": email}
	var respPost interface{}
	if err := hu.Post(context.TODO(), urlPost, &req, &respPost, hu.WithLogTimeCost()); err != nil {
		log.Printf("Post %s err: %s", urlPost, err)
		return
	}
	channel <- respPost
}
func Login(email string, channel chan interface{}) {
	urlPost := "https://jsmao.xyz/auth/login"
	req := map[string]string{"email": email,"password":'vUEFZBD7yjjtS.:'}
	var respPost interface{}
	if err := hu.Post(context.TODO(), urlPost, &req, &respPost, hu.WithLogTimeCost()); err != nil {
		log.Printf("Post %s err: %s", urlPost, err)
		return
	}
	channel <- respPost
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
