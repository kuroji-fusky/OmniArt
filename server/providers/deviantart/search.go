package deviantart

type (
	DASearchPath string
	DAQueryOrder string
)

const (
	QueryOrderWatch      DAQueryOrder = "watch" // Deviants You Watch
	QueryOrderMostRecent DAQueryOrder = "most-recent"
	QueryOrderWeek       DAQueryOrder = "this-week"
	QueryOrderMonth      DAQueryOrder = "this-month"
	QueryOrderYear       DAQueryOrder = "this-year"
	// The "This century" option is selected by default and returns and empty query string

	DA_BASE_SEARCH_PATH DASearchPath = "/search"
	DA_DEVIATIONS       DASearchPath = DA_BASE_SEARCH_PATH + "/deviations"
	DA_JOURNALS         DASearchPath = DA_BASE_SEARCH_PATH + "/journals"
	DA_GROUPS           DASearchPath = DA_BASE_SEARCH_PATH + "/groups"
	DA_ARTISTS          DASearchPath = DA_BASE_SEARCH_PATH + "/artists" // Known as "Deviants" tab
	DA_STATUS_UPDATES   DASearchPath = DA_BASE_SEARCH_PATH + "/status-updates"
	DA_POLLS            DASearchPath = DA_BASE_SEARCH_PATH + "/polls"
	DA_COMMISSIONS      DASearchPath = DA_BASE_SEARCH_PATH + "/commissions"
)

type DASearchQuery struct {
	Query string       `query:"q"`
	Order DAQueryOrder `json:"order,omitempty" query:"order" ts_type:"'watch' | 'most-recent' | 'this-week' | 'this-month' | 'this-year'"`
}

type DASearchQueryAuto struct {
	DASearchQuery
	// The `cursor` parameter is auto-generated at request time, not something you'd manually
	// input into, and it's here for completeness sake
	Cursor string `json:"cursor,omitempty" query:"string,readonly"`
}
