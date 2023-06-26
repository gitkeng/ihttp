package ihttp

type Validatable interface {
	Validate() error
}

type Validator struct{}

func (v *Validator) Validate(i any) error {
	if validatable, ok := i.(Validatable); ok {
		return validatable.Validate()
	}
	var errs Errors
	errs = append(errs, ErrNotValidatable)
	return &errs
}
