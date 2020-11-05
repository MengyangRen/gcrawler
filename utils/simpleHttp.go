package utils

/**
 * To change this template, choose Tools | Templates
 * and open the template in the editor.
 *
 * 基于net/http/goquery封装简单Http客户端
 *
 * @author  m.y <j01428@kok.work>
 * @date
 * @package crawler
 * @description
 *
 * 说明:
 *
 */

import (
	"fmt"
	"gcrawler/model"
	"gcrawler/rule"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

var ()

type SimpleHttp struct {
	Client       http.Client       //定义客户端属性
	Header       map[string]string //请求头
	ResponHeader string            //响应头
	ResponBody   string            //响应body
	Error        error
	Response     *http.Response
}

//create obj
func NewSimpleHttp() *SimpleHttp {
	return &SimpleHttp{
		Client:       http.Client{},
		Header:       make(map[string]string),
		ResponHeader: "",
		ResponBody:   "",
	}
}

//get respon Body
func (this *SimpleHttp) Body() string {
	return this.ResponBody
}

//sethader
//func (this *SimpleHttp) SetHeader(header map[string]string) *SimpleHttp {
func (this *SimpleHttp) SetHeader() *SimpleHttp {
	this.Header = rule.GCrawlerHeader["huoli"]
	return this
}

//get
func (this *SimpleHttp) Get(url string) (*SimpleHttp, *goquery.Document) {

	resp, err := this.Client.Get(url)
	if err != nil {
		Error(fmt.Sprintf("NewSimpleHttp().Get(x) Client.Get Detailed:%q", err))
		this.Error = model.ERROR_CLICENT_FETCH
		return this, nil
	}

	//fmt.Println(resp.Header)
	if resp.StatusCode != http.StatusOK {
		Error(fmt.Sprintf("NewSimpleHttp().Get(x) StatusCode:%d", resp.StatusCode))
		this.Error = model.ERROR_RESPOND_STATUS
		return this, nil
	}

	this.Response = resp

	//body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	Error(fmt.Sprintf("[error] NewSimpleHttp().Get(x) ioutil.ReadAll Detailed:%q", err))
	// 	this.Error = ERROR_CLICENT_FETCH_READ
	// 	return this
	// }
	//this.ResponBody = string(body)
	for key, val := range this.Header {
		resp.Header.Set(key, val)
		//resp.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
	}

	doc, _ := goquery.NewDocumentFromReader(resp.Body)
	defer resp.Body.Close()
	return this, doc
}
