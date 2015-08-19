package postwx

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	tokenurl string = "http://m.6renyou.com/weixin_service/getAccessToken?account_type=1"
	posturl  string = "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token="
)

type Response struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func getAccessToken() ([]byte, error) {
	res, err := http.Get(tokenurl)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func post(body string) (Response, error) {
	var result Response
	token, err := getAccessToken()
	if err != nil {
		return result, err
	}

	res, err := http.Post(fmt.Sprintf("%s%s", posturl, token), "application/x-form-urlencode", strings.NewReader(body))
	if err != nil {
		return result, err
	}

	rev, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(rev, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func formatResponse(res Response) (bool, error) {
	if res.Errcode == 0 {
		return true, nil
	} else {
		return false, errors.New(res.Errmsg)
	}

}
