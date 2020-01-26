package main
import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/bluecough/go-ftd"
	"log"
	"strings"
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

	d.SetId(po.ID + " " + po.Version)
	return resourcePortObjectRead(d, m)
}

func resourcePortObjectRead(d *schema.ResourceData, m interface{}) error {
/*
	cf := m.(*goftd.FTD)
	if d.Get("layer4") == "TCP" {
		_, err := cf.GetTCPPortObjectByID(d.Id())
		if err != nil {
			log.Println("Error: %s", err)
			return err
		}
	}
	_, err := cf.GetUDPPortObjectByID(d.Id())
	if err != nil {
		log.Println("Error: %s", err)
		return err
	}
*/
	return nil
}

func resourcePortObjectUpdate(d *schema.ResourceData, m interface{}) error {
	cf := m.(*goftd.FTD)
	idsplit := strings.Split(d.Id(), " ")
    po := goftd.PortObject{}
    po.ID = idsplit[0]
    po.Version = idsplit[1]
    po.Name = d.Get("name").(string)
    //po.Type = d.Get("type").(string)
    po.Port = d.Get("port").(string)
    log.Println("-------->", resourcePortObject())
	log.Println("-------->", po.Version)
    if d.Get("type") == "UDP" {
    	po.Type = "udpportobject"
	}
	po.Type = "tcpportobject"

    cf.UpdatePortObject(&po)
	d.SetId(po.ID + " " + po.Version)
	return resourcePortObjectRead(d, m)
}

func resourcePortObjectDelete(d *schema.ResourceData, m interface{}) error {
	cf := m.(*goftd.FTD)

	idsplit := strings.Split(d.Id(), " ")

	po := goftd.PortObject{}
	po.ID = idsplit[0]
	po.Version = idsplit[1]
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
