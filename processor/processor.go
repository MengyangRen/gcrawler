package processor

import (
	"gcrawler/processor/demo"
	"gcrawler/processor/huoli"
	"gcrawler/processor/jrskq"
	"gcrawler/thread"
	"gcrawler/utils"
	"os"
)

/**
 * To change this template, choose Tools | Templates
 * and open the template in the editor.
 *
 *  调度处理器
 *
 * @author  m.y <j01428@kok.work>
 * @date
 * @package processor
 * @description
 *
 * 	说明:
 * 	  go build -race .\your_path\main.go  // -race 查是否存在竞争
 */
var (
	Ctn int = 0
)

const (
	SCAN_TASK_MAX_NUM      = 0
	GCRAWLER_SCAN_MAX_RATE = 20
	GCRAWLER_GRAB_MAX_RATE = 2
	HUOLI_TABLE_ID         = 1
	JRSKQ_TABLE_ID         = 5
)

type Processor struct {
	Calls map[string]func()
}

func NewProcessor() *Processor {
	return &Processor{Calls: make(map[string]func())}
}

//boot
func (this *Processor) Boot() {
	for {
		wgw := thread.NewWaitGroupWrapper()
		for _, f := range this.Calls {
			wgw.Wrap(f)
		}
		utils.RandSleep(GCRAWLER_SCAN_MAX_RATE)
		wgw.Wait()
		this.Exit()

	}
}

//register call function
func (this *Processor) Register() *Processor {
	this.Calls = map[string]func(){
		"huoli": huoli.Boot,
		"jrskq": jrskq.Boot,
		"demo":  demo.Boot,
	}
	return this
}

//Exit
func (this *Processor) Exit() {
	if Ctn == SCAN_TASK_MAX_NUM {
		utils.Debug("processor->for->Exit")
		os.Exit(0)
	}
	Ctn++
}
