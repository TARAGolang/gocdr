package testutils

import (
	"github.com/fulldump/golax"

	"github.com/smartdigits/gocdr/constants"
	"github.com/smartdigits/gocdr/model"
)

func GetCdr(c *golax.Context) *model.CDR {
	v, exists := c.Get(constants.CONTEXT_KEY)

	if !exists {
		return nil
	}
	return v.(*model.CDR)
}
