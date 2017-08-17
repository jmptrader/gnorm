// Package pg contains the types for schema 'information_schema'.
package pg

// GENERATED BY XO. DO NOT EDIT.

// ConstraintColumnUsageTable is the database name for the table.
const ConstraintColumnUsageTable = "information_schema.constraint_column_usage"

// ConstraintColumnUsage represents a row from 'information_schema.constraint_column_usage'.
type ConstraintColumnUsage struct {
	TableCatalog      SQLIdentifier `json:"table_catalog"`      // table_catalog
	TableSchema       SQLIdentifier `json:"table_schema"`       // table_schema
	TableName         SQLIdentifier `json:"table_name"`         // table_name
	ColumnName        SQLIdentifier `json:"column_name"`        // column_name
	ConstraintCatalog SQLIdentifier `json:"constraint_catalog"` // constraint_catalog
	ConstraintSchema  SQLIdentifier `json:"constraint_schema"`  // constraint_schema
	ConstraintName    SQLIdentifier `json:"constraint_name"`    // constraint_name
}

// Constants defining each column in the table.
const (
	ConstraintColumnUsageTableCatalogField      = "table_catalog"
	ConstraintColumnUsageTableSchemaField       = "table_schema"
	ConstraintColumnUsageTableNameField         = "table_name"
	ConstraintColumnUsageColumnNameField        = "column_name"
	ConstraintColumnUsageConstraintCatalogField = "constraint_catalog"
	ConstraintColumnUsageConstraintSchemaField  = "constraint_schema"
	ConstraintColumnUsageConstraintNameField    = "constraint_name"
)

// WhereClauses for every type in ConstraintColumnUsage.
var (
	ConstraintColumnUsageTableCatalogWhere      SQLIdentifierField = "table_catalog"
	ConstraintColumnUsageTableSchemaWhere       SQLIdentifierField = "table_schema"
	ConstraintColumnUsageTableNameWhere         SQLIdentifierField = "table_name"
	ConstraintColumnUsageColumnNameWhere        SQLIdentifierField = "column_name"
	ConstraintColumnUsageConstraintCatalogWhere SQLIdentifierField = "constraint_catalog"
	ConstraintColumnUsageConstraintSchemaWhere  SQLIdentifierField = "constraint_schema"
	ConstraintColumnUsageConstraintNameWhere    SQLIdentifierField = "constraint_name"
)

// QueryOneConstraintColumnUsage retrieves a row from 'information_schema.constraint_column_usage' as a ConstraintColumnUsage.
func QueryOneConstraintColumnUsage(db XODB, where WhereClause, order OrderBy) (*ConstraintColumnUsage, error) {
	const origsqlstr = `SELECT ` +
		`table_catalog, table_schema, table_name, column_name, constraint_catalog, constraint_schema, constraint_name ` +
		`FROM information_schema.constraint_column_usage WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + order.String() + " LIMIT 1"

	ccu := &ConstraintColumnUsage{}
	err := db.QueryRow(sqlstr, where.Values()...).Scan(&ccu.TableCatalog, &ccu.TableSchema, &ccu.TableName, &ccu.ColumnName, &ccu.ConstraintCatalog, &ccu.ConstraintSchema, &ccu.ConstraintName)
	if err != nil {
		return nil, err
	}
	return ccu, nil
}

// QueryConstraintColumnUsage retrieves rows from 'information_schema.constraint_column_usage' as a slice of ConstraintColumnUsage.
func QueryConstraintColumnUsage(db XODB, where WhereClause, order OrderBy) ([]*ConstraintColumnUsage, error) {
	const origsqlstr = `SELECT ` +
		`table_catalog, table_schema, table_name, column_name, constraint_catalog, constraint_schema, constraint_name ` +
		`FROM information_schema.constraint_column_usage WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + order.String()

	var vals []*ConstraintColumnUsage
	q, err := db.Query(sqlstr, where.Values()...)
	if err != nil {
		return nil, err
	}
	for q.Next() {
		ccu := ConstraintColumnUsage{}

		err = q.Scan(&ccu.TableCatalog, &ccu.TableSchema, &ccu.TableName, &ccu.ColumnName, &ccu.ConstraintCatalog, &ccu.ConstraintSchema, &ccu.ConstraintName)
		if err != nil {
			return nil, err
		}

		vals = append(vals, &ccu)
	}
	return vals, nil
}
