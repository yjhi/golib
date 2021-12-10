package jutils

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

// make guid
// add by yjh
// 211126
func NewGuid() string {
	uuidUtil, err := uuid.NewRandom()

	if err != nil {
		return ""
	}

	return fmt.Sprintf("%v", uuidUtil)
}

//make guid and ToUpper
//add by yjh 211126
func NewGUID() string {
	uuidUtil, err := uuid.NewRandom()

	if err != nil {
		return ""
	}

	return strings.ToUpper(fmt.Sprintf("%v", uuidUtil))
}
