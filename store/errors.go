package store

import "github.com/go-sql-driver/mysql"

var DuplicateEntryError = mysql.MySQLError{1062, [5]byte{}, ""}
