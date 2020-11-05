package huoli

/**
 * To change this template, choose Tools | Templates
 * and open the template in the editor.
 *
 *  result.go
 *
 * @author  m.y <j01428@kok.work>
 * @date
 * @package processor/huoli
 * @description
 *
 * 说明:
 *  提供回调函数Task     （具体任务） todo
 *  提供回调函数Finished （任务汇总） todo
 *
 *
 */

import (
	"fmt"
	"gcrawler/crawler"
	"gcrawler/model"
	"gcrawler/service"
	"gcrawler/thread"
	"gcrawler/utils"
)

//result halder (type row)
func Task(tid int, w thread.Worker) thread.Result {

	//注意 所有Task函数必须添加 defer utils.CatchPanic(自定义)
	defer utils.CatchPanic("BUSINESS-THROW-PANIC-37406")
	var lp model.LivePram
	utils.Json2Struct([]byte(w.Body), &lp)
	_, html := utils.NewSimpleHttp().SetHeader().Get(lp.LiveAddress)
	utils.RandSleep(3)

	if html == nil {
		return __faild()
	}

	liverUrl := crawler.NewHlCrawler().Live(html)
	//get live detailed

	_, html = utils.NewSimpleHttp().SetHeader().Get(liverUrl)
	if html == nil {
		return __faild()
	}

	//return
	return thread.Result{
		Code:    200,
		Message: "SUCCESS",
		Data: utils.Struct2Json(thread.TaskData{
			Uuid: w.Uuid,
			Type: w.Type,
			Tid:  uint64(tid),
			Body: utils.Map2Json(crawler.NewHlCrawler().VideoAddress(html, lp)),
		}),
	}
}

//callback one
//result collect (type all)
//to db  to redis
func Finished(rc <-chan thread.Result, lenght int) {
	sucNum := 0
	faiNum := 0
	for i := 0; i < lenght; i++ {
		_d := <-rc
		switch _d.Code {
		//suc
		case 200:
			var tt thread.TaskData
			utils.Json2Struct([]byte(_d.Data), &tt)
			mb, _ := utils.Json2Map(tt.Body)
			service.NewLiveService().Store(mb)
			sucNum++
		//faii
		case 500:
			faiNum++
		}
	}
	utils.Debug(fmt.Sprintf("Task.Name:%s", "huoli-gcrawler"))
	utils.Debug(fmt.Sprintf("Task.TotalNum:%d", lenght))
	utils.Debug(fmt.Sprintf("Task.SuccessNum:%d", sucNum))
	utils.Debug(fmt.Sprintf("Task.FaiNum:%d", faiNum))
}

//FAILURE
func __faild() thread.Result {
	return thread.Result{
		Code:    500,
		Message: "FAILURE",
		Data:    fmt.Sprintf("{}"),
	}
}
