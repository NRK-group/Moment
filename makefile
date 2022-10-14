.PHONY: migrateup migratedown createMigration
migrateup:
	migrate -path backend/pkg/db/migrations/sqlite -database sqlite://./backend/social_network.db -verbose up
migratedown:
	migrate -path backend/pkg/db/migrations/sqlite -database sqlite://./backend/social_network.db -verbose down
createMigration:
	migrate create -ext sql -dir backend/pkg/db/migrations/sqlite -seq create_social_network_schema