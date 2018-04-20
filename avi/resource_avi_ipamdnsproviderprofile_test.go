package avi

import (
	"fmt"
	"strings"
	"testing"

	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAVIIPAMDNSProviderProfileBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVIIPAMDNSProviderProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIIPAMDNSProviderProfileConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIIPAMDNSProviderProfileExists("avi_ipamdnsproviderprofile.testipamdnsproviderprofile"),
					resource.TestCheckResourceAttr(
						"avi_ipamdnsproviderprofile.testipamdnsproviderprofile", "name", "ipam-test")),
			},
			{
				Config: testAccUpdatedAVIIPAMDNSProviderProfileConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIIPAMDNSProviderProfileExists("avi_ipamdnsproviderprofile.testipamdnsproviderprofile"),
					resource.TestCheckResourceAttr(
						"avi_ipamdnsproviderprofile.testipamdnsproviderprofile", "name", "ipam-abc")),
			},
		},
	})

}

func testAccCheckAVIIPAMDNSProviderProfileExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No IPAMDNS Provider Profile ID is set")
		}
		url := strings.SplitN(rs.Primary.ID, "/api", 2)[1]
		uuid := strings.Split(url, "#")[0]
		path := "api" + uuid
		err := conn.Get(path, &obj)
		if err != nil {
			return err
		}
		return nil
	}

}

func testAccCheckAVIIPAMDNSProviderProfileDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_ipamdnsproviderprofile" {
			continue
		}
		url := strings.SplitN(rs.Primary.ID, "/api", 2)[1]
		uuid := strings.Split(url, "#")[0]
		path := "api" + uuid
		err := conn.Get(path, &obj)
		if err != nil {
			if strings.Contains(err.Error(), "404") {
				return nil
			}
			return err
		}
		if len(obj.(map[string]interface{})) > 0 {
			return fmt.Errorf("AVI IPAMDNS Provider Profile still exists")
		}
	}
	return nil
}

const testAccAVIIPAMDNSProviderProfileConfig = `
data "avi_tenant" "default_tenant"{
	name= "admin"
}
data "avi_cloud" "default_cloud" {
	name= "Default-Cloud"
}
data "avi_vrfcontext" "global_vrf" {
	name= "global"
}
resource "avi_ipamdnsproviderprofile" "testipamdnsproviderprofile" {
	name = "ipam-test"
	allocate_ip_in_vrf= false
  	internal_profile= {
        ttl= 31
  	}
  	type= "IPAMDNS_TYPE_INTERNAL"
	tenant_ref= "${data.avi_tenant.default_tenant.id}"
}
`

const testAccUpdatedAVIIPAMDNSProviderProfileConfig = `
data "avi_tenant" "default_tenant"{
	name= "admin"
}
data "avi_cloud" "default_cloud" {
	name= "Default-Cloud"
}
data "avi_vrfcontext" "global_vrf" {
	name= "global"
}
resource "avi_ipamdnsproviderprofile" "testipamdnsproviderprofile" {
	name = "ipam-abc"
	allocate_ip_in_vrf= false
    internal_profile= {
        ttl= 31
    }
  	type= "IPAMDNS_TYPE_INTERNAL"
	tenant_ref= "${data.avi_tenant.default_tenant.id}"
}
`
