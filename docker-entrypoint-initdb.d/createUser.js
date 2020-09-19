db.createUser(
    {
        user: "collector",
        pwd: "Ci1aTh1ooshiib6iepha4oongaeSho",
        roles: [
            {
                role: "readWrite",
                db: "stats"
            }
        ]
    }
);