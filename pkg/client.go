package pkg

import (
	"cloud.google.com/go/bigquery"
	"context"
	"fmt"
)

func Client(ctx context.Context, projectID string) (*bigquery.Client, error){
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("Failed to get client.")
	}
	return client, nil
}