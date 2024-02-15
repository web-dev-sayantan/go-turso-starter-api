package data

type Homestay struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Address      string `json:"address"`
	LocationName string `json:"locationName"`
}

type Location struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	State       string  `json:"state"`
	Description string  `json:"description"`
	Lat         float32 `json:"lat"`
	Long        float32 `json:"long"`
	Altitude    uint    `json:"altitude"`
	CoverUrl    string  `json:"coverUrl"`
}
