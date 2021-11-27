package bigquery

import (
	"context"
	"testing"
)

func Test_Export(t *testing.T) {
	mockBqClientImpl := &MockBigqueryClientImpl{nil}
	bq := Exporter{mockBqClientImpl}

	t.Run("Test to put record into bigquery", func(t *testing.T) {
		ctx := context.TODO()
		if err := bq.Export(ctx); err != nil {
			t.Fatal("expect error")
		}
	})
}