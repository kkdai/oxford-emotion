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

}
