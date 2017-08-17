// Package pg contains the types for schema 'information_schema'.
package pg

// GENERATED BY XO. DO NOT EDIT.

// RoleColumnGrantTable is the database name for the table.
const RoleColumnGrantTable = "information_schema.role_column_grants"

// RoleColumnGrant represents a row from 'information_schema.role_column_grants'.
type RoleColumnGrant struct {
	Grantor       SQLIdentifier `json:"grantor"`        // grantor
	Grantee       SQLIdentifier `json:"grantee"`        // grantee
	TableCatalog  SQLIdentifier `json:"table_catalog"`  // table_catalog
	TableSchema   SQLIdentifier `json:"table_schema"`   // table_schema
	TableName     SQLIdentifier `json:"table_name"`     // table_name
	ColumnName    SQLIdentifier `json:"column_name"`    // column_name
	PrivilegeType CharacterData `json:"privilege_type"` // privilege_type
	IsGrantable   YesOrNo       `json:"is_grantable"`   // is_grantable
}

// Constants defining each column in the table.
const (
	RoleColumnGrantGrantorField       = "grantor"
	RoleColumnGrantGranteeField       = "grantee"
	RoleColumnGrantTableCatalogField  = "table_catalog"
	RoleColumnGrantTableSchemaField   = "table_schema"
	RoleColumnGrantTableNameField     = "table_name"
	RoleColumnGrantColumnNameField    = "column_name"
	RoleColumnGrantPrivilegeTypeField = "privilege_type"
	RoleColumnGrantIsGrantableField   = "is_grantable"
)

// WhereClauses for every type in RoleColumnGrant.
var (
	RoleColumnGrantGrantorWhere       SQLIdentifierField = "grantor"
	RoleColumnGrantGranteeWhere       SQLIdentifierField = "grantee"
	RoleColumnGrantTableCatalogWhere  SQLIdentifierField = "table_catalog"
	RoleColumnGrantTableSchemaWhere   SQLIdentifierField = "table_schema"
	RoleColumnGrantTableNameWhere     SQLIdentifierField = "table_name"
	RoleColumnGrantColumnNameWhere    SQLIdentifierField = "column_name"
	RoleColumnGrantPrivilegeTypeWhere CharacterDataField = "privilege_type"
	RoleColumnGrantIsGrantableWhere   YesOrNoField       = "is_grantable"
)

// QueryOneRoleColumnGrant retrieves a row from 'information_schema.role_column_grants' as a RoleColumnGrant.
func QueryOneRoleColumnGrant(db XODB, where WhereClause, order OrderBy) (*RoleColumnGrant, error) {
	const origsqlstr = `SELECT ` +
		`grantor, grantee, table_catalog, table_schema, table_name, column_name, privilege_type, is_grantable ` +
		`FROM information_schema.role_column_grants WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + order.String() + " LIMIT 1"

	rcg := &RoleColumnGrant{}
	err := db.QueryRow(sqlstr, where.Values()...).Scan(&rcg.Grantor, &rcg.Grantee, &rcg.TableCatalog, &rcg.TableSchema, &rcg.TableName, &rcg.ColumnName, &rcg.PrivilegeType, &rcg.IsGrantable)
	if err != nil {
		return nil, err
	}
	return rcg, nil
}

// QueryRoleColumnGrant retrieves rows from 'information_schema.role_column_grants' as a slice of RoleColumnGrant.
func QueryRoleColumnGrant(db XODB, where WhereClause, order OrderBy) ([]*RoleColumnGrant, error) {
	const origsqlstr = `SELECT ` +
		`grantor, grantee, table_catalog, table_schema, table_name, column_name, privilege_type, is_grantable ` +
		`FROM information_schema.role_column_grants WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + order.String()

	var vals []*RoleColumnGrant
	q, err := db.Query(sqlstr, where.Values()...)
	if err != nil {
		return nil, err
	}
	for q.Next() {
		rcg := RoleColumnGrant{}

		err = q.Scan(&rcg.Grantor, &rcg.Grantee, &rcg.TableCatalog, &rcg.TableSchema, &rcg.TableName, &rcg.ColumnName, &rcg.PrivilegeType, &rcg.IsGrantable)
		if err != nil {
			return nil, err
		}

		vals = append(vals, &rcg)
	}
	return vals, nil
}
