package main

import (
	"context"

	"github.com/baselrabia/SQL-Transaction-with-Golang/config"
	"github.com/baselrabia/SQL-Transaction-with-Golang/pkg/sql"
)

func main() {

	// Create a new configuration instance
	config := config.NewEnvConfig()

	pgConn, err := sql.NewPgConnectionManager(config)
	if err != nil {
		// error 
	}

	ctx := context.Background()
	pgConn.GetTransaction().Begin(ctx)

	pgConn.GetQuery().Query(ctx)
	pgConn.GetQuery().Select(ctx)
	pgConn.GetQuery().Exec(ctx)


	pgConn.GetTransaction().Commit(ctx)

	pgConn.GetTransaction().Rollback(ctx)


	


 
}
