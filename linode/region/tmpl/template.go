package tmpl

import (
	"testing"

	"github.com/linode/terraform-provider-linode/linode/acceptance"
)

type TemplateData struct {
	Region string
	Label  string
}

func DataBasic(t *testing.T, id, label string) string {
	return acceptance.ExecuteTemplate(t,
		"region_data_basic", TemplateData{Region: id, Label: label})
}
