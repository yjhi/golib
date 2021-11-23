package jhttp

/*******************************************************************
*
* Add By yjh 2021-06-17
*
*
********************************************************************/

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/jinghui0108/golib/jerrors"
)

func _webService1(url string, fullbody string, errname string) (string, error) {
	res, err := http.Post(url, "text/xml; charset=UTF-8", strings.NewReader(fullbody))
	if nil != err {
		s := fmt.Sprintf("["+errname+"]Http Post Err:%s", err.Error())
		return "", jerrors.NewError(s)
	}

	defer res.Body.Close()

	if 200 != res.StatusCode {

		s := fmt.Sprintf("["+errname+"]WebService Request Fail, Status Is: %d", res.StatusCode)

		return "", jerrors.NewError(s)
	}

	data, err := ioutil.ReadAll(res.Body)
	if nil != err {

		s := fmt.Sprintf("["+errname+"]Read Body err:%s", err)

		return "", jerrors.NewError(s)
	}
	return string(data), nil
}

func _webService(url string, fullbody string, errname string) (string, error) {

	trans := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	httpClient := &http.Client{
		Timeout:   time.Duration(60) * time.Second,
		Transport: trans,
	}

	request, err := http.NewRequest("POST", url, strings.NewReader(fullbody))

	if err != nil {
		return "", jerrors.NewError(err.Error())
	}

	request.Header.Set("Content-Type", "text/xml; charset=UTF-8")

	resp, err := httpClient.Do(request)

	if nil != err {
		s := fmt.Sprintf("["+errname+"]Http Post Err:%s", err.Error())
		return "", jerrors.NewError(s)
	}

	defer resp.Body.Close()

	if 200 != resp.StatusCode {

		s := fmt.Sprintf("["+errname+"]WebService Request Fail, Status Is: %d", resp.StatusCode)

		return "", jerrors.NewError(s)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if nil != err {

		s := fmt.Sprintf("["+errname+"]Read Body err:%s", err)

		return "", jerrors.NewError(s)
	}
	return string(data), nil
}

func WebService11(url string, body string) (string, error) {
	reqBody := `<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" 
  xmlns:xsd="http://www.w3.org/2001/XMLSchema" 
  xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"> 
  <soap:Body>` + body +
		`</soap:Envelope>`

	return _webService(url, reqBody, "Soap1.1")
}

func WebService12(url string, body string) (string, error) {
	reqBody := `<soap12:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" 
  xmlns:xsd="http://www.w3.org/2001/XMLSchema" 
  xmlns:soap12="http://www.w3.org/2003/05/soap-envelope"> 
  <soap12:Body>` + body + `
  </soap12:Body> 
  </soap12:Envelope>`

	return _webService(url, reqBody, "Soap1.2")
}
