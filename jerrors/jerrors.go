package jerrors

/*******************************************************************
*
* Add By yjh 2021-06-17
*
*
********************************************************************/

type Error struct {
	err string
}

func (e *Error) Error() string {
	return e.err
}

func NewError(str string) error {
	return &Error{err: str}
}
