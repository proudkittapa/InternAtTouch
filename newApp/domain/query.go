package domain

type InsertQ struct {
	ID         string   `bson:"_id" json:"id"`
	Name       string   `bson:"name" json:"name" validate:"required"`
	ActualName string   `bson:"actual_name" json:"actual_name" validate:"required"`
	Gender     string   `bson:"gender" json:"gender"`
	BirthDate  int64    `bson:"birth_date" json:"birth_date"`
	Height     int      `bson:"height" json:"height" validate:"gte=0"`
	SuperPower []string `bson:"super_power" json:"super_power"`
	Alive      bool     `bson:"alive" json:"alive"`
}

type DeleteQ struct {
	ID  string  `bson:"_id" json:"id"`
}

type UpdateQ struct {
	ID         string   `bson:"_id" json:"id"`
	Name       string   `bson:"name" json:"name"`
	ActualName string   `bson:"actual_name" json:"actual_name"`
	Gender     string   `bson:"gender" json:"gender"`
	BirthDate  int64    `bson:"birth_date" json:"birth_date"`
	Height     int      `bson:"height" json:"height" validate:"gte=0"`
	SuperPower []string `bson:"super_power" json:"super_power"`
	Alive      bool     `bson:"alive" json:"alive"`
}

type ViewQ struct {
	ID  string  `bson:"_id" json:"id"`
}

type ViewByPageQ struct {
	perPage int
	page 	int
}
