package model

func (cdr *CDR) SetError(code int, desc string) {
	cdr.Response.Error = &Error{
		Code:        code,
		Description: desc,
	}
}

func (cdr *CDR) AddReadAccess(consumerID string) bool {
	if "" == consumerID {
		return false
	}

	for _, element := range cdr.ReadAccess {
		if consumerID == element {
			return false
		}
	}
	cdr.ReadAccess = append(cdr.ReadAccess, consumerID)
	return true
}
