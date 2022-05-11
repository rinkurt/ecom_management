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

// CreateItem implements the ItemServiceImpl interface.
func (s *ItemServiceImpl) CreateItem(ctx context.Context, req *item.CreateItemRequest) (resp *item.CreateItemResponse, err error) {
	return handlers.CreateItem(ctx, req), nil
}

// GetItemList implements the ItemServiceImpl interface.
func (s *ItemServiceImpl) GetItemList(ctx context.Context, req *item.GetItemListRequest) (resp *item.GetItemListResponse, err error) {
	return handlers.GetItemList(ctx, req), nil
}

// GetItemChangeHistory implements the ItemServiceImpl interface.
func (s *ItemServiceImpl) GetItemChangeHistory(ctx context.Context, req *item.GetItemChangeHistoryRequest) (resp *item.GetItemChangeHistoryResponse, err error) {
	return handlers.GetItemChangeHistory(ctx, req), nil
}
