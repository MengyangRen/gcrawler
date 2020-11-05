package crawler

import (
	"gcrawler/helper"
	"gcrawler/rule"
	"time"

	"github.com/PuerkitoBio/goquery"
)

/**
 * To change this template, choose Tools | Templates
 * and open the template in the editor.
 *
 * jrskq爬虫 http://www.jrskq.com/
 *
 * @author  m.y <j01428@kok.work>
 * @date
 * @package crawler
 * @description
 *
 * 说明：
 * 1.实现抽取导航栏数据
 * 2.实现抽取所有比赛数据
 * 3.实现抽取视频播放地址
 *
 */

type JrskqCrawler struct {
	//Response *http.Response
	Document *goquery.Document
}

//func NewHlCrawler(Response *http.Response,Document *goquery.Document) *HlCrawler {
func NewJrskqCrawler() *JrskqCrawler {
	return &JrskqCrawler{}
}

func (this *JrskqCrawler) Match(html *goquery.Document) map[int]map[string]string {
	url := rule.GCrawler["jrskq"]["url"]
	arr := make(map[int]map[string]string)
	html.Find(rule.GCrawler["jrskq"]["tdm"]).Each(func(i int, div *goquery.Selection) {
		div.Find(rule.GCrawler["jrskq"]["ct"]).Each(func(n int, dt *goquery.Selection) {
			dt.Find(rule.GCrawler["jrskq"]["lb"]).Each(func(j int, l *goquery.Selection) {
				ID, _ := l.Attr("id")
				homeImage, _ := l.Find("div").Eq(0).Find("div").Eq(2).Find("p").Eq(0).Find("img").Attr("src")
				teamImage, _ := l.Find("div").Eq(0).Find("div").Eq(2).Find("p").Eq(2).Find("img").Attr("src")
				address, _ := l.Find("div").Eq(4).Find("a").Attr("href")
				arr[n] = map[string]string{
					//主队图片
					"homeIcon": homeImage,
					//客队图片
					"teamIcon": teamImage,
					//体育类型图片
					"sportIcon": ID,
					//直播地址
					"liveAddress": url + address,
					//比赛名称
					"match": l.Find("div").Eq(0).Find("div").Eq(1).Text(),
					//主队名称
					"homeTitle": helper.Space(l.Find("div").Eq(0).Find("div").Eq(2).Find("p").Eq(0).Text()),
					//客队名称
					"teamTitle": helper.Space(l.Find("div").Eq(0).Find("div").Eq(2).Find("p").Eq(2).Text()),
					//比赛时间
					"matchTime": time.Now().Format("2006-01-02") + " " + l.Find("div").Eq(0).Find("div").Eq(0).Text(),
				}
			})

		})
	})
	return arr
}

func (this *JrskqCrawler) Nav()  {}
func (this *JrskqCrawler) List() {}
func (this *JrskqCrawler) Live() {}
