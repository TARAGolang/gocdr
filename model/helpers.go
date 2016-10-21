package model

func (cdr *CDR) SetError(code int, desc string) {
	cdr.Response.Error = &Error{
		Code:        code,
		Description: desc,
	}
}

func (cdr *CDR) AddReadAccess(consumer_id string) bool {
	if "" == consumer_id {
		return false
	}

	for _, element := range cdr.ReadAccess {
		if consumer_id == element {
			return false
		}
	}
	cdr.ReadAccess = append(cdr.ReadAccess, consumer_id)
	return true
}
