package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	auth "github.com/idaaser/oneid-jwt-auth"
	"github.com/spf13/cobra"
)

var (
	configCmd = &cobra.Command{
		Use:   "auth",
		Short: "auth",
		Run: func(_ *cobra.Command, _ []string) {
			if err := _config.init(); err != nil {
				log.Fatalln(err)
			}

			u, err := _config.create()
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println(u)
		},
	}

	_config = &config{}
)

type (
	config struct {
		f    string
		c    *auth.Config
		IDP  idp  `json:"idp"`
		User user `json:"user"`
	}

	idp struct {
		Key           string `json:"key"`
		BaseLoginURL  string `json:"url"`
		Issuer        string `json:"issuer"`
		TokenParam    string `json:"token_param"`
		TokenLifetime int    `json:"token_lifetime"`

		App string `json:"app"`
	}

	user struct {
		ID       string `json:"id"`
		Username string `json:"username"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Mobile   string `json:"mobile"`

		Extension map[string]any `json:"extension"`
	}
)

func (c *config) init() error {
	b, err := os.ReadFile(c.f)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(b, c); err != nil {
		return err
	}

	authc, err := auth.NewConfig(
		c.IDP.BaseLoginURL,
		c.IDP.Key,
		auth.WithIssuer(c.IDP.Issuer),
		auth.WithTokenLifetime(c.IDP.TokenLifetime),
		auth.WithTokenParam(c.IDP.TokenParam),
	)
	if err != nil {
		return err
	}

	c.c = authc
	return nil
}

func (c *config) create() (string, error) {
	return c.c.NewLoginURL(
		auth.Userinfo{
			ID:                c.User.ID,
			Name:              c.User.Name,
			PreferredUsername: c.User.Username,
			Email:             c.User.Email,
			Mobile:            c.User.Mobile,
			Extension:         c.User.Extension,
		},
		c.IDP.App)
}

func init() {
	configCmd.PersistentFlags().StringVarP(&_config.f,
		"config", "c", "", "config file",
	)
	rootCmd.AddCommand(configCmd)
}
