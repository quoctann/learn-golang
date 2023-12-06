# Here is a description from gqlgen about the generated files:

- gqlgen.yml — The gqlgen config file, knobs for controlling the generated code.

- graph/generated/generated.go — The GraphQL execution runtime, the bulk of the generated code.

- graph/model/models_gen.go — Generated models required to build the graph. Often you will override these with your own models. Still very useful for input types.

- graph/schema.graphqls — This is the file where you will add GraphQL schemas.

- graph/schema.resolvers.go — This is where your application code lives. generated.go will call into this to get the data the user has requested.

- server.go — This is a minimal entry point that sets up an http.Handler to the generated GraphQL server. start the server with go run server.go and open your browser and you should see the graphql playground, So setup is right!

# Migrate

go get -u github.com/go-sql-driver/mysql

go build -tags 'mysql' -ldflags="-X main.Version=1.0.0" -o $GOPATH/bin/migrate github.com/golang-migrate/migrate/v4/cmd/migrate/

cd internal/pkg/db/migrations/

migrate create -ext sql -dir mysql -seq create_users_table

migrate create -ext sql -dir mysql -seq create_links_table
