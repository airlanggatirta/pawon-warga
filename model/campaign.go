package model

import "database/sql"

type Campaign struct {
	ID                        uint64          `gorm:"column:id"`
	Title                     string          `gorm:"column:title"`
	UserID                    uint64          `gorm:"column:users_id"`
	CampaignCategoryID        uint64          `gorm:"column:project_categories_id"`
	ShortURL                  sql.NullString  `gorm:"column:short_url"`
	FundingTarget             sql.NullInt64   `gorm:"column:funding_target"` // *uint64
	LocationID                uint64          `gorm:"column:locations_id"`
	ShortDescription          sql.NullString  `gorm:"column:short_description"`
	Description               sql.NullString  `gorm:"column:description"`
	Image                     sql.NullString  `gorm:"column:image"`
	Video                     sql.NullString  `gorm:"column:video"`
	CampaignStatusID          uint64          `gorm:"column:project_statuses_id"`
	IsSuspended               bool            `gorm:"column:is_suspended"`
	Created                   int64           `gorm:"column:created"`
	Updated                   int64           `gorm:"column:updated"`
	Launched                  sql.NullInt64   `gorm:"column:launched"`
	Expired                   sql.NullInt64   `gorm:"column:expired"`
	IsFeatured                bool            `gorm:"column:is_featured"`
	Visits                    uint64          `gorm:"column:visits"`
	PopularityIndex           sql.NullFloat64 `gorm:"column:popularity_index"`
	CampaignPartnerID         sql.NullInt64   `gorm:"column:partners_id"` // *uint64
	PartnerPageStickiness     bool            `gorm:"column:partner_page_stickiness"`
	TrackingScript            sql.NullString  `gorm:"column:tracking_script"`
	DisplayRule               string          `gorm:"column:display_rule"`
	FinalDonationAmount       uint64          `gorm:"column:final_donation_amount"`
	FinalDonationPercentage   float64         `gorm:"column:final_donation_percentage"`
	LockVersion               int64           `gorm:"column:lock_version"`
	FeePercentage             float64         `gorm:"column:fee_percentage"`
	NetDonationAmount         int64           `gorm:"column:net_donation_amount"`
	PlatformFee               int64           `gorm:"column:platform_fee"`
	IsForeverRunning          bool            `gorm:"column:is_forever_running"`
	IsOpenGoal                bool            `gorm:"column:is_open_goal"`
	PayTo                     string          `gorm:"column:pay_to"`
	PayInstantly              bool            `gorm:"column:pay_instantly"`
	IsZakat                   bool            `gorm:"column:is_zakat"`
	IsChecked                 bool            `gorm:"column:is_checked"`
	SearchServerStatus        bool            `gorm:"column:search_server_status"`
	UtmCampaign               sql.NullString  `gorm:"column:utm_campaign"`
	UtmMedium                 sql.NullString  `gorm:"column:utm_medium"`
	UtmSource                 sql.NullString  `gorm:"column:utm_source"`
	FirstDonation             sql.NullInt64   `gorm:"column:first_donation"`
	FiveDonations             sql.NullInt64   `gorm:"column:five_donations"`
	TwentyPercent             sql.NullInt64   `gorm:"column:twenty_percent"`
	Source                    string          `gorm:"column:source"`
	Rating                    sql.NullString  `gorm:"column:rating"`
	CustomCta                 sql.NullString  `gorm:"column:custom_cta"`
	ParentProjectID           sql.NullInt64   `gorm:"column:parent_project_id"` // *uint64
	AgentsID                  sql.NullInt64   `gorm:"column:agents_id"`         // *uint64
	AgentNotifications        sql.NullBool    `gorm:"column:agent_notifications"`
	DonationNotification      bool            `gorm:"column:donation_notification"`
	Beneficiary               sql.NullString  `gorm:"column:beneficiary"`
	IsSameCityWithBeneficiary sql.NullBool    `gorm:"column:is_same_city_with_beneficiary"`
	RequestUserdata           sql.NullBool    `gorm:"column:request_userdata"`
	IsVerified                sql.NullBool    `gorm:"column:is_verified"`
	DonationReceived          int64           `gorm:"-"`
	DonationCount             int64           `gorm:"-"`
	DonationTotal             int64           `gorm:"-"`
	DaysRunning               int64           `gorm:"-"`
	DaysRemaining             int64           `gorm:"-"`
	SecondsRunning            int64           `gorm:"-"`
	IsOnline                  bool            `gorm:"-"`
	IsInstantZakat            bool            `gorm:"-"`
	IsHaveParent              bool            `gorm:"-"`
	HavePartner               bool            `gorm:"-"`
	IsOpenForDonation         bool            `gorm:"-"`
	IsLive                    bool            `gorm:"-"`
	IsExpired                 bool            `gorm:"-"`
	User                      User
	ParentCampaign            *Campaign `gorm:"foreignkey:ID;association_foreignkey:ParentProjectID"`
	CampaignCategory          CampaignCategory
	CampaignPartner           *CampaignPartner
	Location                  Location
	CampaignStatus            CampaignStatus
	Updates                   []CampaignUpdate
	Donation                  []Donation
}

func (Campaign) TableName() string {
	return "projects"
}
