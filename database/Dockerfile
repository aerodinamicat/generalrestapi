FROM postgres:latest

COPY initialQuery.sql /docker-entrypoint-initdb.d/1-initialQuery.sql

EXPOSE 5432

CMD ["postgres"]