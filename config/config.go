package config

import (
	"time"

	"github.com/airlanggatirta/pawon-warga/common"
	"github.com/airlanggatirta/pawon-warga/driver"
	"github.com/spf13/viper"
)

type App struct {
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"version"`
	Host    string `mapstructure:"host"`
	Port    int    `mapstructure:"port"`
}

type Database struct {
	Host                  string        `mapstructure:"host"`
	Port                  int           `mapstructure:"port"`
	Username              string        `mapstructure:"username"`
	Password              string        `mapstructure:"password"`
	Name                  string        `mapstructure:"name"`
	AdditionalParameters  string        `mapstructure:"additional_parameters"`
	MaxOpenConnection     int           `mapstructure:"max_open_connection"`
	MaxIdleConnection     int           `mapstructure:"max_idle_connection"`
	ConnectionMaxLifetime time.Duration `mapstructure:"connection_max_lifetime"`
}

type Cache struct {
	Host               string        `mapstructure:"host"`
	Port               int           `mapstructure:"port"`
	DialConnectTimeout time.Duration `mapstructure:"dial_connect_timeout"`
	ReadTimeout        time.Duration `mapstructure:"read_timeout"`
	WriteTimeout       time.Duration `mapstructure:"write_timeout"`
	MaxIdle            int           `mapstructure:"max_idle"`
	MaxActive          int           `mapstructure:"max_active"`
	IdleTimeout        time.Duration `mapstructure:"idle_timeout"`
	Wait               bool          `mapstructure:"wait"`
	MaxConnLifetime    time.Duration `mapstructure:"max_conn_lifetime"`
	Namespace          string        `mapstructure:"namespace"`
	Password           string        `mapstructure:"password"`
	Items              Items         `mapstructure:"items"`
}

type CacheItem struct {
	Prefix string        `mapstructure:"prefix" validate:"required"`
	TTL    time.Duration `mapstructure:"ttl" validate:"required"`
}

type Items struct {
	Session                           CacheItem `mapstructure:"session"`
	GetAllLoveList                    CacheItem `mapstructure:"get_all_lovelist_service"`
	TempEmailRegistration             CacheItem `mapstructure:"temp_email_registration"`
	TempWaRegistration                CacheItem `mapstructure:"temp_wa_registration"`
	GetAllBanner                      CacheItem `mapstructure:"get_all_banner"`
	GetAllCampaignUpdates             CacheItem `mapstructure:"get_all_campaign_updates"`
	GetLoveLists                      CacheItem `mapstructure:"get_love_lists"`
	GetAllUserLovelist                CacheItem `mapstructure:"get_all_user_lovelist"`
	GetAllUsers                       CacheItem `mapstructure:"get_all_users"`
	GetOneUser                        CacheItem `mapstructure:"get_one_user"`
	GetOneUserByCredentials           CacheItem `mapstructure:"get_one_user_by_credentials"`
	GetOneUserBySecondaryID           CacheItem `mapstructure:"get_one_user_by_secondary_id"`
	GetOneUserByEmail                 CacheItem `mapstructure:"get_one_user_by_email"`
	GetOneUserByWhatsapp              CacheItem `mapstructure:"get_one_user_by_whatsapp"`
	GetUserRegistrationTemp           CacheItem `mapstructure:"get_user_registration_temp"`
	GetSearchUser                     CacheItem `mapstructure:"get_search_user"`
	GetAllCategories                  CacheItem `mapstructure:"get_all_categories"`
	GetOneCampaignCategory            CacheItem `mapstructure:"get_one_campaign_category"`
	GetDonationListByCampaign         CacheItem `mapstructure:"get_donation_list_by_campaign"`
	GetDonationListByUser             CacheItem `mapstructure:"get_donation_list_by_user"`
	GetDonationListByUserRange        CacheItem `mapstructure:"get_donation_list_by_user_range"`
	GetDonatedCampaignId              CacheItem `mapstructure:"get_donated_campaign_id"`
	GetTotalDonationByMonth           CacheItem `mapstructure:"get_total_donation_by_month"`
	GetGoldPrice                      CacheItem `mapstructure:"get_gold_price"`
	GetAllDonationStatus              CacheItem `mapstructure:"get_all_donation_status"`
	GetAllDonatedCampaignId           CacheItem `mapstructure:"get_all_donated_campaign_id"`
	GetCampaignIdByShortUrl           CacheItem `mapstructure:"get_campaign_id_by_short_url"`
	GetAllCampaign                    CacheItem `mapstructure:"get_all_campaign"`
	GetOneCampaign                    CacheItem `mapstructure:"get_one_campaign"`
	FundraiserList                    CacheItem `mapstructure:"fundraiser_list"`
	JeniusTrx                         CacheItem `mapstructure:"jenius_trx"`
	GetAllHomeBanner                  CacheItem `mapstructure:"get_all_home_banner"`
	FlagAssignmentModelId             CacheItem `mapstructure:"flag_assignment_model_id"`
	GetAllCuratedUpdate               CacheItem `mapstructure:"get_all_curated_update"`
	GetCampaignUpdateByCampaignId     CacheItem `mapstructure:"get_campaign_update_by_campaign_id"`
	GetOneFlagByRefModelAndFlagname   CacheItem `mapstructure:"get_one_flag_by_ref_model_and_flagname"`
	GetOnePartner                     CacheItem `mapstructure:"get_one_partner"`
	GetWalletBalance                  CacheItem `mapstructure:"get_wallet_balance"`
	GetWalletTopupStatus              CacheItem `mapstructure:"get_wallet_topup_status"`
	GetWalletTransaction              CacheItem `mapstructure:"get_wallet_transaction"`
	GetAllWalletTransactionType       CacheItem `mapstructure:"get_all_wallet_transaction_type"`
	GetAllCampaignStatus              CacheItem `mapstructure:"get_all_campaign_status"`
	GetAllCampaignReport              CacheItem `mapstructure:"get_all_campaign_report"`
	GetOneCampaignReport              CacheItem `mapstructure:"get_one_campaign_report"`
	GetAllUserStatus                  CacheItem `mapstructure:"get_all_user_status"`
	GetOneUserStatus                  CacheItem `mapstructure:"get_one_user_status"`
	GetAllLocation                    CacheItem `mapstructure:"get_all_location"`
	GetOneLocation                    CacheItem `mapstructure:"get_one_location"`
	GetAllPaymentMethod               CacheItem `mapstructure:"get_all_payment_method"`
	GetOnePaymentMethod               CacheItem `mapstructure:"get_one_payment_method"`
	GetAllPrivilege                   CacheItem `mapstructure:"get_all_privilege"`
	GetOnePrivilege                   CacheItem `mapstructure:"get_one_privilege"`
	OTP                               CacheItem `mapstructure:"otp"`
	RegistrationOTPFlag               CacheItem `mapstructure:"registration_otp_flag"`
	OTPDisbursement                   CacheItem `mapstructure:"otp_disbursement"`
	GetOneUserByEmailWithoutFilter    CacheItem `mapstructure:"get_one_user_by_email_without_filter"`
	GetOneUserByWhatsappWithoutFilter CacheItem `mapstructure:"get_one_user_by_whatsapp_without_filter"`
	GetOneCampaignService             CacheItem `mapstructure:"get_one_campaign_service"`
	GetWalletBalanceService           CacheItem `mapstructure:"get_wallet_balance_service"`
	GetCampaignerPublic               CacheItem `mapstructure:"get_campaigner_public"`
	GetUserSupportedCampaignPublic    CacheItem `mapstructure:"get_user_supported_campaign_public"`
	GetCampaignCategoryIdBySlug       CacheItem `mapstructure:"get_campaign_category_id_by_slug"`
	GetAccessTokenSangu               CacheItem `mapstructure:"get_access_token_sangu"`
	GetCampaignGDV                    CacheItem `mapstructure:"get_campaign_gdv"`
	SendOTPLimit                      CacheItem `mapstructure:"send_otp_limit"`
	GetCampaignDetailSimple           CacheItem `mapstructure:"get_campaign_detail_simple"`
}

type Jwt struct {
	SignKey   string `mapstructure:"sign_key"`
	SecretKey string `mapstructure:"secret_key"`
	Issuer    string `mapstructure:"issuer"`
}

// HmacSignature defines HMAC signature config
type HmacSignature struct {
	SecretKey string `mapstructure:"secret_key"`
}

// KtbsHeader defines config for Ktbs header
type KtbsHeader struct {
	ExcludeMobileHeaderCheck []string `mapstructure:"exclude_mobile_header_check"`
}

type Email struct {
	Header string `mapstructure:"header"`
	Footer string `mapstructure:"footer"`
}

type Firebase struct {
	CredentialPath string `mapstructure:"credential_path"`
}

type ActivationEmail struct {
	Content string `mapstructure:"content"`
}

type ResetPasswordRequestEmail struct {
	Content string `mapstructure:"content"`
}

type Users struct {
	TokenTTL                  time.Duration             `mapstructure:"token_ttl"`
	ActivationURL             string                    `mapstructure:"activation_url"`
	ResetPasswordLink         string                    `mapstructure:"reset_password_link"`
	ActivationKeySalt         string                    `mapstructure:"activation_key_salt"`
	ActivationEmail           ActivationEmail           `mapstructure:"activation_email"`
	ResetPasswordRequestEmail ResetPasswordRequestEmail `mapstructure:"reset_password_request_email"`
}

type Ses struct {
	Region     string `mapstructure:"region"`
	Key        string `mapstructure:"key"`
	Secret     string `mapstructure:"secret"`
	MailSender string `mapstructure:"mail_sender"`
}

type Otp struct {
	Host            string `mapstructure:"host"`
	VerifyNumberURL string `mapstructure:"verify_number_url"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
}

