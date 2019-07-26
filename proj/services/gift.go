package services

import (
	"lottery/proj/dao"
	"lottery/proj/models"
)

type GiftService interface {
	GetAll(useCache bool) []models.LtGift
	CountAll() int64
	//Search(country string) []models.LtGift
	Get(id int, useCache bool) *models.LtGift
	Delete(id int) error
	Update(data *models.LtGift, columns []string) error
	Create(data *models.LtGift) error
}

type giftService struct {
	dao *dao.GiftDao
}

func NewGiftService() GiftService {
	return &giftService{
		dao: dao.NewGiftDao(nil)}
}

func (s *giftService) GetAll(useCache bool) []models.LtGift {}

func (s *giftService) CountAll() int64 {
}

//Search(country string) []models.LtGift
func (s *giftService) Get(id int, useCache bool) *models.LtGift {}

func (s *giftService) Delete(id int) error {}

func (s *giftService) Update(data *models.LtGift, columns []string) error {}

func (s *giftService) Create(data *models.LtGift) error {
}
