package main

/**
 * To change this template, choose Tools | Templates
 * and open the template in the editor.
 *
 * Gcrawlerfu
 *
 * @author  m.y <j01428@kok.work>
 * @date
 * @package crawler
 * @description
 *
 * go build -o gcrawler.exe .\go-env\src\gcrawler
 *
 * main.go
 * 1.初始化的工作
 * 2.开启处理器
 *
 */
import (
	"gcrawler/conf"
	"gcrawler/processor"
)

func init() {
	conf.DbConnection()
}
func main() {
	processor.NewProcessor().Register().Boot()
}
