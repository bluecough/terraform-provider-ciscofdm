package main
import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/bluecough/go-ftd"
	"log"
)

func resourcePortObject() *schema.Resource {
	return &schema.Resource{
		Create: resourcePortObjectCreate,
		Read:   resourcePortObjectRead,
		Update: resourcePortObjectUpdate,
		Delete: resourcePortObjectDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"layer4": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourcePortObjectCreate(d *schema.ResourceData, m interface{}) error {
	cf := m.(*goftd.FTD)
	po := goftd.PortObject{}

	po.Name = d.Get("name").(string)
	po.Type = d.Get("layer4").(string)
	po.Port = d.Get("port").(string)
    var err error
	if po.Type == "UDP" {
		err = cf.CreateUDPPortObject(&po, goftd.DuplicateActionReplace)
	}
	err = cf.CreateTCPPortObject(&po, goftd.DuplicateActionReplace)
	if err != nil {
		log.Println("Error: %s", err)
		return err
	}

	d.SetId(po.ID)
	return resourceServerRead(d, m)
}

func resourcePortObjectRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourcePortObjectUpdate(d *schema.ResourceData, m interface{}) error {
	cf := m.(*goftd.FTD)

	return resourceServerRead(d, m)
}

func resourcePortObjectDelete(d *schema.ResourceData, m interface{}) error {
	cf := m.(*goftd.FTD)
	po := goftd.PortObject{}
	po.ID = d.Id()
	po.Type = d.Get("layer4").(string)
	log.Println("===============> ", po.ID)
	var err error

	if po.Type == "UDP" {
		po.Type = goftd.TypeUDPPortObject
		err = cf.DeletePortObject(&po)
	}
	po.Type = goftd.TypeTCPPortObject
	err = cf.DeletePortObject(&po)

	if err != nil {
		log.Println("Error: %s", err)
		return err
	}
	return nil
}
