package demo

import (
	"fmt"
	"gcrawler/crawler"
	"gcrawler/utils"
)

/**
 * To change this template, choose Tools | Templates
 * and open the template in the editor.
 *
 *  Demo处理器
 *
 * @author  m.y <j01428@kok.work>
 * @date
 * @package processor/demo
 * @description
 *
 *
 */

type DemoProcessor struct{}

func NewDemoProcessor() *DemoProcessor { return &DemoProcessor{} }

func Boot() {
	utils.Debug("processor->for->goroutine->demo->boot")
	//未使用代理
	_, html := utils.NewSimpleHttp().SetHeader("houli").Get("http://www.huolisport.cn")
	cr := crawler.NewHlCrawler().Nav(html)
	utils.Debug(fmt.Sprintf("未使用代理获取的数据:\n%v", cr))

	//使用代理
	// _, phtml := utils.NewSimpleHttp().SetHeader("houli").UseProxy().Get("http://www.huolisport.cn")
	// pcr := crawler.NewHlCrawler().Nav(phtml)
	//utils.Debug(fmt.Sprintf("使用代理获取的数据:\n%v", pcr))

	utils.Debug("demo->boot exit..")
}
