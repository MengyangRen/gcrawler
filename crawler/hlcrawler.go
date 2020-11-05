package crawler

/**
 * To change this template, choose Tools | Templates
 * and open the template in the editor.
 *
 * 火力爬虫 http://www.huolisport.cn/
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
 * _Nav辅助方法
 *_
 */
import (
	"regexp"
	"strings"

	"gcrawler/helper"
	"gcrawler/model"
	"gcrawler/rule"
	"gcrawler/utils"

	"github.com/PuerkitoBio/goquery"
)

type HlCrawler struct {
	//Response *http.Response
	Document *goquery.Document
}

//func NewHlCrawler(Response *http.Response,Document *goquery.Document) *HlCrawler {
func NewHlCrawler() *HlCrawler {
	// return &HlCrawler{
	// 	 Document: Document,
	// 	 Response: Response
	// }
	return &HlCrawler{}
}

//extract nav
func (this *HlCrawler) Nav(html *goquery.Document) map[int]map[string]string {
	//map[int]map[string]string {
	//extract top-nav  url,title
	topNavs := this._Nav(
		html,
		rule.GCrawler["huoli"]["url"],
		rule.GCrawler["huoli"]["nav"],
		rule.GCrawler["huoli"]["href"],
		rule.GCrawler["huoli"]["nav_filter"])

	//extract public-nav  url,title
	publicNavs := this._Nav(
		html,
		rule.GCrawler["huoli"]["url"],
		rule.GCrawler["huoli"]["top_nav"],
		rule.GCrawler["huoli"]["href"],
		rule.GCrawler["huoli"]["top_nav_filter"])

	return utils.MapMerge(topNavs, publicNavs)
}

// extract Nav assist
func (this *HlCrawler) _Nav(
	html *goquery.Document,
	url string,
	navRule string,
	navHrefRule string,
	filter string) map[int]map[string]string {

	arr := make(map[int]map[string]string)
	filterArr := strings.Split(filter, ",")
	html.Find(navRule).Find(navHrefRule).Each(
		func(i int, a *goquery.Selection) {
			title := helper.Space(a.Text())
			if !utils.InStrArray(title, filterArr) {
				v, _ := a.Attr("href")
				arr[i] = map[string]string{
					"title": title,
					"url":   url + v,
				}
			}
		})

	return arr
}

// extract live info
// len(arr)/7 = total page
func (this *HlCrawler) List(html *goquery.Document) map[int]map[string]string {
	return this.Match(html)
}

// extract match info
// len(arr)/7 = total page
func (this *HlCrawler) Match(html *goquery.Document) map[int]map[string]string {
	url := rule.GCrawler["huoli"]["url"]
	arr := make(map[int]map[string]string)
	html.Find(rule.GCrawler["huoli"]["div"]).Each(
		func(n int, div *goquery.Selection) {
			div.Find(rule.GCrawler["huoli"]["ul"]).Each(
				func(i int, ul *goquery.Selection) {
					//time.Sleep(time.Millisecond * time.Duration(rand.Int31n(100)))
					matchTime := this.matchDate(html)[n] + " " + ul.Find("li").Find("span").Eq(1).Text()
					homeIcon, _ := ul.Find("li").Find("span").Find("img").Attr("data-src")
					teamIcon, _ := ul.Find("li").Find("span").Find("img").Eq(1).Attr("data-src")
					sportIcon, _ := ul.Find("img").Attr("src")
					liveAddress, _ := ul.Find("li").Find("a").Attr("href")

					arr[i] = map[string]string{
						//比赛时间
						"matchTime": matchTime,
						//主队图片
						"homeIcon": homeIcon,
						//客队图片
						"teamIcon": teamIcon,
						//体育类型图片
						"sportIcon": url + sportIcon,
						//直播地址
						"liveAddress": liveAddress,
						//比赛名称
						"match": ul.Find("li").Find("span").Eq(0).Text(),
						//主队名称
						"homeTitle": helper.Space(ul.Find("li").Find("span").Eq(3).Text()),
						//客队名称
						"teamTitle": helper.Space(ul.Find("li").Find("span").Eq(5).Text()),
					}
				})

		})
	return arr
}

/**
 * @Description: 取联赛日期
 * @Author: hunter
 * @Date: 2020-10-20 20:18
 * @LastEditTime: 2020-10-20 20:35:00
 * @LastEditors: m.y
 */
func (this *HlCrawler) matchDate(doc *goquery.Document) map[int]string {
	if doc == nil {
		return nil
	}
	matchDateMap := make(map[int]string)
	date := doc.Find(rule.GCrawler["huoli"]["date"])
	date.Find("a").Each(func(i int, selection *goquery.Selection) {
		v, _ := selection.Attr("href")
		matchDateMap[i] = helper.Filter(v, "/?date=", "")
	})
	return matchDateMap
}

/**
 * @Description: 爬取视频播放地址
 * @Author: hunter
 * @Date: 2020-10-20 14:14:29
 * @LastEditTime: 2020-10-20 15:25:21
 * @LastEditors: m.y
 */
func (this *HlCrawler) Live(html *goquery.Document) string {
	//播放地址
	matchVideoAddress, _ := html.Find("iframe").Attr("src")
	return matchVideoAddress
}

/**
 * @Description: 爬取视频播放地址
 * @Author: hunter
 * @Date: 2020-10-20 14:14:29
 * @LastEditTime: 2020-11-2 23:11:21
 * @LastEditors: muyang
 */
func (this *HlCrawler) VideoAddress(html *goquery.Document, lp model.LivePram) map[int]map[string]string {
	playjs := helper.Space(html.Find("script").Text())
	var hrefRegexp = regexp.MustCompile("videoArr=(.*?);")
	match := hrefRegexp.FindAllString(playjs, -1)[1:]
	plays := make(map[int]map[string]string)
	for i := 0; i < len(match); i++ {
		s := strings.Replace(match[i], "videoArr=", " ", -1)
		s = strings.Replace(s, "[", " ", -1)
		s = strings.Replace(s, "'", " ", -1)
		arr := strings.Split(helper.Space(strings.Replace(strings.Replace(s, "[[", " ", -1), ";", " ", -1)), ",")

		// live pc / moblie
		if i == 2 { //pc
			plays[i] = map[string]string{
				"id":            lp.LiveID,
				"pullUrl":       arr[0],
				"videoType":     arr[1],
				"resolution":    arr[2],
				"webUrl":        lp.LiveAddress,
				"matchTime":     lp.MatchTime,
				"equipmentType": "1",
			}
			plays[i+1] = map[string]string{
				"id":            lp.LiveID,
				"pullUrl":       arr[4],
				"videoType":     arr[5],
				"resolution":    arr[6],
				"webUrl":        lp.LiveAddress,
				"matchTime":     lp.MatchTime,
				"equipmentType": "1",
			}
		} else { //moblie
			plays[i] = map[string]string{
				"id":            lp.LiveID,
				"pullUrl":       arr[0],
				"videoType":     arr[1],
				"resolution":    arr[2],
				"webUrl":        lp.LiveAddress,
				"matchTime":     lp.MatchTime,
				"equipmentType": "2",
			}
		}
	}
	return plays
}
