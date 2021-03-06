package postwx

import (
	"errors"
	"fmt"
)

func PostText(openid, content string) (bool, error) {
	msg := `{
		    	"touser":"%s",
				"msgtype":"text",
				"text":
				{
					"content":"%s"
				}
			}`

	res, err := post(fmt.Sprintf(msg, openid, content), posturl)
	if err != nil {
		return false, err
	}
	return formatResponse(res)
}

func PostImage(openid, media_id string) (bool, error) {
	msg := `{
		    	"touser":"%s",
				"msgtype":"image",
				"image":
				{
					"media_id":"%s"
				}
			}`

	res, err := post(fmt.Sprintf(msg, openid, media_id), posturl)
	if err != nil {
		return false, err
	}
	return formatResponse(res)

}

func UploadMedia(filepath, mediaType string) (string, error) {
	res, err := upload(filepath, mediaType)
	if err != nil {
		return "", err
	}

	if result, ok := res.(map[string]interface{}); ok {
		if errmsg, ok := result["errmsg"]; ok {
			msg, _ := errmsg.(string)
			return "", errors.New(msg)
		} else {
			media_id := result["media_id"]
			media, _ := media_id.(string)
			return media, nil
		}
	} else {
		return "", errors.New("Unknwon error")
	}

}

func GetMedia(media_id, savepath string) ([]byte, error) {
	return downloadMedia(media_id, savepath)
}

func PostTpl(data string) (bool, error) {
	res, err := post(data, tplurl)
	if err != nil {
		return false, err
	}
	return formatResponse(res)

}
