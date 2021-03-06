---
layout: "avi"
page_title: "Avi: avi_applicationpersistenceprofile"
sidebar_current: "docs-avi-resource-applicationpersistenceprofile"
description: |-
Creates and manages Avi ApplicationPersistenceProfile.
---

# avi_applicationpersistenceprofile

The ApplicationPersistenceProfile resource allows the creation and management of Avi ApplicationPersistenceProfile

## Example Usage

```hcl
resource "ApplicationPersistenceProfile" "foo" {
    name = "terraform-example-foo"
    tenant = "admin"
}
```

## Argument Reference

The following arguments are supported:

    * `app_cookie_persistence_profile` - (Optional ) argument_description.
        * `description` - (Optional ) argument_description.
        * `hdr_persistence_profile` - (Optional ) argument_description.
        * `http_cookie_persistence_profile` - (Optional ) argument_description.
        * `ip_persistence_profile` - (Optional ) argument_description.
        * `is_federated` - (Optional ) argument_description.
        * `name` - (Required) argument_description.
        * `persistence_type` - (Required) argument_description.
        * `server_hm_down_recovery` - (Optional ) argument_description.
        * `tenant_ref` - (Optional ) argument_description.
        
### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

                                            * `uuid` - argument_description.
    
