package domain

// type Staff struct {
// 	ID        string `bson:"id"`
// 	CompanyID string `bson:"companyID"`
// 	Name      string `bson:"name"`
// 	Tel       string `bson:"tel"`
// 	CreatedAt int64  `bson:"createdAt"`
// 	UpdatedAt int64  `bson:"updatedAt"`
// }

type User struct {
	Name string `bson:"name"`
}

func MakeTestStaff() (user *User) {
	return &User{
		Name: "test",
	}
}
