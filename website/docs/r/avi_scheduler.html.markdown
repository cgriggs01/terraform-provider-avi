---
layout: "avi"
page_title: "Avi: avi_scheduler"
sidebar_current: "docs-avi-resource-scheduler"
description: |-
Creates and manages Avi Scheduler.
---

# avi_scheduler

The Scheduler resource allows the creation and management of Avi Scheduler

## Example Usage

```hcl
resource "Scheduler" "foo" {
    name = "terraform-example-foo"
    tenant = "admin"
}
```

## Argument Reference

The following arguments are supported:

    * `backup_config_ref` - (Optional ) argument_description.
        * `enabled` - (Optional ) argument_description.
        * `end_date_time` - (Optional ) argument_description.
        * `frequency` - (Optional ) argument_description.
        * `frequency_unit` - (Optional ) argument_description.
        * `name` - (Required) argument_description.
        * `run_mode` - (Optional ) argument_description.
        * `run_script_ref` - (Optional ) argument_description.
        * `scheduler_action` - (Optional ) argument_description.
        * `start_date_time` - (Optional ) argument_description.
        * `tenant_ref` - (Optional ) argument_description.
        
### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

                                                * `uuid` - argument_description.
    
