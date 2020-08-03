db.createUser(
    {
        user: "crypto",
        pwd: "test",
        roles: [
            {
                role: "readWrite",
                db: "cryptotweets"
            }
        ]
    }
)
