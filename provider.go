package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	//"github.com/bluecough/go-ftd"
)

func Provider() *schema.Provider {
	return &schema.Provider{

		Schema: map[string]*schema.Schema{
			"api_url": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("FTD_API_URL", nil),
			},

			"username": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("FTD_USERNAME", nil),
			},

			"password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("FTD_PASSWORD", nil),
			},
			"ssl_no_verify": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("FTP_SSL_NO_VERIFY", false),
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"ciscofdm_dummy": resourceServer(),
			"ciscofdm_networkobject" : resourceNetworkObject(),
			"ciscofdm_networkobjectgroup" : resourceNetworkObjectGroup(),
		},
		ConfigureFunc: providerConfigure,
	}
}


func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		APIURL: 		d.Get("api_url").(string),
		Username: 		d.Get("username").(string),
		Password: 		d.Get("password").(string),
		SSLNoVerify: 	d.Get("ssl_no_verify").(string),
	}
	return config.NewClient()
}
