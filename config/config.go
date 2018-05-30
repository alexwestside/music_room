package config

import (
	"github.com/jinzhu/configor"
)

//type SMTPConfig struct {
//	Host     string
//	Port     string
//	User     string
//	Password string
//}

var Config = struct {
	//HTTPS bool `default:"false" env:"HTTPS"`
	//Port  uint `default:"7000" env:"PORT"`
	DB    struct {
		Name     string `env:"DBName" default:"qor_example"`
		Adapter  string `env:"DBAdapter" default:"mysql"`
		Host     string `env:"DBHost" default:"localhost"`
		Port     string `env:"DBPort" default:"3306"`
		User     string `env:"DBUser"`
		Password string `env:"DBPassword"`
	}
	//S3 struct {
	//	AccessKeyID     string `env:"AWS_ACCESS_KEY_ID"`
	//	SecretAccessKey string `env:"AWS_SECRET_ACCESS_KEY"`
	//	Region          string `env:"AWS_Region"`
	//	S3Bucket        string `env:"AWS_Bucket"`
	//}
	//AmazonPay struct {
	//	MerchantID   string `env:"AmazonPayMerchantID"`
	//	AccessKey    string `env:"AmazonPayAccessKey"`
	//	SecretKey    string `env:"AmazonPaySecretKey"`
	//	ClientID     string `env:"AmazonPayClientID"`
	//	ClientSecret string `env:"AmazonPayClientSecret"`
	//	Sandbox      bool   `env:"AmazonPaySandbox"`
	//	CurrencyCode string `env:"AmazonPayCurrencyCode" default:"JPY"`
	//}
	//SMTP         SMTPConfig
	//Github       github.Config
	//Google       google.Config
	//Facebook     facebook.Config
	//Twitter      twitter.Config
	//GoogleAPIKey string `env:"GoogleAPIKey"`
	//BaiduAPIKey  string `env:"BaiduAPIKey"`
}{}


//var (
//	Root           = os.Getenv("GOPATH") + "/src/github.com/qor/qor-example"
//	Mailer         *mailer.Mailer
//	//Render         = render.New()
//	AmazonPay      *amazonpay.AmazonPay
//	PaymentGateway gomerchant.PaymentGateway
//	RedirectBack   = redirect_back.New(&redirect_back.Config{
//		SessionManager:  manager.SessionManager,
//		IgnoredPrefixes: []string{"/auth"},
//	})
//)

func Load() {
	//if err := configor.Load(&Config, "config/database.yml", "config/smtp.yml", "config/application.yml"); err != nil {
	//	panic(err)
	//}

	if err := configor.Load(&Config, "sqlite.yml"); err != nil {
		panic(err)
	}

	//location.GoogleAPIKey = Config.GoogleAPIKey
	//location.BaiduAPIKey = Config.BaiduAPIKey

	//if Config.S3.AccessKeyID != "" {
	//	oss.Storage = s3.New(&s3.Config{
	//		AccessID:  Config.S3.AccessKeyID,
	//		AccessKey: Config.S3.SecretAccessKey,
	//		Region:    Config.S3.Region,
	//		Bucket:    Config.S3.S3Bucket,
	//	})
	//}
	//
	//AmazonPay = amazonpay.New(&amazonpay.Config{
	//	MerchantID: Config.AmazonPay.MerchantID,
	//	AccessKey:  Config.AmazonPay.AccessKey,
	//	SecretKey:  Config.AmazonPay.SecretKey,
	//	Sandbox:    true,
	//	Region:     "jp",
	//})
	//
	//Mailer = mailer.New(&mailer.Config{
	//	Sender: logger.New(&logger.Config{}),
	//})
}