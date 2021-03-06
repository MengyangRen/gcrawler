package demo

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
	"gcrawler/thread"
	"gcrawler/utils"
	"time"
)

//result halder (type row)
func Task(tid int, w thread.Worker) thread.Result {
	//utils.Debug(fmt.Sprintf("process->for->goreunc->demo->result->Task->uuid:%d", w.Uuid))
	//下面代码未make
	defer utils.CatchPanic("BUSINESS-THROW-PANIC-37406")

	return thread.Result{
		Code:    200,
		Message: "SUCCESS",
		Data:    fmt.Sprintf("%T", w),
	}
}

//callback one
//result collect (type all)
//to db  to redis
func Finished(rc <-chan thread.Result, lenght int, start time.Time) {
	sucNum := 0
	faiNum := 0
	for i := 0; i < lenght; i++ {
		_d := <-rc
		switch _d.Code {
		//suc
		case 200:
			sucNum++
		//faii
		case 500:
			faiNum++
		}
	}
	utils.ShowTaskFruitUI(start, "demo-gcrawler", lenght, sucNum, faiNum)
}

//FAILURE
func __faild() thread.Result {
	return thread.Result{
		Code:    500,
		Message: "FAILURE",
		Data:    fmt.Sprintf("{}"),
	}
}
