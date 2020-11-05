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
	"crypto/tls"
	"fmt"
	"gcrawler/model"
	"gcrawler/rule"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

var ()

type SimpleHttp struct {
	Client       *http.Client      //定义客户端属性
	Header       map[string]string //请求头
	ResponHeader string            //响应头
	ResponBody   string            //响应body
	Error        error
	Response     *http.Response
}

//create obj
func NewSimpleHttp() *SimpleHttp {
	return &SimpleHttp{
		Client:       &http.Client{},
		Header:       make(map[string]string),
		ResponHeader: "",
		ResponBody:   "",
	}
}

func (this *SimpleHttp) UseProxy() *SimpleHttp {
	//忽略https认证
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	proxyUrl, err := url.Parse(NewIPProxyPool().Rand())
	if err == nil { // seting proxy
		tr.Proxy = http.ProxyURL(proxyUrl)
	}

	this.Client = &http.Client{
		Transport: tr,
	}
	return this
}

//get respon Body
func (this *SimpleHttp) Body() string {
	return this.ResponBody
}

//sethader
//func (this *SimpleHttp) SetHeader(header map[string]string) *SimpleHttp {
func (this *SimpleHttp) SetHeader(h string) *SimpleHttp {
	if h != "" {
		this.Header = rule.GCrawlerHeader[h]
		return this
	}
	this.Header = rule.GCrawlerHeader["default"]
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

	for key, val := range this.Header {
		resp.Header.Set(key, val)
	}

	doc, _ := goquery.NewDocumentFromReader(resp.Body)
	defer resp.Body.Close()
	return this, doc
}

// //use Proxy
// func (this *SimpleHttp) UseProxyGet(u string, proxy string) error {
// 	req, err := http.NewRequest(http.MethodGet, u, nil)
// 	if err != nil {
// 		return err
// 	}
// 	tr := &http.Transport{TLSClientConfig: &tls.Config{
// 		InsecureSkipVerify: true,
// 	}}
// 	if proxy != "" {
// 		proxyUrl, err := url.Parse(proxy)
// 		if err == nil { // 使用传入代理
// 			tr.Proxy = http.ProxyURL(proxyUrl)
// 		}
// 	}
// 	r, err := (&http.Client{Transport: tr}).Do(req)
// 	if err != nil {
// 		return err
// 	}
// 	if r != nil {
// 		defer r.Body.Close()
// 	}
// 	b, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println("useProxyURL:", string(b))
// 	return nil
// }
