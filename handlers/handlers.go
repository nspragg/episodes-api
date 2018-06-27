package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"

	"github.com/nspragg/episodes-api/dao"

	"github.com/gorilla/mux"
	"github.com/nspragg/episodes-api/model"
	// translation "github.com/nspragg/episodes-api/translations"
)

var translationTemplate = regexp.MustCompile(`#{(.*?)}`)

const (
	keyPrefix     = "episodes:"
	version       = "2.0"
	schema        = "/ibl/v1/schema/ibl.json"
	defaultRights = "web"
)

// Parameters ..
type Parameters struct {
	id     string
	rights string
}

// GetEpisodes an episode handler
func GetEpisodes(router *mux.Router, client dao.Dao) func(writer http.ResponseWriter, request *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		parms, err := getParms(r)

		key := createCacheKey(parms.id, parms.rights)
		reply, err := client.Find(key)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		response, err := buildResponse(reply, parms.rights)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, string(response))
	}
}

func createCacheKey(pid string, rights string) string {
	return keyPrefix + pid + ":" + rights
}

func getParms(r *http.Request) (*Parameters, error) {
	params := mux.Vars(r)
	if len(params) == 0 {
		return nil, errors.New("Mandatory args absent. 1 - 20 pids expected")
	}
	parameters := new(Parameters)
	parameters.id = params["id"]

	query := r.URL.Query()
	rights := query.Get("rights")
	if rights == "" {
		rights = defaultRights

	}
	parameters.rights = rights

	return parameters, nil
}

func buildResponse(episodesToUnmarshall []string, rights string) (string, error) {
	episodes := make([]model.Episode, 0)
	for _, data := range episodesToUnmarshall {
		// translation.Translate(data, rights)
		var episode model.Episode
		if err := json.Unmarshal([]byte(data), &episode); err != nil {
			return "", err
		}
		episodes = append(episodes, episode)
	}

	ibl := model.Root{
		version,
		schema,
		episodes,
	}

	result, err := json.Marshal(ibl)
	if err != nil {
		return "", err
	}
	return string(result), nil
}
