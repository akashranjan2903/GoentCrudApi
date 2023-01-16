package db

import (
	"context"
	"log"

	_ "github.com/gocrud/ent/runtime"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/gocrud/config"
	"github.com/gocrud/ent"
	_ "github.com/lib/pq"
)

// func main() {
// 	client, err := ent.Open("postgres", "host=localhost port=8080 user=postgres dbname=DemoDatabase password=@kash123 sslmode=disable")
// 	if err != nil {
// 		log.Fatalf("failed opening connection to postgres: %v", err)
// 	}
// 	defer client.Close()
// 	// Run the auto migration tool.
// 	if err := client.Schema.Create(context.Background()); err != nil {
// 		log.Fatalf("failed creating schema resources: %v", err)
// 	}
// }

func NewSqlClient() *ent.Client {
	conf := config.GetConfig()

	// Set ent debug mode if enabled in config
	entOptions := []ent.Option{}
	if conf.App.Env == "dev" {
		entOptions = append(entOptions, ent.Debug())
	}

	// Open database connection
	client, err := ent.Open(dialect.Postgres, conf.Database.Dsn, entOptions...)
	if err != nil {
		log.Fatal(err)
	}

	opts := []schema.MigrateOption{
		schema.WithDropColumn(true),
		schema.WithDropIndex(true),
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background(), opts...); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Run the Client.
	return client
}
