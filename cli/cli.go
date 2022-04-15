package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mr-doggy/doggy-sdk-go/doggy"
	"github.com/spf13/viper"
)

type Session struct {
	AccessToken string `json:"access_token" yaml:"access_token" mapstructure:"access_token"`
	ExpiresIn   int64  `json:"expires_in" yaml:"expires_in" mapstructure:"expires_in"`
	CreatedAt   int64  `json:"created_at" yaml:"created_at" mapstructure:"created_at"`
	TokenType   string `json:"token_type" yaml:"token_type" mapstructure:"token_type"`
	Scope       string `json:"scope" yaml:"scope" mapstructure:"scope"`
	PhoneNumber string `json:"phone_number" yaml:"phone_number" mapstructure:"phone_number"`
}

type Config struct {
	Session *Session `json:"session" yaml:"session" mapstructure:"session"`
}

func (session *Session) Valid() bool {
	if session.AccessToken == "" {
		return false
	}
	if session.ExpiresIn <= 0 {
		return false
	}
	if session.CreatedAt <= 0 {
		return false
	}
	return session.CreatedAt+session.ExpiresIn > time.Now().Unix()
}

type DoggyCli struct {
	api     *doggy.APIClient
	session *Session
	config  *Config
}

func NewDoggyCli() *DoggyCli {
	cfg := doggy.NewConfiguration()
	cfg.Scheme = "https"
	cfg.Host = "api.doggy.code2code.cn"
	api := doggy.NewAPIClient(cfg)
	config := &Config{}
	err := viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
	if config.Session.Valid() {
		api.GetConfig().AddDefaultHeader("Authorization", config.Session.TokenType+" "+config.Session.AccessToken)
	}
	return &DoggyCli{
		api:     api,
		session: config.Session,
		config:  config,
	}
}

func (cli *DoggyCli) Login() error {
	if cli.session.Valid() {
		return fmt.Errorf("已经登录， 无需重复登录")
	}
	var phoneNumber string
	if err := survey.AskOne(&survey.Input{Message: "请输入手机号码+86:"}, &phoneNumber); err != nil {
		return err
	}
	if !strings.HasPrefix("+86", phoneNumber) {
		phoneNumber = "+86" + phoneNumber
	}
	_, err := cli.api.SmsApi.ApiAppSmsSendLoginCodePost(context.Background()).
		SendSmsCodeDto(doggy.SendSmsCodeDto{
			PhoneNumber: *doggy.NewNullableString(&phoneNumber),
		}).Execute()
	if err != nil {
		return err
	}
	var smsCode string
	if err := survey.AskOne(&survey.Input{Message: "请输入短信验证码:"}, &smsCode); err != nil {
		return err
	}
	v := url.Values{}
	v.Set("grant_type", "sms")
	v.Set("scope", "Doggy")
	v.Set("client_id", "Doggy_Sms_GrantType")
	v.Set("client_secret", "1q2w3e*")
	v.Set("phoneNumber", phoneNumber)
	v.Set("smsCode", smsCode)

	rsp, err := cli.api.GetConfig().HTTPClient.PostForm(cli.api.GetConfig().Scheme+"://"+cli.api.GetConfig().Host+"/connect/token", v)
	if err != nil {
		return err
	}
	bts, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return err
	}
	session := &Session{}
	err = json.Unmarshal(bts, &session)
	if err != nil {
		return err
	}
	cli.session = session
	session.PhoneNumber = phoneNumber
	session.CreatedAt = time.Now().Unix()
	viper.Set("session", session)
	return viper.WriteConfig()
}

func (cli *DoggyCli) Logout() error {
	viper.Set("session", nil)
	return viper.WriteConfig()
}
