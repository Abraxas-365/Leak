package auth

import (
	"github/abraxas-365/Leak/pkg/auth/application"
	"github/abraxas-365/Leak/pkg/auth/infrastructure/rest"
	user "github/abraxas-365/Leak/pkg/user/application"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/redirect/v2"
	"golang.org/x/oauth2"
)

func ModuleFactory(app *fiber.App, Userapp user.Application) {
	//Oauth2 config
	conf := &oauth2.Config{
		ClientID:     "651680703153478",
		ClientSecret: "33bfc8401b26f33e2f8a9d32ff2f471d",
		RedirectURL:  "http://google.com/",
		Scopes:       []string{"user_profile", "user_media"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://api.instagram.com/oauth/authorize",
			TokenURL: "https://api.instagram.com/oauth/access_token",
		},
	}

	/*
		Redirect to istagram to login
		after login this will redirect to a page specify in the config
		and will send in the URL the token as parameter
	*/
	app.Use(redirect.New(redirect.Config{
		Rules: map[string]string{
			"/instagram": conf.AuthCodeURL("randomstate"),
		},
		StatusCode: 301,
	}))
	application := application.AppFactory(Userapp)
	handler := rest.HandlerFabric(application)
	rest.Routes(app, handler)

}
