#!/bin/bash

set -e
set -u

function create_database() {
	local database=$1
	echo "  Creating database '$database'"
	psql -v ON_ERROR_STOP=1 --username "postgres" <<-EOSQL
	    CREATE DATABASE $database;
	    \c $database;
	    CREATE SCHEMA cezedd;
      GRANT ALL PRIVILEGES ON DATABASE $database TO postgres;
      GRANT ALL ON ALL TABLES IN SCHEMA cezedd TO postgres;
      GRANT USAGE ON SCHEMA cezedd TO postgres;
      ALTER ROLE postgres SET search_path = public;
EOSQL
}

if [ -n "$POSTGRES_MULTIPLE_DATABASES" ]; then
	echo "Multiple database creation requested: $POSTGRES_MULTIPLE_DATABASES"
	for db in $(echo $POSTGRES_MULTIPLE_DATABASES | tr ',' ' '); do
		create_database $db
	done
	echo "Multiple databases created"
fi

#cezedd-test