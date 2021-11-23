package jhttp

/*******************************************************************
*
* Add By yjh 2021-07-28
*
*
********************************************************************/

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func Get(url string) (string, error) {
	resp, err1 := http.Get(url)

	if err1 != nil {
		return "", err1
	}

	defer resp.Body.Close()

	body, err2 := ioutil.ReadAll(resp.Body)

	if err2 != nil {
		return "", err2
	}
	return string(body), nil
}

func PostJson(url string, d string) (string, error) {
	return Post(url, "application/json", d)
}

func Post(url string, t string, d string) (string, error) {
	resp, err := http.Post(url, "application/json", strings.NewReader(d))

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err2 := ioutil.ReadAll(resp.Body)

	if err2 != nil {
		return "", err2
	}
	return string(body), nil

}
