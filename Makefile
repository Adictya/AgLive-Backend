server:
	gin --appPort 8080 --port 8000

createDb:
	docker exec -it postgres createdb --username=root --owner=root agoraLive

dropDb:
	docker exec -it postgres dropDb agoraLive

migrateUp:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/agoraLive?sslmode=disable" -verbose up

migrateDown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/agoraLive?sslmode=disable" -verbose down

migrateUp1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/agoraLive?sslmode=disable" -verbose up 1

migrateDown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/agoraLive?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

server:
	gin --appPort 8080 --port 8000

db:
	docker start postgres -a

.Phony:
	postgres createdb dropDb migrateUp migrateUp1 migrateDown migrateDown1 sqlc server db
