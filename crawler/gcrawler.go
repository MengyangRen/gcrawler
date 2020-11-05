package crawler

/**
 * To change this template, choose Tools | Templates
 * and open the template in the editor.
 *
 * 定义爬虫接口
 *
 * @author  m.y <j01428@kok.work>
 * @date
 * @package crawler
 * @description
 *
 * 说明：
 *  1.所有爬虫必须实现该接口中方法
 *  2._Nav为其他辅助方法进行实现
 */

type Gcrawler interface {
	Nav()  //抽取导航栏
	List() //抽取列表，返回url
	Live() //抽取直播地址
	Match()
}
