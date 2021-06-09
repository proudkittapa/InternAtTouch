package validator

import "github.com/go-playground/validator/v10"

func (v *GoPlayGroundValidator) IsProud(structLV validator.StructLevel) {
	input := structLV.Current().Interface().(Person)
	v.checkName(structLV, input.Name)
}
