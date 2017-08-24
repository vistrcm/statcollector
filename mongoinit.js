use admin

db.createUser(
    {
        user: "admin",
        pwd: "Uqua7aev7ailit8Oovie5goo8uaZeu",
        roles: [ { role: "userAdminAnyDatabase", db: "admin" } ]
    }
)

----

use stats

db.createUser(
    {
        user: "collector",
        pwd: "Ci1aTh1ooshiib6iepha4oongaeSho",
        roles: [ { role: "readWrite", db: "stats" }]
    }
)