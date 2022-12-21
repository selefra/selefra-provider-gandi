package gandi_client

import (
	"context"
	"errors"
	"github.com/selefra/selefra-provider-gandi/constants"
	"os"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"

	config "github.com/go-gandi/go-gandi/config"
)

func Connect(ctx context.Context, gandiConfig *Config) (*config.Config, error) {

	key := os.Getenv(constants.GANDIKEY)
	if gandiConfig.Key != constants.Constants_0 {
		key = gandiConfig.Key
	}
	if key == constants.Constants_1 {
		return nil, errors.New(constants.KeymustbesetintheconnectionconfigurationEdityourconnectionconfigurationfileorsettheGANDIKEYenvironmentvariableandthenrestartselefra)
	}
	return &config.Config{APIKey: key}, nil
}

func ExtractorTimestamp(path string) schema.ColumnValueExtractor {
	return column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
		v, err := column_value_extractor.StructSelector(path).Extract(ctx, clientMeta, client, task, row, column, result)
		return v, err
	})
}
