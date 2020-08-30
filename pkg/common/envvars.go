package common

import (
	"fmt"
	"log"
	"os"
)

const (
	EnvJWTAudience       = "JWT_AUDIENCE"
	EnvJWTExpiration     = "JWT_EXPIRATION"
	EnvJWTIssuer         = "JWT_ISSUER"
	EnvJWTSecret         = "JWT_SECRET"
	EnvJWTSubject        = "JWT_SUBJECT"
	EnvGoogleOauthClient = "GOOGLE_OAUTH_CLIENT"
	EnvGoogleOauthSecret = "GOOGLE_OAUTH_SECRET"
	EnvServerHostname    = "SERVER_HOSTNAME"
	EnvFrontendHostname  = "FRONTEND_HOSTNAME"
)

func AssertEnvVarsSet() {
	allVars := []string{
		EnvJWTAudience,
		EnvJWTExpiration,
		EnvJWTIssuer,
		EnvJWTSecret,
		EnvJWTSubject,
		EnvGoogleOauthClient,
		EnvGoogleOauthSecret,
		EnvServerHostname,
		EnvFrontendHostname,
	}

	fail := false

	for _, envVar := range allVars {
		if os.Getenv(envVar) == "" {
			fmt.Println("*** Environment var '" + envVar + "' must be set ***")
			fail = true
		}
	}

	if fail {
		log.Fatal("Some required env vars missing! Copy `.env.example` to `.env` and fill in the values.")
	}
}
