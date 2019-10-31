package handler

import (
	"github.com/jinzhu/gorm"

	"github.com/airlanggatirta/pawon-warga/common"
)

var errorMapEnglish map[error]string
var errorMapBahasa map[error]string

func init() {
	errorMapEnglish = make(map[error]string)
	errorMapBahasa = make(map[error]string)

	errorMapEnglish[gorm.ErrRecordNotFound] = "Data not found"
	errorMapBahasa[gorm.ErrRecordNotFound] = "Data tidak ditemukan"

	errorMapEnglish[common.ErrWrongUsernamePassword] = "Wrong username or password"
	errorMapBahasa[common.ErrWrongUsernamePassword] = "Username atau password salah"

	errorMapEnglish[common.ErrInvalidHeader] = "Invalid header"
	errorMapBahasa[common.ErrInvalidHeader] = "Header tidak valid"

	errorMapEnglish[common.ErrInvalidSignature] = "Invalid signature header"
	errorMapBahasa[common.ErrInvalidSignature] = "Signature header tidak valid"

	errorMapEnglish[common.ErrInvalidRequestEntity] = "Invalid request entity"
	errorMapBahasa[common.ErrInvalidRequestEntity] = "Entitas permintaan tidak valid"

	errorMapEnglish[common.ErrMissingRequestEntity] = "Missing request entity"
	errorMapBahasa[common.ErrMissingRequestEntity] = "Entitas permintaan tidak ditemukan"

	errorMapEnglish[common.ErrSocialMediaInvalidToken] = "Invalid token"
	errorMapBahasa[common.ErrSocialMediaInvalidToken] = "Token tidak valid"

	errorMapEnglish[common.ErrSocialMediaInvalidUser] = "Invalid user"
	errorMapBahasa[common.ErrSocialMediaInvalidUser] = "Pengguna tidak valid"

	errorMapEnglish[common.ErrSocialMediaFailedToRegister] = "Failed to register new user"
	errorMapBahasa[common.ErrSocialMediaFailedToRegister] = "Gagal dalam mendaftarkan pengguna baru"

	errorMapEnglish[common.ErrSocialMediaFailedToGenerateToken] = "Failed to generate new token"
	errorMapBahasa[common.ErrSocialMediaFailedToGenerateToken] = "Gagal dalam membuat token baru"

	errorMapEnglish[common.ErrSocialMediaFailedToUpdateToken] = "Failed to update new token"
	errorMapBahasa[common.ErrSocialMediaFailedToUpdateToken] = "Gagal dalam menyimpan token baru"

	errorMapEnglish[common.ErrInvalidZakatInput] = "All inputs must be integers"
	errorMapBahasa[common.ErrInvalidZakatInput] = "Semua input harus berupa integer"

	errorMapEnglish[common.ErrMonthlyIncomeCannotBeEmpty] = "Monthly income cannot be empty"
	errorMapBahasa[common.ErrMonthlyIncomeCannotBeEmpty] = "Pendapatan bulanan tidak boleh kosong"

	errorMapEnglish[common.ErrAllInputsCannotBeEmpty] = "All inputs cannot be empty"
	errorMapBahasa[common.ErrAllInputsCannotBeEmpty] = "Semua input tidak boleh kosong"

	errorMapEnglish[common.ErrTopupConfirmationAmountDidNotMatch] = "Top up confirmation amount did not match"
	errorMapBahasa[common.ErrTopupConfirmationAmountDidNotMatch] = "Nominal konfirmasi top up tidak sesuai"

	errorMapEnglish[common.ErrWalletTopupStillPending] = "Unable do top up while previous transaction still pending"
	errorMapBahasa[common.ErrWalletTopupStillPending] = "Top up tidak dapat dilakukan ketika transaksi sebelumnya masih pending"

	errorMapEnglish[common.ErrWallletTopupAmountInsufficent] = "Topup nominal is less than 10000"
	errorMapBahasa[common.ErrWallletTopupAmountInsufficent] = "Nominal topup kurang dari 10000"

	errorMapEnglish[common.ErrWallletTopupAmountNotInThousand] = "Topup amount should be multiplication of a thousand"
	errorMapBahasa[common.ErrWallletTopupAmountNotInThousand] = "Jumlah topup harus kelipatan dari seribu"

	errorMapEnglish[common.ErrTopupPaymentMethodDisallowed] = "Unable to top up using the selected payment method"
	errorMapBahasa[common.ErrTopupPaymentMethodDisallowed] = "Top up tidak dapat dilakukan dengan metode pembayaran yang dipilih"

	errorMapEnglish[common.ErrTopupConfirmationInvalidInput] = "Top up confirmation amount must be integer"
	errorMapBahasa[common.ErrTopupConfirmationInvalidInput] = "Nominal konfirmasi top up harus berupa integer"

	errorMapEnglish[common.ErrTopupConfirmationAmountCannotBeEmpty] = "Top up confirmation amount cannot be empty"
	errorMapBahasa[common.ErrTopupConfirmationAmountCannotBeEmpty] = "Nominal konfirmasi top up tidak boleh kosong"

	errorMapEnglish[common.ErrTopupOwnership] = "Cannot see topup data from another user"
	errorMapBahasa[common.ErrTopupOwnership] = "Tidak dapat melihat data topup pengguna lain"

	errorMapEnglish[common.ErrTopupDetailIDInvalid] = "Top up ID must be an integer"
	errorMapBahasa[common.ErrTopupDetailIDInvalid] = "Top up ID harus berupa angka"

	errorMapEnglish[common.ErrDonationEmptyUserID] = "Empty user ID"
	errorMapBahasa[common.ErrDonationEmptyUserID] = "User ID tidak boleh kosong"

	errorMapEnglish[common.ErrDonationTooLarge] = "Your donation too large"
	errorMapBahasa[common.ErrDonationEmptyUserID] = "Donasi Anda terlalu besar"

	errorMapEnglish[common.ErrDonationEmptyCampaignID] = "Empty campaign ID"
	errorMapBahasa[common.ErrDonationEmptyCampaignID] = "Campaign ID tidak boleh kosong"

	errorMapEnglish[common.ErrDonationEmptyAmount] = "Empty donation amount"
	errorMapBahasa[common.ErrDonationEmptyAmount] = "Nominal donasi tidak boleh kosong"

	errorMapEnglish[common.ErrDonationAmountMinumum10K] = "Minimum donation amount is 10K"
	errorMapBahasa[common.ErrDonationAmountMinumum10K] = "Nominal donasi minimal Rp10.000"

	errorMapEnglish[common.ErrDonationAmountMinumum1K] = "Minimum donation amount is 1K"
	errorMapBahasa[common.ErrDonationAmountMinumum1K] = "Nominal donasi minimal Rp1.000"

	errorMapEnglish[common.ErrDonationAmountMinimum] = "Donation amount less than minimum amount"
	errorMapBahasa[common.ErrDonationAmountMinimum] = "Nominal donasi tidak sesuai dengan minimum donasi"

	errorMapEnglish[common.ErrDonationEmptyPaymentMethodID] = "Empty payment method ID"
	errorMapBahasa[common.ErrDonationEmptyPaymentMethodID] = "Metode pembayaran tidak boleh kosong"

	errorMapEnglish[common.ErrDuplicateDonation] = "Cannot create donation with same amount within one minute"
	errorMapBahasa[common.ErrDuplicateDonation] = "Tidak bisa membuat donasi dengan nominal yang sama dalam waktu satu menit"

	errorMapEnglish[common.ErrInsufficientWalletBalance] = "Insufficient wallet balance"
	errorMapBahasa[common.ErrInsufficientWalletBalance] = "Saldo dompet tidak cukup"

	errorMapEnglish[common.ErrDonationMustBeMultiplesOfThousands] = "Donation amount must be multiples of thousands"
	errorMapBahasa[common.ErrDonationMustBeMultiplesOfThousands] = "Nominal donasi harus dalam kelipatan ribuan"

	errorMapEnglish[common.ErrInactivePaymentMethod] = "Inactive payment method. Please choose another payment method"
	errorMapBahasa[common.ErrInactivePaymentMethod] = "Metode pembayaran sedang tidak aktif. Silakan pilih metode pembayaran lain"

	errorMapEnglish[common.ErrDonatedCampaignIsNotOpenForDonation] = "Campaign is no longer open for donation"
	errorMapBahasa[common.ErrDonatedCampaignIsNotOpenForDonation] = "Campaign sudah tidak menerima donasi"

	errorMapEnglish[common.ErrUserAlreadyExists] = "User already exists"
	errorMapBahasa[common.ErrUserAlreadyExists] = "Email/Nomor WhatsApp sudah terdaftar"

	errorMapEnglish[common.ErrUserAlreadyReportCampaign] = "You already reported this campaign"
	errorMapBahasa[common.ErrUserAlreadyReportCampaign] = "Anda sudah pernah melaporkan campaign ini"

	errorMapEnglish[common.ErrDuplicateLovelist] = "Campaign telah ada di lovelist"
	errorMapBahasa[common.ErrDuplicateLovelist] = "Campaign has been added on your lovelist"

	errorMapEnglish[common.ErrInvalidTimeDonationReminder] = "Time format invalid"
	errorMapBahasa[common.ErrInvalidTimeDonationReminder] = "Format jam tidak valid"

	errorMapEnglish[common.ErrLoveListPermissionDenied] = "User ID tidak valid untuk melakukan process tersebut"
	errorMapBahasa[common.ErrLoveListPermissionDenied] = "Invalid user ID for this operation"

	errorMapEnglish[common.ErrEmptyFrequencyDonationReminder] = "Empty frequency"
	errorMapBahasa[common.ErrEmptyFrequencyDonationReminder] = "Frekuensi tidak boleh kosong"

	errorMapEnglish[common.ErrEmptyTimeDonationReminder] = "Empty time"
	errorMapBahasa[common.ErrEmptyTimeDonationReminder] = "Waktu tidak boleh kosong"

	errorMapEnglish[common.ErrEmptyTimeLocationDonationReminder] = "Empty time location"
	errorMapBahasa[common.ErrEmptyTimeLocationDonationReminder] = "Lokasi waktu tidak boleh kosong"

	errorMapEnglish[common.ErrTranslateDayDonationReminder] = "No index for translation day"
	errorMapBahasa[common.ErrTranslateDayDonationReminder] = "Tidak ada indeks untuk terjemahan hari"

	errorMapEnglish[common.ErrInvalidFrequencyDonationReminder] = "Invalid frequency, frequency value must [monthly, weekly, daily]"
	errorMapBahasa[common.ErrInvalidFrequencyDonationReminder] = "Frekuensi tidak valid, nilai frekuensi harus [monthly, weekly, daily]"

	errorMapEnglish[common.ErrInvalidDayFrequencyMonthlyDonationReminder] = "Invalid day for frequency monthly, day value must 1-31"
	errorMapBahasa[common.ErrInvalidDayFrequencyMonthlyDonationReminder] = "Hari tidak valid untuk frekuensi bulanan, nilai hari harus 1-31"

	errorMapEnglish[common.ErrInvalidDayFrequencyWeeklyDonationReminder] = "Invalid day for frequency weekly, day value must type string and value senin-minggu"
	errorMapBahasa[common.ErrInvalidDayFrequencyWeeklyDonationReminder] = "Hari tidak valid untuk frekuensi mingguan, nilai hari harus mengetik string dan nilai senin-minggu"

	errorMapEnglish[common.ErrInvalidDayFrequencyDailyDonationReminder] = "Invalid day for frequency daily, day value must empty"
	errorMapBahasa[common.ErrInvalidDayFrequencyDailyDonationReminder] = "Hari tidak valid untuk frekuensi harian, nilai hari harus kosong"

	errorMapEnglish[common.ErrInvalidTimeLocationDonationReminder] = "Invalid time location"
	errorMapBahasa[common.ErrInvalidTimeLocationDonationReminder] = "Lokasi waktu tidak valid"

	errorMapEnglish[common.ErrEmptyMonthTotalDonation] = "Selected month cannot be empty"
	errorMapBahasa[common.ErrEmptyMonthTotalDonation] = "Bulan tidak boleh kosong"

	errorMapEnglish[common.ErrEmptyYearTotalDonation] = "Selected year cannot be empty"
	errorMapBahasa[common.ErrEmptyYearTotalDonation] = "Tahun tidak boleh kosong"

	errorMapEnglish[common.ErrFailedToCreateJeniusTrx] = "Failed to create Jenius transaction"
	errorMapBahasa[common.ErrFailedToCreateJeniusTrx] = "Gagal membuat transaksi Jenius"

	errorMapEnglish[common.ErrAccountRegisteredBySocialMedia] = "Account registered by social media, please login with social media"
	errorMapBahasa[common.ErrAccountRegisteredBySocialMedia] = "Akun terdaftar dengan media sosial, silakan masuk dengan media sosial"

	errorMapEnglish[common.ErrSessionExpired] = "Session has expired"
	errorMapBahasa[common.ErrSessionExpired] = "Sesi telah kedaluwarsa"

	errorMapEnglish[common.ErrInvalidFullName] = "Invalid fullname"
	errorMapBahasa[common.ErrInvalidFullName] = "Nama lengkap tidak valid"

	errorMapEnglish[common.ErrInsufficientCharacterOfReport] = "Report must be 50 character or more"
	errorMapBahasa[common.ErrInsufficientCharacterOfReport] = "Report harus memiliki 50 karakter atau lebih"

	errorMapEnglish[common.ErrEmptyRedirectCallbackDonation] = "Empty redirect callback"
	errorMapBahasa[common.ErrEmptyRedirectCallbackDonation] = "Redirect callback tidak boleh kosong"

	errorMapEnglish[common.ErrWhatsappNumberNotRegisteredInKitabisa] = common.ErrWhatsappNumberNotRegisteredInKitabisa.Error()
	errorMapBahasa[common.ErrWhatsappNumberNotRegisteredInKitabisa] = "Nomor Whatsapp ini tidak terdaftar di Kitabisa"

	errorMapEnglish[common.ErrManualDonationConfirmationEmptyDonationID] = common.ErrManualDonationConfirmationEmptyDonationID.Error()
	errorMapBahasa[common.ErrManualDonationConfirmationEmptyDonationID] = "Donasi ID tidak boleh kosng"

	errorMapEnglish[common.ErrManualDonationConfirmationEmptyConfirmationCode] = common.ErrManualDonationConfirmationEmptyConfirmationCode.Error()
	errorMapBahasa[common.ErrManualDonationConfirmationEmptyConfirmationCode] = "Kode konfirmasi tidak boleh kosong"

	errorMapEnglish[common.ErrManualDonationConfirmationEmptyBankName] = common.ErrManualDonationConfirmationEmptyBankName.Error()
	errorMapBahasa[common.ErrManualDonationConfirmationEmptyBankName] = "Nama bank tidak boleh kosong"

	errorMapEnglish[common.ErrManualDonationConfirmationEmptyBankAccHolder] = common.ErrManualDonationConfirmationEmptyBankAccHolder.Error()
	errorMapBahasa[common.ErrManualDonationConfirmationEmptyBankAccHolder] = "Pemilik akun bank tidak boleh kosong"

	errorMapEnglish[common.ErrManualDonationConfirmationEmptyAmount] = common.ErrManualDonationConfirmationEmptyAmount.Error()
	errorMapBahasa[common.ErrManualDonationConfirmationEmptyAmount] = "Nominal tidak boleh kosng"

	errorMapEnglish[common.ErrManualDonationConfirmationEmptyTransferAt] = common.ErrManualDonationConfirmationEmptyTransferAt.Error()
	errorMapBahasa[common.ErrManualDonationConfirmationEmptyTransferAt] = "Waktu transfer tidak boleh kosong"

	errorMapEnglish[common.ErrGetDonationMultipleNotParent] = "Donation is not multiple donation"
	errorMapBahasa[common.ErrGetDonationMultipleNotParent] = "Donasi bukan merupakan multiple donasi"

	errorMapEnglish[common.ErrCreateDonationMultiplePaymentMethodNotSupported] = common.ErrCreateDonationMultiplePaymentMethodNotSupported.Error()
	errorMapBahasa[common.ErrCreateDonationMultiplePaymentMethodNotSupported] = "Metode pembayaran tidak mendukung"

	errorMapEnglish[common.ErrManualDonationConfirmationInvalidCode] = common.ErrManualDonationConfirmationInvalidCode.Error()
	errorMapBahasa[common.ErrManualDonationConfirmationInvalidCode] = "Kode konfirmasi donasi manual Anda salah"

	errorMapEnglish[common.ErrManualDonationConfirmationAlreadyConfirmed] = common.ErrManualDonationConfirmationAlreadyConfirmed.Error()
	errorMapBahasa[common.ErrManualDonationConfirmationAlreadyConfirmed] = "Kamu sudah mengirim permintaan verifikasi"

	errorMapEnglish[common.ErrWrongOTPCode] = common.ErrWrongOTPCode.Error()
	errorMapBahasa[common.ErrWrongOTPCode] = "Kode verifikasi Anda salah"

	errorMapEnglish[common.ErrOTPTooManyAttempts] = common.ErrOTPTooManyAttempts.Error()
	errorMapBahasa[common.ErrOTPTooManyAttempts] = "Total submit OTP Anda telah melebihi batas 3x"

	errorMapEnglish[common.ErrOTPExpired] = common.ErrOTPExpired.Error()
	errorMapBahasa[common.ErrOTPExpired] = "OTP telah kadaluarsa"

	errorMapEnglish[common.ErrOTPAlreadyVerified] = common.ErrOTPAlreadyVerified.Error()
	errorMapBahasa[common.ErrOTPAlreadyVerified] = "OTP sebelumnya telah diverifikasi"

	errorMapEnglish[common.ErrOTPTooManyRequest] = common.ErrOTPTooManyRequest.Error()
	errorMapBahasa[common.ErrOTPTooManyRequest] = "Total request kode OTP telah melebihi batas"

	errorMapEnglish[common.ErrOTPHasBeenSent] = common.ErrOTPHasBeenSent.Error()
	errorMapBahasa[common.ErrOTPHasBeenSent] = "Permintaan OTP telah terkirim, silakan coba beberapa saat lagi"

	errorMapEnglish[common.ErrCreateDonationNonLoginEmptyName] = common.ErrCreateDonationNonLoginEmptyName.Error()
	errorMapBahasa[common.ErrCreateDonationNonLoginEmptyName] = "Nama tidak boleh kosong"

	errorMapEnglish[common.ErrCreateDonationNonLoginEmptyEmailWaNumber] = common.ErrCreateDonationNonLoginEmptyEmailWaNumber.Error()
	errorMapBahasa[common.ErrCreateDonationNonLoginEmptyEmailWaNumber] = "Email atau nomor WhatsApp tidak boleh kosong"

	errorMapEnglish[common.ErrCreateDonationNonLoginInvalidEmailWaNumber] = common.ErrCreateDonationNonLoginInvalidEmailWaNumber.Error()
	errorMapBahasa[common.ErrCreateDonationNonLoginInvalidEmailWaNumber] = "Email atau nomor WhatsApp tidak valid"

	errorMapEnglish[common.ErrWhatsappNumberNotRegistered] = common.ErrWhatsappNumberNotRegistered.Error()
	errorMapBahasa[common.ErrWhatsappNumberNotRegistered] = "Nomor WhatsApp tidak terdaftar"

	errorMapEnglish[common.ErrGetCampaingGDVEmptyParams] = "Empty parameter"
	errorMapBahasa[common.ErrGetCampaingGDVEmptyParams] = "Parameter tidak boleh kosong"

	errorMapEnglish[common.ErrGetCampaignGDVInvalidParamsValue] = "Invalid parameter value"
	errorMapBahasa[common.ErrGetCampaignGDVInvalidParamsValue] = "Value parameter tidak valid"

	errorMapEnglish[common.ErrInvalidDisbursementID] = common.ErrInvalidDisbursementID.Error()
	errorMapBahasa[common.ErrInvalidDisbursementID] = "disbursement_id Anda invalid"

	errorMapEnglish[common.ErrCannotCancelDisbursement] = common.ErrCannotCancelDisbursement.Error()
	errorMapBahasa[common.ErrCannotCancelDisbursement] = "gagal melakukan cancel disbursement. status disbursement harus pending."
	errorMapEnglish[common.ErrRequestDisburseEmptyBankAccountNumber] = "Empty bank account number"
	errorMapBahasa[common.ErrRequestDisburseEmptyBankAccountNumber] = "Nomor akun bank tidak boleh kosong"

	errorMapEnglish[common.ErrRequestDisburseEmptyBankCode] = "Empty bank code"
	errorMapBahasa[common.ErrRequestDisburseEmptyBankCode] = "Kode bank tidak boleh kosong"

	errorMapEnglish[common.ErrRequestDisburseEmptyAmount] = "Empty amount"
	errorMapBahasa[common.ErrRequestDisburseEmptyAmount] = "Nominal tidak boleh kosong"

	errorMapEnglish[common.ErrDisburseWalletAmountNotAllowed] = "Disburse amount not in range"
	errorMapBahasa[common.ErrDisburseWalletAmountNotAllowed] = "Jumlah pencairan tidak dalam kisaran"

	errorMapEnglish[common.ErrWalletDisburseStillPending] = "Unable do disburse while previous transaction still pending"
	errorMapBahasa[common.ErrWalletDisburseStillPending] = "Pencairan dana tidak dapat dilakukan ketika transaksi sebelumnya masih pending"

	errorMapEnglish[common.ErrWalletDisburseInvalidActivationCode] = "Invalid activation code"
	errorMapBahasa[common.ErrWalletDisburseInvalidActivationCode] = "Kode aktivasi tidak valid"
}
