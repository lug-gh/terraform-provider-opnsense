package opnsense

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider function to return a schema.Provider.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_url": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The URL of the OPNSense API endpoint.",
			},
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The API key for authenticating with the OPNSense instance.",
			},
			"api_secret": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The API secret for authenticating with the OPNSense instance.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"opnsense_haproxy_real_server": resourceRealServer(),
		},
	}
}
