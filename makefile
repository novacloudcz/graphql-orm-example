initialize:
	go run github.com/novacloudcz/graphql-orm init

generate:
	go run github.com/novacloudcz/graphql-orm

run:
	DATABASE_URL=sqlite3://test.db PORT=8080 go run *.go

heroku-deploy:
	heroku container:push --app graphql-orm-example web
	heroku container:release --app graphql-orm-example web