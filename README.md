# statcollector
Simple web interface to mongo. Currently it just saving records, may be useful to collect analytics data.

**Does not provide an interface to retrieve data for now.**

## interface
Main endpoint is `/{collection_name}` for example `/users`.

POST requests to this endpoint accept JSON data.
Application annotates data with a timestamp and saves into the corresponding collection.

Database to store data is specified by `mongoUrl` parameter. Default `mongodb://mongo/stats`. 

## usage example
Here short explanation how to start service and add data to the collections.

### start instance
```bash
$ docker-compose up
```
This will start collector application on port 8080 and MongoDB to store the data.

Check data in the database by starting mongo CLI inside mongo container:
```bash
$ docker-compose exec mongo mongo
MongoDB shell version v4.0.0
connecting to: mongodb://127.0.0.1:27017
MongoDB server version: 4.0.0
....

> show databases;
admin   0.000GB
config  0.000GB
local   0.000GB
>
```
You can see default mongo setup here.

Let's send next JSON into collection "users"
```json
{
  "name": "vasily",
  "action": "registered",
  "id": 100500
}
```
will use [httpie](https://httpie.org/) for this:
```bash
$ http -v http://localhost:8080/users\?sendback\=true name=vasily action=registered id=100500
  POST /users HTTP/1.1
  Accept: application/json, */*
  Accept-Encoding: gzip, deflate
  Connection: keep-alive
  Content-Length: 58
  Content-Type: application/json
  Host: localhost:8080
  User-Agent: HTTPie/0.9.9
  
  {
      "action": "registered",
      "id": "100500",
      "name": "vasily"
  }
  
  HTTP/1.1 201 Created
  Content-Length: 266
  Content-Type: application/json; charset=UTF-8
  Date: Fri, 13 Jul 2018 02:54:56 GMT
  
  {
      "data": {
          "action": "registered",
          "id": "100500",
          "name": "vasily"
      },
      "raw": "eyJuYW1lIjogInZhc2lseSIsICJhY3Rpb24iOiAicmVnaXN0ZXJlZCIsICJpZCI6ICIxMDA1MDAifQ==",
      "string": "{\"name\": \"vasily\", \"action\": \"registered\", \"id\": \"100500\"}",
      "timestamp": 1531450496716258600
  }
```
Note `sendback` parameter is set to `true`. Without it, server returns an empty response body.


Check data again:
```bash
$ docker-compose exec mongo mongo
MongoDB shell version v4.0.0
connecting to: mongodb://127.0.0.1:27017
MongoDB server version: 4.0.0
...

> show databases;
admin   0.000GB
config  0.000GB
local   0.000GB
stats   0.000GB
>
```

New database `stats` created. Let's examine it in mongo console:
```
> use stats;
switched to db stats
> show collections;
users
> db.users.find().pretty()
{
	"_id" : ObjectId("5b4814803c37e523005bfaed"),
	"timestamp" : NumberLong("1531450496716258600"),
	"raw" : BinData(0,"eyJuYW1lIjogInZhc2lseSIsICJhY3Rpb24iOiAicmVnaXN0ZXJlZCIsICJpZCI6ICIxMDA1MDAifQ=="),
	"string" : "{\"name\": \"vasily\", \"action\": \"registered\", \"id\": \"100500\"}",
	"data" : {
		"name" : "vasily",
		"action" : "registered",
		"id" : "100500"
	}
}
>
```

## setup mongo auth
Start mongo with additional environment variables to set admin password:
```bash
docker-compose run -e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=somepass mongo
```

Wait for next lines
```text
MongoDB init process complete; ready for start up.
...
...
...
2018-07-21T00:57:28.027+0000 I NETWORK  [initandlisten] waiting for connections on port 27017
```

Connect in a separate terminal. We are using `statcollector_mongo_run_1` - id of the container from docker run above.
```bash
docker exec -it statcollector_mongo_run_1 mongo -u admin -p somepass admin
```

In the mongo shell create user for database `stats`
```
> use stats;
switched to db stats
> db.createUser(
      {
        user: "statsusr",
        pwd: "aicaiR1iivahToh7cei7reeseeyaer",
        roles: [ { role: "readWrite", db: "stats" }]
      }
    )
Successfully added user: {
	"user" : "statsusr",
	"roles" : [
		{
			"role" : "readWrite",
			"db" : "stats"
		}
	]
}
> exit
bye
```

This will close mongo shell.

Please stop mongo started on the first step by pressing `Control-C` in the first terminal.

After that action, you can use `docker-compose up` to start the application.
