{{/*
This template is rendered at the top of every file generated by xo.
Unused imports here will be stripped out by go-imports.
*/}}// Package {{ .Package }} contains the types for schema '{{ schema .Schema }}'.
package {{ .Package }}

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"database/sql/driver"
	"encoding/csv"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/lib/pq"
	"github.com/lib/pq/hstore"
)

