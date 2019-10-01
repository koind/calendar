db-create-table:
	 psql 'host=localhost user=calendar password=123123 dbname=calendar' < scripts/001.sql

db-set-index:
	psql 'host=localhost user=calendar password=123123 dbname=calendar' < scripts/002.sql