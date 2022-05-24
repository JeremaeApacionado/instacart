package models
// Page struct to store in database
type Search struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
   }