package validator

import (
	"context"
	"fmt"
	"log"
	"regexp"

	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp2/domain"

	"github.com/go-playground/validator/v10"
)

func (v *GoPlayGroundValidator) checkName(structLV validator.StructLevel, name string) {
	fmt.Println("here")
	if name != "Proud" {
		structLV.ReportError("not Proud", "name", "name", "unique", "")
	}
}

func (v *GoPlayGroundValidator) checkTH(structLV validator.StructLevel, name string) {
	re := regexp.MustCompile("[A-Za-z]+")
	ok := re.MatchString(name)
	if !ok {
		structLV.ReportError(name, "err validation", "err validation", "match", "")
	}
}

func (v *GoPlayGroundValidator) checkUserNameUnique(ctx context.Context, structLV validator.StructLevel, name string) (user *domain.UpdateStruct) {
	a, err := v.elasRepo.CheckExistName(ctx, name)
	if err != nil {
		structLV.ReportError(err, "err validation", "err validation", "error from database", "")
	}
	if a == true {
		structLV.ReportError(name, "name", "name", "unique", "")
	}
	return user
}

func (v *GoPlayGroundValidator) checkUserActualNameUnique(ctx context.Context, structLV validator.StructLevel, name string) (user *domain.UpdateStruct) {
	a, err := v.elasRepo.CheckExistActualName(ctx, name)
	if err != nil {
		structLV.ReportError(err, "err validation", "err validation", "error from database", "")
	}
	if a == true {
		structLV.ReportError(name, "actual_name", "actual_name", "unique", "")
	}
	return user
}

func (v *GoPlayGroundValidator) checkUserNameUniqueUpdate(ctx context.Context, structLV validator.StructLevel, name string, actualName string, id string) (user *domain.UpdateStruct) {
	n, err := v.elasRepo.CheckExistName(ctx, name)
	log.Println("qq", err)
	an, _ := v.elasRepo.CheckExistActualName(ctx, actualName)
	log.Println("qq1")
	if n == true { //jer name
		if an == true { //
			temp, _ := v.elasRepo.View(id, ctx)
			if temp.Name != name {
				structLV.ReportError(actualName, "actual_name", "actual_name", "unique", "")
			}
		}
	}
	if an == true { //jer name
		if n == true { //
			temp, _ := v.elasRepo.View(id, ctx)
			if temp.ActualName != actualName {
				structLV.ReportError(name, "name", "name", "unique", "")
			}
		}
	}
	return user
}

//func (v *GoPlayGroundValidator) checkUserIDUnique(ctx context.Context, structLV validator.StructLevel, id string) (user *domain.UpdateQ) {
//	a, _ := v.userRepo.CheckExistID(ctx, id)
//	if a == true {
//		structLV.ReportError(id, "id", "id", "unique", "")
//	}
//	return user
//}
