package model

type Partner struct {
	ID         			uint64 	`gorm:"column:id"`
	UserID     			uint64 	`gorm:"column:users_id"`
	Name 				string 	`gorm:"column:name"`
	Description    		string	`gorm:"column:description"`
	Url    				string 	`gorm:"column:url"`
	Image				string  `gorm:"column:image"`
	Slug    			string  `gorm:"column:slug"`
	IsActive    		int  	`gorm:"column:is_active"`
	Created    			int64  	`gorm:"column:created"`
	Updated    			int64  	`gorm:"column:updated"`
	DisplayOrder    	int  	`gorm:"column:display_order"`
	ProjectsOrdering	string  `gorm:"column:projects_ordering"`
	CtaText    			string  `gorm:"column:cta_text"`
	CtaUrl    			string  `gorm:"column:cta_url"`
	PartnerType    		string  `gorm:"column:partner_type"`
	BannerImage			string  `gorm:"column:banner_image"`
	User				User
}