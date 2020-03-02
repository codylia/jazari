package app

import (
	"strings"

	v "github.com/oussama4/validate/v4"
	vs "github.com/oussama4/validate/v4/validators"
)

type applicationForm struct {
	FullName,
	BirthPlace,
	Job,
	Educ,
	Address,
	Phone,
	Cin,
	Email,
	TajweedLevel,
	HifdAmount,
	Reason,
	ImgName,
	BirthDate string
}

func (af *applicationForm) IsValid(errors *v.Errors) {

	errs := v.Validate(
		&vs.StringLengthInRange{
			Name:    "fullName",
			Field:   strings.TrimSpace(af.FullName),
			Min:     6,
			Max:     50,
			Message: "المرجو إعادة كتابة الإسم",
		},
		&vs.StringLengthInRange{
			Name:    "birthPlace",
			Field:   strings.TrimSpace(af.BirthPlace),
			Min:     3,
			Max:     30,
			Message: "المرجو إعادة كتابة مكان لازدياد",
		},
		&vs.StringLengthInRange{
			Name:    "job",
			Field:   strings.TrimSpace(af.Job),
			Min:     4,
			Max:     30,
			Message: "المرجو إعادة كتابة المهنة",
		},
		&vs.StringLengthInRange{
			Name:    "educ",
			Field:   strings.TrimSpace(af.Educ),
			Min:     4,
			Max:     30,
			Message: "المرجو إعادة كتابة المستوى الدراسي",
		},
		&vs.RegexMatch{
			Name:    "phone",
			Field:   strings.TrimSpace(af.Phone),
			Expr:    "^(\\+|0)?([0-9]){6,13}$",
			Message: "المرجو إعادة كتابة رقم الهاتف بشكل صحيح",
		},
		&vs.EmailIsPresent{
			Name:    "email",
			Field:   af.Email,
			Message: "المرجو إعادة كتابة رقم البريد الإلكتروني بشكل صحيح",
		},
		&vs.StringInclusion{
			Name:    "tajweedLevel",
			Field:   af.TajweedLevel,
			List:    []string{"1", "2", "3", "4"},
			Message: "المرجو إختيار مستوى التجويد",
		},
		&vs.StringLengthInRange{
			Name:    "hifdAmount",
			Field:   strings.TrimSpace(af.HifdAmount),
			Min:     5,
			Max:     300,
			Message: "المرجو ملئ خانة الحفظ",
		},
		&vs.StringLengthInRange{
			Name:    "reason",
			Field:   strings.TrimSpace(af.Reason),
			Min:     6,
			Max:     300,
			Message: "المرجو ملئ خانة القرار",
		},
		&vs.StringIsPresent{
			Name:    "birthDate",
			Field:   af.BirthDate,
			Message: "المرجو ملئ تاريخ الازدياد",
		},
		&vs.RegexMatch{
			Name:    "cin",
			Field:   strings.TrimSpace(af.Cin),
			Expr:    "^([a-zA-z]){1}([0-9]){6}|([a-zA-Z]){2}([0-9]){5}$",
			Message: "المرجو إعادة كتابة بطاقة التعريف بشكل صحيح",
		},
		&vs.StringLengthInRange{
			Name:    "address",
			Field:   strings.TrimSpace(af.Address),
			Min:     10,
			Max:     100,
			Message: "المرجو ملئ خانة العنوان",
		},
	)
	errors.Append(errs)
}
