package Database

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Superhero_q struct {
	ID 			int		`bson:"ID"`
	Name  		string	`bson:"Name"`
	Actual_name string	`bson:"Actual_name"`
	Gender 		string	`bson:"Gender"`
	Age 		int		`bson:"Age"`
	Super_power string	`bson:"Super_power"`
}


func max_id() int {
	var result bson.M
	var curr_id Superhero_q
	opts := options.FindOne().SetSort(bson.D{{"ID", -1}})
	err := Coll.FindOne(Ctx, bson.D{{}}, opts).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	bsonBytes, _ := bson.Marshal(result)
	bson.Unmarshal(bsonBytes, &curr_id)
	//fmt.Println("Current ID", curr_id.ID)
	//defer client.Disconnect(ctx)
	return curr_id.ID
}

//insert() receive all or not all at least have to be name receive all as a struct
//update by name and update by id
//delete by ID// Delete by name // view by ID

