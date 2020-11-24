package model

import (
	"fmt"

	b "github.com/orca-zhang/borm"
)

// 播放资源表
type Living struct {
	ID         uint64 `borm:"id"`          // id
	LiveID     uint64 `borm:"live_id"`     // lives表id
	SitesID    int8   `borm:"sites_id"`    // 线路1,线路2;source_sites表id
	Resolution int8   `borm:"resolution"`  // 分辨率;1=标清,2=高清,3=超清
	Status     int8   `borm:"status"`      // 1=显示,2=隐藏
	PullUrl    string `borm:"pull_url"`    // 播流地址
	CreateTime int64  `borm:"create_time"` // 创建时间
	WebUrl     string `borm:"web_url"`     // 用于爬取的网页地址
	MatchAt    string `borm:"match_at"`    // 比赛时间;只抓取3小时以内的比赛
	//	EquipmentType int8   `borm:"equipment_type"` // 设备类型;1=PC,2=移动
	VideoType int8 `borm:"video_type"` // 视频格式;1=flv，2=m3u8
	JumpPage  int8 `borm:"jump_page"`  // 1=无需跳转,2=需要跳转
}

type livingAddress struct {
	ID      uint64 `borm:"id"`       // id
	PullUrl string `borm:"pull_url"` // 播流地址
}

func AddLivingAddress(data *Living) error {
	// 查询数据库中是否存在这条播放资源
	id := getLivingID(data)

	if id == 0 {
		t := b.Table(db, "live_play").Debug()

		_, err := t.Insert(data)

		if err != nil {
			fmt.Printf("添加播放资源信息出现异常: %s", err)
			return nil
		}
		return nil
	}

	// 更新播放资源
	err := updateLivingAddress(id, data.PullUrl)
	if err != nil {
		return err
	}

	return nil
}

func getLivingID(data *Living) uint64 {

	var liveData Living

	_, err := b.Table(db, "live_play").Debug().Select(&liveData,
		b.Where("live_id=? and pull_url=? and video_type=?", data.LiveID, data.PullUrl, data.VideoType))

	if err != nil {
		return 0
	}

	if liveData.ID != 0 {
		return liveData.ID
	}
	return 0
}

func updateLivingAddress(id uint64, pullUrl string) error {
	t := b.Table(db, "live_play").Debug()
	model := livingAddress{
		ID:      id,
		PullUrl: pullUrl,
	}
	_, err := t.Update(&model, b.Where(b.Eq("id", id)))
	if err != nil {
		fmt.Printf("更新播放资源信息出现异常: %s", err)
		return err
	}
	return nil
}
