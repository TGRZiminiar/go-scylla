## Docker pull
```
docker pull scylladb/scylla:5.2
docker compose up -d
```
## Show status
After the setup is done you will see two node running on DC1
```
docker exec -it scylla-node1 nodetool status
```


## Access CQl to generate some keyspace and table
```
docker exec -it scylla-node1 cqlsh
```

## Create datacenter and initial table
// users = keyspace
```
CREATE KEYSPACE users WITH REPLICATION = { 'class' : 'NetworkTopologyStrategy','DC1' : 3};
use users;
CREATE TABLE IF NOT EXISTS users.userData (
    id UUID PRIMARY KEY,
    name TEXT,
    email TEXT
);
describe users.userdata;
```

## Check that datacenter is working
If the data of the tables users show that mean datacenter work
```
docker exec -it scylla-node2 cqlsh
describe users.userdata;
```

## To test it with the api you need everything to be on the same network
```

// server image
docker run -d --net=test-scylla_web --name some-go-app go-app

// for sending api
docker run -it --name my-ubuntu-container --network test-scylla_web ubuntu:noble-20240114

// save user in db
curl --location 'some-go-app:5000/users_v1/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name":"mix1",
    "email":"mix1@gmail.com"
}'

// get users in db
curl --location 'some-go-app:5000/users_v1/get-users'
```


