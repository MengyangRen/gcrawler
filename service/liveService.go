package service

/**
 * To change this template, choose Tools | Templates
 * and open the template in the editor.
 *
 * 直播源服务
 *
 * @author  m.y <j01428@kok.work>
 * @date
 * @package service
 * @description
 *
 * 说明:
 * 提供 Store()
 * 相关表 live_play
 *
 */

import (
	"gcrawler/helper"
	"gcrawler/model"
	"strconv"
	"time"
)

//create for MatchService
type LiveService struct{}

func NewLiveService() *LiveService {
	return &LiveService{}
}

//create live
func (this *LiveService) Store(plays map[int]map[string]string) bool {
	for _, item := range plays {
		intNum, _ := strconv.Atoi(item["id"])
		_as := this._store(item)
		//equipmentType, _ := strconv.ParseInt(item["equipmentType"], 10, 8)

		err := model.AddLivingAddress(&model.Living{
			ID:         helper.GenId(),
			LiveID:     uint64(intNum),
			SitesID:    1,
			Resolution: _as["resolution"],
			Status:     1,
			PullUrl:    item["pullUrl"],
			CreateTime: time.Now().Unix(),
			WebUrl:     item["webUrl"],
			MatchAt:    item["matchTime"],
			VideoType:  _as["videoType"],
			//EquipmentType: int8(equipmentType),
			JumpPage: 1,
		})

		if err != nil {
			continue //写入日志
		}
	}
	return true
}
func (this *LiveService) _store(m map[string]string) map[string]int8 {
	_as := make(map[string]int8)
	switch m["resolution"] {
	case "标清":
		_as["resolution"] = 1
	case "高清":
		_as["resolution"] = 2
	case "超清":
		_as["resolution"] = 3
	case "超清1080P":
		_as["resolution"] = 3
	case "超清720P":
		_as["resolution"] = 3
	default:
		_as["resolution"] = 1
	}

	switch m["videoType"] {
	case model.FlvCode:
		_as["videoType"] = 1
	case model.M3u8Code:
		_as["videoType"] = 2
	default:
		_as["videoType"] = 1
	}
	return _as
}
