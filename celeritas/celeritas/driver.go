package celeritas

import (
    "database/sql"
    "github.com/jackc/pgconn"
    "github.com/jackc/pgx/v4"
    "github.com/jackc/pgx/v4/stdlib"
)

func (c *Celeritas) OpenDB(dbType, dsn string) (*sql.DB, error) {
    if dbType == "postgres" || dbType = "postgresql" {
        dbType = "pgx"
    }

    db, err := sql.Open(dbType, dsn)
    if err != nil {
        return nil, err
    }

    err = db.Ping()
    if err != nil {
        return nil, err
    }

    return db, nil
}
