package utils

import (
	"gcrawler/rule"
	"math/rand"
)

/**
 * To change this template, choose Tools | Templates
 * and open the template in the editor.
 *
 * IP代理池
 *
 * @author  m.y <j01428@kok.work>
 * @date
 * @package crawler
 * @description
 *
 */

type IPProxyPool struct {
	Pool []string
}

func NewIPProxyPool() *IPProxyPool {
	IPP := new(IPProxyPool)
	IPP.Pool = make([]string, 10)
	IPP.Pool = rule.IPPool
	return IPP
}

//rand ip
func (this *IPProxyPool) Rand() string {
	return this.Pool[rand.Intn(this.Size())]
}

//get lenght
func (this *IPProxyPool) Size() int {
	return len(this.Pool)
}
