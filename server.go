package main

import (
	"context"
	// "go-ent-gqlgen/ent/migrate"
	"go-ent-gqlgen/infra"
	"go-ent-gqlgen/resolver"

	"log"
	"net/http"
	"os"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/debug"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/alecthomas/kong"
)

const defaultPort = "3000"

func main() {
	var cli struct {
		Addr  string `name:"address" default:":3000" help:"Address to listen on."`
		Debug bool   `name:"debug" help:"Enable debugging mode."`
	}
	kong.Parse(&cli)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	client := infra.InitDB()
	defer client.Close()

	if err := client.Schema.Create(
		context.Background(),
		// migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatal("running schema migration", err)
	}

	srv := handler.NewDefaultServer(resolver.NewSchema(client))
	srv.Use(entgql.Transactioner{TxOpener: client})

	if cli.Debug {
		srv.Use(&debug.Tracer{})
	}

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
