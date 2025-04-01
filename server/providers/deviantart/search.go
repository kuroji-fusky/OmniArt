package deviantart

const (
	QueryOrderWatch      = "watch" // Deviants You Watch
	QueryOrderMostRecent = "most-recent"
	QueryOrderWeek       = "this-week"
	QueryOrderMonth      = "this-month"
	QueryOrderYear       = "this-year"
	// The "This century" option is selected by default and returns and empty query string

	DA_BASE_SEARCH_PATH = "/search"
	DA_DEVIATIONS       = DA_BASE_SEARCH_PATH + "/deviations"
	DA_JOURNALS         = DA_BASE_SEARCH_PATH + "/journals"
	DA_GROUPS           = DA_BASE_SEARCH_PATH + "/groups"
	DA_ARTISTS          = DA_BASE_SEARCH_PATH + "/artists" // Known as "Deviants" tab
	DA_STATUS_UPDATES   = DA_BASE_SEARCH_PATH + "/status-updates"
	DA_POLLS            = DA_BASE_SEARCH_PATH + "/polls"
	DA_COMMISSIONS      = DA_BASE_SEARCH_PATH + "/commissions"
)

type DASearchQuery struct {
	Query  string `query:"q"`
	Cursor string `json:"cursor,omitempty" query:"string,readonly"`
	Order  string `json:"order,omitempty" query:"order" ts_type:"'watch' | 'most-recent' | 'this-week' | 'this-month' | 'this-year'"`
}
