package tables

import (
	"context"

	"github.com/go-gandi/go-gandi"
	"github.com/go-gandi/go-gandi/domain"
	"github.com/selefra/selefra-provider-gandi/gandi_client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableGandiDomainGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGandiDomainGenerator{}

func (x *TableGandiDomainGenerator) GetTableName() string {
	return "gandi_domain"
}

func (x *TableGandiDomainGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGandiDomainGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGandiDomainGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableGandiDomainGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			config, err := gandi_client.Connect(ctx, taskClient.(*gandi_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			client := gandi.NewDomainClient(*config)
			domains, err := client.ListDomains()
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			for _, domain := range domains {
				resultChannel <- domain
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func getLiveDNS(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {
	domain := result.(domain.ListResponse)

	config, err := gandi_client.Connect(ctx, taskClient.(*gandi_client.Client).Config)
	if err != nil {

		return nil, err
	}

	client := gandi.NewDomainClient(*config)
	liveDNS, err := client.GetLiveDNS(domain.FQDN)
	if err != nil {

		return nil, err
	}
	return liveDNS, nil
}

func (x *TableGandiDomainGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGandiDomainGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("Unique id of the domain.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tld").ColumnType(schema.ColumnTypeString).Description("The top-level domain.").
			Extractor(column_value_extractor.StructSelector("TLD")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("owner").ColumnType(schema.ColumnTypeString).Description("The username of the owner.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dates_created_at").ColumnType(schema.ColumnTypeTimestamp).Description("The date the domain started to be handled by Gandi.").
			Extractor(gandi_client.ExtractorTimestamp("Dates.CreatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("livednssec_available").ColumnType(schema.ColumnTypeBool).Description("Indicates if DNSSEC with liveDNS may be applied to this domain.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getLiveDNS(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("LiveDNSSECAvailable")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dates_pending_delete_ends_at").ColumnType(schema.ColumnTypeTimestamp).Description("The date from which the domain will be available after a deletion.").
			Extractor(gandi_client.ExtractorTimestamp("Dates.PendingDeleteEndsAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dates_registry_ends_at").ColumnType(schema.ColumnTypeTimestamp).Description("The date the domain will end at the registry.").
			Extractor(gandi_client.ExtractorTimestamp("Dates.RegistryEndsAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("fqdn").ColumnType(schema.ColumnTypeString).Description("Fully qualified domain name, written in its native alphabet (IDN).").
			Extractor(column_value_extractor.StructSelector("FQDN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("domain_owner").ColumnType(schema.ColumnTypeString).Description("The full name of the owner.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_renew").ColumnType(schema.ColumnTypeBool).Description("Automatic renewal status.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sharing_id").ColumnType(schema.ColumnTypeString).Description("The id of the organization.").
			Extractor(column_value_extractor.StructSelector("SharingID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeJSON).Description("The status of the domain.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dates_deletes_at").ColumnType(schema.ColumnTypeTimestamp).Description("The date on which the domain will be deleted at the registry.").
			Extractor(gandi_client.ExtractorTimestamp("Dates.DeletesAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("fqdn_unicode").ColumnType(schema.ColumnTypeString).Description("Fully qualified domain name, written in unicode.").
			Extractor(column_value_extractor.StructSelector("FQDNUnicode")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dates_hold_begins_at").ColumnType(schema.ColumnTypeTimestamp).Description("The date from which the domain is held.").
			Extractor(gandi_client.ExtractorTimestamp("Dates.HoldBeginsAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("livedns_current").ColumnType(schema.ColumnTypeString).Description("Type of nameservers currently set. classic corresponds to Gandi's classic nameservers, livedns is for the new, default, Gandi nameservers, premium_dns indicates the presence of Gandi's Premium DNS nameserver and the corresponding service subscription, and other is for custom nameservers.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getLiveDNS(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("Current")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dnssec_available").ColumnType(schema.ColumnTypeBool).Description("Indicates if DNSSEC may be applied to the domain.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getLiveDNS(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("DNSSECAvailable")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("orga_owner").ColumnType(schema.ColumnTypeString).Description("The username of the organization owner.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("Tags associated to this domain.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dates_registry_created_at").ColumnType(schema.ColumnTypeTimestamp).Description("The date the domain was created on the registry.").
			Extractor(gandi_client.ExtractorTimestamp("Dates.RegistryCreatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dates_updated_at").ColumnType(schema.ColumnTypeTimestamp).Description("The last update date.").
			Extractor(gandi_client.ExtractorTimestamp("Dates.UpdatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dates_hold_ends_at").ColumnType(schema.ColumnTypeTimestamp).Description("The date from which the domain can’t be renewed anymore (the domain can be restored if the registry supports redemption period otherwise the domain might be destroyed at Gandi at that date).").
			Extractor(gandi_client.ExtractorTimestamp("Dates.HoldEndsAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("nameservers").ColumnType(schema.ColumnTypeJSON).Description("List of current nameservers.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				r, err := getLiveDNS(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				extractor := column_value_extractor.StructSelector("Nameservers")
				return extractor.Extract(ctx, clientMeta, taskClient, task, row, column, r)
			})).Build(),
	}
}

func (x *TableGandiDomainGenerator) GetSubTables() []*schema.Table {
	return nil
}
