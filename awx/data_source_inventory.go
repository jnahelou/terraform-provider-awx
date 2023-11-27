package awx

import (
	"context"
	"strconv"

	awx "github.com/denouche/goawx/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceInventory() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceInventoriesRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"organization_id": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func dataSourceInventoriesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*awx.AWX)
	params := make(map[string]string)
	if groupName, okName := d.GetOk("name"); okName {
		params["name"] = groupName.(string)
	}

	if groupID, okGroupID := d.GetOk("id"); okGroupID {
		params["id"] = strconv.Itoa(groupID.(int))
	}

	if organizationID, okIOrgID := d.GetOk("organization_id"); okIOrgID {
		params["organization"] = strconv.Itoa(organizationID.(int))
	}
	if len(params) == 0 {
		return buildDiagnosticsMessage(
			"Get: Missing Parameters",
			"Please use one of the selectors (name or id)",
		)
	}
	inventories, _, err := client.InventoriesService.ListInventories(params)
	if err != nil {
		return buildDiagnosticsMessage(
			"Get: Fail to fetch Inventory",
			"Fail to find the inventory got: %s",
			err.Error(),
		)
	}
	if len(inventories) > 1 {
		return buildDiagnosticsMessage(
			"Get: find more than one Element",
			"The Query Returns more than one Inventory, %d",
			len(inventories),
		)
	}

	if len(inventories) == 0 {
		return buildDiagnosticsMessage(
			"Get: Inventory does not exist",
			"The Query Returns no Inventory matching filter %v",
			params,
		)
	}

	inventory := inventories[0]
	d = setInventoryResourceData(d, inventory)
	return diags
}
