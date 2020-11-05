package haokan

import (
	"fmt"
	"gcrawler/utils"
)

/**
 * To change this template, choose Tools | Templates
 * and open the template in the editor.
 *
 *  Jrs处理器
 *
 * @author  m.y <j01428@kok.work>
 * @date
 * @package processor/jrs
 * @description
 *
 *
 */

type JrsProcessor struct{}

func NewJrsProcessor() *JrsProcessor {
	return &JrsProcessor{}
}

func Boot() {
	utils.Debug(fmt.Sprintf("processor->goroutine(->%s->%s):%s", "jrs", "Boot", "jrs.."))
}
