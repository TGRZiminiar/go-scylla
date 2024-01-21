docker pull scylladb/scylla:5.2

docker run --name scylla-local -p 9042:9042 -d scylladb/scylla

## Show status
docker exec -it db-scylla nodetool status

## Access CQl
docker exec -it db-scylla cqlsh

users = keyspace
CREATE KEYSPACE users WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1};

CREATE TABLE IF NOT EXISTS users.userData (
    id UUID PRIMARY KEY,
    name TEXT,
    email TEXT
);
<!-- insert into userData ("first_name","last_name","address","picture_location") VALUES ('Bob','Loblaw','1313 Mockingbird Lane', 'http://www.facebook.com/bobloblaw'); -->
