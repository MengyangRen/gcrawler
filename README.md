# gcrawler

### 简介
  Gcrawler服务设计的初衷为团队的打造团队基础服务。<br>
  为了解决团队快速构建爬虫<br>
  为了解决Gcrawler整合其他语言的优势（图形识别，验证码拖动,CEF控件）<br>
  为了解决多爬出的管理，监控，报警 <br>

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

```mermaid
graph LR
A[pricessor] -->B(crawler)
    B --> C[工作线程worker/task]
    C -->D[task] -->G
    C -->E[finshed]-->G
    G[数据存储服务/mysql/redis/mongodb]
    
```

###思考
 对于刚入Go坑不足2周的小萌新来说<br/>
 Go语言在多线程编程方面是非常有优势的<br/>     
 Go语言去整合不同语言的特定功能就再合适不过了

 
