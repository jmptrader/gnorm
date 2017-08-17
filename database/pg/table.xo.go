// Package pg contains the types for schema 'information_schema'.
package pg

// GENERATED BY XO. DO NOT EDIT.

// TableTable is the database name for the table.
const TableTable = "information_schema.tables"

// Table represents a row from 'information_schema.tables'.
type Table struct {
	TableCatalog              SQLIdentifier `json:"table_catalog"`                // table_catalog
	TableSchema               SQLIdentifier `json:"table_schema"`                 // table_schema
	TableName                 SQLIdentifier `json:"table_name"`                   // table_name
	TableType                 CharacterData `json:"table_type"`                   // table_type
	SelfReferencingColumnName SQLIdentifier `json:"self_referencing_column_name"` // self_referencing_column_name
	ReferenceGeneration       CharacterData `json:"reference_generation"`         // reference_generation
	UserDefinedTypeCatalog    SQLIdentifier `json:"user_defined_type_catalog"`    // user_defined_type_catalog
	UserDefinedTypeSchema     SQLIdentifier `json:"user_defined_type_schema"`     // user_defined_type_schema
	UserDefinedTypeName       SQLIdentifier `json:"user_defined_type_name"`       // user_defined_type_name
	IsInsertableInto          YesOrNo       `json:"is_insertable_into"`           // is_insertable_into
	IsTyped                   YesOrNo       `json:"is_typed"`                     // is_typed
	CommitAction              CharacterData `json:"commit_action"`                // commit_action
}

// Constants defining each column in the table.
const (
	TableTableCatalogField              = "table_catalog"
	TableTableSchemaField               = "table_schema"
	TableTableNameField                 = "table_name"
	TableTableTypeField                 = "table_type"
	TableSelfReferencingColumnNameField = "self_referencing_column_name"
	TableReferenceGenerationField       = "reference_generation"
	TableUserDefinedTypeCatalogField    = "user_defined_type_catalog"
	TableUserDefinedTypeSchemaField     = "user_defined_type_schema"
	TableUserDefinedTypeNameField       = "user_defined_type_name"
	TableIsInsertableIntoField          = "is_insertable_into"
	TableIsTypedField                   = "is_typed"
	TableCommitActionField              = "commit_action"
)

// WhereClauses for every type in Table.
var (
	TableTableCatalogWhere              SQLIdentifierField = "table_catalog"
	TableTableSchemaWhere               SQLIdentifierField = "table_schema"
	TableTableNameWhere                 SQLIdentifierField = "table_name"
	TableTableTypeWhere                 CharacterDataField = "table_type"
	TableSelfReferencingColumnNameWhere SQLIdentifierField = "self_referencing_column_name"
	TableReferenceGenerationWhere       CharacterDataField = "reference_generation"
	TableUserDefinedTypeCatalogWhere    SQLIdentifierField = "user_defined_type_catalog"
	TableUserDefinedTypeSchemaWhere     SQLIdentifierField = "user_defined_type_schema"
	TableUserDefinedTypeNameWhere       SQLIdentifierField = "user_defined_type_name"
	TableIsInsertableIntoWhere          YesOrNoField       = "is_insertable_into"
	TableIsTypedWhere                   YesOrNoField       = "is_typed"
	TableCommitActionWhere              CharacterDataField = "commit_action"
)

// QueryOneTable retrieves a row from 'information_schema.tables' as a Table.
func QueryOneTable(db XODB, where WhereClause, order OrderBy) (*Table, error) {
	const origsqlstr = `SELECT ` +
		`table_catalog, table_schema, table_name, table_type, self_referencing_column_name, reference_generation, user_defined_type_catalog, user_defined_type_schema, user_defined_type_name, is_insertable_into, is_typed, commit_action ` +
		`FROM information_schema.tables WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + order.String() + " LIMIT 1"

	t := &Table{}
	err := db.QueryRow(sqlstr, where.Values()...).Scan(&t.TableCatalog, &t.TableSchema, &t.TableName, &t.TableType, &t.SelfReferencingColumnName, &t.ReferenceGeneration, &t.UserDefinedTypeCatalog, &t.UserDefinedTypeSchema, &t.UserDefinedTypeName, &t.IsInsertableInto, &t.IsTyped, &t.CommitAction)
	if err != nil {
		return nil, err
	}
	return t, nil
}

// QueryTable retrieves rows from 'information_schema.tables' as a slice of Table.
func QueryTable(db XODB, where WhereClause, order OrderBy) ([]*Table, error) {
	const origsqlstr = `SELECT ` +
		`table_catalog, table_schema, table_name, table_type, self_referencing_column_name, reference_generation, user_defined_type_catalog, user_defined_type_schema, user_defined_type_name, is_insertable_into, is_typed, commit_action ` +
		`FROM information_schema.tables WHERE (`

	idx := 1
	sqlstr := origsqlstr + where.String(&idx) + ") " + order.String()

	var vals []*Table
	q, err := db.Query(sqlstr, where.Values()...)
	if err != nil {
		return nil, err
	}
	for q.Next() {
		t := Table{}

		err = q.Scan(&t.TableCatalog, &t.TableSchema, &t.TableName, &t.TableType, &t.SelfReferencingColumnName, &t.ReferenceGeneration, &t.UserDefinedTypeCatalog, &t.UserDefinedTypeSchema, &t.UserDefinedTypeName, &t.IsInsertableInto, &t.IsTyped, &t.CommitAction)
		if err != nil {
			return nil, err
		}

		vals = append(vals, &t)
	}
	return vals, nil
}