type OtpDisbursement struct {
	MaxRetry   int `mapstructure:"max_retry"`
	MaxRequest int `mapstructure:"max_request"`
}

type Zakat struct {
	RicePrice int `mapstructure:"rice_price"`
	GoldPrice int `mapstructure:"gold_price"`
}

type CustomerIO struct {
	SiteID           string `mapstructure:"site_id"`
	APIKey           string `mapstructure:"api_key"`
	CustomerEndpoint string `mapstructure:"customer_endpoint"`
	CustomerIdPrefix string `mapstructure:"customer_id_prefix"`
}

type Sns struct {
	Key                   string `mapstructure:"key"`
	Secret                string `mapstructure:"secret"`
	Region                string `mapstructure:"region"`
	SnsSmsServiceTopicArn string `mapstructure:"sns_sms_service_topic_arn"`
}

type CampaignerNotificationGetDonation struct {
	Content string `mapstructure:"content"`
}

type DonorNotificationWallet struct {
	Content string `mapstructure:"content"`
}

type Donations struct {
	CampaignerNotificationGetDonation CampaignerNotificationGetDonation `mapstructure:"campaigner_notification_get_donation"`
	DonorNotificationWallet           DonorNotificationWallet           `mapstructure:"donor_notification_wallet"`
}

type Link struct {
	BaseURL      string `mapstructure:"base_url"`
	AssetsURL    string `mapstructure:"assets_url"`
	NewAssetsURL string `mapstructure:"new_assets_url"`
}

