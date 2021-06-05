package main

import (
	"fmt"
	goftd "github.com/bluecough/go-ftd"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceDeployObject() *schema.Resource {
	return &schema.Resource{
		Create: resourceDeployCreate,
//		Delete: resourceDeployDelete,

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
				Type:     schema.TypeBool,
				Required: true,
			},
		},
	}
}

func resourceDeployCreate(d *schema.ResourceData, m interface{}) error {
	cf := m.(*goftd.FTD)

	n := new(goftd.DeployObject)
	n.Name = d.Get("name")
	n.Value = d.Get("value")

	err := cf.PostDeploy(n, 1)
	if err != nil {
		fmt.Errorf("error: %s\n", err)
	}
	return resourceDeployCreate(d, m)
}