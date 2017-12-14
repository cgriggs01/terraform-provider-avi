/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceHdrMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"hdr": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"match_case": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "INSENSITIVE",
			},
			"match_criteria": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}
