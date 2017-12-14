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

func ResourcePKIProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"ca_certs": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceSSLCertificateSchema()},
		"created_by": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"crl_check": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"crls": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceCRLSchema()},
		"ignore_peer_chain": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"is_federated": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true},
		"tenant_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true},
		"validate_only_leaf_crl": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
	}
}

func resourceAviPKIProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviPKIProfileCreate,
		Read:   ResourceAviPKIProfileRead,
		Update: resourceAviPKIProfileUpdate,
		Delete: resourceAviPKIProfileDelete,
		Schema: ResourcePKIProfileSchema(),
	}
}

func ResourceAviPKIProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourcePKIProfileSchema()
	log.Printf("[INFO] ResourceAviPKIProfileRead Avi Client %v\n", d)
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/pkiprofile/" + uuid.(string)
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
	log.Printf("ResourceAviPKIProfileRead CURRENT obj %v\n", d)

	log.Printf("ResourceAviPKIProfileRead Read API obj %v\n", obj)
	if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
		log.Printf("[INFO] ResourceAviPKIProfileRead Converted obj %v\n", tObj)
		//err = d.Set("obj", tObj)
		if err != nil {
			log.Printf("[ERROR] in setting read object %v\n", err)
		}
	}
	log.Printf("[INFO] ResourceAviPKIProfileRead Updated %v\n", d)
	return nil
}

func resourceAviPKIProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourcePKIProfileSchema()
	err := ApiCreateOrUpdate(d, meta, "pkiprofile", s)
	log.Printf("[DEBUG] created object %v: %v", "pkiprofile", d)
	if err == nil {
		err = ResourceAviPKIProfileRead(d, meta)
	}
	log.Printf("[DEBUG] created object %v: %v", "pkiprofile", d)
	return err
}

func resourceAviPKIProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourcePKIProfileSchema()
	err := ApiCreateOrUpdate(d, meta, "pkiprofile", s)
	if err == nil {
		err = ResourceAviPKIProfileRead(d, meta)
	}
	log.Printf("[DEBUG] updated object %v: %v", "pkiprofile", d)
	return err
}

func resourceAviPKIProfileDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "pkiprofile"
	log.Println("[INFO] ResourceAviPKIProfileRead Avi Client")
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviPKIProfileDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
