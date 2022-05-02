package main

import (
	"context"
	"item_info/handlers"
	"item_info/kitex_gen/item"
)

// ItemServiceImpl implements the last service interface defined in the IDL.
type ItemServiceImpl struct{}

// GetItem implements the ItemServiceImpl interface.
func (s *ItemServiceImpl) GetItem(ctx context.Context, req *item.GetItemRequest) (resp *item.GetItemResponse, err error) {
	return handlers.GetItem(ctx, req), nil
}

// UpdateItem implements the ItemServiceImpl interface.
func (s *ItemServiceImpl) UpdateItem(ctx context.Context, req *item.UpdateItemRequest) (resp *item.UpdateItemResponse, err error) {
	return handlers.UpdateItem(ctx, req), nil
}

// TODO: CreateItem
