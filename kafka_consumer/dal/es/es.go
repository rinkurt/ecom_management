package es

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/pkg/errors"
	"log"
)

var cli *elasticsearch.Client

func Init() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	}

	c, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	cli = c
}

func Upsert(ctx context.Context, index, id string, val interface{}) error {
	body := &bytes.Buffer{}
	err := json.NewEncoder(body).Encode(val)
	if err != nil {
		return err
	}

	resp, err := cli.Index(index, body, cli.Index.WithDocumentID(id), cli.Index.WithContext(ctx))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.IsError() {
		return errors.New(resp.String())
	}

	return nil
}

// Search returns all items as map[string]interface{}
func Search(ctx context.Context, index string, query map[string]interface{}) ([]map[string]interface{}, error) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	resp, err := cli.Search(
		cli.Search.WithContext(ctx),
		cli.Search.WithIndex(index),
		cli.Search.WithBody(&buf),
		cli.Search.WithPretty(),
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.IsError() {
		return nil, errors.New(resp.String())
	}

	var r map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return nil, err
	}

	res := make([]map[string]interface{}, 0)
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		source := hit.(map[string]interface{})["_source"].(map[string]interface{})
		res = append(res, source)
	}

	return res, nil
}
