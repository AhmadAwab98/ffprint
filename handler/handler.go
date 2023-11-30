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
	"strconv"
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
	var finalResponse models.BodyResponse
	// var cache models.BodyResponse

	var response models.InterimBodyResponse

	// if exists, _ := rdb.Exists(ctx, responseCache).Result(); exists == 1 {
	// 	// get the response from cache and convert it into json format
	// 	responsecached, _ := rdb.HGetAll(ctx,responseCache).Result()
	// 	_ = json.Unmarshal([]byte(responsecached["cachedData"]), &cache)
	// 	JSONresponse1, _ := json.MarshalIndent(cache, "", "	")
	// 	// respond to client using cache
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.Write([]byte(JSONresponse1))
	// 	return
	// }

	// prepare response
	recPrepareResponse(request.Path, &response, &finalResponse)

	file, _ := os.Stat(request.Path)

	finalResponse.Status = "success"
	finalResponse.Path = request.Path
	finalResponse.Contents = append(finalResponse.Contents, response.Contents[0])
	finalResponse.Size = int(file.Size()) / 1024


	// converting to JSON
	JSONresponse, errjson := json.Marshal(finalResponse)
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
	w.Write([]byte(JSONresponse))
}

func recPrepareResponse(Path string, response *models.InterimBodyResponse, finalResponse *models.BodyResponse) {
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
		finalResponse.TFolders++
		response.Type = "folder"
		response.Name = name[len(name) - 1]
		response.Path = Path
		response.Size = strconv.Itoa(int(info.Size()))
		response.LastModified = info.ModTime().Format("2006-01-02T15:04:05Z")
		files, err := os.ReadDir(Path)
		if err != nil {
			log.Println(err)
			return
		}
		for _, file := range files {
			// call the function recursively on every directory
			var subResponse models.InterimBodyResponse
			recPrepareResponse(filepath.Join(Path, file.Name()), &subResponse, finalResponse)
			response.Contents = append(response.Contents, subResponse)
		}
	default:
		// add files to content
		finalResponse.TFiles++
		response.Name = name[len(name) - 1]
		response.Type = "file"
		response.Path = Path
		response.Size = strconv.Itoa(int(info.Size()))
		response.LastModified = info.ModTime().Format("2006-01-02T15:04:05Z")
	}
}
