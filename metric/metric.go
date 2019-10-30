package metric

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	totalSuccessLoginCounterVec = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "coreapi",
			Subsystem: "user",
			Name:      "success_login",
			Help:      "Total number of success login",
		},
		[]string{"type"},
	)
	totalOtpSentCounterVec = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "coreapi",
			Subsystem: "otp",
			Name:      "otp_sent",
			Help:      "Total number of OTP",
		},
		[]string{"type"},
	)
	totalJWTInvalidCounterVec = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "coreapi",
			Subsystem: "jwt",
			Name:      "token_invalid",
			Help:      "Total number of Invalid JWT Token",
		},
		[]string{"type"},
	)
	campaignSearchDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "coreapi",
			Subsystem: "campaign",
			Name:      "cloudsearch_duration_seconds",
			Help:      "Campaign cloudserach duration",
			Buckets:   []float64{1, 2, 5, 10, 30},
		},
		[]string{"type"},
	)
	otpServiceDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "coreapi",
			Subsystem: "otp",
			Name:      "service_duration_seconds",
			Help:      "OTP Service duration",
			Buckets:   []float64{1, 2, 5, 10, 30},
		},
		[]string{"type"},
	)
	totalJwtTokenCreatedCounterVec = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "coreapi",
			Subsystem: "user",
			Name:      "jwt_token",
			Help:      "Total number of jwt token created",
		},
		[]string{"login"},
	)
	totalSnsEventCounterVec = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "coreapi",
			Subsystem: "aws",
			Name:      "sns_published",
			Help:      "Total number of event published",
		},
		[]string{"event_name"},
	)
	totalSesEmailCounterVec = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "coreapi",
			Subsystem: "aws",
			Name:      "ses_sent",
			Help:      "Total number of email sent",
		},
		[]string{"type"},
	)
	firebaseServiceDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "coreapi",
			Subsystem: "firebase",
			Name:      "service_duration_seconds",
			Help:      "Firebase Service duration",
			Buckets:   []float64{1, 2, 5, 10, 30},
		},
		[]string{"type"},
	)
	sanguServiceDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "coreapi",
			Subsystem: "sangu",
			Name:      "service_duration_seconds",
			Help:      "Sangu Service duration",
			Buckets:   []float64{1, 2, 5, 10, 30},
		},
		[]string{"type"},
	)
	midtransServiceDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "coreapi",
			Subsystem: "midtrans",
			Name:      "midtrans_duration_seconds",
			Help:      "Midtrans Service duration",
			Buckets:   []float64{1, 2, 5, 10, 30},
		},
		[]string{"type"},
	)
	imagineServiceDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "coreapi",
			Subsystem: "imagine",
			Name:      "imagine_duration_seconds",
			Help:      "Imagine Service duration",
			Buckets:   []float64{1, 2, 5, 10, 30},
		},
		[]string{"type"},
	)
	snsServiceDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "coreapi",
			Subsystem: "aws",
			Name:      "aws_sns_duration_seconds",
			Help:      "AWS SNS duration",
			Buckets:   []float64{1, 2, 5, 10, 30},
		},
		[]string{"type"},
	)
	sesServiceDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "coreapi",
			Subsystem: "aws",
			Name:      "aws_ses_duration_seconds",
			Help:      "AWS SES duration",
			Buckets:   []float64{1, 2, 5, 10, 30},
		},
		[]string{"type"},
	)
)

func RegisterMetric() {
	prometheus.MustRegister(totalSuccessLoginCounterVec)
	prometheus.MustRegister(totalJwtTokenCreatedCounterVec)
	prometheus.MustRegister(campaignSearchDuration)
	prometheus.MustRegister(otpServiceDuration)
	prometheus.MustRegister(totalOtpSentCounterVec)
	prometheus.MustRegister(firebaseServiceDuration)
	prometheus.MustRegister(sanguServiceDuration)
	prometheus.MustRegister(midtransServiceDuration)
	prometheus.MustRegister(imagineServiceDuration)
	prometheus.MustRegister(snsServiceDuration)
	prometheus.MustRegister(sesServiceDuration)
	prometheus.MustRegister(totalSnsEventCounterVec)
	prometheus.MustRegister(totalSesEmailCounterVec)
	prometheus.MustRegister(totalJWTInvalidCounterVec)
}

func CountSuccessLogin(loginType string) {
	totalSuccessLoginCounterVec.WithLabelValues(loginType).Inc()
}

func CountJwtTokenCreated() {
	totalJwtTokenCreatedCounterVec.WithLabelValues("").Inc()
}

func CountJwtTokenInvalid(errType string) {
	totalJWTInvalidCounterVec.WithLabelValues(errType).Inc()
}

func ObserveCampaignSearchDuration(searchType string, duration float64) {
	campaignSearchDuration.WithLabelValues(searchType).Observe(duration)
}

func ObserveOtpDuration(otpType string, duration float64) {
	otpServiceDuration.WithLabelValues(otpType).Observe(duration)
}

func CountOtpSent(otpType string) {
	totalOtpSentCounterVec.WithLabelValues(otpType).Inc()
}

func ObserveFirebaseDuration(otpType string, duration float64) {
	firebaseServiceDuration.WithLabelValues(otpType).Observe(duration)
}

func ObserveSanguDuration(otpType string, duration float64) {
	sanguServiceDuration.WithLabelValues(otpType).Observe(duration)
}

func ObserveMidtransDuration(otpType string, duration float64) {
	midtransServiceDuration.WithLabelValues(otpType).Observe(duration)
}

func ObserveImagineDuration(otpType string, duration float64) {
	imagineServiceDuration.WithLabelValues(otpType).Observe(duration)
}

func ObserveSnsDuration(eventType string, duration float64) {
	snsServiceDuration.WithLabelValues(eventType).Observe(duration)
}

func CountSnsPublished(eventName string) {
	totalSnsEventCounterVec.WithLabelValues(eventName).Inc()
}

func ObserveSesDuration(otpType string, duration float64) {
	sesServiceDuration.WithLabelValues(otpType).Observe(duration)
}

func CountSesPublished(emailType string) {
	totalSesEmailCounterVec.WithLabelValues(emailType).Inc()
}
