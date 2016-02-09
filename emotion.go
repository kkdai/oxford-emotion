package emotion

import (
	"bytes"
	"errors"
)

type Emotion struct {
	client *Client
}

func NewEmotion(key string) *Emotion {
	if len(key) == 0 {
		return nil
	}

	f := new(Emotion)
	f.client = NewClient(key)
	return f
}

func (e *Emotion) detect(data *bytes.Buffer, useJson bool) ([]byte, *ErrorResponse) {
	url := getEmotionURL()
	return f.client.Connect("POST", url, data, useJson)
}

//Detect Emotion with input URL
func (e *Emotion) EmotionUrl(url string) ([]byte, *ErrorResponse) {
	data := getUrlByteBuffer(url)
	return f.detect(data, true)
}

//Detect Emotion with input image file path
func (e *Emotion) EmotionFile(filePath string) ([]byte, *ErrorResponse) {
	data, err := getFileByteBuffer(filePath)
	if err != nil {
		return nil, &ErrorResponse{Err: errors.New("File path err:" + err.Error())}
	}
	return f.detect(data, false)
}
