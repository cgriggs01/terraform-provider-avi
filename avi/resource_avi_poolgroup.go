/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strings"
)

func ResourcePoolGroupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cloud_config_cksum": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"cloud_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "/api/cloud?name=Default-Cloud"},
		"created_by": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"deployment_policy_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"description": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"fail_action": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Set:      func(v interface{}) int { return 0 }, Elem: ResourceFailActionSchema()},
		"implicit_priority_labels": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"members": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourcePoolGroupMemberSchema()},
		"min_servers": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true},
		"priority_labels_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"tenant_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true},
	}
}

func resourceAviPoolGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviPoolGroupCreate,
		Read:   ResourceAviPoolGroupRead,
		Update: resourceAviPoolGroupUpdate,
		Delete: resourceAviPoolGroupDelete,
		Schema: ResourcePoolGroupSchema(),
	}
}

func ResourceAviPoolGroupRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourcePoolGroupSchema()
	log.Printf("[INFO] ResourceAviPoolGroupRead Avi Client %v\n", d)
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/poolgroup/" + uuid.(string)
		err := client.AviSession.Get(path, &obj)
		if err != nil {
			d.SetId("")
			return nil
		}
	} else {
		d.SetId("")
		return nil
	}
	// no need to set the ID
	log.Printf("ResourceAviPoolGroupRead CURRENT obj %v\n", d)

	log.Printf("ResourceAviPoolGroupRead Read API obj %v\n", obj)
	if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
		log.Printf("[INFO] ResourceAviPoolGroupRead Converted obj %v\n", tObj)
		//err = d.Set("obj", tObj)
		if err != nil {
			log.Printf("[ERROR] in setting read object %v\n", err)
		}
	}
	log.Printf("[INFO] ResourceAviPoolGroupRead Updated %v\n", d)
	return nil
}

func resourceAviPoolGroupCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourcePoolGroupSchema()
	err := ApiCreateOrUpdate(d, meta, "poolgroup", s)
	log.Printf("[DEBUG] created object %v: %v", "poolgroup", d)
	if err == nil {
		err = ResourceAviPoolGroupRead(d, meta)
	}
	log.Printf("[DEBUG] created object %v: %v", "poolgroup", d)
	return err
}

func resourceAviPoolGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourcePoolGroupSchema()
	err := ApiCreateOrUpdate(d, meta, "poolgroup", s)
	if err == nil {
		err = ResourceAviPoolGroupRead(d, meta)
	}
	log.Printf("[DEBUG] updated object %v: %v", "poolgroup", d)
	return err
}

func resourceAviPoolGroupDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "poolgroup"
	log.Println("[INFO] ResourceAviPoolGroupRead Avi Client")
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviPoolGroupDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
