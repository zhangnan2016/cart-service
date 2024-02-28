package service

import (
	"github.com/zhangnan2016/cart-service/domain/model"
	"github.com/zhangnan2016/cart-service/domain/repository"
)

type ICartDataService interface {
	AddCart(*model.Cart) (int64, error)
	DeleteCart(int64) error
	UpdateCart(*model.Cart) error
	FindCartByID(int64) (*model.Cart, error)
	FindAllCart(int64) ([]model.Cart, error)

	CleanCart(int64) error
	DecrNum(int64, int64) error
	IncrNum(int64, int64) error
}

// NewCartDataService 创建
func NewCartDataService(cartRepository repository.ICartRepository) ICartDataService {
	return &CartDataService{cartRepository}
}

type CartDataService struct {
	CartRepository repository.ICartRepository
}

// AddCart 插入
func (u *CartDataService) AddCart(cart *model.Cart) (int64, error) {
	return u.CartRepository.CreateCart(cart)
}

// DeleteCart 删除
func (u *CartDataService) DeleteCart(cartID int64) error {
	return u.CartRepository.DeleteCartByID(cartID)
}

// UpdateCart 更新
func (u *CartDataService) UpdateCart(cart *model.Cart) error {
	return u.CartRepository.UpdateCart(cart)
}

// FindCartByID 查找
func (u *CartDataService) FindCartByID(cartID int64) (*model.Cart, error) {
	return u.CartRepository.FindCartByID(cartID)
}

// FindAllCart 查找
func (u *CartDataService) FindAllCart(userID int64) ([]model.Cart, error) {
	return u.CartRepository.FindAll(userID)
}

func (u *CartDataService) CleanCart(userID int64) error {
	return u.CartRepository.CleanCart(userID)
}

func (u *CartDataService) DecrNum(cartID int64, num int64) error {
	return u.CartRepository.DecrNum(cartID, num)
}

func (u *CartDataService) IncrNum(cartID int64, num int64) error {
	return u.CartRepository.IncrNum(cartID, num)
}
