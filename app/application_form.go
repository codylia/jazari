package app

import (
	"time"
	v "github.com/oussama4/validate/v4"
	vs "github.com/oussama4/validate/v4/validators"
)

type applicationForm struct {
	fullName,
	birthPlace,
	job,
	educationLevel,
	address,
	phone,
	cin,
	email,
	tajweedLevel,
	hifdAmount,
	reason string
	birthDate time.Time
}

func (af *applicationForm) IsValid(errors *v.Errors)  {
	return
}