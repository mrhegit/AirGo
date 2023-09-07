package service

import (
	"AirGo/global"
	"AirGo/model"
)

func NewMonitor(isp *model.ISP) error {
	return global.DB.Create(&isp).Error
}
func UpdateMonitor(isp *model.ISP) error {
	return global.DB.Save(&isp).Error
}
func DeleteMonitor(isp *model.ISP) error {
	return global.DB.Where("user_id = ?", isp.UserID).Delete(&isp).Error

}
func GetMonitorByUserID(userID int64) (*model.ISP, error) {
	var isp model.ISP
	err := global.DB.First(&isp, userID).Error
	return &isp, err
}
