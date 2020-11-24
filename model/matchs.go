package model

import (
	b "github.com/orca-zhang/borm"
)

// 赛事名称
type Match struct {
	ID      uint64 `borm:"id"`       // 主键ID
	ClassID int8   `borm:"class_id"` // live_class表的id
	name    string `borm:"name"`     // 赛事名称
}

type MatchParam struct {
	SportImage string `json:"sport_image"` //体育类型图片
	MatchName  string `json:"match_name"`  //比赛名称
	HomeImage  string `json:"home_image"`  //主队图片
	TeamImage  string `json:"team_image"`  //客队图片
	HomeTeam   string `json:"home_team"`   //主队名称
	AwayTeam   string `json:"away_team"`   //客队名称
	Address    string `json:"address"`     //直播地址
	MatchTime  string `json:"match_time"`  //比赛时间
}

type MatchListData struct {
	D []MatchParam `json:"d"`
}

// 体育类型
const ()

// Is there a race match
func IsExist(name string) bool {
	var count int64
	b.Table(db, "matchs").Select(&count, b.Fields("count(1)"), b.Where(b.Eq("name", name)))
	if count < 1 {
		return false
	}
	return true
}

func InsertMatch(img string, name string) {
	err := matchFindOne(name)
	if err != nil {
		return
	}

	data := Match{}
	data.name = name
	data.ClassID = ClassID(img)

	t := b.Table(db, "matchs").Debug()
	_, err = t.Insert(&data)
	if err != nil {
		return
	}
}

func matchFindOne(name string) error {
	var count int

	t := b.Table(db, "matchs").Debug()

	_, err := t.Select(&count, b.Fields("count(1)"), b.Where(b.Eq("name", name)))

	if err != nil {
		return err
	}

	if count < 1 {
		return nil
	}

	return err
}

func GetMatchID(name string) uint64 {
	data := Match{}

	t := b.Table(db, "matchs").Debug()
	_, err := t.Select(&data, b.Where(b.Eq("name", name)))

	if err != nil {
		return 0
	}

	return data.ID
}

func ClassID(img string) int8 {
	switch img {
	case "http://www.huolisport.cn/images/foot.png?v=1.3.3.12":
		return Soccer

	case "http://www.huolisport.cn/images/basketball.png?v=1.3.3.12":
		return Basketball
	case "18":
		return Soccer
	case "17":
		return Basketball
	}
	return 0
}
