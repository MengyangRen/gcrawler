# gcrawler

### 简介
  Gcrawler服务开发的初衷，为团队构建基础爬虫服务。<br>
  为了解决团队快速构建爬虫<br>
  为了解决Gcrawler整合其他语言的优势（图形识别，验证码拖动,CEF控件）<br>
  为了解决多爬虫的管理，监控，报警 <br>

### 环境需求
* \>= go 1.14

### 组成说明
 * 1.processor服务
 * 2.crawler服务
 * 3.worker/task线程服务   
 * 4.数据存储服务
 * 5.抽取规则服务
 * 6.动态IP代理池
 * 7.其他语言组件服务(验证码识别(Python)，CEF控件(C#)) （暂无）
 * 8.监控与报警服务 


### 基础工作流程
![example-0](https://github.com/MengyangRen/gcrawler/blob/main/doc/example-0.png)

###其他
 对于刚入Golang坑普通开发者，可能对于Golang的一些特性不甚了解，可能实现有些不太优雅，个人也会持续完善
 Go语言在多线程编程方面是非常有优势的,去整合不同语言的特定功能就再合适不过了

 