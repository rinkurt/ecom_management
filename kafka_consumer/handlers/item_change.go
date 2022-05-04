package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sanity-io/litter"
	"kafka_consumer/dal/es"
	"kafka_consumer/model"
)

func ItemChangeHandler(ctx context.Context, val []byte) {
	ic := &model.ItemChange{}
	err := json.Unmarshal(val, &ic)
	if err != nil {
		fmt.Printf("Error: unmarshal failed: %v\n", err)
		return
	}

	fmt.Println("item change:", litter.Sdump(ic))

	err = es.UpsertItem(ctx, ic.ItemAfter)
	if err != nil {
		fmt.Printf("Error: es upsert failed: %v\n", err)
		return
	}
}
