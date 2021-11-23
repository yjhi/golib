package jhttp

import (
	"fmt"
	"math/rand"
	"time"
)

type HttpUA struct {
	SysName []string
}

func NewUA() *HttpUA {
	h := &HttpUA{}

	h.SysName = append(h.SysName, "Windows NT 6.1; Win64; x64")
	h.SysName = append(h.SysName, "Windows NT 6.2; Win64; x64")
	h.SysName = append(h.SysName, "Windows NT 10.0; WOW64")

	rand.Seed(time.Now().UnixNano())

	return h
}

func (h *HttpUA) AddAgentSysName(str string) {
	h.SysName = append(h.SysName, str)
}

func (h *HttpUA) RandEdgeVerUA() string {
	aminver := rand.Intn(1000) + 3000
	return "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0." + fmt.Sprintf("%d", aminver) + ".101 Safari/537.36 Edg/91.0.864.48"
}

func (h *HttpUA) RandChromeVerUA() string {

	aminver := rand.Intn(1000) + 3000
	return "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0." + fmt.Sprintf("%d", aminver) + ".100 Safari/537.36"
}

func (h *HttpUA) RandEdgeSysNameAndVerUA() string {

	amax := len(h.SysName) - 1
	aindex := rand.Intn(amax)

	aminver := rand.Intn(1000) + 3000
	return "Mozilla/5.0 (" + h.SysName[aindex] + ") AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0." + fmt.Sprintf("%d", aminver) + ".101 Safari/537.36 Edg/91.0.864.48"
}

func (h *HttpUA) RandChromeSysNameAndVerUA() string {

	amax := len(h.SysName) - 1
	aindex := rand.Intn(amax)

	aminver := rand.Intn(1000) + 3000
	return "Mozilla/5.0 (" + h.SysName[aindex] + ") AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0." + fmt.Sprintf("%d", aminver) + ".101 Safari/537.36"
}
