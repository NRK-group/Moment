# start with base image
FROM mysql:8.0.32

# import data into container
# All scripts in docker-entrypoint-initdb.d/ are automatically executed during container startup
COPY ./db/migrations/*.up.sql /docker-entrypoint-initdb.d/