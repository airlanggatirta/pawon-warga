package model

import (
	"database/sql"
)

type User struct {
	ID                    uint64         `gorm:"column:id"`
	SecondaryUserID       string         `gorm:"column:secondary_id"`
	PrivilegeID           uint64         `gorm:"column:privileges_id"`
	UserStatusID          uint64         `gorm:"column:user_statuses_id"`
	Email                 sql.NullString `gorm:"column:email" valid:"email"`
	Password              string         `gorm:"column:password" valid:"length(8|255)"`
	FullName              string         `gorm:"column:full_name"`
	PhoneNumber           sql.NullString `gorm:"column:phone_number"`
	Image                 sql.NullString `gorm:"column:image"`
	Biography             sql.NullString `gorm:"column:biography"`
	LocationID            sql.NullInt64  `gorm:"column:locations_id"` // *uint64
	Address               string         `gorm:"column:address"`
	Created               int64          `gorm:"column:created"`
	Updated               sql.NullInt64  `gorm:"column:updated"`
	LastSeenOnline        sql.NullInt64  `gorm:"column:last_seen_online"`
	ActivationKey         string         `gorm:"column:activation_key"`
	WalletKey             sql.NullString `gorm:"column:wallet_key"`
	ResetPasswordKey      sql.NullString `gorm:"column:reset_password_key"`
	ResetPasswordExpired  sql.NullInt64  `gorm:"column:reset_password_expired"`
	OrganizationStatus    string         `gorm:"column:organization_status"`
	APIToken              sql.NullString `gorm:"column:api_token"`
	FbUID                 sql.NullString `gorm:"column:fb_uid"`
	CurrentDonationCount  int64          `gorm:"column:current_donation_count"`
	CurrentDonationAmount int64          `gorm:"column:current_donation_amount"`
	VerificationIDScan    sql.NullString `gorm:"column:verification_id_scan"`
	VerificationImage     sql.NullString `gorm:"column:verification_image"`
	VerificationStatus    string         `gorm:"column:verification_status"`
	VerificationChannel   sql.NullString `gorm:"column:verification_channel"`
	VerificationTimestamp sql.NullInt64  `gorm:"column:verification_timestamp"`
	MaxLiveCampaign       int64          `gorm:"column:max_live_campaign"`
	FacebookURL           sql.NullString `gorm:"column:facebook_url"`
	LinkedinURL           sql.NullString `gorm:"column:linkedin_url"`
	WebsiteURL            sql.NullString `gorm:"column:website_url"`
	InstagramAccount      sql.NullString `gorm:"column:instagram_account"`
	IsPremium             sql.NullBool   `gorm:"column:is_premium"`
	Gender                sql.NullString `gorm:"column:gender"`
	ZakatInfo             sql.NullBool   `gorm:"column:zakat_info"`
	Birthday              int64          `gorm:"column:birthday"`
	ReminderDate          sql.NullInt64  `gorm:"column:reminder_date"`
	GoogleID              sql.NullString `gorm:"column:google_id"`
	WaPhoneNumber         sql.NullString `gorm:"column:wa_phone_number"`
	//Verified                int64     `gorm:"column:verified"`
	WaVerificationStatus         sql.NullBool   `gorm:"column:wa_verification_status"`
	EmailVerificationStatus      sql.NullBool   `gorm:"column:email_verification_status"`
	Npwz                         sql.NullString `gorm:"column:npwz"`
	TempID                       string         `gorm:"-"`
	FirebaseUID                  sql.NullString `gorm:"column:firebase_uid"`
	LastDonationPaymentMethodsID int64          `gorm:"column:last_donation_payment_methods_id"`
	LastTopupPaymentMethodsID    int64          `gorm:"column:last_topup_payment_methods_id"`
	IsVerified                   bool           `gorm:"-"`
	CampaignCreatedCount         int64          `gorm:"-"`
	Privilege                    Privilege
	UserStatus                   UserStatus
	Location                     Location
}

func (User) TableName() string {
	return "users"
}
