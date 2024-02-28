package handler

import (
	"context"

	"github.com/zhangnan2016/cart-service/common"
	"github.com/zhangnan2016/cart-service/domain/model"
	"github.com/zhangnan2016/cart-service/domain/service"
	cart "github.com/zhangnan2016/cart-service/proto/cart"
)

type Cart struct {
	CartDataService service.ICartDataService
}

// AddCart 添加购物车
func (c *Cart) AddCart(ctx context.Context, request *cart.CartInfo, response *cart.ResponseAdd) (err error) {
	cart := &model.Cart{}
	common.SwapTo(request, cart)
	response.CartId, err = c.CartDataService.AddCart(cart)
	return err
}

// CleanCart 清空购物车
func (c *Cart) CleanCart(ctx context.Context, request *cart.Clean, response *cart.Response) error {
	if err := c.CartDataService.CleanCart(request.UserId); err != nil {
		return err
	}
	response.Meg = "购物车清空成功"
	return nil
}

// Incr 添加购物车数量
func (c *Cart) Incr(ctx context.Context, request *cart.Item, response *cart.Response) error {
	if err := c.CartDataService.IncrNum(request.Id, request.ChangeNum); err != nil {
		return err
	}
	response.Meg = "购物车添加成功"
	return nil
}

// Decr 购物车减少商品数量
func (c *Cart) Decr(ctx context.Context, request *cart.Item, response *cart.Response) error {
	if err := c.CartDataService.DecrNum(request.Id, request.ChangeNum); err != nil {
		return err
	}
	response.Meg = "购物程减少成功"
	return nil
}

// DeleteItemByID 删除购物车
func (c *Cart) DeleteItemByID(ctx context.Context, request *cart.CartID, response *cart.Response) error {
	if err := c.CartDataService.DeleteCart(request.Id); err != nil {
		return err
	}
	response.Meg = "购物车删除成功"
	return nil
}

// GetAll 查询用户所有的购物车信息
func (c *Cart) GetAll(ctx context.Context, request *cart.CartFindAll, response *cart.CartAll) error {
	cartAll, err := c.CartDataService.FindAllCart(request.UserId)
	if err != nil {
		return err
	}
	for _, v := range cartAll {
		cart := &cart.CartInfo{}
		if err := common.SwapTo(v, cart); err != nil {
			return err
		}
		response.CartInfo = append(response.CartInfo, cart)
	}
	return nil
}
