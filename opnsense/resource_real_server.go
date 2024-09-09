package opnsense

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRealServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceRealServerCreate,
		Read:   resourceRealServerRead,
		Update: resourceRealServerUpdate,
		Delete: resourceRealServerDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"address": {
				Type:     schema.TypeString,
				Required: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func resourceRealServerCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*OPNSenseClient)

	realServer := map[string]interface{}{
		"name":    d.Get("name").(string),
		"address": d.Get("address").(string),
		"port":    d.Get("port").(int),
	}

	_, err := client.DoRequest("POST", "haproxy/realserver/add", realServer)
	if err != nil {
		return err
	}

	d.SetId(realServer["name"].(string))
	return resourceRealServerRead(d, meta)
}

func resourceRealServerRead(d *schema.ResourceData, meta interface{}) error {
	// Reading not necessary for now, as this is just a simple implementation.
	return nil
}

func resourceRealServerUpdate(d *schema.ResourceData, meta interface{}) error {
	// Update logic can be similar to create.
	return resourceRealServerCreate(d, meta)
}

func resourceRealServerDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*OPNSenseClient)
	_, err := client.DoRequest("DELETE", "haproxy/realserver/delete/"+d.Id(), nil)
	return err
}
