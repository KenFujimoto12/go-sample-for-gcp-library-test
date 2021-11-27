package bigquery

import (
	"cloud.google.com/go/bigquery"
	"context"
	"fmt"
)

type Item struct {
	Name string
	Gender string
}

type (
	BigqueryClient interface {
		putRecord(ctx context.Context, dataset string, table string, item []Item) error
	}

	Exporter struct {
		Bq BigqueryClient
	}

	BigqueryClientImpl struct {
		BqClient *bigquery.Client
	}

	MockBigqueryClientImpl struct {
		BqClient *bigquery.Client
	}
)

func (b *Exporter) Export(ctx context.Context) error {
	dataset := "test"
	table := "test-table"

	item := []Item{
		{
			Name: "kendric",
			Gender: "male",
		},
	}

	if err := b.Bq.putRecord(ctx, dataset, table, item); err != nil {
		return err
	}

	fmt.Println("Finished to insert record to bigquery.")

	return nil
}

func (b *BigqueryClientImpl) putRecord(ctx context.Context, dataset string, table string, item []Item) error {
	if err := b.BqClient.Dataset(dataset).Table(table).Inserter().Put(ctx, item); err != nil {
		return fmt.Errorf("Failed to insert record to Bigquery.")
	}
	return nil
}