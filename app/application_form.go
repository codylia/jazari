package app

import (
	"time"
	"strings"

	v "github.com/oussama4/validate/v4"
	vs "github.com/oussama4/validate/v4/validators"
)

type applicationForm struct {
	fullName,
	birthPlace,
	educAndJob,
	address,
	phone,
	cin,
	email,
	tajweedLevel,
	hifdAmount,
	reason string
	birthDate time.Time
}

func (af *applicationForm) IsValid(errors *v.Errors)   {

	err := v.Validate(
		&vs.StringLengthInRange{
			Name : "fullName",
			Field : strings.TrimSpace(af.fullName),
			Min : 6,
			Max : 50,
			Message : "المرجو إعادة كتابة الإسم",
		},
		&vs.StringLengthInRange{
			Name : "birthPlace",
			Field : strings.TrimSpace(af.birthPlace),
			Min : 3,
			Max : 30,
			Message : "المرجو إعادة كتابة مكان لازدياد",
		},
		&vs.StringLengthInRange{
			Name : "educAndJob",
			Field : strings.TrimSpace(af.educAndJob),
			Min : 4,
			Max : 30,
			Message : "المرجو إعادة كتابة المهنة/المستوى الدراسي",
		},
		&vs.RegexMatch{
			Name : "phone",
			Field : strings.TrimSpace(af.phone),
			Expr : "^(\\+)?([0-9]){6,13}$",
			Message : "المرجو إعادة كتابة رقم الهاتف بشكل صحيح",
		},
		&vs.RegexMatch{
			Name : "cin",
			Field : strings.TrimSpace(af.cin),
			Expr : "^([a-z]){1,2}([0-9]{6})$",
			Message : "المرجو إعادة كتابة رقم الهاتف بشكل صحيح",
		},

	)
	return err 
}