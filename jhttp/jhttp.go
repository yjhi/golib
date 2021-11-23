package jhttp

/*******************************************************************
*
* Add By yjh 2021-06-17
*
*
********************************************************************/

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/jinghui0108/golib/jerrors"
)

type Http struct {
	Client  *http.Client
	Error   string
	Request *http.Request
	IsOk    bool
}

func Create(url string, secondtimeout int, way string) *Http {
	return New(url, secondtimeout, way)
}

func New(url string, secondtimeout int, way string) *Http {

	h := &Http{
		Client:  nil,
		Request: nil,
		Error:   "",
		IsOk:    false,
	}

	h.Client = &http.Client{
		Timeout: time.Duration(secondtimeout) * time.Second,
	}

	var err error

	h.Request, err = http.NewRequest(way, url, nil)

	if err != nil {
		h.Error = err.Error()
		return h
	}

	h.IsOk = true

	return h
}

func CreateGet(url string, secondtimeout int) *Http {

	return Create(url, secondtimeout, "GET")
}

func CreatePOST(url string, secondtimeout int) *Http {

	return Create(url, secondtimeout, "POST")
}

func (h *Http) AddCookie(name string, value string, path string, domain string) *Http {

	if h.IsOk {
		h.Request.AddCookie(&http.Cookie{
			Name:   name,
			Value:  value,
			Path:   path,
			Domain: domain,
		})
	}

	return h
}

func (h *Http) SetHeader(name string, value string) *Http {
	if h.IsOk {
		h.Request.Header.Set(name, value)
	}

	return h
}

func (h *Http) Do() (string, error) {

	if h.IsOk {
		resp, err := h.Client.Do(h.Request)
		if err != nil {
			return "", err
		}

		defer resp.Body.Close()
		bodyContent, errret := ioutil.ReadAll(resp.Body)

		if errret != nil {
			return "", errret
		}

		return string(bodyContent), nil

	} else {

		return "", jerrors.NewError("Http 初始化失败")
	}

}

func (h *Http) DoWithoutBody() error {

	if h.IsOk {
		_, err := h.Client.Do(h.Request)
		if err != nil {
			return err
		}
		return nil

	} else {

		return jerrors.NewError("Http 初始化失败")
	}

}

func (h *Http) DoWithResp() (*HttpResp, error) {

	s := &HttpResp{
		Response: nil,
	}

	var err error

	s.Response, err = h.Client.Do(h.Request)

	if err != nil {
		return nil, err
	}

	return s, nil
}
