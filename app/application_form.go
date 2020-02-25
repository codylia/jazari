package app

import (
	"strings"
	"time"

	v "github.com/oussama4/validate/v4"
	vs "github.com/oussama4/validate/v4/validators"
)

type applicationForm struct {
	FullName,
	BirthPlace,
	EducAndJob,
	Address,
	Phone,
	Cin,
	Email,
	TajweedLevel,
	HifdAmount,
	Reason,
	ImgName string
	BirthDate time.Time
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
			Name:    "educAndJob",
			Field:   strings.TrimSpace(af.EducAndJob),
			Min:     4,
			Max:     30,
			Message: "المرجو إعادة كتابة المهنة/المستوى الدراسي",
		},
		&vs.RegexMatch{
			Name:    "phone",
			Field:   strings.TrimSpace(af.Phone),
			Expr:    "^(\\+)?([0-9]){6,13}$",
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
			List:    []string{"مستحسن", "متوسط", "ضعيف", "ضعيف جدا"},
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
			Min:     10,
			Max:     300,
			Message: "المرجو ملئ خانة القرار",
		},
		&vs.TimeIsPresent{
			Name:    "birthDate",
			Field:   af.BirthDate,
			Message: "المرجو ملئ تاريخ الازدياد",
		},
	)

	errors.Append(errs)
}
