---
layout: "avi"
page_title: "AVI: avi_customipamdnsprofile"
sidebar_current: "docs-avi-datasource-customipamdnsprofile"
description: |-
Get information of Avi CustomIpamDnsProfile.
---

# avi_customipamdnsprofile

This data source is used to to get avi_customipamdnsprofile objects.

## Example Usage

```hcl
data "CustomIpamDnsProfile" "foo_CustomIpamDnsProfile" {
    uuid = "CustomIpamDnsProfile-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search CustomIpamDnsProfile by name.
* `uuid` - (Optional) Search CustomIpamDnsProfile by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `name` - Name of the custom ipam dns profile.
* `script_params` - Parameters that are always passed to the ipam/dns script.
* `script_uri` - Script uri of form controller //ipamdnsscripts/<file-name>.
* `tenant_ref` - It is a reference to an object of type tenant.
* `uuid` - Field introduced in 17.1.1.

