package service

import (
	"AirGo/global"
	"AirGo/model"
)

func NewPictureUrl(id int64, url, subject string) error {
	var p = model.Gallery{
		UserID:     id,
		PictureUrl: url,
		Subject:    subject,
	}
	return global.DB.Create(&p).Error
}

// 获取图库列表，最近10条
func GetPictureList(params *model.PaginationParams) (*[]model.Gallery, error) {
	var picArr []model.Gallery
	var err error
	if params.Search != "" {
		err = global.DB.Where("subject like ?", "%"+params.Search+"%").Limit(10).Order("id desc").Find(&picArr).Error
	} else {
		err = global.DB.Limit(10).Order("id desc").Find(&picArr).Error
	}

	return &picArr, err
}
