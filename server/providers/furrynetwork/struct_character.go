package furrynetwork

type CharacterAPIResponse struct {
	Characters
	Stats ProfileStats `json:"stats"`
}

type (
	Characters struct {
		ID                         int      `json:"id"`
		Name                       string   `json:"name"`
		DisplayName                string   `json:"display_name"`
		DefaultCharacter           bool     `json:"default_character"`
		Private                    bool     `json:"private"`
		NoIndex                    bool     `json:"noindex"`
		Avatar                     string   `json:"avatar"`
		AvatarExplicit             int      `json:"avatar_explicit"`
		Banner                     *string  `json:"banner"`
		BannerExplicit             int      `json:"banner_explicit"`
		Staff                      bool     `json:"staff"`
		DateCreated                string   `json:"created"`
		DateDeleted                *string  `json:"deleted"`
		LastLogin                  string   `json:"last_login"`
		AcceptingCommissions       bool     `json:"accepting_commissions"`
		AdminBlock                 bool     `json:"admin_block"`
		PriceSheet                 *string  `json:"pricesheet"`
		PriceSheetNSFW             *string  `json:"pricesheet_nsfw"`
		PriceSheetInstructions     *string  `json:"pricesheet_instructions"`
		PriceSheetInstructionsNSFW *string  `json:"pricesheet_instructions_nsfw"`
		CommissionMedia            []int    `json:"commission_media"`
		CommissionStipulations     *string  `json:"commission_stipulations"`
		Promotes                   int      `json:"promotes"`
		SpecialtyTagLimit          int      `json:"specialty_tag_limit"`
		Rating                     int      `json:"rating"`
		EarliestCommissionType     *string  `json:"earliest_commission_type"`
		TicketID                   *int     `json:"ticket_id"`
		Points                     *int     `json:"points"`
		PaymentEmail               *string  `json:"payment_email"`
		Enabled                    bool     `json:"enabled"`
		IsAuthenticatedUser        bool     `json:"isAuthenticatedUser"`
		Following                  bool     `json:"following"`
		FollowingCommissions       bool     `json:"followingCommissions"`
		Avatars                    Avatars  `json:"avatars"`
		Banners                    *Banners `json:"banners"`
		CommissionTypes            []string `json:"commission_types"`
		CommissionAverage          int      `json:"commission_average"`
		SpecialtyTags              []string `json:"specialtyTags"`
		Reviews                    []string `json:"reviews"`
		ReviewCount                int      `json:"review_count"`
	}

	Avatars struct {
		Avatar   string `json:"avatar"`
		Original string `json:"original"`
		Small    string `json:"small"`
		Tiny     string `json:"tiny"`
	}

	Banners struct {
		Banner    string `json:"banner"`
		HalfStrip string `json:"halfStrip"`
		Original  string `json:"original"`
		Strip     string `json:"strip"`
	}
)
