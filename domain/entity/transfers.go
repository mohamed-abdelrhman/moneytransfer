package entity

type Transfer struct {
	ID        string  `json:"id"`
	OriginID      string `json:"origin_id"`
	DestinationID      string `json:"destination_id"`
	Amount     int `json:"amount"`
}

type Transfers []Transfer
