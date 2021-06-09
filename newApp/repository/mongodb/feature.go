package mongodb

import (
	goxid "github.com/touchtechnologies-product/xid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"reflect"
)

func udstr(id string, key string, valStr string) {
	_, err := Coll.UpdateOne(
		Ctx,
		bson.M{"_id": id},
		bson.D{
			{"$set", bson.D{{key, valStr}}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}

func (figure *InsertQ)RunQ(){
	initID := goxid.New()
	idGen := initID.Gen()
	_, err := Coll.InsertOne(Ctx, bson.D{
		{"_id", idGen},
		{"name", figure.Name},
		{"actual_name", figure.ActualName},
		{"gender", figure.Gender},
		{"birth_date", figure.BirthDate},
		{"height", figure.Height},
		{"super_power", figure.SuperPower},
		{"alive", figure.Alive},
	})
	if err != nil {
		log.Fatal("err3: ", err)
	}
}

func (d *DeleteQ)RunQ() {
	_, err := Coll.DeleteOne(Ctx, bson.M{"_id": d.ID})
	if err != nil {
		log.Fatal(err)
	}
}

func (figure *UpdateQ)RunQ() {
	v := ViewQ{ID : figure.ID}
	origin := v.RunView()
	if figure.Name != origin.Name {
		udstr(figure.ID, "name", figure.Name)
	}
	if figure.ActualName != origin.ActualName {
		udstr(figure.ID, "actual_name", figure.ActualName)
	}
	if figure.Gender != "" {
		udstr(figure.ID, "gender", figure.Gender)
	}
	if figure.Height != origin.Height {
		_, err := Coll.UpdateOne(
			Ctx,
			bson.M{"_id": figure.ID},
			bson.D{
				{"$set", bson.D{{"height", figure.Height}}},
			},
		)
		if err != nil {
			log.Fatal(err)
		}
	}
	if !reflect.DeepEqual(figure.SuperPower, origin.SuperPower) {
		_, err := Coll.UpdateOne(
			Ctx,
			bson.M{"_id": figure.ID},
			bson.D{
				{"$set", bson.D{{"super_power", bson.A{figure.SuperPower}}}},
			},
		)
		if err != nil {
			log.Fatal(err)
		}
	}
	if figure.BirthDate != origin.BirthDate {
		_, err := Coll.UpdateOne(
			Ctx,
			bson.M{"_id": figure.ID},
			bson.D{
				{"$set", bson.D{{"birth_date", figure.BirthDate}}},
			},
		)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (v *ViewQ)RunView() SuperheroQ {
	var resultBson bson.D
	var resultStruct SuperheroQ
	err := Coll.FindOne(Ctx, bson.D{{"_id", v.ID}}).Decode(&resultBson)
	if err != nil {
		log.Fatal(err)
	}
	bsonBytes, _ := bson.Marshal(resultBson)
	bson.Unmarshal(bsonBytes, &resultStruct)
	return resultStruct
}

func (v *ViewByPageQ)RunViewAll()[]SuperheroQ {
	skip := int64(v.page * v.perPage)
	limit := int64(v.perPage)
	opts := options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
	}

	cursor, err := Coll.Find(nil, bson.M{}, &opts)
	var display []SuperheroQ
	for cursor.Next(Ctx) {
		var resultBson bson.D
		var resultStruct SuperheroQ
		if err = cursor.Decode(&resultBson); err != nil {
			log.Fatal(err)
		}
		bsonBytes, _ := bson.Marshal(resultBson)
		bson.Unmarshal(bsonBytes, &resultStruct)
		display = append(display, resultStruct)
	}
	return display
}
