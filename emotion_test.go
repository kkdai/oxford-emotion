package emotion_test

import (
	"fmt"
	"os"
	"testing"

	. "github.com/kkdai/oxford-emotion"
)

var API_KEY string

const ()

func init() {
	API_KEY = os.Getenv("MSFT_KEY")
	if API_KEY == "" {
		fmt.Println("Please export your key to environment first, `export MSFT_KEY=12234`")
	}
}

func TestEmotion(t *testing.T) {
	if API_KEY == "" {
		return
	}

	e := NewEmotion(API_KEY)
	if e == nil {
		t.Error("Cannot connect to server")
	}

	res, err := e.EmotionUrl("http://www.skanaa.com/assets/images/news/20151112/5644565fa81bb78c538b4567.jpg")

	if err != nil {
		t.Error("Error happen on :", err.Err)
	}
	fmt.Println("Got response:", string(res))
}
