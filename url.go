package emotion

import "fmt"

const (
	EMOTION_URL string = "https://api.projectoxford.ai/emotion/v1.0/"
	EMOTION_API string = "recognize"
)

func getDetectURL(option *DetectParameters) string {
	apiURL := FACE_URL + DETECT_API
	if option == nil {
		return apiURL
	}

	opURL := fmt.Sprintf("%s?returnFaceId=%s&returnFaceLandmarks=%s&returnFaceAttributes=%s",
		apiURL,
		getBooleanString(option.RceturnFaceIdcdd),
		getBooleanString(option.ReturnFaceLandmarks),
		option.ReturnFaceAttributes)

	return opURL
}

func getEmotionURL() string {
	return EMOTION_URL + EMOTION_API
}
