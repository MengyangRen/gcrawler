package thread

/**
* To change this template, choose Tools | Templates
* and open the template in the editor.
*
*  同步阻塞
*
* @author  m.y <j01428@kok.work>
* @date
* @package thread
* @description
*
* 说明:
*  在Go语言中，如果使用goroutine，经常需要阻塞主进程来等待goroutine的结束
*
   使用sync.WaitGroup实现
*    waiteGroup顾名思义，是等待一组行为执行结束，利用wg.Add()来添加group，
*  利用wg.Done或wg.Add(-1)来移除，wg.Wait()一直阻塞直到group完全释放
*
*  使用channel实现
*
*/

import (
	"sync"
)

type WaitGroupWrapper struct {
	sync.WaitGroup
}

func NewWaitGroupWrapper() *WaitGroupWrapper {
	return &WaitGroupWrapper{}
}
func (w *WaitGroupWrapper) Wrap(cb func()) {
	w.Add(1)
	go func() {
		cb()
		w.Done()
	}()
}
func (w *WaitGroupWrapper) _Wrap(cb func()) {
	ch := make(chan int, 1)
	go func() {
		cb()
		ch <- 1
	}()
	<-ch
}
