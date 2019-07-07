FROM postgres:11

# Custom initialization scripts
COPY ./create_db.sh /docker-entrypoint-initdb.d/20-create_db.sh
COPY schema.sql /schema.sql
COPY test_data.sql /test_data.sql

RUN chmod +x /docker-entrypoint-initdb.d/20-create_db.sh
