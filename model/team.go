package model

import (
	"fmt"
	b "github.com/orca-zhang/borm"
	"time"
)

// 球队信息
type Team struct {
	ID       int    `borm:"id"`        // 主键
	ClassID  int8   `borm:"class_id"`  // live_class表的code
	CnName   string `borm:"cn_name"`   // 中文名
	EnName   string `borm:"en_name"`   // 英文名
	ShowName string `borm:"show_name"` // 显示的名字
	Icon     string `borm:"icon"`      //队伍图标
}

type MatchInfo struct {
	Name      string //赛事名称
	HomeName  string //主队名称
	AwayName  string //客队名称
	MatchTime string //比赛时间
}


/**
 * @Description: 添加球队信息
 * @Author: hunter
 * @Date: 2020-10-21 20:43:20
 * @LastEditTime: 2020-10-21 20:48:39
 * @LastEditors: hunter
 */
func InsertTeam(data *Team)  {

	t := b.Table(db,"team").Debug()

	_,err := t.Insert(data)
	if err != nil {
		fmt.Printf("程序出现异常: %s",err)
		return
	}
}


/**
 * @Description: 获取球队ID
 * @Author: hunter
 * @Date: 2020-10-21 20:49:08
 * @LastEditTime: 2020-10-21 20:55:25
 * @LastEditors: hunter
 */
func GetTeamID(name string)  int {

	data := Team{}
	t,err := b.Table(db,"team").Debug().Select(&data,b.Where(b.Eq("cn_name",name)),b.Limit(1))

	if err != nil {
		fmt.Printf("程序出现异常: %s",err)
		return 0
	}

	if t != 1 {
		return 0
	}
	return data.ID

}

/**
 * @Description: 查询比赛信息是否有重复
 * @Author: hunter
 * @Date: 2020-10-22 12:16:06
 * @LastEditTime: 2020-10-22 12:58:11
 * @LastEditors: hunter
 */
func IsRepeatMatch(data *MatchInfo) bool {
	var count int
	_,err := b.Table(db,"matchs").Debug().Select(&count,b.Fields("count(1)"),b.Where(b.Eq("name",data.Name)))

	if err != nil {
		fmt.Printf("程序出现异常: %s",err)
		return false
	}

	if count > 0 {

		date,err := time.Parse("2006-01-02 15:04", data.MatchTime)
		if err != nil {
			fmt.Printf("程序出现异常: %s",err)
			return false
		}

		//主队ID
		homeID := GetTeamID(data.HomeName)
		//客队ID
		awayID := GetTeamID(data.AwayName)

		_,err = b.Table(db,"lives").Debug().Select(&count,b.Fields("count(1)"),
				 b.Where(b.Eq("home_team_id",homeID),b.Eq("away_team_id",awayID),
				 b.Eq("start_time",date)))

		if err != nil {
			fmt.Printf("程序出现异常: %s",err)
			return false
		}
		if count > 0 {
			return false
		}
	}
	return true
}