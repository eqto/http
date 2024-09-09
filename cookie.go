package http

import (
	"bytes"
)

type Cookie struct {
	Value string
}

func (c *Cookie) parse(key string, val []byte) bool {
	if !bytes.HasPrefix(val, []byte(key+`=`)) {
		return false
	}
	val = val[len(key)+1:]
	split := bytes.Split(val, []byte(`;`))
	c.Value = string(bytes.TrimSpace(split[0]))
	return true
}
