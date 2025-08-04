package schemas


type TrendingItem struct {
	ID     string  `json:"id"`     // Unique identifier for the trending item
	Name   string  `json:"name"`   // Name of the trending item
	Weight float64 `json:"weight"` // Weight of the trending item
	Images []Image `json:"images"` // List of images associated with the trending item
}


type UserTrending struct {
	ID           string         `json:"id"`
	User         string         `json:"user"`
	DurationType string         `json:"duration_type"`
	Songs        []TrendingItem `json:"songs"`
	Artists      []TrendingItem `json:"artists"`
	CreatedAt    string         `json:"createdAt"`
	UpdatedAt    string         `json:"updatedAt"`
}

