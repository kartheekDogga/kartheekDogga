package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/sirupsen/logrus"
	"github.com/sourcegraph/conc/pool"
)

type MigrationClient struct {
	dbClient    *dynamodb.Client
	OriginTable string
	TargetTable string
	// MutationFunc func(map[string]types.AttributeValue) (map[string]types.AttributeValue, error)
	log logrus.FieldLogger
}

func (mc *MigrationClient) batchWriteItemsV3(ctx context.Context, items []map[string]types.AttributeValue) error {
	p := pool.New().WithContext(ctx).WithMaxGoroutines(5).WithFirstError()
	const batchSize = 25
	batchID := 1
	for i := 0; i < len(items); i += batchSize {
		end := i + batchSize
		if end > len(items) {
			end = len(items)
		}

		batch := items[i:end]
		writeRequests := make([]types.WriteRequest, len(batch))
		for j, item := range batch {
			writeRequests[j] = types.WriteRequest{
				PutRequest: &types.PutRequest{
					Item: item,
				},
			}
		}
		p.Go(func(ctx context.Context) error {
			return mc.writeSub(ctx, writeRequests, batchID)
		})
		batchID++

	}

	return p.Wait()
}

func (mc *MigrationClient) writeSub(ctx context.Context, writeRequests []types.WriteRequest, batchID int) error {
	mc.log.Infof("received write operation. batchID: %v batchlen: %v", batchID, len(writeRequests))
	input := &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]types.WriteRequest{
			mc.TargetTable: writeRequests,
		},
	}

	for {
		result, err := mc.dbClient.BatchWriteItem(ctx, input)
		if err != nil {
			return fmt.Errorf("failed to batch write items (batchID: %v ): %w", batchID, err)
		}

		// Check for unprocessed items
		if len(result.UnprocessedItems) == 0 {
			break
		}

		// Log unprocessed items
		mc.log.Errorf("Unprocessed items: %v", result.UnprocessedItems)

		// Retry unprocessed items
		mc.log.Infof("Retrying unprocessed items... (BatchID: %v)", batchID)
		input.RequestItems = result.UnprocessedItems

		// Optional: Add a delay before retrying
		// time.Sleep(1 * time.Second)
	}

	return nil
}
