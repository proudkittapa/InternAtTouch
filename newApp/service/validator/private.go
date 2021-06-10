package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func (v *GoPlayGroundValidator) checkName(structLV validator.StructLevel, name string) {
	fmt.Println("here")
	if name != "Proud" {
		structLV.ReportError("not Proud", "name", "name", "unique", "")
	}
}

// func (v *GoPlayGroundValidator) CheckExistName(ctx context.Context, structLV validator.StructLevel, name string) (bool, error) {
// 	log.Println("checkexistname")
// 	count, err := v.userRepo.View(ctx, bson.D{{"name", name}})
// 	if count < 1 {
// 		return false, err
// 	}
// 	return true, err
// }
