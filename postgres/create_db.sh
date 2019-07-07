#!/bin/bash
set -e

POSTGRES="psql --username ${POSTGRES_USER}"

echo "Creating database: ${DB_NAME}"

$POSTGRES <<EOSQL
CREATE DATABASE ${DB_NAME} OWNER ${POSTGRES_USER};
EOSQL

echo "Creating schema..."
psql -d ${DB_NAME} -a -U${POSTGRES_USER} -f /schema.sql

if [ ${POPULATE_DB} -eq 1 ]; then
    echo "Populating database..."
    psql -d ${DB_NAME} -a  -U${POSTGRES_USER} -f /test_data.sql
fi
