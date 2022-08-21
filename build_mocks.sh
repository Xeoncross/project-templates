#!/bin/bash

# mockgen database/db/querier.go Querier -destination=../mocks/mock_querier.go -package=mocks

# mockgen -source=database/db/querier.go Querier -package=mocks -destination=internal/mocks/mock_querier.go

mockgen -source=./database/db/querier.go -destination=./internal/mocks/querier.go -package=mocks
