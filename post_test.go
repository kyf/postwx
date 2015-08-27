package postwx

import (
	"fmt"
	"testing"
)

var (
	openid string = "o5voKuA644vh_pUxYW3h0N9XuQ4M"
)

func TestPostText(t *testing.T) {
	content := `post text ...`
	_, err := PostText(openid, content)
	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestPostImage(t *testing.T) {
	_, err := PostImage(openid, "F5HOIcB87Aw3KyAIsLzwntp9t5nDoVN8GUZktX0syoiFmvZHTJd2Egv_MyUwnpmv")
	if err != nil {
		t.Errorf("%v", err)
	}
}

/*
func TestUploadMedia(t *testing.T) {
	filepath := "/work/gopro/src/6renyou/postwx/gtl.jpg"
	mediaType := "image"
	media_id, err := UploadMedia(filepath, mediaType)
	if err != nil {
		t.Errorf("%v", err)
	} else {
		fmt.Println("media_id is ", media_id)
	}
}

func TestGetMedia(t *testing.T) {
	media_id := "2XORNpyFZfl8iwUAHARfZzyxVcNubUgijT4e8dW0kbUfytyFnKyKhmEzKxQ--_b9"
	savepath := "/home/kyf/media/"
	fullpath, err := GetMedia(media_id, fmt.Sprintf("%s%s", savepath, media_id))
	if err != nil {
		t.Errorf("%v", err)
	} else {
		fmt.Println("file save success!", string(fullpath))
	}
}

*/

func TestPostTpl(t *testing.T) {
	tpl_id := "e5sFqp2BHA4OhbzOpzeqmi0ir6lT9sA3DanMOYOPhRI"
	url := "http://www.6renyou.com/"
	color := map[string]string{
		"top":    "#FF0000",
		"first":  "#000000",
		"data":   "#3eb166",
		"remark": "#939393",
	}

	data := `{
		"touser":"%s",
		"template_id":"%s",
		"url":"%s",
		"topcolor":"%s",
		"data":{
			"first":{
				"value":"",
				"color":"%s"
			},
			"keyword1":{
				"value":"微信客户端",
				"color":"%s"
			},
			"keyword2": {
				"value":"王子于线路(上海出发|台湾台北5天4晚自由行）申请了团长",
				"color":"%s"
			},
			"keyword3": {
				"value":"2015年7月16 11:12:12",
				"color":"%s"
			},
			"remark":{
				"value":"",
				"color":"%s"
			}
		}
	}`

	d := fmt.Sprintf(data, openid, tpl_id, url, color["top"], color["first"], color["data"], color["data"], color["data"], color["remark"])
	_, err := PostTpl(d)
	if err != nil {
		t.Errorf("%v", err)
	}

}
