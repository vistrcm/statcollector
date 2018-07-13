# statcollector
Simple web interface to mongo. Can save records for now.
Currently, it just saving records, may be useful to collect analytics data.

**Does not provide an interface to retrieve data for now.**

## interface
Main endpoint is `/{collection_name}` for example `/users/`.

POST requests to this endpoint accept JSON data.
Application annotates data with a timestamp and saves into the corresponding collection.

## usage example
Here short explanation how to start service and add data to the collections.

### start instance
```bash
$ docker-compose up
```
This will start collector application on port 8080 and MongoDB to store the data.

Check data in the database:
```bash
$ mongo
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
$ http http://localhost:8080/users name=vasily action=registered id=100500
HTTP/1.1 201 Created
Content-Length: 266
Content-Type: application/json; charset=UTF-8
Date: Fri, 13 Jul 2018 02:22:13 GMT

{
    "data": {
        "action": "registered",
        "id": "100500",
        "name": "vasily"
    },
    "raw": "eyJuYW1lIjogInZhc2lseSIsICJhY3Rpb24iOiAicmVnaXN0ZXJlZCIsICJpZCI6ICIxMDA1MDAifQ==",
    "string": "{\"name\": \"vasily\", \"action\": \"registered\", \"id\": \"100500\"}",
    "timestamp": 1531448533554435600
}
```

Check data again:
```bash
$ mongo
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
	"_id" : ObjectId("5b480cd5b2af390b507f5501"),
	"timestamp" : NumberLong("1531448533554435600"),
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
