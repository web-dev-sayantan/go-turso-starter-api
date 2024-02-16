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

type Room struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Category        string `json:"category"`
	BaseOccupancy   int16  `json:"baseOccupancy"`
	ExtraOccupancy  int16  `json:"extraOccupancy"`
	ToiletAttached  bool   `json:"toiletAttached"`
	BalconyAttached bool   `json:"balconyAttached"`
	KitchenAttached bool   `json:"kitchenAttached"`
	AirConditioned  bool   `json:"airConditioned"`
	IsDorm          bool   `json:"isDorm"`
	Recommended     bool   `json:"recommended"`
	HomestayName    string `json:"HomestayName"`
}
