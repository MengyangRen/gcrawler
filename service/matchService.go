package service

/**
 * To change this template, choose Tools | Templates
 * and open the template in the editor.
 *
 * 定义比赛相关服务
 *
 * @author  m.y <j01428@kok.work>
 * @date
 * @package crawler
 * @description
 *
 * 说明:
 * 提供 Store()
 * 添加流程 team->live->live_list
 *
 */

import (
	"fmt"
	"gcrawler/contrib/jingo"
	"gcrawler/helper"
	"gcrawler/model"
)

//create for MatchService
type MatchService struct{}

func NewMatchService() *MatchService {
	return &MatchService{}
}

//Get 3 hours of live broadcast
func (this *MatchService) Get3HoursLiveItems(wid uint8) model.LiveListData {
	data, err := model.GetLiveList(wid)
	buf := jingo.NewBufferFromPool()
	length := len(data.D)
	if length > 0 && err == nil {
		jingo.NewStructEncoder(model.LiveListData{}).Marshal(&data, buf)
		return data
	}
	return data
}

//create for Store
func (this *MatchService) Store(items map[int]map[string]string, wid uint8) {
	for _, item := range items {
		this._store(item, wid)
	}
}

//relation table(match,team,lives)
func (this *MatchService) _store(item map[string]string, wid uint8) error {
	// to or false
	if !model.IsExist(item["match"]) {
		model.InsertMatch(item["sportIcon"], item["match"])
	}

	// match info
	mifo := &model.MatchInfo{
		Name:      item["match"],
		HomeName:  item["homeTitle"],
		AwayName:  item["teamTitle"],
		MatchTime: item["matchTime"],
	}

	mId := model.GetMatchID(item["match"])

	//if model.IsUniqueMatch(mifo) {
	//IsRepeatMatch方法 建议优化(涉及新建字段)
	//使用时候注意 是否重复比赛 {true 表示不重复 false 表示重复}
	//建议修改成   是否重复比赛 {true 表示重复   false 表示不重复} 逻辑合理点
	if model.IsRepeatMatch(mifo) {
		// //主队
		// hId := this.AddTeam(item["sportIcon"], item["homeTitle"], item["homeIcon"])
		// // if homeID <= 0 {
		// // 	return model.ERROR_TEAM_TALBE_INSTER
		// // }
		// aId := this.AddTeam(item["sportIcon"], item["teamTitle"], item["teamIcon"])
		this.addLives(
			//主队
			this.addTeam(item["sportIcon"], item["homeTitle"], item["homeIcon"]),
			//客队
			this.addTeam(item["sportIcon"], item["teamTitle"], item["teamIcon"]),
			//赛事ID
			mId,
			//比赛时间
			item["matchTime"],
			//icon
			item["sportIcon"],
			//直播地址
			item["liveAddress"],
			//来源网站id
			wid,
		)
	}

	return nil
}

//add Team table
func (this *MatchService) addTeam(sportIcon string, name string, icon string) int {

	team := &model.Team{
		ClassID:  model.ClassID(sportIcon),
		CnName:   name,
		EnName:   "",
		ShowName: name,
		Icon:     icon,
	}
	teamID := model.GetTeamID(team.CnName) //获取主队ID
	if teamID == 0 {
		model.InsertTeam(team)
		return model.GetTeamID(team.CnName)
	}
	return 0
}

//add live table
func (this *MatchService) addLives(hId int, aId int,
	mId uint64, startTime string, sportIcon string, liveAddress string, wid uint8) {
	lives := &model.Lives{
		MatchID:    mId,
		StartTime:  startTime,
		HomeTeamID: hId,
		AwayTeamID: aId,
	}
	model.AddLives(lives, sportIcon)
	// to table live_list
	this.addLiveList(hId, aId, mId, startTime, sportIcon, liveAddress, wid)
}

//add live-list table
func (this *MatchService) addLiveList(hId int, aId int,
	mId uint64, startTime string, sportIcon string, liveAddress string, wid uint8) {
	list := &model.LivePram{
		ID:          helper.GenId(),
		LiveID:      fmt.Sprintf("%d", mId),
		WebSiteID:   wid,
		LiveAddress: liveAddress,
		MatchTime:   startTime,
	}
	//utils.Debug(*list)
	model.AddLiveList(list)
}
