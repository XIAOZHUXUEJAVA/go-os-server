package model

type Rss struct {
	Id          int    `gorm:"type:int(11) auto_increment; primary_key; comment: 'Rss编号'" json:"rss_id"  `
	Title       string `gorm:"type:varchar(255); comment: '标题';" json:"rss_title"  `
	Description string `gorm:"type:varchar(255); comment: '描述';" json:"rss_description"  `
}

func GetAllRss() (rss []*Rss, err error) {
	if err = db.Find(&rss).Error; err != nil {
		return nil, err
	}
	return
}
