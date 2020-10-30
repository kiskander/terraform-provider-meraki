package meraki

// Parameter 	Type 	Description
// name 	string 	The name of the new network
// productTypes 	array 	The product type(s) of the new network. Valid types are wireless, appliance, switch, systemsManager, camera, cellularGateway, environmental. If more than one type is included, the network will be a combined network.
// tags 	array 	A list of tags to be applied to the network
// timeZone 	string 	The timezone of the network. For a list of allowed timezones, please see the 'TZ' column in the table in this article.
// copyFromNetworkId 	string 	The ID of the network to copy configuration from. Other provided parameters will override the copied configuration, except type which must match this network's type exactly.

// func resourceMerakiOrganizationNetwork() *schema.Resource {
// 	return &schema.Resource{
// 		Read: resourceMerakiOrganizationNetworkRead,
// 		Schema: map[string]*schema.Schema{
// 			"organization_id": {
// 				Type:     schema.TypeString,
// 				Required: true,
// 				// ForceNew: true,
// 			},
// 			"name": {
// 				Type:     schema.TypeString,
// 				Required: true,
// 				// ForceNew: true,
// 			},
// 			"product_types": {
// 				Type:     schema.TypeSet,
// 				Required: true,
// 				Elem: &schema.Schema{
// 					Type: schema.TypeString,
// 				},
// 			},
// 			"tags": {
// 				Type:     schema.TypeSet,
// 				Optional: true,
// 				Elem: &schema.Schema{
// 					Type: schema.TypeString,
// 				},
// 			},
// 			"timezone": {
// 				Type:     schema.TypeString,
// 				Required: true,
// 			},
// 			// "copy_from_network_id": {
// 			// Type:     schema.TypeString,
// 			// Required: true,
// 			// },
// 		},
// 	}
// }
