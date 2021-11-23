package jhttp

import (
	"net/http"
)

/*******************************************************************
*
* Add By yjh 2021-06-17
*
*
********************************************************************/

type HttpResp struct {
	Response *http.Response
}

func (h *HttpResp) GetCookie(name string) string {
	for _, cookie := range h.Response.Cookies() {
		if cookie.Name == name {
			return cookie.Value
		}
	}

	return ""
}

func (h *HttpResp) Close() {

	defer h.Response.Body.Close()
}
