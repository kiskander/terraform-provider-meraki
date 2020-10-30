package meraki

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDevice() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDeviceRead,
		Schema: map[string]*schema.Schema{
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
						"networkid": &schema.Schema{
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
						"lanip": &schema.Schema{
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

	orgID := d.Get("org_id").(string)

	client := m.(*Client)
	devices, diags := client.GetDevices(ctx, &GetDevicesInput{OrgID: orgID})
	if diags != nil {
		return diags
	}

	if err := d.Set("devices", devices); err != nil {
		return diag.FromErr(err)
	}
	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
