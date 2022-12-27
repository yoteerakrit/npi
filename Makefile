test:
	echo "Hello World"
migrate-db:
	go run script/migrationdb/main.go migrate

drop-db:
	go run script/migrationdb/main.go drop