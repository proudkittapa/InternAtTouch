package validator

import (
	"context"
	"log"

	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/domain"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/user/userin"
	"github.com/go-playground/validator/v10"
)

func (v *GoPlayGroundValidator) UserUpdateStructLevelValidation(structLV validator.StructLevel) {
	ctx := context.Background()
	input := structLV.Current().Interface().(userin.UpdateInput)
	v.checkUserNameUniqueUpdate(ctx, structLV, input.Name, input.ActualName, input.ID)
	v.checkUserIDUnique(ctx, structLV, input.ID)
	// v.checkName(structLV, input.Name)
}

func (v *GoPlayGroundValidator) checkUserNameUniqueUpdate(ctx context.Context, structLV validator.StructLevel, name string, actualName string, id string) (user *domain.UpdateQ) {
	n, err := v.userRepo.CheckExistName(ctx, name)
	log.Println("qq", err)
	an, _ := v.userRepo.CheckExistActualName(ctx, actualName)
	log.Println("qq1")
	if n == true { //jer name
		if an == true { //
			temp, _ := v.userRepo.View(ctx, id)
			if temp.Name != name {
				structLV.ReportError(actualName, "actual_name", "actual_name", "unique", "")
			}
		}
	}
	if an == true { //jer name
		if n == true { //
			temp, _ := v.userRepo.View(ctx, id)
			if temp.ActualName != actualName {
				structLV.ReportError(name, "name", "name", "unique", "")
			}
		}
	}
	return user
}

func (v *GoPlayGroundValidator) checkUserIDUnique(ctx context.Context, structLV validator.StructLevel, id string) (user *domain.UpdateQ) {
	a, _ := v.userRepo.CheckExistID(ctx, id)
	if a == true {
		structLV.ReportError(id, "id", "id", "unique", "")
	}
	return user
}
