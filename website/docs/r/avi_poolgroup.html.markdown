---
layout: "avi"
page_title: "Avi: avi_poolgroup"
sidebar_current: "docs-avi-resource-poolgroup"
description: |-
Creates and manages Avi PoolGroup.
---

# avi_poolgroup

The PoolGroup resource allows the creation and management of Avi PoolGroup

## Example Usage

```hcl
resource "PoolGroup" "foo" {
    name = "terraform-example-foo"
    tenant = "admin"
}
```

## Argument Reference

The following arguments are supported:

    * `cloud_config_cksum` - (Optional ) argument_description.
        * `cloud_ref` - (Optional ) argument_description.
        * `created_by` - (Optional ) argument_description.
        * `deployment_policy_ref` - (Optional ) argument_description.
        * `description` - (Optional ) argument_description.
        * `fail_action` - (Optional ) argument_description.
        * `implicit_priority_labels` - (Optional ) argument_description.
        * `members` - (Optional ) argument_description.
        * `min_servers` - (Optional ) argument_description.
        * `name` - (Required) argument_description.
        * `priority_labels_ref` - (Optional ) argument_description.
        * `tenant_ref` - (Optional ) argument_description.
        
### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

                                                    * `uuid` - argument_description.
    
