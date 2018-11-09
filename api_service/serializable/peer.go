package serializable

type Node struct {
	ID             string `json:"id"`
	Addr           string `json:"addr"`
	LastActiveTime string `json:"lastActiveTime"`
}

type Heartbeat struct {
	Delay  float64 `json:"delay"`
	Health uint8   `json:"health"`
}

type Peer struct {
	Guid      string    `json:"guid"`
	Node      Node      `json:"node"`
	Heartbeat Heartbeat `json:"heartbeat"`
}
