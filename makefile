.PHONY: migrateup migratedown createMigration run backendDocker frontendDocker
createMigration:
	migrate create -ext sql -dir backend/pkg/db/migrations/sqlite -seq create_social_network_schema
rungo:
	cd backend && go run server.go
rundev:
	cd frontend && npm run dev
runtest:
	rm -rf backend/Test/social_network_test.db
	cd backend/Test && go test -v .
backendDocker:
	docker build -t golang-server ./backend
	docker run --rm -it  -p 5070:5070/tcp golang-server:latest
frontendDocker:
	docker build -t react-frontend-app ./frontend
	docker run --rm -it -p 80:80/tcp -p 8070:8070/tcp react-frontend-app:latest