type Jenius struct {
	CreateTrxEndpoint string `mapstructure:"create_trx_endpoint"`
	Callback          string `mapstructure:"callback"`
}

type Gopay struct {
	CreateTrxEndpoint string `mapstructure:"create_trx_endpoint"`
	Callback          string `mapstructure:"callback"`
}

type VA struct {
	CreateTrxEndpoint string        `mapstructure:"create_trx_endpoint"`
	Callback          string        `mapstructure:"callback"`
	Expired           time.Duration `mapstructure:"expired"`
}

type LinkAja struct {
	CreateTrxEndpoint string `mapstructure:"create_trx_endpoint"`
	Callback          string `mapstructure:"callback"`
	MinimumAmount     int64  `mapstructure:"minimum_amount"`
}

type Dana struct {
	CreateTrxEndpoint string        `mapstructure:"create_trx_endpoint"`
	Callback          string        `mapstructure:"callback"`
	Expired           time.Duration `mapstructure:"expired"`
	MinimumAmount     int64         `mapstructure:"minimum_amount"`
}

type Sangu struct {
	SanguBaseURL           string  `mapstructure:"sangu_base_url"`
	SanguUsername          string  `mapstructure:"sangu_username"`
	SanguPassword          string  `mapstructure:"sangu_password"`
	SanguGetTokenEndpoint  string  `mapstructure:"sangu_get_token_endpoint"`
	SanguCreateTrxEndpoint string  `mapstructure:"sangu_create_trx_endpoint"`
	Jenius                 Jenius  `mapstructure:"jenius"`
	Gopay                  Gopay   `mapstructure:"gopay"`
	VA                     VA      `mapstructure:"va"`
	LinkAja                LinkAja `mapstructure:"linkaja"`
	Dana                   Dana    `mapstructure:"dana"`
}

type Cloudsearch struct {
	Region         string `mapstructure:"region"`
	SearchEndpoint string `mapstructure:"search_endpoint"`
	Key            string `mapstructure:"key"`
	Secret         string `mapstructure:"secret"`
}

