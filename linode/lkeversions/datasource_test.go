package lkeversions_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/linode/terraform-provider-linode/linode/acceptance"
	"github.com/linode/terraform-provider-linode/linode/lkeversions/tmpl"
)

func TestAccDataSourceLinodeLkeVersions_basic(t *testing.T) {
	t.Parallel()

	resourceName := "data.linode_lke_versions.foobar"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acceptance.PreCheck(t) },
		Providers: acceptance.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: tmpl.DataBasic(t),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "versions.0.id"),
				),
			},
		},
	})
}
