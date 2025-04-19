package furrynetwork

type FNQuerySize int

const (
	QuerySize12 FNQuerySize = 12
	QuerySize24 FNQuerySize = 24
	QuerySize36 FNQuerySize = 36
	QuerySize48 FNQuerySize = 48
	QuerySize60 FNQuerySize = 60
	QuerySize72 FNQuerySize = 72
)

type FurryNetworkParams struct {
	Size      FNQuerySize `url:"size"`
	From      *int        `url:"from,omitempty"`
	Time      *string     `url:"time,omitempty"`
	Character string      `url:"character"`
	Types     []string    `url:"types"`
	Sort      string      `url:"sort"`
}
