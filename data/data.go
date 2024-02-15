package data

type Homestay struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Address    string `json:"address"`
	LocationId int    `json:"locationId"`
}

type Location struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Lat         float32 `json:"lat"`
	Long        float32 `json:"long"`
	Altitude    int     `json:"altitude"`
}
