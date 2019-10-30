package common

import "github.com/kitabisa/pawon-warga/model"

func GetVerificationInfo(user model.User) (verificationStatus bool, verificationBadge string) {
	if user.VerificationStatus == "VERIFIED" {
		verificationStatus = true

		switch user.OrganizationStatus {
		case "PERSONAL", "PENDING":
			verificationBadge = "https://assets.kitabisa.com/images/icon__verified-user.svg"
		case "ORGANIZATION":
			verificationBadge = "https://assets.kitabisa.com/images/icon__verified-org.svg"
		case "CORPORATE":
			verificationBadge = "https://assets.kitabisa.com/images/icon__verified-corp.svg"
		}
	}

	return verificationStatus, verificationBadge
}
