package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/airlanggatirta/pawon-warga/common"
	"github.com/airlanggatirta/pawon-warga/config"
	"github.com/airlanggatirta/pawon-warga/driver"
	"github.com/airlanggatirta/pawon-warga/handler"
	"github.com/airlanggatirta/pawon-warga/metric"
	"github.com/airlanggatirta/pawon-warga/middleware"
	"github.com/airlanggatirta/pawon-warga/repository"
	"github.com/airlanggatirta/pawon-warga/service"
	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	metrics "github.com/slok/go-http-metrics/metrics/prometheus"
	metricMdlwr "github.com/slok/go-http-metrics/middleware"
	negronimiddleware "github.com/slok/go-http-metrics/middleware/negroni"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/urfave/negroni"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pawon-warga",
	Short: "Admin management system on pawon",
	Long: `The direct user of pawon is the admin of kitabisa website. Previously on Aurum
the admin is also user of kitabisa. This time we make it a bit 
different by separating the administration system between admin 
kitabisa user.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		InitApp()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.github.com/airlanggatirta/pawon-warga.toml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".pawon-warga" (without extension).
		viper.AddConfigPath(".")
		viper.AddConfigPath(home)
		viper.AddConfigPath("./params")
		viper.AddConfigPath("/opt/pawon-warga/bin")
		viper.AddConfigPath("/opt/pawon-warga/bin/params")
		viper.AddConfigPath("/etc/pawon-warga")
		viper.AddConfigPath("/usr/local/etc/pawon-warga")
		viper.AddConfigPath("/etc/pawon-warga")
		viper.SetConfigName("pawon-warga")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func InitApp() {
	config, err := config.NewAppConfig()
	if err != nil {
		log.Fatalf("Config error : %s", err)
	}

	metric.RegisterMetric()

	cache, err := driver.NewCache(config.GetCacheOption())
	if err != nil {
		log.Fatalf("%s : %v", "Cache error", err)
	}

	db, err := driver.NewMysqlDatabase(config.GetDBOption())
	if err != nil {
		log.Fatalf("%s : %v", "DB error", err)
		panic(err)
	}
	defer db.Close()

	stdOutLogger := log.New()
	stdOutLogger.SetOutput(os.Stdout)
	stdOutLogger.SetLevel(log.DebugLevel)

	stdErrLogger := log.New()
	stdErrLogger.SetOutput(os.Stderr)
	stdOutLogger.SetReportCaller(true)
	stdErrLogger.SetLevel(log.DebugLevel)

	logger := common.NewAPILogger(stdOutLogger, stdErrLogger)

	repo := WiringUpRepository(db, cache, logger, config)

	handler.BaseURL = config.Link.BaseURL
	common.ImageURL = config.Imagine.ImageEndpoint
	common.ImgixURL = config.Imgix.ImageEndpoint
	common.SantetBaseURL = config.Santet.BaseURL
	common.SantetBasicAuthToggle = config.Santet.BasicAuthToggle
	common.SantetUsername = config.Santet.BasicAuthUsername
	common.SantetPassword = config.Santet.BasicAuthPassword
	common.WhatsappBaseURL = config.Otp.VerifyNumberURL

	service, err := WiringUpServiceV3(repo, db, cache, logger, config)
	if err != nil {
		panic(err)
	}

	middleware, err := middleware.NewMiddlewareBuilder().
		SetJwtSignKey(config.GetJWTConfig().SignKey).
		SetHmacConfig(config.HmacSignature).
		SetKtbsHeaderConfig(config.KtbsHeader).
		SetTokenService(service.Token).
		SetUserService(service.User).
		SetOTPService(service.OTP).
		Build()
	if err != nil {
		log.Fatalln(err)
	}

	urlHandler := handler.NewHandler(service, config.GetJWTConfig().SignKey, logger)
	urlHandlerV3 := v3Handler.NewHandlerV3(v3Service)
	metricMiddleware := metricMdlwr.New(metricMdlwr.Config{
		Recorder: metrics.NewRecorder(metrics.Config{}),
	})

	r := mux.NewRouter()

	r.Use(middleware.CommonHeaderMiddleware)
	r.Handle("/login", handler.HttpHandler{logger, urlHandler.Login}).Methods(http.MethodPost)
	r.Handle("/user/registration", handler.HttpHandler{logger, urlHandler.Register}).Methods(http.MethodPost)
	r.Handle("/user/registration/temp", handler.HttpHandler{logger, urlHandler.RegisterTemp}).Methods(http.MethodPost)
	r.Handle("/user/registration/temp/{username}", handler.HttpHandler{logger, urlHandler.GetRegisterTemp}).Methods(http.MethodGet)
	r.Handle("/activation/{user_id}/{activation_key}", handler.HttpHandler{logger, urlHandler.ActivateUser}).Methods(http.MethodGet)
	r.Handle("/user/registration/otp/{number}", handler.HttpHandler{logger, urlHandler.SendOTP}).Methods(http.MethodGet)
	r.Handle("/user/registration/otp", handler.HttpHandler{logger, urlHandler.ValidateOTP}).Methods(http.MethodPost)
	r.Handle("/user/reset-password/otp", handler.HttpHandler{logger, urlHandler.ValidateOTPUpdatePassword}).Methods(http.MethodPost)
	r.Handle("/user/registration/otp/check", handler.HttpHandler{logger, urlHandler.ValidateWhatsappNumber}).Methods(http.MethodGet)
	r.Handle("/user/update/password", handler.HttpHandler{logger, urlHandler.UpdatePassword}).Methods(http.MethodPost)
	r.Handle("/user/reset_password/{user_email}", handler.HttpHandler{logger, urlHandler.ResetPasswordRequest}).Methods(http.MethodGet)
	r.Handle("/user/reset_password/email", handler.HttpHandler{logger, urlHandler.ResetPasswordByEmail}).Methods(http.MethodPost)
	r.Handle("/token/validate", handler.HttpHandler{logger, urlHandler.ValidateTokenHandler}).Methods(http.MethodPost)

	nonLoginDonationNegroni := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
		negroni.HandlerFunc(middleware.NonLoginMiddleware),
	)

	nonLoginDonationRouter := mux.NewRouter()
	nonLoginDonationRouter.Handle("/donations/create_multiple", handler.HttpHandler{logger, urlHandler.CreateDonationMultiple}).Methods(http.MethodPost)

	r.Path("/donations/create_multiple").Handler(nonLoginDonationNegroni.With(
		negroni.Wrap(nonLoginDonationRouter),
	))

	withAuthRouter := mux.NewRouter()

	commonNegroni := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
		negroni.HandlerFunc(middleware.JWTMiddleware),
	)

	cors := cors.New(cors.Options{
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Authorization", "Content-Type", "Version"},
	})

	r.Handle("/metrics", promhttp.Handler()).Methods(http.MethodGet)

	KtbsHeaderNegroni := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
		negroni.HandlerFunc(middleware.KtbsHeader),
		negroni.HandlerFunc(middleware.HmacSignature),
	)

	// V3 endpoints
	v3Route := r.PathPrefix("/v3").Subrouter()
	v3Route.Handle("/user/reset_password/otp/{number}", handler.HttpHandler{logger, urlHandlerV3.ResetPasswordRequestOTP}).Methods(http.MethodGet)
	v3Route.Handle("/user/registration/otp/{number}", handler.HttpHandler{logger, urlHandlerV3.SendOTP}).Methods(http.MethodGet)
	v3Route.Handle("/user/registration/temp", handler.HttpHandler{logger, urlHandlerV3.RegisterTemp}).Methods(http.MethodPost)

	n := negroni.Classic()

	n.Use(cors)
	n.UseHandler(r)
	n.Use(negronimiddleware.Handler("/metrics", metricMiddleware))

	var srv http.Server
	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		stdOutLogger.Printf("Server shutdown.. \n")
		// We received an interrupt signal, shut down.
		if err := srv.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			stdErrLogger.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	srv.Addr = fmt.Sprintf("%s:%d", config.App.Host, config.App.Port)
	srv.Handler = n
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		stdErrLogger.Printf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
	stdOutLogger.Printf("Bye. \n")
}

// WiringUpRepositoryV3 will bootstrapping repository V3
func WiringUpRepository(db *gorm.DB, cache *redis.Pool, logger *common.APILogger, appConfig *config.Config) *v3Repository.RepositoryV3 {
	v3RepoOptions := v3Repository.RepositoryV3Option{
		DB:        db,
		Cache:     cache,
		Logger:    logger,
		AppConfig: appConfig,
	}

	otpRepoV3 := v3Repository.NewOTPRepository(v3RepoOptions)

	repo := v3Repository.NewRepositoryV3()
	repo.SetOTPRepository(otpRepoV3)

	return repo
}

// WiringUpServiceV3 will bootstrapping service V3
func WiringUpService(repo *repository.Repository, db *gorm.DB, cache *redis.Pool, logger *common.APILogger, appConfig *config.Config) (*service.Service, error) {

	serviceOption := service.ServiceOption{
		DB:        db,
		Logger:    logger,
		Cache:     cache,
		AppConfig: appConfig,
	}

	/* userService, err := service.NewUserServiceBuilder(serviceOption).
		Build()
	if err != nil {
		log.Fatalln(err)
	} */

	service := &service.Service{
		//User: userService,
	}

	return service, nil
}
