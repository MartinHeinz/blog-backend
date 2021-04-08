
## PostgreSQL Database Image

_Notes:_

* Use `POPULATE_DB` variable in `.env` file to change whether database is populated with data or not

---------------------------------

* When repopulating database make sure to delete `go-vue-blog_data` volume otherwise updates of schema or test data won't 
take effect.

---------------------------------

_Repopulating steps:_

```
docker-compose rm  # Remove old containers...
```
_Set `.env` variable to `1` or `0`._
```
docker-compose build  # Builds ONLY blog_db image
docker-compose up blog_db  # Start only blog_db to create schema and optionally populate with test data
```
_Stop docker-compose (CTRL-C)_
```
docker-compose up  # Start all services this time
curl https://localhost:8080/api/v1/<endpoint>/ | jq  # Check presence of populated data...
```
