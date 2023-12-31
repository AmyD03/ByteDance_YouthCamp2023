package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"encoding/json"
	"bytes"
)
//构造结构体，与json一一对应
type DictRequest struct{
	TransType string `json:"trans_type"`
	Source	  string `json:"source"`
	UserID    string `json:"user_id"`
}
type DictResponse struct {
	Rc int `json:"rc"`
	Wiki struct {
	} `json:"wiki"`
	Dictionary struct {
		Prons struct {
			EnUs string `json:"en-us"`
			En string `json:"en"`
		} `json:"prons"`
		Explanations []string `json:"explanations"`
		Synonym []string `json:"synonym"`
		Antonym []string `json:"antonym"`
		WqxExample [][]string `json:"wqx_example"`
		Entry string `json:"entry"`
		Type string `json:"type"`
		Related []interface{} `json:"related"`
		Source string `json:"source"`
	} `json:"dictionary"`
}
func main() {
	client := &http.Client{}
	var data = strings.NewReader(`{"trans_type":"en2zh","source":"good","user_id":"64c09d58d8f63e0019cb1709"}`)
	request := DictRequest{TransType: "en2zh",Source:"good"}
	buf, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	var data = bytes.NewReader(buf)
	//创建请求
	req, err := http.NewRequest("POST", "https://api.interpreter.caiyunai.com/v1/dict", data)
	if err != nil {
		log.Fatal(err)
	}
	//设置请求头
	req.Header.Set("authority", "api.interpreter.caiyunai.com")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("app-name", "xy")
	req.Header.Set("content-type", "application/json;charset=UTF-8")
	req.Header.Set("device-id", "")
	req.Header.Set("origin", "https://fanyi.caiyunapp.com")
	req.Header.Set("os-type", "web")
	req.Header.Set("os-version", "")
	req.Header.Set("referer", "https://fanyi.caiyunapp.com/")
	req.Header.Set("sec-ch-ua", `"Not/A)Brand";v="99", "Microsoft Edge";v="115", "Chromium";v="115"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "cross-site")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36 Edg/115.0.1901.183")
	req.Header.Set("x-authorization", "token:qgemv4jr1y38jyq6vhvi")
	//发起请求
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	//读取响应
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var dictResponse DictResponse
	err = json.Unmarshal(bodyText,&dictResponse)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", dictResponse)
}