#!/bin/bash

# Regenerate our database mock
mockgen -source=./database/db/querier.go -destination=./internal/mocks/querier.go -package=mocks
