initialize:
	go run github.com/novacloudcz/graphql-orm init

generate:
	go run github.com/novacloudcz/graphql-orm

run:
	DATABASE_URL=sqlite3://test.db MEMBER_PROVIDER_URL=http://localhost:8002/graphql PORT=8080 go run *.go
