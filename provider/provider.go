package provider

import (
	"context"
	"github.com/selefra/selefra-provider-gandi/constants"
	"os"

	"github.com/selefra/selefra-provider-sdk/provider"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/spf13/viper"

	"github.com/selefra/selefra-provider-gandi/gandi_client"
)

var Version = constants.V

func GetProvider() *provider.Provider {
	return &provider.Provider{
		Name:      constants.Gandi,
		Version:   Version,
		TableList: GenTables(),
		ClientMeta: schema.ClientMeta{
			InitClient: func(ctx context.Context, clientMeta *schema.ClientMeta, config *viper.Viper) ([]any, *schema.Diagnostics) {
				var gandiConfig gandi_client.Config
				err := config.Unmarshal(&gandiConfig)
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorMsg(constants.Analysisconfigerrs, err.Error())
				}

				if gandiConfig.Key == "" {
					gandiConfig.Key = os.Getenv("GANDI_KEY")
				}

				if gandiConfig.Key == "" {
					return nil, schema.NewDiagnostics().AddErrorMsg("missing key in configuration")
				}

				clients, err := gandi_client.NewClients(gandiConfig)

				if err != nil {
					clientMeta.ErrorF(constants.Newclientserrs, err.Error())
					return nil, schema.NewDiagnostics().AddError(err)
				}

				if len(clients) == 0 {
					return nil, schema.NewDiagnostics().AddErrorMsg(constants.Accountinformationnotfound)
				}

				res := make([]interface{}, 0, len(clients))
				for i := range clients {
					res = append(res, clients[i])
				}
				return res, nil
			},
		},
		ConfigMeta: provider.ConfigMeta{
			GetDefaultConfigTemplate: func(ctx context.Context) string {
				return `key: <Your Gandi KEY>`
			},
			Validation: func(ctx context.Context, config *viper.Viper) *schema.Diagnostics {
				var gandiConfig gandi_client.Config
				err := config.Unmarshal(&gandiConfig)
				if err != nil {
					return schema.NewDiagnostics().AddErrorMsg(constants.Analysisconfigerrs, err.Error())
				}
				return nil
			},
		},
		TransformerMeta: schema.TransformerMeta{
			DefaultColumnValueConvertorBlackList: []string{
				constants.Constants_2,
				constants.NA,
				constants.Notsupported,
			},
			DataSourcePullResultAutoExpand: true,
		},
		ErrorsHandlerMeta: schema.ErrorsHandlerMeta{

			IgnoredErrors: []schema.IgnoredError{schema.IgnoredErrorOnSaveResult},
		},
	}
}
