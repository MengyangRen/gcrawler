package jrskq

/**
 * To change this template, choose Tools | Templates
 * and open the template in the editor.
 *
 *  Jrskq处理器
 *
 * @author  m.y <j01428@kok.work>
 * @date
 * @package processor/jrskq
 * @description
 *
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
)

const (
	JRSKQ_TABLE_ID = 5
)

type JrskqProcessor struct{}

func NewJrskqProcessor() *JrskqProcessor { return &JrskqProcessor{} }
func (this *JrskqProcessor) _hlth() {
	ws, length, _ := this._ghl()
	if length <= 0 {
		return
	}
	utils.ReadMemoryStats()
	thread.NewThread(length, Task, Finished).Run(length/2, ws)
	utils.ReadMemoryStats()
}
func (this *JrskqProcessor) _ghl() (ws []thread.Worker, length int, r error) {
	_hlt := service.NewMatchService().Get3HoursLiveItems(JRSKQ_TABLE_ID)
	lenght := len(_hlt.D)
	//数据为空清空直接 false
	if lenght <= 0 {
		return nil, 0, model.ERROR_TASK_3HOURS_DATA_EMPTY
	}
	for i := 0; i < lenght; i++ {
		worker := thread.Worker{
			Uuid: helper.GenId(),
			Type: thread.WORKER_GRAWLER_3HOURS_JRSKQ,
		}
		worker.Body = utils.Struct2Json(_hlt.D[i])
		ws = append(ws, worker)
	}
	return ws, lenght, nil
}

func (this *JrskqProcessor) _cr() {
	
	_, html := utils.NewSimpleHttp().SetHeader("jrskq").Get(rule.GCrawler["jrskq"]["url"])

	if html == nil {
		return
	}
	match := crawler.NewJrskqCrawler().Match(html)
	service.NewMatchService().Store(match, JRSKQ_TABLE_ID)
}
func (this *JrskqProcessor) Scan() {
	jrskq := NewJrskqProcessor()
	jrskq._cr()
	jrskq._ghl()
}

func Boot() {
	utils.Debug("processor->for->goroutine->jrskq->boot->Scan")
	NewJrskqProcessor().Scan()
	utils.Debug("jrskq->scan exit")
}
