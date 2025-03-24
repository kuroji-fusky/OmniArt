package furrynetwork

const (
	QuerySize12 = 12
	QuerySize24 = 24
	QuerySize36 = 36
	QuerySize48 = 48
	QuerySize60 = 60
	QuerySize72 = 72
)

type FurryNetworkParams struct {
	Size      int      `url:"size"`
	From      int      `url:"from,omitempty"`
	Time      string   `url:"time,omitempty"`
	Character string   `url:"character"`
	Types     []string `url:"types"`
	Sort      string   `url:"sort"`
}
