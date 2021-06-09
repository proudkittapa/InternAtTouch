package validator

import (
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/user/userin"
	"github.com/go-playground/validator/v10"
)

type GoPlayGroundValidator struct {
	validate *validator.Validate
}

func New() (v *GoPlayGroundValidator) {
	v = &GoPlayGroundValidator{
		validate: validator.New(),
	}
	v.validate.RegisterStructValidation(v.IsProud, &userin.CreateInput{})
	// v.validate.RegisterStructValidation(v.CompanyCreateStructLevelValidation, &companyin.CreateInput{})
	// v.validate.RegisterStructValidation(v.CompanyUpdateStructLevelValidation, &companyin.UpdateInput{})
	// v.validate.RegisterStructValidation(v.PageOptionStructLevelValidation, &domain.PageOption{})

	return v
}

func (v *GoPlayGroundValidator) Validate(item interface{}) (err error) {
	return v.validate.Struct(item)
}

// type Person struct {
// 	Name string `json:"name" validate:"required"`
// }

// type Err struct {
// 	Code  int
// 	Cause string
// }

// func IsValid(v ValidatorInterface) {
// 	v.ValidInter()
// }
// func (e Err) ValidInter(structLV validator.StructLevel, input Person) {
// 	if input.Name != "Proud" {
// 		e.Cause = "not Proud"
// 		e.Code = 400
// 	}
// }
