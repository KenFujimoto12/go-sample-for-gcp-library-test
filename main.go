package main

import (
	"context"
	"fmt"
	"go-sample/bigquery"
	"go-sample/pkg"
)

func main () {
	ctx := context.TODO()
	projectID := "xxx"

	client, err := pkg.Client(ctx, projectID)
	if err != nil {
		fmt.Println(err)
	}

	bqClientImpl := &bigquery.BigqueryClientImpl{client}
	bq := bigquery.Exporter{bqClientImpl}

	if err := bq.Export(ctx); err != nil {
		fmt.Println(err)
	}
}