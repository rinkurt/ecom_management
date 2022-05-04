package es

import (
	"context"
	"github.com/aquasecurity/esquery"
	"github.com/spf13/cast"
	"item_info/kitex_gen/item"
	"item_info/model"
	"item_info/tools"
)

const IndexItem = "item"

func UpsertItem(ctx context.Context, it *model.Item) error {
	return Upsert(ctx, IndexItem, cast.ToString(it.ItemId), it)
}

func SearchItem(ctx context.Context, req *item.GetItemListRequest) ([]*model.Item, error) {
	q := esquery.Bool()
	if req.IsSetItemId() {
		q = q.Filter(esquery.Terms("item_id", ToAnySlice(req.GetItemId())...))
	}
	if req.IsSetTitle() {
		q = q.Filter(esquery.Term("title", req.GetTitle()))
	}
	if req.IsSetCreateTimeFrom() {
		q = q.Filter(esquery.Range("create_time").Gte(req.GetCreateTimeFrom()))
	}
	if req.IsSetCreateTimeTo() {
		q = q.Filter(esquery.Range("create_time").Lte(req.GetCreateTimeTo()))
	}
	if req.IsSetLabel() {
		q = q.Filter(esquery.Terms("label", ToAnySlice(req.GetLabel())...))
	}
	if req.IsSetUserId() {
		q = q.Filter(esquery.Terms("user_id", ToAnySlice(req.GetUserId())...))
	}
	if req.IsSetStatus() {
		q = q.Filter(esquery.Terms("status", ToAnySlice(req.GetStatus())...))
	}
	if req.IsSetRate() {
		q = q.Filter(esquery.Terms("rate", ToAnySlice(req.GetRate())...))
	}
	if req.IsSetContentLevel() {
		q = q.Filter(esquery.Terms("content_level", ToAnySlice(req.GetContentLevel())...))
	}

	order := esquery.OrderDesc
	if req.GetOrder() == 1 {
		order = esquery.OrderAsc
	}

	searchReq := esquery.Search().Query(q).
		From(uint64(req.GetOffset())).
		Size(uint64(req.GetSize())).
		Sort(req.GetOrderBy(), order)
	resp, err := Search(ctx, IndexItem, searchReq)
	if err != nil {
		return nil, err
	}

	var res []*model.Item
	err = tools.ConvertMapSlice(resp, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
