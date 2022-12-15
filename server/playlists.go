package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/audaykumar/spoti-go/config"
)

func (s *Server) handlePlaylists(w http.ResponseWriter, req *http.Request) {
	log.Println("handlePlaylists start")
	defer func() {
		log.Println("handlePlaylists Done")
	}()

	url := fmt.Sprintf("https://api.spotify.com/v1/users/%s/playlists", s.userID)
	profileRequest, _ := http.NewRequest("GET", url, nil)

	profileResponse, err := s.httpClient.Do(profileRequest)
	if err != nil && profileResponse.StatusCode != 200 {
		errMsg := "An error occurred while trying to retrieve the playlists"
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
	var playlists Playlists
	err = json.Unmarshal(respBody, &playlists)
	if err != nil {
		errMsg := "There was an error attempting to unmarshal the data from the organisation endpoint."
		w.Write([]byte(errMsg + " Please check the log for details."))
		log.Println(errMsg)
		log.Println(err)
		if config.DebugMode {
			log.Println("Response Body:")
			fmt.Println(string(respBody))
		}
		return
	}

	w.Write([]byte("<h1>Playlists</h1>"))
	w.Write([]byte("<ul>"))
	for _, playlist := range playlists.Items {
		w.Write([]byte("<li> Name: " + playlist.Name + "   Total Tracks: " + fmt.Sprint(playlist.Tracks.Total) + "</li>"))
	}
	w.Write([]byte("</ul>"))
	// strResp := fmt.Sprintf("%+v", playlists)
	// w.Write([]byte(strResp))

}

type Playlists struct {
	Items    []Items `json:"items,omitempty"`
	Href     string  `json:"href,omitempty"`
	Limit    int     `json:"limit,omitempty"`
	Next     string  `json:"next,omitempty"`
	Offset   int     `json:"offset,omitempty"`
	Previous string  `json:"previous,omitempty"`
	Total    int     `json:"total,omitempty"`
}

type Owner struct {
	ExternalUrls ExternalUrls `json:"external_urls,omitempty"`
	Followers    Followers    `json:"followers,omitempty"`
	Href         string       `json:"href,omitempty"`
	ID           string       `json:"id,omitempty"`
	Type         string       `json:"type,omitempty"`
	URI          string       `json:"uri,omitempty"`
	DisplayName  string       `json:"display_name,omitempty"`
}
type Tracks struct {
	Href  string `json:"href,omitempty"`
	Total int    `json:"total,omitempty"`
}

type Items struct {
	Collaborative bool         `json:"collaborative,omitempty"`
	Description   string       `json:"description,omitempty"`
	ExternalUrls  ExternalUrls `json:"external_urls,omitempty"`
	Href          string       `json:"href,omitempty"`
	ID            string       `json:"id,omitempty"`
	Images        []Images     `json:"images,omitempty"`
	Name          string       `json:"name,omitempty"`
	Owner         Owner        `json:"owner,omitempty"`
	Public        bool         `json:"public,omitempty"`
	SnapshotID    string       `json:"snapshot_id,omitempty"`
	Tracks        Tracks       `json:"tracks,omitempty"`
	Type          string       `json:"type,omitempty"`
	URI           string       `json:"uri,omitempty"`
}
