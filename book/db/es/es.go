package es

import "github.com/elastic/go-elasticsearch/v7"

var es *elasticsearch.Client

func InitEs() error {
	var err error
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://111.231.76.156:9200",
		},
	}
	es, err = elasticsearch.NewClient(cfg)
	return err
}
