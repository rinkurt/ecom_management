package tools

import (
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/jinzhu/copier"
	"github.com/sanity-io/litter"
	"item_info/kitex_gen/item"
	"item_info/model"
	"testing"
)

func TestCopy(t *testing.T) {
	req := &item.UpdateItemRequest{ItemId: 123, Title: thrift.StringPtr("abc")}
	it := &model.Item{ItemId: 123, Title: "old", UserId: 456}

	copier.CopyWithOption(it, req, copier.Option{IgnoreEmpty: true})

	fmt.Println(litter.Sdump(it))
}
