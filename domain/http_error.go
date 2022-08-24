package domain

type HttpError interface {
	Error() string
	GetCode() int
}

// bad request error
type badRequestError struct {
	Err        error
	StatusCode int
}

func NewBadRequestError(err error) HttpError {
	return &badRequestError{
		Err:        err,
		StatusCode: 400,
	}
}

func (b *badRequestError) Error() string {
	return b.Err.Error()
}

func (b *badRequestError) GetCode() int {
	return b.StatusCode
}

// unauthentiacted error
type unAuthenticatedError struct {
	Err        error
	StatusCode int
}

func NewUnAuthenticatedError(err error) HttpError {
	return &unAuthenticatedError{
		Err:        err,
		StatusCode: 401,
	}
}

func (u *unAuthenticatedError) Error() string {
	return u.Err.Error()
}

func (u *unAuthenticatedError) GetCode() int {
	return u.StatusCode
}

// forbidden error
type forbiddenError struct {
	Err        error
	StatusCode int
}

func NewForbiddenError(err error) HttpError {
	return &forbiddenError{
		Err:        err,
		StatusCode: 403,
	}
}

func (f *forbiddenError) Error() string {
	return f.Err.Error()
}

func (f *forbiddenError) GetCode() int {
	return f.StatusCode
}

// not found error
type notFoundError struct {
	Err        error
	StatusCode int
}

func NewNotFoundError(err error) HttpError {
	return &notFoundError{
		Err:        err,
		StatusCode: 404,
	}
}

func (n *notFoundError) Error() string {
	return n.Err.Error()
}

func (n *notFoundError) GetCode() int {
	return n.StatusCode
}

// internal server error
type internalServerError struct {
	Err        error
	StatusCode int
}

func NewInternalServerError(err error) HttpError {
	return &internalServerError{
		Err:        err,
		StatusCode: 500,
	}
}

func (i *internalServerError) Error() string {
	return i.Err.Error()
}

func (i *internalServerError) GetCode() int {
	return i.StatusCode
}

// handle http error
func HandleHttpError(err HttpError, out *CommonResult) bool {
	if err != nil {
		out.SetError(err.GetCode(), err.Error())
		return true
	}
	return false
}
