package router

import (
	"strings"
	"time"
)

type CustomTime time.Time

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*ct = CustomTime(t)
	return nil
}
