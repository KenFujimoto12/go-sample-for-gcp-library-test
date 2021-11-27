package bigquery

import (
	"context"
)

func (m *MockBigqueryClientImpl) putRecord(ctx context.Context, dataset string, table string, item []Item) error {
	return nil
}