package thread

/**
 * To change this template, choose Tools | Templates
 * and open the template in the editor.
 *
 * 线程服务
 *
 * @author  m.y <j01428@kok.work>
 * @date
 * @package thread
 * @description
 *
 */
import (
	"fmt"
	"gcrawler/utils"
	"reflect"
	"time"
)

//definition type
var (
	WORKER_GRAWLER_3HOURS_HUOLI = "WORKER_GRAWLER_3HOURS_HUOLI"
	WORKER_GRAWLER_3HOURS_JRSKQ = "WORKER_GRAWLER_3HOURS_JRSKQ"
)

//definition  Result
type Result struct {
	Code    int
	Message string
	Data    string //Json{"k":"val"}
}

//definition IsEmpty()
func (this *Result) IsEmpty() bool {
	return reflect.DeepEqual(this, Result{})
}

type TaskData struct {
	Uuid uint64 `json:"uuid"`
	Tid  uint64 `json:"taskid"`
	Type string `json:"type"`
	Body string `json:"body"`
}

//definition worker
type Worker struct {
	Uuid uint64
	Type string
	Body string
}

func (this *Worker) IsEmpty() bool {
	return reflect.DeepEqual(this, Worker{})
}

//definition thread
type Thread struct {
	WorkerChan       chan Worker
	ResultChan       chan Result
	WorkerNum        int
	WorkerFinishChan chan bool
	//register callback
	TaskCallBack     func(tid int, w Worker) Result
	FinishedCallBack func(rc <-chan Result, lenght int)
}

//create new thread
func NewThread(
	WorkerNum int,
	Task func(tid int, w Worker) Result,
	Finished func(rc <-chan Result, lenght int)) *Thread {
	return &Thread{
		WorkerNum:        WorkerNum,
		WorkerChan:       make(chan Worker, WorkerNum),
		ResultChan:       make(chan Result, WorkerNum),
		WorkerFinishChan: make(chan bool, WorkerNum),
		TaskCallBack:     Task,
		FinishedCallBack: Finished,
	}
}

//produce (worker<-)
func (this *Thread) Produce(ls []Worker) {
	for _, l := range ls {
		this.WorkerChan <- l
	}
	close(this.WorkerChan)
}

//Consume(<-worker)
//Consume(ResultChan<-)
//Conusme(WorkerFinishChan<-)
func (this *Thread) Consume(consumeNum int) {
	for i := 0; i < consumeNum; i++ {
		go func(i int) {
			defer func() {
				if err := recover(); err != nil {
					utils.CatchPanic("THREAD-THROW-PANIC-37405")
				}
			}()
			for {
				v, ok := <-this.WorkerChan
				if ok {
					result := this.TaskCallBack(i, v)
					if !result.IsEmpty() {
						this.ResultChan <- this.__faild()
					} else {
						this.ResultChan <- result
					}
				}
				if !ok { //线程未从管道中抢到任务，直接退出
					break
				}
				this.WorkerFinishChan <- true
			}
		}(i)
	}
}

//start
func (this *Thread) Run(consumeNum int, w []Worker) {
	start := time.Now().Unix()
	this.Produce(w)
	this.Consume(consumeNum)
	for i := 0; i < this.WorkerNum; i++ {
		<-this.WorkerFinishChan
	}
	end := time.Now().Unix()
	utils.Debug(fmt.Sprintf("Thread.UseTime:%d", end-start))
	utils.Debug(fmt.Sprintf("ResultChan.len:%d", len(this.ResultChan)))
	this.FinishedCallBack(this.ResultChan, len(this.ResultChan))
	this.Close()
}

//close
func (this *Thread) Close() {
	close(this.ResultChan)
	close(this.WorkerFinishChan)
}

//失败
func (this *Thread) __faild() Result {
	return Result{
		Code:    500,
		Message: "FAILURE",
		Data:    fmt.Sprintf("{}"),
	}
}
