package provider

import (
	"github.com/selefra/selefra-provider-gandi/tables"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

func GenTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&tables.TableGandiCertificateGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableGandiDomainGenerator{}),
	}
}
