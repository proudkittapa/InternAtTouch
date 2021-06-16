package domain

type InsertStruct struct {
	Name      		string   `bson:"name" json:"name" validate:"required"`
	ActualName 		string   `bson:"actual_name" json:"actual_name"`
	ActualLastName  string   `bson:"actual_lastname" json:"actual_lastname"`
	Gender     		string   `bson:"gender" json:"gender"`
	BirthDate  		int64    `bson:"birth_date" json:"birth_date"`
	Height     		int      `bson:"height" json:"height" validate:"gte=0"`
	SuperPower 		[]string `bson:"super_power" json:"super_power"`
	Alive      		bool     `bson:"alive" json:"alive"`
	Universe 		string	 `bson:"universe" json:"universe"`
	Movies			[]string `bson:"movies" json:"movies"`
	Enemies			[]string `bson:"enemies" json:"enemies"`
	FamilyMember	[]string `bson:"family_member" json:"family_member"`
	About			string	 `bson:"about" json:"about"`
}

type UpdateStruct struct {
	ID         string   `bson:"_id" json:"id"`
	Name      		string   `bson:"name" json:"name" validate:"required"`
	ActualName 		string   `bson:"actual_name" json:"actual_name"`
	ActualLastName  string   `bson:"actual_lastname" json:"actual_lastname"`
	Gender     		string   `bson:"gender" json:"gender"`
	BirthDate  		int64    `bson:"birth_date" json:"birth_date"`
	Height     		int      `bson:"height" json:"height" validate:"gte=0"`
	SuperPower 		[]string `bson:"super_power" json:"super_power"`
	Alive      		bool     `bson:"alive" json:"alive"`
	Universe 		string	 `bson:"universe" json:"universe"`
	Movies			[]string `bson:"movies" json:"movies"`
	Enemies			[]string `bson:"enemies" json:"enemies"`
	FamilyMember	[]string `bson:"family_member" json:"family_member"`
	About			string	 `bson:"about" json:"about"`
}

type DeleteQ struct {
	ID string `bson:"_id" json:"id"`
}