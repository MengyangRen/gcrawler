package model

import (
	"gcrawler/helper"
	"time"

	b "github.com/orca-zhang/borm"
)

// 比赛信息
type Lives struct {
	ID         uint64 `borm:"id"`           // id
	ClassID    int8   `borm:"class_id"`     // live_class表的code
	MatchID    uint64 `borm:"match_id"`     // matchs表的id
	StartTime  string `borm:"start_time"`   // 比赛开始时间
	IsHot      int8   `borm:"is_hot"`       // 推荐热门;1=推荐,2=不推荐
	HotTag     int8   `borm:"hot_tag"`      // 火热标签;1=有,2=没有
	HomeTeamID int    `borm:"home_team_id"` // 主队id
	AwayTeamID int    `borm:"away_team_id"` // 客队id
	CreateTime int64  `borm:"create_time"`  // 创建时间
}

// 比赛信息
type LivesPram struct {
	ID         string `borm:"id"`           // id
	StartTime  string `borm:"start_time"`   // 比赛开始时间
	HomeTeamID int    `borm:"home_team_id"` // 主队id
	AwayTeamID int    `borm:"away_team_id"` // 客队id
}

/**
 * @Description: 添加比赛信息
 * @Author: hunter
 * @Date: 2020-10-21 19:37:00
 * @LastEditTime: 2020-10-21 19:39:15
 * @LastEditors: hunter
 */
func AddLives(data *Lives, sportIcon string) uint64 {
	data.ID = helper.GenId()
	data.ClassID = ClassID(sportIcon)
	data.IsHot = 2
	data.HotTag = 2
	data.CreateTime = time.Now().Unix()
	t := b.Table(db, "lives").Debug()
	_, err := t.Insert(data)
	if err != nil {
		return 0
	}

	return data.ID
}

/**
 * @Description: 获取比赛ID
 * @Author: hunter
 * @Date: 2020-10-30 16:00:05
 * @LastEditTime: 2020-10-30 16:02:15
 * @LastEditors: hunter
 */
func GetLives(data *LivesPram) string {

	liveData := LivesPram{}

	_, err := b.Table(db, "live_list").Debug().Select(&liveData,
		b.Where("home_team_id=? and away_team_id=? and start_time=?", data.HomeTeamID, data.AwayTeamID, data.StartTime))

	if err != nil {
		return ""
	}

	if liveData.ID != "" {
		return liveData.ID
	}
	return ""
}
