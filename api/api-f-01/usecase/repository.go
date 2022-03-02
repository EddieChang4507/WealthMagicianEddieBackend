package usecase

import (
	"eddie-1.0/api/api-f-01/domain"
	"github.com/jinzhu/gorm"
)

// Repository 數據接口
type Repository interface {
	GetDB() *gorm.DB
	GetPrice(db *gorm.DB, date string) domain.CandlestickChart
}
