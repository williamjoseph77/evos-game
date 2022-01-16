package postgres

import (
	"context"
	"log"
	"os"

	"github.com/go-pg/pg/v10"
)

type DBLogger struct{}

func (d DBLogger) BeforeQuery(ctx context.Context, evt *pg.QueryEvent) (context.Context, error) {
	q, err := evt.FormattedQuery()
	if err != nil {
		return nil, err
	}

	if evt.Err != nil {
		log.Printf("Error %s executing query:\n%s\n", evt.Err, q)
	} else {
		log.Printf("%s", q)
	}

	return ctx, nil
}

func (d DBLogger) AfterQuery(ctx context.Context, evt *pg.QueryEvent) error {
	return nil
}

func Connect() *pg.DB {
	return pg.Connect(&pg.Options{
		Addr:     os.Getenv("DB_ADDRESS"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_DATABASE"),
	})
}
