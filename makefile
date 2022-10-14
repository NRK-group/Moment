.PHONY: migrateup migratedown
migrateup:
	migrate -path backend/pkg/db/migrations/sqlite -database sqlite://./backend/social_network.db -verbose up
migratedown:
	migrate -path backend/pkg/db/migrations/sqlite -database sqlite://./backend/social_network.db -verbose down