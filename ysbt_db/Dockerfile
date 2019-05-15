FROM postgres:11-alpine

ENV POSTGRES_USER ysbt_db
ENV POSTGRES_PASSWORD password

COPY *.sql /docker-entrypoint-initdb.d/
