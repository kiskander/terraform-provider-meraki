package meraki

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDevice() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDeviceRead,
		Schema: map[string]*schema.Schema{
			"org_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: false,
				Required: true,
			},
			"devices": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"lat": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"lng": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"notes": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"network_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"serial": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"model": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"mac": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"lan_ip": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"firmware": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceDeviceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	OrganizationID := d.Get("org_id").(string)
	if OrganizationID == "" {
		return diag.FromErr(errors.New("missing org_id"))
	}

	client := m.(*Client)
	devices, diags := client.GetOrganizationDevices(ctx, &GetOrganizationDevicesInput{OrganizationID: OrganizationID})
	if diags != nil {
		return diags
	}

	items := make([]map[string]interface{}, 0)
	for _, device := range devices {
		items = append(items, map[string]interface{}{
			"name":       device.Name,
			"lat":        device.Lat,
			"lng":        device.Lng,
			"address":    device.Address,
			"notes":      device.Notes,
			"network_id": device.NetworkID,
			"serial":     device.Serial,
			"model":      device.Model,
			"mac":        device.Mac,
			"lan_ip":     device.LanIP,
			"firmware":   device.Firmware,
		})
	}

	if err := d.Set("devices", items); err != nil {
		return diag.FromErr(err)
	}
	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
