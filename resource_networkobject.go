package main

import (
	"fmt"
	"github.com/bluecough/go-ftd"
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
	cf := m.(*goftd.FTD)

	n := new(goftd.NetworkObject)
	n.Name = d.Get("name").(string)
	n.SubType = d.Get("subtype").(string)
	n.Value = d.Get("value").(string)

	err := cf.CreateNetworkObject(n, goftd.DuplicateActionReplace)
	if err != nil {
		fmt.Errorf("error: %s\n", err)
	}
	d.SetId(n.ID)
	return resourceServerRead(d, m)
}

func resourceNetworkObjectRead(d *schema.ResourceData, m interface{}) error {
	cf := m.(*goftd.FTD)
	cf.GetNetworkObjectByID(d.Id())
	return nil
}

func resourceNetworkObjectUpdate(d *schema.ResourceData, m interface{}) error {
	cf := m.(*goftd.FTD)

	existing := new(goftd.NetworkObject)
	existing, err := cf.GetNetworkObjectByID(d.Id())

	n := new(goftd.NetworkObject)
	n.ID = d.Id()
	n.Name = d.Get("name").(string)
	n.SubType = d.Get("subtype").(string)
	n.Value = d.Get("value").(string)
	n.Type = "networkobject"
	n.Version = existing.Version

	err = cf.UpdateNetworkObject(n)
	if err != nil {
		fmt.Errorf("error: %s\n", err)
	}

	//d.SetId(n.ID)
	return resourceServerRead(d, m)
}

func resourceNetworkObjectDelete(d *schema.ResourceData, m interface{}) error {

	cf := m.(*goftd.FTD)
	id := d.Id()

	err := cf.DeleteNetworkObjectByID(id)
	if err != nil {
		return err
	}

	return nil
}
