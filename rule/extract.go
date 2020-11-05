package rule

/**
 * To change this template, choose Tools | Templates
 * and open the template in the editor.
 *
 * 定义抽取规则文件
 *
 * @author  m.y <j01428@kok.work>
 * @date
 * @package crawler
 * @description
 *
 * 说明:
 * 1.定义不同站点爬虫抽离规则
 * 2.考虑该规则可以入库,json,yaml
 */
var (

	//抽取规则浏览器头规则
	GCrawler = map[string]map[string]string{

		//火力直播
		"huoli": map[string]string{
			"url":            "http://www.huolisport.cn",
			"nav":            "[class='public-nav']",
			"top_nav":        "[class='am-nav am-nav-pills am-topbar-nav match-list']",
			"div":            "[class*=am-tab-panel ]",
			"ul":             "[class='hvr-shadow hvr-pop']",
			"date":           "[class='am-tabs-nav am-nav am-nav-tabs']",
			"href":           "a",
			"nav_filter":     "首页,体育资讯,CCTV5",
			"top_nav_filter": "所有联赛,展开",
		},

		//火速直播
		"jrskq": map[string]string{
			"url": "http://www.jrskq.com",
			"lb":  "[class='listBox']",
			"ct":  "[class='contenTab']",
			"tdm": "[class='todayMatch match']",
		},
	}

	//爬虫请求头
	GCrawlerHeader = map[string]map[string]string{
		"huoli": map[string]string{
			"Host": "www.huolisport.cn",
			//"X-Requested-With":          "XMLHttpRequest",
			"Upgrade-Insecure-Requests": "1",
			//	"Referer":                   "",
			"Accept-Language": "zh-CN,zh;q=0.9,en;q=0.8",
			"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36",
			"Cookie":          "__51cke__=; Hm_lvt_5364f18b8f6e3730ab236adafd0d5d57=1603109810,1603945791,1604043217; td_cookie=2692290499; __tins__20760639=%7B%22sid%22%3A%201604043217141%2C%20%22vd%22%3A%207%2C%20%22expires%22%3A%201604045110366%7D; __51laig__=7; Hm_lpvt_5364f18b8f6e3730ab236adafd0d5d57=1604043311",
		},

		"jrskq": map[string]string{
			"Host": "www.jrskq.com",
			//"X-Requested-With":          "XMLHttpRequest",
			"Upgrade-Insecure-Requests": "1",
			//	"Referer":                   "",
			"Accept-Language": "zh-CN,zh;q=0.9,en;q=0.8",
			"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36",
			"Cookie":          "__51cke__=; Hm_lvt_5364f18b8f6e3730ab236adafd0d5d57=1603109810,1603945791,1604043217; td_cookie=2692290499; __tins__20760639=%7B%22sid%22%3A%201604043217141%2C%20%22vd%22%3A%207%2C%20%22expires%22%3A%201604045110366%7D; __51laig__=7; Hm_lpvt_5364f18b8f6e3730ab236adafd0d5d57=1604043311",
		},

		"default": map[string]string{
			"Host": "www.baidu.com",
			//"X-Requested-With":          "XMLHttpRequest",
			"Upgrade-Insecure-Requests": "1",
			"Referer":                   "http://www.baidu.com",
			"Accept-Language":           "zh-CN,zh;q=0.9,en;q=0.8",
			"User-Agent":                "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36",
			"Cookie":                    "__51cke__=; Hm_lvt_5364f18b8f6e3730ab236adafd0d5d57=1603109810,1603945791,1604043217; td_cookie=2692290499; __tins__20760639=%7B%22sid%22%3A%201604043217141%2C%20%22vd%22%3A%207%2C%20%22expires%22%3A%201604045110366%7D; __51laig__=7; Hm_lpvt_5364f18b8f6e3730ab236adafd0d5d57=1604043311",
		},
	}
)
