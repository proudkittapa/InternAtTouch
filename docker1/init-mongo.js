db.createUser(
    {
        user: "kittapa",
        pwd: "hello",
        roles:[
            {
            role: "readWrite", 
            db: "kittapa"
            }
        ]
    }
)