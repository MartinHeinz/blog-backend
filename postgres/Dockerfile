FROM postgres:11

# Custom initialization scripts
COPY ./create_db.sh /docker-entrypoint-initdb.d/20-create_db.sh
RUN chmod +x /docker-entrypoint-initdb.d/20-create_db.sh
