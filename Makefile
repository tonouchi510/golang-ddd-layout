# note: call scripts from /scripts

init-db:
	mysql -uroot -p golang_ddd_layout < init/ddl.sql

sqlboiler:
	sqlboiler mysql --config configs/sqlboiler.toml --wipe

run-test:
	go test -v -cover -covermode=atomic ./internal/...