type Midtrans struct {
	ServerKey        string `mapstructure:"server_key"`
	ClientKey        string `mapstructure:"client_key"`
	EnvType          string `mapstructure:"env_type"`
	DeeplinkDonation string `mapstructure:"deeplink_donation"`
	DeeplinkZakat    string `mapstructure:"deeplink_zakat"`
	DonateBaseURL    string `mapstructure:"donate_base_url"`
}

type Clevertap struct {
	BaseURL   string `mapstructure:"base_url"`
	AccountID string `mapstructure:"account_id"`
	Passcode  string `mapstructure:"passcode"`
}

type Imagine struct {
	BaseURL       string `mapstructure:"base_url"`
	Username      string `mapstructure:"username"`
	Password      string `mapstructure:"password"`
	ImageEndpoint string `mapstructure:"image_endpoint"`
}

// Imgix defines imgix struct config
type Imgix struct {
	ImageEndpoint string `mapstructure:"image_endpoint"`
}

// Santet defines santet configurations
type Santet struct {
	BaseURL           string `mapstructure:"base_url"`
	BasicAuthToggle   bool   `mapstructure:"basic_auth_toggle"`
	BasicAuthUsername string `mapstructure:"basic_auth_username"`
	BasicAuthPassword string `mapstructure:"basic_auth_password"`
}

type Urunan struct {
	BaseURL    string `mapstructure:"base_url"`
	Toggle     bool   `mapstructure:"toggle"`
	AppID      string `mapstructure:"app_id"`
	Key        string `mapstructure:"key"`
	CampaignID string `mapstructure:"campaign_id"`
	TrxType    string `mapstructure:"trx_type"`
}

type Config struct {
	App             App
	Database        Database
	Cache           Cache
	Jwt             Jwt
	HmacSignature   HmacSignature `mapstructure:"hmac_signature"`
	KtbsHeader      KtbsHeader    `mapstructure:"ktbs_header"`
	Email           Email
	Firebase        Firebase
	Users           Users
	Ses             Ses
	Otp             Otp
	OtpDisbursement OtpDisbursement
	Zakat           Zakat
	CustomerIO      CustomerIO
	Sns             Sns
	Donations       Donations
	Link            Link
	Sangu           Sangu
	Cloudsearch     Cloudsearch
	Midtrans        Midtrans
	Clevertap       Clevertap
	Imagine         Imagine
	Imgix           Imgix
	Santet          Santet
	Urunan          Urunan
}

func NewAppConfig() (conf *Config, err error) {
	err = viper.Unmarshal(&conf)
	if err != nil {
		return
	}

	err = common.Validate.Struct(conf)
	if err != nil {
		return
	}

	return
}

func (c *Config) GetDBOption() driver.DBOption {
	return driver.DBOption{
		Host:                 c.Database.Host,
		Port:                 c.Database.Port,
		Username:             c.Database.Username,
		Password:             c.Database.Password,
		DBName:               c.Database.Name,
		AdditionalParameters: c.Database.AdditionalParameters,
		ConnectionSetting: driver.ConnectionSetting{
			MaxOpenConns:    c.Database.MaxOpenConnection,
			MaxIdleConns:    c.Database.MaxIdleConnection,
			ConnMaxLifetime: c.Database.ConnectionMaxLifetime,
		},
	}
}

func (c *Config) GetCacheOption() driver.CacheOption {
	return driver.CacheOption{
		Host:               c.Cache.Host,
		Port:               c.Cache.Port,
		DialConnectTimeout: c.Cache.DialConnectTimeout,
		WriteTimeout:       c.Cache.WriteTimeout,
		ReadTimeout:        c.Cache.ReadTimeout,
		Namespace:          c.Cache.Namespace,
		Password:           c.Cache.Password,
		MaxIdle:            c.Cache.MaxIdle,
		MaxActive:          c.Cache.MaxActive,
		IdleTimeout:        c.Cache.IdleTimeout,
		Wait:               c.Cache.Wait,
		MaxConnLifetime:    c.Cache.MaxConnLifetime,
	}
}

func (c *Config) GetJWTConfig() JWT {
	return JWT{
		Issuer:    c.Jwt.Issuer,
		SecretKey: c.Jwt.SecretKey,
		SignKey:   c.Jwt.SignKey,
	}
}

func (c *Config) GetFirebaseCredentialPath() string {
	return c.Firebase.CredentialPath
}

func (c *Config) GetCSDOption() driver.CloudSearchOption {
	return driver.CloudSearchOption{
		Key:      c.Cloudsearch.Key,
		Secret:   c.Cloudsearch.Secret,
		Endpoint: c.Cloudsearch.SearchEndpoint,
		Region:   c.Cloudsearch.Region,
	}
}
