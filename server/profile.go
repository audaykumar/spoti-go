package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/audaykumar/spoti-go/config"
)

func (s *Server) handleProfile(w http.ResponseWriter, req *http.Request) {
	log.Println("handleProfile start")
	defer func() {
		log.Println("handleProfile Done")
	}()
	profileRequest, _ := http.NewRequest("GET", "https://api.spotify.com/v1/me", nil)

	profileResponse, err := s.httpClient.Do(profileRequest)
	if err != nil && profileResponse.StatusCode != 200 {
		errMsg := "An error occurred while trying to retrieve the profile"
		w.Write([]byte(errMsg + " Please check the log for details"))
		log.Println(errMsg)
		log.Println(err)
	}

	respBody, err := io.ReadAll(profileResponse.Body)
	if err != nil {
		errMsg := "An error occurred while trying to read the response from the Spotify API"
		w.Write([]byte(errMsg + " Please check the log for details"))
		log.Println(errMsg)
		log.Fatalln(err)
		return
	}
	var profile Profile
	err = json.Unmarshal(respBody, &profile)
	if err != nil {
		errMsg := "unmarshal error"
		w.Write([]byte(errMsg + " Please check the log for details."))
		log.Println(errMsg)
		log.Println(err)
		if config.DebugMode {
			log.Println("Response Body:")
			fmt.Println(string(respBody))
		}
		return
	}

	// w.Write([]byte("<h1>Profile</h1>"))
	// line := fmt.Sprintf(
	// 	"<p>%s</p><br><p>%s</p>",
	// 	profile.DisplayName,
	// 	profile.ID,
	// )

	s.userID = profile.ID
	strResp := fmt.Sprintf("%+v", profile)
	w.Write([]byte(strResp))

}

type Profile struct {
	Country         string          `json:"country,omitempty"`
	DisplayName     string          `json:"display_name,omitempty"`
	Email           string          `json:"email,omitempty"`
	ExplicitContent ExplicitContent `json:"explicit_content,omitempty"`
	ExternalUrls    ExternalUrls    `json:"external_urls,omitempty"`
	Followers       Followers       `json:"followers,omitempty"`
	Href            string          `json:"href,omitempty"`
	ID              string          `json:"id,omitempty"`
	Images          []Images        `json:"images,omitempty"`
	Product         string          `json:"product,omitempty"`
	Type            string          `json:"type,omitempty"`
	URI             string          `json:"uri,omitempty"`
}
type ExplicitContent struct {
	FilterEnabled bool `json:"filter_enabled,omitempty"`
	FilterLocked  bool `json:"filter_locked,omitempty"`
}
type ExternalUrls struct {
	Spotify string `json:"spotify,omitempty"`
}
type Followers struct {
	Href  string `json:"href,omitempty"`
	Total int    `json:"total,omitempty"`
}
type Images struct {
	URL    string `json:"url,omitempty"`
	Height int    `json:"height,omitempty"`
	Width  int    `json:"width,omitempty"`
}
