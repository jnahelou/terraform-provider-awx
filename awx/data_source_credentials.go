/*
Use this data source to query Credential by ID.

Example Usage

```hcl
*TBD*
```

*/
package awx

import (
	"context"
	"fmt"
	"strconv"
	"time"

	awx "github.com/denouche/goawx/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCredentials() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCredentialsRead,
		Schema: map[string]*schema.Schema{
			"credentials": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"username": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"kind": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"credential_type_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceCredentialsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*awx.AWX)

	creds, err := client.CredentialsService.ListCredentials(map[string]string{})
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to fetch credentials",
			Detail:   fmt.Sprintf("Unable to fetch credentials from AWX API with error: %s.", err),
		})
		return diags
	}

	parsedCreds := make([]map[string]interface{}, 0)
	for _, c := range creds {
		fmt.Printf("%+v\n", c)
		parsedCreds = append(parsedCreds, map[string]interface{}{
			"id":                 c.ID,
			"name":               c.Name,
			"username":           c.Inputs["username"],
			"kind":               c.Kind,
			"credential_type_id": c.CredentialTypeID,
		})
	}

	err = d.Set("credentials", parsedCreds)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
