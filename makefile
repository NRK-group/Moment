.PHONY: createMigration runTest build openPPT
createMigration:
	migrate create -ext sql -dir backend/pkg/db/migrations/sqlite -seq create_social_network_schema
runTest:
	rm -rf backend/Test/social_network_test.db
	cd backend/Test && go test -v .
build:
	docker compose up --build
openPPT:
	marp slides/slides.md --preview