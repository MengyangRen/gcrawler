package haokan

/**
 * To change this template, choose Tools | Templates
 * and open the template in the editor.
 *
 *  result.go
 *
 * @author  m.y <j01428@kok.work>
 * @date
 * @package processor/jrs
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
)

//result halder (type row)
func Task(tid int, w thread.Worker) thread.Result {
	//return
	return thread.Result{
		Code:    200,
		Message: "SUCCESS",
		Data:    fmt.Sprintf("{}"),
	}
}

//callback one
//result collect (type all)
//to db  to redis
func Finished(rc <-chan thread.Result, lenght int) {}

//FAILURE
func __faild() thread.Result {
	return thread.Result{
		Code:    500,
		Message: "FAILURE",
		Data:    fmt.Sprintf("{}"),
	}
}
