package postwx

import (
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
	_, err := PostImage(openid, "kyf")
	if err != nil {
		t.Errorf("%v", err)
	}
}
