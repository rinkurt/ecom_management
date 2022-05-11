package es

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/aquasecurity/esquery"
	"github.com/bitly/go-simplejson"
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
func Search(ctx context.Context, index string, req *esquery.SearchRequest) ([]*simplejson.Json, int64, error) {
	resp, err := req.Run(
		cli,
		cli.Search.WithContext(ctx),
		cli.Search.WithIndex(index),
		cli.Search.WithPretty(),
	)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()
	if resp.IsError() {
		return nil, 0, errors.New(resp.String())
	}

	r, err := simplejson.NewFromReader(resp.Body)
	if err != nil {
		return nil, 0, err
	}

	res := make([]*simplejson.Json, 0)
	hits := r.GetPath("hits", "hits")
	for i := range hits.MustArray() {
		source := hits.GetIndex(i).Get("_source")
		res = append(res, source)
	}

	total := r.GetPath("hits", "total", "value").MustInt64()
	return res, total, nil
}
