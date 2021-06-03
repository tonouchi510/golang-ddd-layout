# note: call scripts from /scripts

init-db:
	mysql -uroot -p golang_ddd_layout < init/ddl.sql

sqlboiler:
	sqlboiler mysql --config configs/sqlboiler.toml --wipe --add-panic-variants