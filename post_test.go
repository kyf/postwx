package postwx

import (
	"fmt"
	"testing"
)

var (
	openid string = "o5voKuA644vh_pUxYW3h0N9XuQ4M"
)

func TestPostText(t *testing.T) {
	_, err := PostText(openid, "post text testing ...")
	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestPostImage(t *testing.T) {
	_, err := PostImage(openid, "Y0EtjhANujDworpFLdm7p-1UfPW1H89lu-WU0dRvZzhfzknDKqmke3htopGc-ku8")
	if err != nil {
		t.Errorf("%v", err)
	}
}

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
