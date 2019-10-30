package common

// APICode defines api code response.
// APICode definition can be found here https://app.gitbook.com/@kitabisa-engineering/s/backend/standardization-1/api-status-code/error-list-core-api-10
var APICode = map[error]int{
	// general
	ErrInvalidRequestEntity: 101001,
	ErrMissingRequestEntity: 101002,
	ErrInvalidPlatformValue: 101004,
	ErrInvalidHeader:        101005,
	ErrInvalidSignature:     101006,

	// token & authentication
	ErrInvalidToken:          101100,
	ErrSessionExpired:        101101,
	ErrUnsufficientPrivilege: 101102,
	ErrAccessDenied:          101103,

	// user
	ErrWrongUsernamePassword:          101200,
	ErrInvalidUserID:                  101201,
	ErrInvalidActivationKey:           101202,
	ErrInvalidEmail:                   101203,
	ErrUserAlreadyExists:              101204,
	ErrEmptyPassword:                  101205,
	ErrPasswordDoesNotMeetCriteria:    101206,
	ErrInvalidFullName:                101207,
	ErrWhatsappNumberNotRegistered:    101208,
	ErrAccountRegisteredBySocialMedia: 101209,

	// user social media
	ErrSocialMediaInvalidToken:          101300,
	ErrSocialMediaInvalidUser:           101301,
	ErrSocialMediaFailedToRegister:      101302,
	ErrSocialMediaFailedToGenerateToken: 101303,
	ErrSocialMediaFailedToUpdateToken:   101304,

	// reset password
	ErrResetPasswordInvalidResetPasswordKey: 101400,
	ErrResetPasswordExpiredResetPasswordKey: 101401,
	ErrResetPasswordEmptyUserID:             101402,
	ErrResetPasswordEmptyResetPasswordKey:   101403,
	ErrResetPasswordEmptyPassword:           101404,

	// OTP
	ErrWrongOTPCode:                          101500,
	ErrOTPTooManyAttempts:                    101501,
	ErrInvalidWhatsappNumber:                 101502,
	ErrWhatsappNumberNotRegisteredInKitabisa: 101503,
	ErrOTPExpired:                            101504,
	ErrOTPAlreadyVerified:                    101505,
	ErrOTPTooManyRequest:                     101506,

	// pending registration
	ErrPendingRegistration: 101600,
	ErrEmptyUserID:         101601,
	ErrInvalidUserIDType:   101602,

	// campaign
	ErrInvalidCampaignID:             101700,
	ErrUserAlreadyReportCampaign:     101701,
	ErrCampaignMedicalDocsNotFound:   101702,
	ErrInsufficientCharacterOfReport: 101703,

	// lovelist
	ErrLoveListPermissionDenied: 101800,
	ErrDuplicateLovelist:        101801,

	// zakat
	ErrInvalidZakatInput:                    101900,
	ErrMonthlyIncomeCannotBeEmpty:           101901,
	ErrTopupConfirmationAmountCannotBeEmpty: 101902,
	ErrAllInputsCannotBeEmpty:               101903,

	// topup
	ErrTopupConfirmationAmountDidNotMatch: 102000,
	ErrTopupConfirmationInvalidInput:      102001,
	ErrTopupDetailIDInvalid:               102002,
	ErrTopupOwnership:                     102003,

	// donation
	ErrDonationEmptyUserID:                 102100,
	ErrDonationEmptyCampaignID:             102101,
	ErrDonationEmptyAmount:                 102102,
	ErrDonationAmountMinumum10K:            102103,
	ErrDonationAmountMinumum1K:             102104,
	ErrDonationEmptyPaymentMethodID:        102105,
	ErrDonationTooLarge:                    102106,
	ErrDuplicateDonation:                   102107,
	ErrDonationMustBeMultiplesOfThousands:  102108,
	ErrInactivePaymentMethod:               102109,
	ErrDonatedCampaignIsNotOpenForDonation: 102110,
	ErrEmptyRedirectCallbackDonation:       102111,
	ErrDonationStreakEmptyUnixDateTime:     102112,
	ErrEmptyMonthTotalDonation:             102113,
	ErrEmptyYearTotalDonation:              102114,
	ErrDonationAmountMinimum:               102115,
	ErrEmptySuccessUrlDonation:             102116,
	ErrEmptyFailUrlDonation:                102117,
	ErrEmptyCashtagDonation:                102118,
	ErrEmptyPayReturnUrlDonation:           102119,

	// multiple donation
	ErrCreateDonationMultipleEmptyPayload:              102200,
	ErrCreateDonationMultiplePaymentMethodNotSupported: 102201,
	ErrGetDonationMultipleNotParent:                    102202,
	ErrMultipleDonationDataCannotEmpty:                 102203,

	// non-login donation
	ErrCreateDonationNonLoginEmptyName:                 102300,
	ErrCreateDonationNonLoginEmptyEmailWaNumber:        102301,
	ErrCreateDonationNonLoginInvalidEmailWaNumber:      102302,
	ErrCreateDonationNonLoginNotSupportedPaymentMethod: 102303,

	// wallet
	ErrWalletTopupStillPending:               102400,
	ErrWallletTopupAmountInsufficent:         102401,
	ErrWallletTopupAmountNotInThousand:       102402,
	ErrInsufficientWalletBalance:             102403,
	ErrDisburseWalletAmountNotAllowed:        102404,
	ErrInvalidBankCode:                       102405,
	ErrInvalidDisbursementID:                 102406,
	ErrCannotCancelDisbursement:              102407,
	ErrWalletDisburseStillPending:            102408,
	ErrRequestDisburseEmptyBankAccountNumber: 102409,
	ErrRequestDisburseEmptyBankCode:          102410,
	ErrRequestDisburseEmptyAmount:            102411,
	ErrWalletDisburseInvalidActivationCode:   102412,

	// donation reminder
	ErrInvalidTimeDonationReminder:                102500,
	ErrTranslateDayDonationReminder:               102501,
	ErrEmptyFrequencyDonationReminder:             102502,
	ErrEmptyTimeDonationReminder:                  102503,
	ErrEmptyTimeLocationDonationReminder:          102504,
	ErrInvalidFrequencyDonationReminder:           102505,
	ErrInvalidDayFrequencyMonthlyDonationReminder: 102506,
	ErrInvalidDayFrequencyWeeklyDonationReminder:  102507,
	ErrInvalidDayFrequencyDailyDonationReminder:   102508,
	ErrInvalidTimeLocationDonationReminder:        102509,

	// payment
	ErrTopupPaymentMethodDisallowed: 102600,
	ErrFailedToCreateJeniusTrx:      102601,

	// manual donation confirmation
	ErrManualDonationConfirmationEmptyDonationID:       102700,
	ErrManualDonationConfirmationEmptyConfirmationCode: 102701,
	ErrManualDonationConfirmationEmptyBankName:         102702,
	ErrManualDonationConfirmationEmptyBankAccHolder:    102703,
	ErrManualDonationConfirmationEmptyAmount:           102704,
	ErrManualDonationConfirmationEmptyTransferAt:       102705,
	ErrManualDonationConfirmationAlreadyConfirmed:      102706,
	ErrManualDonationConfirmationInvalidCode:           102707,

	// sangu
	ErrSanguCallbackEmptyVendorTrxId:     102800,
	ErrSanguCallbackEmptyPaymentMethodId: 102801,

	// campaign GDV
	ErrGetCampaingGDVEmptyParams:        103000,
	ErrGetCampaignGDVInvalidParamsValue: 103001,
}

// GetAPICode will return api code according to response
func GetAPICode(err error) int {
	var code int
	var ok bool

	if err == nil {
		code = 101000
	} else {
		code, ok = APICode[err]
		if !ok {
			code = 101003
		}
	}

	return code
}
