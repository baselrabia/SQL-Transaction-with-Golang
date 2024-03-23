package sql

import (
	"fmt"

	"github.com/baselrabia/SQL-Transaction-with-Golang/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PgConnectionManager struct {
	db *sqlx.DB
}

func NewPgConnectionManager(config config.EnvConfig) (*PgConnectionManager, error) {

	// Construct the database connection string using values from EnvConfig
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.PgHost,
		config.PgPort,
		config.PgUser,
		config.PgPassword,
		config.PgDatabase,
		config.PgSslMode,
	)
	// Open a connection to the database
	db, err := sqlx.Open("postgres", connectionString)
	if err != nil {
		//log error
		return nil, err
	}

	// Ping the database to verify the connection
	if err := db.Ping(); err != nil {
		//log the error
		return nil, err
	}

	return &PgConnectionManager{
		db: db,
	}, nil
}

func (cm *PgConnectionManager) Close() error {
	return cm.db.Close()
}

func (cm *PgConnectionManager) GetQuery() *SingleInstruction {
	return NewSingleInstruction(cm.db)
}

func (cm *PgConnectionManager) GetTransaction() *MultiInstruction {
	return NewMultiInstruction(cm.db)
}
