package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceNetworkObject() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkObjectCreate,
		Read:   resourceNetworkObjectRead,
		Update: resourceNetworkObjectUpdate,
		Delete: resourceNetworkObjectDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"subtype": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceNetworkObjectCreate(d *schema.ResourceData, m interface{}) error {

	return resourceServerRead(d, m)
}

func resourceNetworkObjectRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceNetworkObjectUpdate(d *schema.ResourceData, m interface{}) error {

	return resourceServerRead(d, m)
}

func resourceNetworkObjectDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
