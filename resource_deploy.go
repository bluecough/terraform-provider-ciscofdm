package main

import (
	"fmt"
	goftd "github.com/bluecough/go-ftd"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceDeployObject() *schema.Resource {
	return &schema.Resource{
		Create: resourceDeployCreate,
		Read:   resourceDeployRead,
		Update: resourceDeployUpdate,
		Delete: resourceDeployDelete,

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
	n.Name = d.Get("name").(string)
	n.SubType = d.Get("subtype").(string)
	n.Value = d.Get("value").(string)

	err := cf.PostDeploy(n, 1)
	if err != nil {
		fmt.Errorf("error: %s\n", err)
	}
	return resourceDeployCreate(d, m)
}
func resourceDeployRead(d *schema.ResourceData, m interface{}) error {
	//cf := m.(*goftd.FTD)
	log.Println("GS DEBUG === NetworkObjectRead== id ", d.Id())
	//cf.GetNetworkObjectByID(d.Id())
	return nil
}
func resourceDeployUpdate(d *schema.ResourceData, m interface{}) error {
	//cf := m.(*goftd.FTD)
	log.Println("GS DEBUG === NetworkObjectRead== id ", d.Id())
	//cf.GetNetworkObjectByID(d.Id())
	return nil
}
func resourceDeployDelete(d *schema.ResourceData, m interface{}) error {
	cf := m.(*goftd.FTD)

	n := new(goftd.DeployObject)
	n.Name = d.Get("name").(string)
	n.SubType = d.Get("subtype").(string)
	n.Value = d.Get("value").(string)

	err := cf.PostDeploy(n, 1)
	if err != nil {
		fmt.Errorf("error: %s\n", err)
	}
	return resourceDeployCreate(d, m)
}
