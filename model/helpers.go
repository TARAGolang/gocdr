package model

// SetError overwrite error code and description in a CDR
func (cdr *CDR) SetError(code int, desc string) {
	cdr.Response.Error = &Error{
		Code:        code,
		Description: desc,
	}
}

// AddReadAccess add a consumerID to read access list for this CDR
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
