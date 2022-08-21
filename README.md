## Echo + sqlc Go usage

This is an example of using [sqlc](https://github.com/kyleconroy/sqlc) for model
generation and [echo](https://echo.labstack.com/) for the web framework. 

This includes a [service/](./internal/service/) folder in which you would define
your core business services without the implementation details. That means no
SQL, no http handlers, no cli flags, no redis cache wrappers, etc. Doing so 
means it's easy to reuse the same logic for any kind of transports or interfaces
such as gRPC, REST, GraphQL, unit tests, CLI, or carrier pigeon. See 
[SOLID design in Go](https://dave.cheney.net/2016/08/20/solid-go-design).

Mocking is also also demonstrated using [gomock](https://github.com/golang/mock)
in the hopes that the unit tests would be easier to construct without needing a
live database connection. You can also write mocks by hand. The important thing
is that you rely on the interfaces _not_ the concret classes (use db.Querier vs 
db.Queries).
