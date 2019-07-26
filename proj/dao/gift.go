package dao

import (
	"log"
	"lottery/proj/models"

	"github.com/go-xorm/xorm"
)

type GiftDao struct {
	engine *xorm.Engine
}

func NewGiftDao(engine *xorm.Engine) *GiftDao {
	return &GiftDao{
		engine: engine}
}

func (d *GiftDao) Get(id int) *models.LtGift {
	data := &models.LtGift{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	}

	data.Id = 0
	return data
}

func (d *GiftDao) GetAll() []models.LtGift {
	list := make([]models.LtGift, 0)

	err := d.engine.
		Asc("sys_status").
		Asc("displayorder").
		Find(&list)

	if err != nil {
		log.Println("gift get all error:", err)

		return list
	}

	return list
}

func (d *GiftDao) CountAll() int64 {
	num, err := d.engine.Count(&models.LtGift{})

	if err != nil {
		return 0
	}

	return num
}

func (d *GiftDao) Delete(id int) error {
	data := &models.LtGift{Id: id, SysStatus: 1}
	_, err := d.engine.Id(data.Id).Update(data)
	return err
}

func (d *GiftDao) Update(data *models.LtGift, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

func (d *GiftDao) Create(data *models.LtGift) error {
	_, err := d.engine.Insert(data)
	return err
}
