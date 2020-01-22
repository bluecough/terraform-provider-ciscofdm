package main

import (
	"github.com/bluecough/go-ftd"
)

type Config struct {
	APIURL 		string
	Username 	string
	Password	string
	SSLNoVerify	string
}

func (c *Config) NewClient() (*goftd.FTD, error) {
	params := make(map[string]string)
	params["grant_type"] = "password"
	params["username"] = c.Username
	params["password"] = c.Password
	params["insecure"] = "true"

	return goftd.NewFTD(c.APIURL, params)
}