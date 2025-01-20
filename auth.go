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
		f      string
		signer *auth.Signer
		IDP    idp               `json:"idp"`
		User   auth.Userinfo     `json:"user"`
		Params map[string]string `json:"params"`
	}

	idp struct {
		Key           string `json:"key"`
		BaseLoginURL  string `json:"url"`
		Issuer        string `json:"issuer"`
		TokenParam    string `json:"token_param"`
		TokenLifetime int    `json:"token_lifetime"`

		App string `json:"app"`
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

	signer, err := auth.NewSigner(
		c.IDP.Key,
		c.IDP.Issuer,
		c.IDP.BaseLoginURL,
		auth.WithTokenLifetime(c.IDP.TokenLifetime),
	)
	if err != nil {
		return err
	}

	c.signer = signer
	return nil
}

func (c *config) paramsKV() []string {
	kvs := make([]string, 0, 2*len(c.Params))
	for k, v := range c.Params {
		kvs = append(kvs, k, v)
	}
	return kvs
}

func (c *config) create() (string, error) {
	return c.signer.NewLoginURL(c.User, c.IDP.App, c.paramsKV()...)
}

func init() {
	configCmd.PersistentFlags().StringVarP(&_config.f,
		"config", "c", "", "config file",
	)
	rootCmd.AddCommand(configCmd)
}
