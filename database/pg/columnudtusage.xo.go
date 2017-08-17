// Package pg contains the types for schema 'information_schema'.
package pg

// GENERATED BY XO. DO NOT EDIT.

// ColumnUdtUsageTable is the database name for the table.
const ColumnUdtUsageTable = "information_schema.column_udt_usage"

// ColumnUdtUsage represents a row from 'information_schema.column_udt_usage'.
type ColumnUdtUsage struct {
	UdtCatalog   SQLIdentifier `json:"udt_catalog"`   // udt_catalog
	UdtSchema    SQLIdentifier `json:"udt_schema"`    // udt_schema
	UdtName      SQLIdentifier `json:"udt_name"`      // udt_name
	TableCatalog SQLIdentifier `json:"table_catalog"` // table_catalog
	TableSchema  SQLIdentifier `json:"table_schema"`  // table_schema
	TableName    SQLIdentifier `json:"table_name"`    // table_name
	ColumnName   SQLIdentifier `json:"column_name"`   // column_name
}

// Constants defining each column in the table.
const (
	ColumnUdtUsageUdtCatalogField   = "udt_catalog"
	ColumnUdtUsageUdtSchemaField    = "udt_schema"
	ColumnUdtUsageUdtNameField      = "udt_name"
	ColumnUdtUsageTableCatalogField = "table_catalog"
	ColumnUdtUsageTableSchemaField  = "table_schema"
	ColumnUdtUsageTableNameField    = "table_name"
	ColumnUdtUsageColumnNameField   = "column_name"
)

// WhereClauses for every type in ColumnUdtUsage.
var (
	ColumnUdtUsageUdtCatalogWhere   SQLIdentifierField = "udt_catalog"
	ColumnUdtUsageUdtSchemaWhere    SQLIdentifierField = "udt_schema"
	ColumnUdtUsageUdtNameWhere      SQLIdentifierField = "udt_name"
	ColumnUdtUsageTableCatalogWhere SQLIdentifierField = "table_catalog"
	ColumnUdtUsageTableSchemaWhere  SQLIdentifierField = "table_schema"
	ColumnUdtUsageTableNameWhere    SQLIdentifierField = "table_name"
	ColumnUdtUsageColumnNameWhere   SQLIdentifierField = "column_name"
)

// QueryOneColumnUdtUsage retrieves a row from 'information_schema.column_udt_usage' as a ColumnUdtUsage.
func QueryOneColumnUdtUsage(db XODB, where WhereClause, order OrderBy) (*ColumnUdtUsage, error) {
	const origsqlstr = `SELECT ` +
		`udt_catalog, udt_schema, udt_name, table_catalog, table_schema, table_name, column_name ` +
		`FROM information_schema.column_udt_usage WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + order.String() + " LIMIT 1"

	cuu := &ColumnUdtUsage{}
	err := db.QueryRow(sqlstr, where.Values()...).Scan(&cuu.UdtCatalog, &cuu.UdtSchema, &cuu.UdtName, &cuu.TableCatalog, &cuu.TableSchema, &cuu.TableName, &cuu.ColumnName)
	if err != nil {
		return nil, err
	}
	return cuu, nil
}

// QueryColumnUdtUsage retrieves rows from 'information_schema.column_udt_usage' as a slice of ColumnUdtUsage.
func QueryColumnUdtUsage(db XODB, where WhereClause, order OrderBy) ([]*ColumnUdtUsage, error) {
	const origsqlstr = `SELECT ` +
		`udt_catalog, udt_schema, udt_name, table_catalog, table_schema, table_name, column_name ` +
		`FROM information_schema.column_udt_usage WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + order.String()

	var vals []*ColumnUdtUsage
	q, err := db.Query(sqlstr, where.Values()...)
	if err != nil {
		return nil, err
	}
	for q.Next() {
		cuu := ColumnUdtUsage{}

		err = q.Scan(&cuu.UdtCatalog, &cuu.UdtSchema, &cuu.UdtName, &cuu.TableCatalog, &cuu.TableSchema, &cuu.TableName, &cuu.ColumnName)
		if err != nil {
			return nil, err
		}

		vals = append(vals, &cuu)
	}
	return vals, nil
}
