package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourcePortObjectGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourcePortObjectGroupCreate,
		Read:   resourcePortObjectGroupRead,
		Update: resourcePortObjectGroupUpdate,
		Delete: resourcePortObjectGroupDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourcePortObjectGroupCreate(d *schema.ResourceData, m interface{}) error {
	address := d.Get("address").(string)
	d.SetId(address)
	return resourceServerRead(d, m)
}

func resourcePortObjectGroupRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourcePortObjectGroupUpdate(d *schema.ResourceData, m interface{}) error {
	address := d.Get("address").(string)
	d.SetId(address)
	return resourceServerRead(d, m)
}

func resourcePortObjectGroupDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}