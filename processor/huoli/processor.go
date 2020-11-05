package huoli

/**
 * To change this template, choose Tools | Templates
 * and open the template in the editor.
 *
 *  火力处理器
 *
 * @author  m.y <j01428@kok.work>
 * @date
 * @package processor/huoli
 * @description
 *
 *
 * 爬虫:
 *	_, html := utils.NewSimpleHttp().SetHeader().Get(rule.GCrawler["huoli"]["url"])
 *	//抽取导航栏
 *	crawler.NewHlCrawler().Nav(html)
 *	//抽取比赛数据
 *	crawler.NewHlCrawler().List(html)
 *	crawler.NewHlCrawler().Match(html)
 *	//抽取直播源
 *	 _,liveHtml := utils.NewSimpleHttp().SetHeader().Get(liveUrl)
 *	crawler.NewHlCrawler().Live(liveHtml)
 *
 * 数据服务:
 *   MatchService
 *   比赛数据存储
 *   service.NewMatchService().Store(match)
 *   service.NewMatchService().Get3HoursLiveItems()
 *
 */
import (
	"gcrawler/crawler"
	"gcrawler/helper"
	"gcrawler/model"
	"gcrawler/rule"
	"gcrawler/service"
	"gcrawler/thread"
	"gcrawler/utils"
	"math/rand"
	"time"
)

const (
	HUOLI_TABLE_ID = 1
)

type HlProcessor struct{}

func NewHlProcessor() *HlProcessor {
	return &HlProcessor{}
}

func (this *HlProcessor) Scan() {
	this._cr()
	//开始工作
	this._hlth()
	//随机睡眠(1-10秒) 让出CPU使用权
	time.Sleep(time.Duration(rand.Intn(20)) * time.Second)
	// for {
	// 	WaitGroupWrapper := &thread.WaitGroupWrapper{}
	// 	WaitGroupWrapper.Wrap(func() {
	// 		utils.Debug("processor->for->goroutine->huoli->boot->Scan..")
	// 		this._cr()
	// 		//开始工作
	// 		this._hlth()
	// 		//随机睡眠(1-10秒) 让出CPU使用权
	// 		time.Sleep(time.Duration(rand.Intn(20)) * time.Second)
	// 	})
	// 	WaitGroupWrapper.Wait()
	// }
}

func (this *HlProcessor) _cr() {

	_, html := utils.NewSimpleHttp().SetHeader("houli").Get(rule.GCrawler["huoli"]["url"])
	//一定要判断
	if html == nil {
		return
	}

	crawler.NewHlCrawler().Nav(html)
	match := crawler.NewHlCrawler().Match(html)
	service.NewMatchService().Store(match, HUOLI_TABLE_ID)
}

func (this *HlProcessor) _hlth() {
	ws, length, _ := this._ghl()
	if length <= 0 {
		return
	}
	utils.ReadMemoryStats()
	thread.NewThread(length, Task, Finished).Run(length/2, ws)
	utils.ReadMemoryStats()
}
func (this *HlProcessor) _ghl() (ws []thread.Worker, length int, r error) {
	_hlt := service.NewMatchService().Get3HoursLiveItems(HUOLI_TABLE_ID)
	lenght := len(_hlt.D)
	//数据为空清空直接 false
	if lenght <= 0 {
		return nil, 0, model.ERROR_TASK_3HOURS_DATA_EMPTY
	}
	for i := 0; i < lenght; i++ {
		worker := thread.Worker{
			Uuid: helper.GenId(),
			Type: thread.WORKER_GRAWLER_3HOURS_HUOLI,
		}
		worker.Body = utils.Struct2Json(_hlt.D[i])
		ws = append(ws, worker)
	}
	return ws, lenght, nil
}

//boot
func Boot() {
	utils.Debug("processor->for->goroutine->huoli->boot->Scan")
	NewHlProcessor().Scan()
	utils.Debug("huoli->scan exit")
}
