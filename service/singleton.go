package service

import "sync"

var expressOrderServiceSingleton *ExpressOrderService
var onceExpressOrderService sync.Once

func NewExpressOrderService() *ExpressOrderService {
	onceExpressOrderService.Do(func() {
		expressOrderServiceSingleton = &ExpressOrderService{}
	})
	return expressOrderServiceSingleton
}
