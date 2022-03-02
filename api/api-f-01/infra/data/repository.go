package data

import (
	"eddie-1.0/api/api-f-01/domain"
	"eddie-1.0/api/api-f-01/usecase"
	"github.com/jinzhu/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) usecase.Repository {
	return &repository{db: db}
}

func (rep *repository) GetDB() *gorm.DB {
	return rep.db
}

func (rep *repository) GetPrice(db *gorm.DB, date string) (outputObj domain.CandlestickChart) {

	sql := `
            select
            	dateprice.openprice ,
            	dateprice.highestprice ,
            	dateprice.closingprice ,
            	dateprice.lowestprice
            from
            	dateprice
            where
            	dateprice.date = ? -- 日期
`
	db.Raw(sql).Scan(&outputObj)
	return
}
