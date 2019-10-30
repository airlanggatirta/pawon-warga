package model

import (
	"database/sql"
	"strings"

	"github.com/go-sql-driver/mysql"
)

type Donation struct {
	ID                          uint64         `gorm:"column:id"`
	CampaignID                  uint64         `gorm:"column:projects_id"`
	UserID                      uint64         `gorm:"column:users_id"`
	Amount                      int64          `gorm:"column:amount"`
	UniqueCode                  int64          `gorm:"column:unique_code"`
	Fee                         int64          `gorm:"column:fee"`
	BankChargeFee               int64          `gorm:"column:bank_charge_fee"`
	PaymentMethodID             sql.NullInt64  `gorm:"column:payment_methods_id"`
	DonationStatusID            uint64         `gorm:"column:donation_statuses_id"`
	ReferenceID                 sql.NullInt64  `gorm:"column:reference_id"`
	IsAnonymous                 bool           `gorm:"column:is_anonymous"`
	Created                     int64          `gorm:"column:created"`
	Expire                      sql.NullInt64  `gorm:"column:expire"`
	Confirmed                   sql.NullInt64  `gorm:"column:confirmed"`
	Verified                    sql.NullInt64  `gorm:"column:verified"`
	Paid                        sql.NullInt64  `gorm:"column:paid"`
	Updated                     sql.NullInt64  `gorm:"column:updated"`
	ConfirmationCode            sql.NullString `gorm:"column:confirmation_code"`
	McBankName                  sql.NullString `gorm:"column:mc_bank_name"`
	McBankAccHolder             sql.NullString `gorm:"column:mc_bank_acc_holder"`
	McAmount                    sql.NullInt64  `gorm:"column:mc_amount"`
	McTransferAt                mysql.NullTime `gorm:"column:mc_transfer_at"`
	McNotes                     sql.NullString `gorm:"column:mc_notes"`
	McImage                     sql.NullString `gorm:"column:mc_image"`
	McPotRequestSent            bool           `gorm:"column:mc_pot_request_sent"`
	RefererID                   sql.NullInt64  `gorm:"column:referer_id"` // *uint64
	TempTotalDonation           sql.NullInt64  `gorm:"column:temp_total_donation"`
	TempTotalInvoiced           sql.NullInt64  `gorm:"column:temp_total_invoiced"`
	ReminderSent                bool           `gorm:"column:reminder_sent"`
	UtmSource                   sql.NullString `gorm:"column:utm_source"`
	UtmMedium                   sql.NullString `gorm:"column:utm_medium"`
	UtmCampaign                 sql.NullString `gorm:"column:utm_campaign"`
	LockVersion                 int64          `gorm:"column:lock_version"`
	PhoneNumber                 sql.NullString `gorm:"column:phone_number"`
	FeePercentage               float64        `gorm:"column:fee_percentage"`
	MarketingFeePercentage      float64        `gorm:"column:marketing_fee_percentage"`
	NetAmount                   int64          `gorm:"column:net_amount"`
	PlatformFee                 int64          `gorm:"column:platform_fee"`
	MarketingFee                int64          `gorm:"column:marketing_fee"`
	Source                      string         `gorm:"column:source"`
	Comment                     sql.NullString `gorm:"column:comment"`
	ShareUserdataAgreement      sql.NullBool   `gorm:"column:share_userdata_agreement"`
	CampaignerRelation          sql.NullBool   `gorm:"column:campaigner_relation"`
	CampaignMarketingCampaignID sql.NullInt64  `gorm:"column:project_marketing_campaign_id"` // *uint64
	TransferredFrom             sql.NullInt64  `gorm:"column:transferred_from"`              // *uint64
	RedonateID                  sql.NullInt64  `gorm:"column:redonate_id"`                   // *uint64
	NthDonation                 int64          `gorm:"column:nth_donation"`
	SendNotifications           bool           `gorm:"column:send_notifications"`
	ChildProjectID              sql.NullInt64  `gorm:"column:child_project_id"` // *uint64
	VaNumber                    sql.NullString `gorm:"column:va_number"`
	VendorTrxID                 sql.NullString `gorm:"column:vendor_trx_id"`
	ParentDonationID            sql.NullInt64  `gorm:"column:parent_donation_id"`
	ParentStatus                int            `gorm:"column:parent_status"`
	Cashtag                     string         `gorm:"-"`
	IsZakat                     bool           `gorm:"-"`
	PromoCode                   string         `gorm:"-"`
	RedirectCallback            string         `gorm:"-"`
	RedirectCallbackParams      string         `gorm:"-"`
	SanguGopayGenereteQRCode    string         `gorm:"-"`
	SanguGopayDeeplinkRedirect  string         `gorm:"-"`
	SanguLinkAjaSuccessUrl      string         `gorm:"-"`
	SanguLinkAjaFailUrl         string         `gorm:"-"`
	SanguLinkAjaPgpToken        string         `gorm:"-"`
	SanguLinkCheckoutUrl        string         `gorm:"-"`
	SanguDanaPayReturnUrl       string         `gorm:"-"`
	SanguDanaCheckoutUrl        string         `gorm:"-"`
	ImplementGopaySangu         bool           `gorm:"-"`
	DependentId                 string         `gorm:"-"` // For urunan/proteksi
	DependentStatus             int            `gorm:"-"` // For urunan/proteksi
	PaymentMethod               PaymentMethod
	CampaignMarketingCampaign   CampaignMarketingCampaign
	Campaign                    Campaign
	User                        User
	DonationStatus              DonationStatus
}

func (Donation) TableName() string {
	return "donations"
}

func (d Donation) ReferenceName() string {
	return strings.Title(d.TableName())
}
