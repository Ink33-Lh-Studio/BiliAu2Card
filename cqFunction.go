package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/tidwall/gjson"
)

func cqSendGroupMsg(id, msg string) {
	conf := readConfig()
	cqAddr := gjson.Get(conf, "CoolQ.0.Api.HttpAPIAddr").String()
	cqToken := gjson.Get(conf, "CoolQ.0.Api.HttpAPIToken").String()
	getWbeContent(cqAddr + "/send_group_msg?access_token=" + cqToken + "&group_id=" + id + "&message=" + url.QueryEscape(msg))
}

func cqSendPrivateMsg(id, msg string) {
	conf := readConfig()
	cqAddr := gjson.Get(conf, "CoolQ.0.Api.HttpAPIAddr").String()
	cqToken := gjson.Get(conf, "CoolQ.0.Api.HttpAPIToken").String()
	getWbeContent(cqAddr + "/send_private_msg?access_token=" + cqToken + "&user_id=" + id + "&message=" + url.QueryEscape(msg))
}

func readConfig() string {
	file, err := ioutil.ReadFile("conf.json")
	if err != nil {
		log.Fatal(err)
	}
	result := string(file)
	return result
}

func getWbeContent(url string) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4117.2 Safari/537.36")
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
		os.Exit(20)
	}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
		os.Exit(21)
	}
	defer response.Body.Close()
	//fmt.Println(response)
	//fmt.Printf("%T",response)

}
