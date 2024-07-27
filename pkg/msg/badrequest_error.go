package msg

type BadRequestError struct {
	Message string
}

func (badRequestError BadRequestError) Error() string {
	return badRequestError.Message
}
