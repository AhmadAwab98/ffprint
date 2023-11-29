package handler

import (
	"ffprint/models"
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
	"log"
)

func ListHandler(w http.ResponseWriter, r *http.Request) {

	// decode JSON into BodyRequest
	decoder := json.NewDecoder(r.Body)
	var request models.BodyRequest
	err:= decoder.Decode(&request)
	if err != nil {
		log.Println(err)
		return
	}
	responseCache := GetMD5Hash(request.Path)

	var response, cache models.BodyResponse

	if exists, _ := rdb.Exists(ctx, responseCache).Result(); exists == 1 {
		// get the response from cache and convert it into json format
		responsecached, _ := rdb.HGetAll(ctx,responseCache).Result()
		_ = json.Unmarshal([]byte(responsecached["cachedData"]), &cache)
		JSONresponse1, _ := json.MarshalIndent(cache, "", "	")
		// respond to client using cache
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(JSONresponse1))
	}

	// prepare response
	recPrepareResponse(request.Path, &response)

	// converting to JSON
	JSONresponse, errjson := json.Marshal(response)
	if errjson != nil {
		log.Println(err)
		return
	}

	// caching response for 5 minutes
	err = rdb.HSet(ctx, responseCache, "cachedData", JSONresponse).Err()
	err = rdb.Expire(ctx, responseCache, 5*time.Minute).Err()

	if(err != nil) {
		log.Println(err)
		return
	}

	// respond to client using actual data
	w.Header().Set("Content-Type", "application/json")
}

func recPrepareResponse(Path string, response *models.BodyResponse) {
	info, err := os.Stat(Path)

	// getting the name of the directory/file
	name := strings.Split(Path, "/")
	switch {
	case err != nil:
		// if broken or not opening return
		log.Println(err)
		return
	case info.IsDir():
		// add the folder to content 
		response.Name = name[len(name) - 1]
		files, err := os.ReadDir(Path)
		if err != nil {
			log.Println(err)
			return
		}
		for _, file := range files {
			// call the function recursively on every directory
			var subResponse models.BodyResponse
			recPrepareResponse(filepath.Join(Path, file.Name()), &subResponse)
			response.Contents = append(response.Contents, subResponse)
		}
	default:
		// add files to content
		response.Name = name[len(name) - 1]
	}
}
