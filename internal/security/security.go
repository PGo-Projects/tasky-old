package security

import (
	"encoding/base64"

	asecurity "github.com/PGo-Projects/authboss-security"
	"github.com/PGo-Projects/tasky/internal/config"
	"github.com/PGo-Projects/tasky/internal/tasky"
	"github.com/go-chi/chi"
	"github.com/gorilla/schema"
	"github.com/spf13/viper"
)

var schemaDec = schema.NewDecoder()

func MustSetupSecurity() {
	usernameRule := asecurity.Rules{
		FieldName: "username", Required: true,
		MinLength: 4,
	}
	passwordRule := asecurity.Rules{
		FieldName: "password", Required: true,
		MinLength: 12,
	}

	authBossHTTPBodyReader := asecurity.NewHTTPBodyReader(usernameRule, passwordRule)
	authBossConfig := asecurity.NewAuthBossConfigBuilder("tasky", authBossHTTPBodyReader).
		WithRootURL("http://lvh.me:8080").
		WithAuthLoginOK("/tasky").
		WithRegisterOK("/tasky").
		Build()

	sessionStore, cookieStore := mustSetupSessionCookieStores()

	asecurity.MustSetupAuthbossWithStores(authBossConfig, "tasky", sessionStore, cookieStore)
	schemaDec.IgnoreUnknownKeys(true)
}

func mustSetupSessionCookieStores() (asecurity.SessionStorer, asecurity.CookieStorer) {
	k := viper.GetString(config.CookieStoreKey)
	cookieStoreKey, err := base64.StdEncoding.DecodeString(k)
	if err != nil {
		panic(err)
	}
	k = viper.GetString(config.SessionStoreKey)
	sessionStoreKey, err := base64.StdEncoding.DecodeString(k)
	if err != nil {
		panic(err)
	}
	k = viper.GetString(config.EncryptionKey)
	encryptionKey, err := base64.StdEncoding.DecodeString(k)
	if err != nil {
		panic(err)
	}

	sessionStore := asecurity.NewSessionStorerBuilder("tasky", sessionStoreKey, encryptionKey).
		WithMaxAge(3600).
		WithSecure(false).
		WithHttpOnly(false).
		Build()
	cookieStore := asecurity.NewCookieStorer(cookieStoreKey, encryptionKey).
		WithMaxAge(3600).
		WithSecure(false).
		WithHttpOnly(false)
	return sessionStore, cookieStore
}

func SetupSecurityRouter(mux chi.Router) {
	asecurity.RegisterAuthRoutes(mux, "/auth")

	config := asecurity.ProtectedRoutesConfig{
		RedirectToLoginPage: true,
		ForceFullAuth:       false,
		Force2FA:            false,
	}
	asecurity.RegisterProtectedRoutes(mux, config, tasky.ProtectedRoutes)
}
