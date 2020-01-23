package main

import (
	"github.com/golang/glog"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/bluecough/go-ftd"
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
		glog.Errorf("error: %s\n", err)
	}
	d.SetId(n.ID)
	return resourceServerRead(d, m)
}

func resourceNetworkObjectRead(d *schema.ResourceData, m interface{}) error {
	cf := m.(*goftd.FTD)
	cf.GetNetworkObjects(100)
	return nil
}

func resourceNetworkObjectUpdate(d *schema.ResourceData, m interface{}) error {

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
