// Package pg contains the types for schema 'information_schema'.
package pg

// GENERATED BY XO. DO NOT EDIT.

// ForeignTableTable is the database name for the table.
const ForeignTableTable = "information_schema.foreign_tables"

// ForeignTable represents a row from 'information_schema.foreign_tables'.
type ForeignTable struct {
	ForeignTableCatalog  SQLIdentifier `json:"foreign_table_catalog"`  // foreign_table_catalog
	ForeignTableSchema   SQLIdentifier `json:"foreign_table_schema"`   // foreign_table_schema
	ForeignTableName     SQLIdentifier `json:"foreign_table_name"`     // foreign_table_name
	ForeignServerCatalog SQLIdentifier `json:"foreign_server_catalog"` // foreign_server_catalog
	ForeignServerName    SQLIdentifier `json:"foreign_server_name"`    // foreign_server_name
}

// Constants defining each column in the table.
const (
	ForeignTableForeignTableCatalogField  = "foreign_table_catalog"
	ForeignTableForeignTableSchemaField   = "foreign_table_schema"
	ForeignTableForeignTableNameField     = "foreign_table_name"
	ForeignTableForeignServerCatalogField = "foreign_server_catalog"
	ForeignTableForeignServerNameField    = "foreign_server_name"
)

// WhereClauses for every type in ForeignTable.
var (
	ForeignTableForeignTableCatalogWhere  SQLIdentifierField = "foreign_table_catalog"
	ForeignTableForeignTableSchemaWhere   SQLIdentifierField = "foreign_table_schema"
	ForeignTableForeignTableNameWhere     SQLIdentifierField = "foreign_table_name"
	ForeignTableForeignServerCatalogWhere SQLIdentifierField = "foreign_server_catalog"
	ForeignTableForeignServerNameWhere    SQLIdentifierField = "foreign_server_name"
)

// QueryOneForeignTable retrieves a row from 'information_schema.foreign_tables' as a ForeignTable.
func QueryOneForeignTable(db XODB, where WhereClause, order OrderBy) (*ForeignTable, error) {
	const origsqlstr = `SELECT ` +
		`foreign_table_catalog, foreign_table_schema, foreign_table_name, foreign_server_catalog, foreign_server_name ` +
		`FROM information_schema.foreign_tables WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + order.String() + " LIMIT 1"

	ft := &ForeignTable{}
	err := db.QueryRow(sqlstr, where.Values()...).Scan(&ft.ForeignTableCatalog, &ft.ForeignTableSchema, &ft.ForeignTableName, &ft.ForeignServerCatalog, &ft.ForeignServerName)
	if err != nil {
		return nil, err
	}
	return ft, nil
}

// QueryForeignTable retrieves rows from 'information_schema.foreign_tables' as a slice of ForeignTable.
func QueryForeignTable(db XODB, where WhereClause, order OrderBy) ([]*ForeignTable, error) {
	const origsqlstr = `SELECT ` +
		`foreign_table_catalog, foreign_table_schema, foreign_table_name, foreign_server_catalog, foreign_server_name ` +
		`FROM information_schema.foreign_tables WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + order.String()

	var vals []*ForeignTable
	q, err := db.Query(sqlstr, where.Values()...)
	if err != nil {
		return nil, err
	}
	for q.Next() {
		ft := ForeignTable{}

		err = q.Scan(&ft.ForeignTableCatalog, &ft.ForeignTableSchema, &ft.ForeignTableName, &ft.ForeignServerCatalog, &ft.ForeignServerName)
		if err != nil {
			return nil, err
		}

		vals = append(vals, &ft)
	}
	return vals, nil
}
