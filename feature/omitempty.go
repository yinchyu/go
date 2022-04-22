package feature

type address struct {
	Street     string      `json:"street"`
	Ste        string      `json:"suite,omitempty"`
	City       string      `json:"city"`
	State      string      `json:"state"`
	Zipcode    string      `json:"zipcode"`
	Coordinate *coordinate `json:"coordinate,omitempty"`
}

type coordinate struct {
	Lat float64 `json:"latitude"`
	Lng float64 `json:"longitude"`
}
