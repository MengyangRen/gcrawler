package model

import (
	"fmt"

	b "github.com/orca-zhang/borm"
)

// 播放链接地址
type LivePram struct {
	ID          uint64 `borm:"id" json:"id"`                     // 主键
	LiveID      string `borm:"live_id" json:"live_id"`           // lives表id
	LiveAddress string `borm:"live_address" json:"list_address"` // 播放地址
	MatchTime   string `borm:"match_time" json:"match_time"`     // 比赛时间
	WebSiteID   uint8  `borm:"website_id" json:"website_id"`     // 来源网站id
}

type LiveListData struct {
	D []LivePram `json:"d"`
}

func AddLiveList(data *LivePram) {

	t := b.Table(db, "live_list").Debug()
	err := liveListFindOne(data.LiveAddress)
	if err != nil {
		UpdateLiveList(data)
		return
	}

	_, err = t.Insert(data)
	if err != nil {
		fmt.Printf("添加播放维护地址出现异常: %s", err)
		return
	}
}

func UpdateLiveList(data *LivePram) {
	_, err := b.Table(db, "live_list").Debug().Update(&data, b.Where(b.Eq("live_id=?", data.LiveID)))

	if err != nil {
		fmt.Printf("更新播放维护地址出现异常: %s", err)
		return
	}
}

func liveListFindOne(url string) error {
	var count int

	t := b.Table(db, "live_list").Debug()

	_, err := t.Select(&count, b.Fields("count(1)"), b.Where(b.Eq("live_address", url)))

	if err != nil {
		return err
	}

	if count < 1 {
		return nil
	}

	return err
}

func GetLiveList(wid uint8) (LiveListData, error) {
	var data LiveListData

	_, err := b.Table(db, "live_list").Debug().Select(&data.D,
		b.Where("TIMESTAMPDIFF(MINUTE,match_time,now()) <= ? and TIMESTAMPDIFF(MINUTE,match_time,now()) > ? and live_address <> '' and website_id =?", 180, 0, wid))

	if err != nil {
		return data, err
	}

	return data, nil
}
