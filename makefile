 include .env

  create_migration:
    migrate create -ext=sql -dir=internal/database/migrations -seq init

  migrate_up:
    migrate -path=infra/sql/migrations -database "${CONNECTION_STRING}" -verbose up

  migrate_down:
    migrate -path=infra/sql/migrations -database "${CONNECTION_STRING}" -verbose down

  .PHONY: create_migration migrate_up migrate_down