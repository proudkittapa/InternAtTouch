package validator

import (
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/user/userin"
	"github.com/go-playground/validator/v10"
)

func (v *GoPlayGroundValidator) IsProud(structLV validator.StructLevel) {
	input := structLV.Current().Interface().(userin.CreateInput)
	v.checkName(structLV, input.Name)
}
