/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviVirtualService() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviVirtualServiceRead,
		Schema: map[string]*schema.Schema{
			"active_standby_se_tag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "ACTIVE_STANDBY_SE_1"},
			"analytics_policy": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAnalyticsPolicySchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"analytics_profile_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"apic_contract_graph": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"application_profile_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"bulk_sync_kvcache": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"client_auth": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHTTPClientAuthenticationParamsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"close_client_conn_on_config_update": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"cloud_config_cksum": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cloud_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cloud_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "CLOUD_NONE"},
			"connections_rate_limit": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRateProfileSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"content_rewrite": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceContentRewriteProfileSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"created_by": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"delay_fairness": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dns_info": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDnsInfoSchema(),
			},
			"dns_policies": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDnsPoliciesSchema(),
			},
			"east_west_placement": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"enable_autogw": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"enable_rhi": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enable_rhi_snat": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"error_page_profile_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"flow_dist": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "LOAD_AWARE"},
			"flow_label_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "NO_LABEL"},
			"fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"host_name_xlate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_policies": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceHTTPPoliciesSchema(),
			},
			"ign_pool_net_reach": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"l4_policies": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceL4PoliciesSchema(),
			},
			"limit_doser": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"max_cps_per_client": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"microservice_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_profile_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_security_policy_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nsx_securitygroup": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"performance_limits": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcePerformanceLimitsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"pool_group_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pool_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"remove_listening_port_on_vs_down": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"requests_rate_limit": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRateProfileSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"scaleout_ecmp": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"se_group_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_network_profile_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_metadata": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_pool_select": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceServicePoolSelectorSchema(),
			},
			"services": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceServiceSchema(),
			},
			"sideband_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSidebandProfileSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"snat_ip": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"ssl_key_and_certificate_refs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"ssl_profile_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_sess_cache_avg_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1024,
			},
			"static_dns_records": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDnsRecordSchema(),
			},
			"tenant_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"traffic_clone_profile_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"traffic_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "VS_TYPE_NORMAL"},
			"use_bridge_ip_as_vip": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"use_vip_as_snat": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vh_domain_name": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vh_parent_vs_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vip": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVipSchema(),
			},
			"vrf_context_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vs_datascripts": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceVSDataScriptsSchema(),
			},
			"vsvip_cloud_config_cksum": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vsvip_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"waf_policy_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
		},
	}
}
