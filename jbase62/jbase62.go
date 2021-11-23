package jbase62

/**************************************************

Created By YJH 2021-06-30
Add Base62 Algo.

/**************************************************/

import (
	"bytes"
	"math"
)

const encodeStd = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

type Encoding struct {
	encodeMap [62]byte
	decodeMap map[byte]int
}

func NewEncoding() *Encoding {

	e := &Encoding{}

	copy(e.encodeMap[:], encodeStd)

	e.decodeMap = make(map[byte]int)

	for i, v := range e.encodeMap {
		e.decodeMap[v] = i
	}

	return e

}

func (en *Encoding) ToBase62(indata int64) string {
	var outStr bytes.Buffer

	last := indata

	for last != 0 {

		i := last % 62

		outStr.WriteByte(en.encodeMap[i])
		last = (last - i) / 62

	}

	s := outStr.String()

	return s
}

func (en *Encoding) FromBase62(indata string) int64 {
	var ret int64
	slen := len(indata)

	for i := 0; i < slen; i++ {
		ret += int64(en.decodeMap[indata[i]] * int(math.Pow(62, float64(i))))
	}

	return ret

}
