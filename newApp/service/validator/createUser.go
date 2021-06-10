package validator

import (
	"context"
	"regexp"

	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/domain"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/user/userin"
	"github.com/go-playground/validator/v10"
)

func (v *GoPlayGroundValidator) IsProud(structLV validator.StructLevel) {
	input := structLV.Current().Interface().(userin.CreateInput)
	v.checkName(structLV, input.Name)
}

func (v *GoPlayGroundValidator) UserCreateStructLevelValidation(structLV validator.StructLevel) {
	ctx := context.Background()
	input := structLV.Current().Interface().(userin.CreateInput)
	//v.checkTH(structLV, input.Name)
	// v.checkName(structLV, input.Name)
	v.checkUserNameUnique(ctx, structLV, input.Name)
	v.checkUserActualNameUnique(ctx, structLV, input.ActualName)
}

func (v *GoPlayGroundValidator) checkTH(structLV validator.StructLevel, name string) {
	re := regexp.MustCompile("[A-Za-z]+")
	ok := re.MatchString(name)
	if !ok {
		structLV.ReportError(name, "err validation", "err validation", "match", "")
	}
}

func (v *GoPlayGroundValidator) checkUserNameUnique(ctx context.Context, structLV validator.StructLevel, name string) (user *domain.InsertQ) {
	a, err := v.userRepo.CheckExistName(ctx, name)
	if err != nil {
		structLV.ReportError(err, "err validation", "err validation", "error from database", "")
	}
	if a == true {
		structLV.ReportError(name, "name", "name", "unique", "")
	}
	return user
}

func (v *GoPlayGroundValidator) checkUserActualNameUnique(ctx context.Context, structLV validator.StructLevel, name string) (user *domain.InsertQ) {
	a, _ := v.userRepo.CheckExistActualName(ctx, name)
	// if err != nil {
	// 	structLV.ReportError(err, "err validation", "err validation", "error from database", "")
	// }
	if a == true {
		structLV.ReportError(name, "actual_name", "actual_name", "unique", "")
	}
	return user
}
