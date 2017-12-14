/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourcensxVersionsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"version": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourcensxVersionSchema(),
			},
		},
	}
}
