package tables

import (
	"context"

	"github.com/go-gandi/go-gandi"
	"github.com/selefra/selefra-provider-gandi/gandi_client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableGandiCertificateGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGandiCertificateGenerator{}

func (x *TableGandiCertificateGenerator) GetTableName() string {
	return "gandi_certificate"
}

func (x *TableGandiCertificateGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGandiCertificateGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGandiCertificateGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableGandiCertificateGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			config, err := gandi_client.Connect(ctx, taskClient.(*gandi_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			client := gandi.NewCertificateClient(*config)
			certificates, err := client.ListCertificates()
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			for _, certificate := range certificates {
				resultChannel <- certificate
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TableGandiCertificateGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGandiCertificateGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("cn").ColumnType(schema.ColumnTypeString).Description("Common Name.").
			Extractor(column_value_extractor.StructSelector("CN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("package_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Package.Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Description("One of: 'pending', 'valid', 'revoked', 'replaced', 'replaced_rev', 'expired'.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("expiration").ColumnType(schema.ColumnTypeTimestamp).Description("Expiration date.").
			Extractor(gandi_client.ExtractorTimestamp("Dates.EndsAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("UUID.").Build(),
	}
}

func (x *TableGandiCertificateGenerator) GetSubTables() []*schema.Table {
	return nil
}
