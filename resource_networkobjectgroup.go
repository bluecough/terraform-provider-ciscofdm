package main

import (
//	"fmt"
//	"github.com/bluecough/go-ftd"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceNetworkObjectGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkObjectGroupCreate,
		Read:   resourceNetworkObjectGroupRead,
		Update: resourceNetworkObjectGroupUpdate,
		Delete: resourceNetworkObjectGroupDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"objects": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource {
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

		},
	}
}

func resourceNetworkObjectGroupCreate(d *schema.ResourceData, m interface{}) error {

	return resourceServerRead(d, m)
}

func resourceNetworkObjectGroupRead(d *schema.ResourceData, m interface{}) error {

	return nil
}

func resourceNetworkObjectGroupUpdate(d *schema.ResourceData, m interface{}) error {

	return resourceServerRead(d, m)
}

func resourceNetworkObjectGroupDelete(d *schema.ResourceData, m interface{}) error {


	return nil
}
