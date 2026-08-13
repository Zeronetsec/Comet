// https://github.com/Zeronetsec/Comet

package sqlscan

var sqlSignatures = []string{
    "SQL syntax",
    "mysql_fetch",
    "MySQLSyntaxErrorException",
    "valid MySQL result",
    "PostgreSQL query failed",
    "PG::SyntaxError",
    "Microsoft OLE DB Provider for SQL Server",
    "Unclosed quotation mark after the character string",
    "SQLServerException",
    "Oracle error",
    "ORA-01756",
    "SQLite/JDBCDriver",
    "System.Data.SQLite.SQLiteException",
    "Warning: sqlite_",
}

// Copyright (c) 2026 Zeronetsec