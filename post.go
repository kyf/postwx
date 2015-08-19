package postwx

import (
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

	res, err := post(fmt.Sprintf(msg, openid, content))
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

	res, err := post(fmt.Sprintf(msg, openid, media_id))
	if err != nil {
		return false, err
	}
	return formatResponse(res)

}
