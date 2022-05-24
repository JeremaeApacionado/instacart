package route

import (
	"encoding/json"
	"instacart/models"
	"log"
	"net/http"
)

// SearchResult struct to handle search queries
type SearchResult struct {
	Products []models.Product `json:"product"`
	Input string `json:"input"`
   }
   func searchHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	searchInput := r.Form.Get("input")
   log.Print("Querying database for: ", searchInput)
   
   searchResult := SearchResult{
	 Input: searchInput,
	}
   jsonData, err := json.Marshal(searchResult)
	if err != nil {
	 log.Print("JSON executing error: ", err)
	 return
	}
   w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
   }

func SearchContent(searchInput string) {
	panic("unimplemented")
}