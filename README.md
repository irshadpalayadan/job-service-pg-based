# job-service-pg-based

steps to initialize the project

1. go mod init github.com/irshadpalayadan/job-service-pg-based
2. go get github.com/99designs/gqlgen
3. printf '// +build tools\npackage tools\nimport (_ "github.com/99designs/gqlgen"\n _ "github.com/99designs/gqlgen/graphql/introspection")' | gofmt > tools.go
4. go mod tidy
5. go run github.com/99designs/gqlgen init
6. go run server.go
7. remove the example todo code
8. add the relevant schema changes
9. go run github.com/99designs/gqlgen generate


