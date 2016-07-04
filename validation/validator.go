package validation

type Validator interface  {
	Validate(i interface{}) (bool,[]string)
}
