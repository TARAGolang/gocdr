package gocdr

import (
	"github.com/fulldump/golax"

	"github.com/smartdigits/gocdr/constants"
	"github.com/smartdigits/gocdr/model"
)

// GetCdr retrieve object from context.
//
// Typical usage:
//
// ```go
// func MyHandler(c *golax.Context) {
// 	// ...
// 	cdr := gocdr.GetCdr(c)
// 	// ...
// }
// ```
func GetCdr(c *golax.Context) *model.CDR {
	v, exists := c.Get(constants.CONTEXT_KEY)

	if !exists {
		return nil
	}
	return v.(*model.CDR)
}
