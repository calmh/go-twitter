package twitter

import (
	"encoding/base64"
	"net/http"

	"github.com/dghubble/sling"
)

type MediaService struct {
	sling *sling.Sling
}

func newMediaService(sling *sling.Sling) *MediaService {
	return &MediaService{
		sling: sling.Base("https://upload.twitter.com/1.1/media/"),
	}
}

type MediaParams struct {
	Media string `url:"media_data"`
}

type MediaResponse struct {
	MediaID int64 `json:"media_id"`
}

func (s *MediaService) UploadMedia(media []byte) (*MediaResponse, *http.Response, error) {
	params := &MediaParams{
		Media: base64.StdEncoding.EncodeToString(media),
	}

	mediaresp := new(MediaResponse)
	apiError := new(APIError)
	resp, err := s.sling.New().Post("upload.json").BodyForm(params).Receive(mediaresp, apiError)
	return mediaresp, resp, relevantError(err, *apiError)
}
