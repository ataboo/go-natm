package common

import (
	"fmt"
	"log"
	"os"
)

const (
	EnvJWTAudience       = "JWT_AUDIENCE"
	EnvJWTIssuer         = "JWT_ISSUER"
	EnvJWTSecret         = "JWT_SECRET"
	EnvJWTSubject        = "JWT_SUBJECT"
	EnvJWTRefreshExpMins = "JWT_REFRESH_EXP_MINS"
	EnvJWTIssueExpMins   = "JWT_ISSUE_EXP_MINS"
	EnvGoogleOauthClient = "GOOGLE_OAUTH_CLIENT"
	EnvGoogleOauthSecret = "GOOGLE_OAUTH_SECRET"
	EnvServerHostname    = "SERVER_HOSTNAME"
	EnvFrontendHostname  = "FRONTEND_HOSTNAME"
)

func AssertEnvVarsSet() {
	allVars := []string{
		EnvJWTAudience,
		EnvJWTIssuer,
		EnvJWTSecret,
		EnvJWTSubject,
		EnvJWTRefreshExpMins,
		EnvJWTIssueExpMins,
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
