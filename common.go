package postwx

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var (
	tokenurl string = "http://m.6renyou.com/weixin_service/getAccessToken?account_type=1"

	posturl string = "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token="

	uploadurl string = "https://api.weixin.qq.com/cgi-bin/media/upload?access_token=%s&type=%s"

	mediaurl string = "https://api.weixin.qq.com/cgi-bin/media/get?access_token=%s&media_id=%s"

	mediatype map[string]string = map[string]string{
		"image/jpeg": ".jpg",
		"image/gif":  ".gif",
		"image/png":  ".png",
	}
)

type Response struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type UploadResponse struct {
	Type       string `json:"type"`
	Media_id   string `json:"media_id"`
	Created_at int64  `json:"created_at"`
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

	res, err := http.Post(fmt.Sprintf("%s%s", posturl, token), "application/x-www-form-urlencoded", strings.NewReader(body))
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

func upload(fpath, mediaType string) (interface{}, error) {
	var result interface{}
	file, err := os.Open(fpath)
	if err != nil {
		return result, err
	}

	token, err := getAccessToken()
	if err != nil {
		return result, err
	}

	url := fmt.Sprintf(uploadurl, token, mediaType)
	buffer := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buffer)
	part, err := writer.CreateFormFile("media", filepath.Base(fpath))
	if err != nil {
		return result, err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return result, err
	}

	err = writer.Close()
	if err != nil {
		return result, err
	}

	bodyType := writer.FormDataContentType()
	res, err := http.Post(url, bodyType, buffer)
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

func downloadMedia(media_id, savepath string) error {
	token, err := getAccessToken()
	if err != nil {
		return err
	}

	url := fmt.Sprintf(mediaurl, token, media_id)
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	contentType := res.Header.Get("Content-Type")
	if strings.EqualFold(contentType, "text/plain") {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		var response Response
		err = json.Unmarshal(body, &response)
		if err != nil {
			return err
		}

		_, err = formatResponse(response)
		return err
	} else {
		ext := mediatype[contentType]

		file, err := os.Create(fmt.Sprintf("%s%s", savepath, ext))
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(file, res.Body)
		if err != nil {
			return err
		}
		return nil
	}
}
