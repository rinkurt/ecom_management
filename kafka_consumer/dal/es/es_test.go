package es

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/sanity-io/litter"
	"kafka_consumer/model"
	"strings"
	"testing"
	"time"
)

func TestCreateIndex(t *testing.T) {
	Init()
	resp, err := cli.Indices.Create("item", cli.Indices.Create.WithBody(strings.NewReader(`
	{
		"mappings": {
			"properties": {
				"title": {
					"type": "text",
					"analyzer": "ik_max_word",
					"search_analyzer": "ik_max_word"
				}
			}
		}
	}
	`)))
	fmt.Println(err)
	fmt.Println(litter.Sdump(resp))
}

func TestUpsertDocument(t *testing.T) {
	Init()
	body := &bytes.Buffer{}
	err := json.NewEncoder(body).Encode(&model.Item{
		ItemId:       1234,
		UserId:       12345,
		Title:        "eeeee",
		VideoUrl:     "www",
		Label:        "nnn",
		Status:       1,
		Rate:         2,
		IsEcom:       false,
		ContentLevel: 0,
		CreateTime:   time.Now().Unix(),
		UpdateTime:   time.Now().Unix(),
	})
	fmt.Println(err)
	resp, err := cli.Index("item", body, cli.Index.WithDocumentID("1234"))
	fmt.Println(err)
	fmt.Println(litter.Sdump(resp))
}

func TestSearch(t *testing.T) {
	Init()
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"item_id": 1234,
			},
		},
	}
	res, err := Search(context.Background(), "item", query)
	fmt.Println(err)
	fmt.Println(res)
}
