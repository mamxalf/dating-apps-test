package infras

import (
	"fmt"
	"time"

	"dating-apps/configs"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

// TransactionBlock contains a transaction block
type TransactionBlock func(db *sqlx.Tx, c chan error)

// PostgresConn wraps a pair of read/write PostgreSQL connections.
type PostgresConn struct {
	Read  *sqlx.DB
	Write *sqlx.DB
}

// ProvidePostgresConn is the provider for PostgresConn.
func ProvidePostgresConn(config *configs.Config) *PostgresConn {
	// actually you can split between read and write depends your requirements :ok:
	return &PostgresConn{
		Read:  CreatePostgresConn("read", config),
		Write: CreatePostgresConn("write", config),
	}
}

// CreatePostgresConn creates a database connection for read access.
func CreatePostgresConn(connType string, config *configs.Config) *sqlx.DB {
	return CreatePostgresDBConnection(
		connType,
		config.DB.PG.User,
		config.DB.PG.Password,
		config.DB.PG.Host,
		config.DB.PG.Port,
		config.DB.PG.Name,
		config.DB.PG.SSLMode,
		config.DB.PG.MaxConnLifetime,
		config.DB.PG.MaxIdleConn,
		config.DB.PG.MaxOpenConn)
}

// CreatePostgresDBConnection creates a database connection.
func CreatePostgresDBConnection(
	connType, username, password, host, port, dbName, sslmode string,
	maxConnLifetime time.Duration,
	maxIdleConn, maxOpenConn int,
) *sqlx.DB {
	conn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host,
		port,
		username,
		password,
		dbName,
		sslmode)

	if password == "" {
		conn = fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s sslmode=%s",
			host,
			port,
			username,
			dbName,
			sslmode)
	}

	db, err := sqlx.Connect("postgres", conn)
	if err != nil {
		log.
			Fatal().
			Err(err).
			Str("type", connType).
			Str("host", host).
			Str("port", port).
			Str("dbName", dbName).
			Msg("Failed connecting to Postgres database")
	} else {
		log.
			Info().
			Str("type", connType).
			Str("host", host).
			Str("port", port).
			Str("dbName", dbName).
			Msg("Connected to Postgres database")
	}

	db.SetConnMaxLifetime(maxConnLifetime)
	db.SetMaxIdleConns(maxIdleConn)
	db.SetMaxOpenConns(maxOpenConn)

	return db
}

// WithTransaction performs queries with transaction
func (m *PostgresConn) WithTransaction(block TransactionBlock) (err error) {
	e := make(chan error)
	tx, err := m.Write.Beginx()
	if err != nil {
		log.Err(err).Msg("error begin transaction")
		return
	}
	go block(tx, e)
	err = <-e
	if err != nil {
		if errTx := tx.Rollback(); errTx != nil {
			log.Err(errTx).Msg("error rollback transaction")
		}
		return
	}
	err = tx.Commit()
	return
}
