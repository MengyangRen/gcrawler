package model

import "errors"

/**
 * To change this template, choose Tools | Templates
 * and open the template in the editor.
 *
 * 定义全局异常错误码
 *
 * @author  m.y <j01428@kok.work>
 * @date
 * @package crawler
 * @description
 *
 * 说明:
 * 1.定义异常码返回，基于报警服务，xxx-xxx-8xxx ,xxx-xxx-7xxx 为异常报警->tg
 * 2.所有错误码对于信息在生成环境中,自动录入日志，不包括debug
 */
var (
	ERROR_CLICENT_FETCH          = errors.New("CLICENT-GET-7503")            //客户端连接失败
	ERROR_CLICENT_FETCH_READ     = errors.New("CLICENT-GET-7503")            //客户端流读取失败
	ERROR_RESPOND_STATUS         = errors.New("CLICENT-GET-8500")            //http协议错误
	ERROR_TEAM_TALBE_INSTER      = errors.New("CLICENT-GET-9500")            //Team表入库失败
	ERROR_TASK_3HOURS_DATA_EMPTY = errors.New("TASK-FETCH-3HOURS-DATA-9600") //最近3小时的直播源数据为空
)
