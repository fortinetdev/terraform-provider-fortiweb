// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure server_policy/server_pool.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceServerPolicyServerPool() *schema.Resource {
	return &schema.Resource{
		Read:   resourceServerPolicyServerPoolRead,
		Update: resourceServerPolicyServerPoolUpdate,
		Create: resourceServerPolicyServerPoolCreate,
		Delete: resourceServerPolicyServerPoolDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"server_balance": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"health": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"hlck_sip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"hlck_sip6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"lb_algo": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"persistence": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"flag": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"server_pool_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"sub_table_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"sub_table_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"http_reuse": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"reuse_conn_total_time": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"reuse_conn_idle_time": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"reuse_conn_max_request": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"reuse_conn_max_count": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"adfs_server_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"panic_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"panic_threshold": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"server_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"adfs_username": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"adfs_password": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"sdn_addr_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"sdn": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"filter": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"server_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"backup_server": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"proxy_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"proxy_protocol_version": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ssl": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"implicit_ssl": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ssl_quiet_shutdown": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ssl_session_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"server_side_sni": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"multi_certificate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"certificate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"certificate_group": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"certificate_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"lets_certificate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"intermediate_certificate_group": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"certificate_verify": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"client_certificate_proxy": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"client_certificate_proxy_sign_ca": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"client_certificate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"server_certificate_verify": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"server_certificate_verify_policy": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"server_certificate_verify_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"session_ticket_reuse": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"session_id_reuse": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"sni": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"sni_certificate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"sni_strict": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"urlcert": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"urlcert_group": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"urlcert_hlen": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"use_ciphers_group": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ssl_ciphers_group": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"tls_v10": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"tls_v11": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"tls_v12": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"tls_v13": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ssl_noreg": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ssl_cipher": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ssl_custom_cipher": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"tls13_custom_cipher": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"rfc7919_comply": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"supported_groups": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"tls_pqc_support": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"tls_pqc_groups": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"hsts_header": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"hsts_max_age": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"hsts_include_subdomains": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"hsts_preload": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"hpkp_header": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"client_certificate_forwarding": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"client_certificate_forwarding_sub_header": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"client_certificate_forwarding_cert_header": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"health_check_inherit": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"conn_limit": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"recover": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"warm_up": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"warm_rate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"http2": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"http2_window_size": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"hlck_domain": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"http_https_adaptive": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"http_port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"https_port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"enforce_trust_establishment": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceServerPolicyServerPoolCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	mkey := ""

	t := d.Get("mkey")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing ServerPolicyServerPool: type error")
	}

	obj, err := getObjectServerPolicyServerPool(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating ServerPolicyServerPool resource while getting object: %v", err)
	}

	_, err = c.CreateServerPolicyServerPool(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating ServerPolicyServerPool resource: %v", err)
	}

	d.SetId(mkey)

	return resourceServerPolicyServerPoolRead(d, m)
}

func resourceServerPolicyServerPoolUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectServerPolicyServerPool(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating ServerPolicyServerPool resource while getting object: %v", err)
	}

	_, err = c.UpdateServerPolicyServerPool(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating ServerPolicyServerPool resource: %v", err)
	}

	return resourceServerPolicyServerPoolRead(d, m)
}

func resourceServerPolicyServerPoolDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteServerPolicyServerPool(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting ServerPolicyServerPool resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceServerPolicyServerPoolRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadServerPolicyServerPool(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading ServerPolicyServerPool resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectServerPolicyServerPool(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading ServerPolicyServerPool resource from API: %v", err)
	}

	return nil
}

func flattenServerPolicyServerPool(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectServerPolicyServerPool(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("type", flattenServerPolicyServerPool(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}
	if err = d.Set("protocol", flattenServerPolicyServerPool(o["protocol"], d, "protocol", sv)); err != nil {
		if !fortiAPIPatch(o["protocol"]) {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}
	if err = d.Set("server_balance", flattenServerPolicyServerPool(o["server-balance"], d, "server_balance", sv)); err != nil {
		if !fortiAPIPatch(o["server-balance"]) {
			return fmt.Errorf("Error reading server_balance: %v", err)
		}
	}
	if err = d.Set("health", flattenServerPolicyServerPool(o["health"], d, "health", sv)); err != nil {
		if !fortiAPIPatch(o["health"]) {
			return fmt.Errorf("Error reading health: %v", err)
		}
	}
	if err = d.Set("hlck_sip", flattenServerPolicyServerPool(o["hlck-sip"], d, "hlck_sip", sv)); err != nil {
		if !fortiAPIPatch(o["hlck-sip"]) {
			return fmt.Errorf("Error reading hlck_sip: %v", err)
		}
	}
	if err = d.Set("hlck_sip6", flattenServerPolicyServerPool(o["hlck-sip6"], d, "hlck_sip6", sv)); err != nil {
		if !fortiAPIPatch(o["hlck-sip6"]) {
			return fmt.Errorf("Error reading hlck_sip6: %v", err)
		}
	}
	if err = d.Set("lb_algo", flattenServerPolicyServerPool(o["lb-algo"], d, "lb_algo", sv)); err != nil {
		if !fortiAPIPatch(o["lb-algo"]) {
			return fmt.Errorf("Error reading lb_algo: %v", err)
		}
	}
	if err = d.Set("persistence", flattenServerPolicyServerPool(o["persistence"], d, "persistence", sv)); err != nil {
		if !fortiAPIPatch(o["persistence"]) {
			return fmt.Errorf("Error reading persistence: %v", err)
		}
	}
	if err = d.Set("comment", flattenServerPolicyServerPool(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}
	if err = d.Set("flag", flattenServerPolicyServerPool(o["flag"], d, "flag", sv)); err != nil {
		if !fortiAPIPatch(o["flag"]) {
			return fmt.Errorf("Error reading flag: %v", err)
		}
	}
	if err = d.Set("server_pool_id", flattenServerPolicyServerPool(o["server-pool-id"], d, "server_pool_id", sv)); err != nil {
		if !fortiAPIPatch(o["server-pool-id"]) {
			return fmt.Errorf("Error reading server_pool_id: %v", err)
		}
	}
	if err = d.Set("sub_table_id", flattenServerPolicyServerPool(o["sub_table_id"], d, "sub_table_id", sv)); err != nil {
		if !fortiAPIPatch(o["sub_table_id"]) {
			return fmt.Errorf("Error reading sub_table_id: %v", err)
		}
	}
	if err = d.Set("sub_table_action", flattenServerPolicyServerPool(o["sub_table_action"], d, "sub_table_action", sv)); err != nil {
		if !fortiAPIPatch(o["sub_table_action"]) {
			return fmt.Errorf("Error reading sub_table_action: %v", err)
		}
	}
	if err = d.Set("http_reuse", flattenServerPolicyServerPool(o["http-reuse"], d, "http_reuse", sv)); err != nil {
		if !fortiAPIPatch(o["http-reuse"]) {
			return fmt.Errorf("Error reading http_reuse: %v", err)
		}
	}
	if err = d.Set("reuse_conn_total_time", flattenServerPolicyServerPool(o["reuse-conn-total-time"], d, "reuse_conn_total_time", sv)); err != nil {
		if !fortiAPIPatch(o["reuse-conn-total-time"]) {
			return fmt.Errorf("Error reading reuse_conn_total_time: %v", err)
		}
	}
	if err = d.Set("reuse_conn_idle_time", flattenServerPolicyServerPool(o["reuse-conn-idle-time"], d, "reuse_conn_idle_time", sv)); err != nil {
		if !fortiAPIPatch(o["reuse-conn-idle-time"]) {
			return fmt.Errorf("Error reading reuse_conn_idle_time: %v", err)
		}
	}
	if err = d.Set("reuse_conn_max_request", flattenServerPolicyServerPool(o["reuse-conn-max-request"], d, "reuse_conn_max_request", sv)); err != nil {
		if !fortiAPIPatch(o["reuse-conn-max-request"]) {
			return fmt.Errorf("Error reading reuse_conn_max_request: %v", err)
		}
	}
	if err = d.Set("reuse_conn_max_count", flattenServerPolicyServerPool(o["reuse-conn-max-count"], d, "reuse_conn_max_count", sv)); err != nil {
		if !fortiAPIPatch(o["reuse-conn-max-count"]) {
			return fmt.Errorf("Error reading reuse_conn_max_count: %v", err)
		}
	}
	if err = d.Set("adfs_server_name", flattenServerPolicyServerPool(o["adfs-server-name"], d, "adfs_server_name", sv)); err != nil {
		if !fortiAPIPatch(o["adfs-server-name"]) {
			return fmt.Errorf("Error reading adfs_server_name: %v", err)
		}
	}
	if err = d.Set("panic_mode", flattenServerPolicyServerPool(o["panic-mode"], d, "panic_mode", sv)); err != nil {
		if !fortiAPIPatch(o["panic-mode"]) {
			return fmt.Errorf("Error reading panic_mode: %v", err)
		}
	}
	if err = d.Set("panic_threshold", flattenServerPolicyServerPool(o["panic-threshold"], d, "panic_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["panic-threshold"]) {
			return fmt.Errorf("Error reading panic_threshold: %v", err)
		}
	}
	if err = d.Set("server_type", flattenServerPolicyServerPool(o["server-type"], d, "server_type", sv)); err != nil {
		if !fortiAPIPatch(o["server-type"]) {
			return fmt.Errorf("Error reading server_type: %v", err)
		}
	}
	if err = d.Set("ip", flattenServerPolicyServerPool(o["ip"], d, "ip", sv)); err != nil {
		if !fortiAPIPatch(o["ip"]) {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}
	if err = d.Set("domain", flattenServerPolicyServerPool(o["domain"], d, "domain", sv)); err != nil {
		if !fortiAPIPatch(o["domain"]) {
			return fmt.Errorf("Error reading domain: %v", err)
		}
	}
	if err = d.Set("adfs_username", flattenServerPolicyServerPool(o["adfs-username"], d, "adfs_username", sv)); err != nil {
		if !fortiAPIPatch(o["adfs-username"]) {
			return fmt.Errorf("Error reading adfs_username: %v", err)
		}
	}
	if err = d.Set("adfs_password", flattenServerPolicyServerPool(o["adfs-password"], d, "adfs_password", sv)); err != nil {
		if !fortiAPIPatch(o["adfs-password"]) {
			return fmt.Errorf("Error reading adfs_password: %v", err)
		}
	}
	if err = d.Set("sdn_addr_type", flattenServerPolicyServerPool(o["sdn-addr-type"], d, "sdn_addr_type", sv)); err != nil {
		if !fortiAPIPatch(o["sdn-addr-type"]) {
			return fmt.Errorf("Error reading sdn_addr_type: %v", err)
		}
	}
	if err = d.Set("sdn", flattenServerPolicyServerPool(o["sdn"], d, "sdn", sv)); err != nil {
		if !fortiAPIPatch(o["sdn"]) {
			return fmt.Errorf("Error reading sdn: %v", err)
		}
	}
	if err = d.Set("filter", flattenServerPolicyServerPool(o["filter"], d, "filter", sv)); err != nil {
		if !fortiAPIPatch(o["filter"]) {
			return fmt.Errorf("Error reading filter: %v", err)
		}
	}
	if err = d.Set("port", flattenServerPolicyServerPool(o["port"], d, "port", sv)); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}
	if err = d.Set("weight", flattenServerPolicyServerPool(o["weight"], d, "weight", sv)); err != nil {
		if !fortiAPIPatch(o["weight"]) {
			return fmt.Errorf("Error reading weight: %v", err)
		}
	}
	if err = d.Set("status", flattenServerPolicyServerPool(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}
	if err = d.Set("server_id", flattenServerPolicyServerPool(o["server-id"], d, "server_id", sv)); err != nil {
		if !fortiAPIPatch(o["server-id"]) {
			return fmt.Errorf("Error reading server_id: %v", err)
		}
	}
	if err = d.Set("backup_server", flattenServerPolicyServerPool(o["backup-server"], d, "backup_server", sv)); err != nil {
		if !fortiAPIPatch(o["backup-server"]) {
			return fmt.Errorf("Error reading backup_server: %v", err)
		}
	}
	if err = d.Set("proxy_protocol", flattenServerPolicyServerPool(o["proxy-protocol"], d, "proxy_protocol", sv)); err != nil {
		if !fortiAPIPatch(o["proxy-protocol"]) {
			return fmt.Errorf("Error reading proxy_protocol: %v", err)
		}
	}
	if err = d.Set("proxy_protocol_version", flattenServerPolicyServerPool(o["proxy-protocol-version"], d, "proxy_protocol_version", sv)); err != nil {
		if !fortiAPIPatch(o["proxy-protocol-version"]) {
			return fmt.Errorf("Error reading proxy_protocol_version: %v", err)
		}
	}
	if err = d.Set("ssl", flattenServerPolicyServerPool(o["ssl"], d, "ssl", sv)); err != nil {
		if !fortiAPIPatch(o["ssl"]) {
			return fmt.Errorf("Error reading ssl: %v", err)
		}
	}
	if err = d.Set("implicit_ssl", flattenServerPolicyServerPool(o["implicit_ssl"], d, "implicit_ssl", sv)); err != nil {
		if !fortiAPIPatch(o["implicit_ssl"]) {
			return fmt.Errorf("Error reading implicit_ssl: %v", err)
		}
	}
	if err = d.Set("ssl_quiet_shutdown", flattenServerPolicyServerPool(o["ssl-quiet-shutdown"], d, "ssl_quiet_shutdown", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-quiet-shutdown"]) {
			return fmt.Errorf("Error reading ssl_quiet_shutdown: %v", err)
		}
	}
	if err = d.Set("ssl_session_timeout", flattenServerPolicyServerPool(o["ssl-session-timeout"], d, "ssl_session_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-session-timeout"]) {
			return fmt.Errorf("Error reading ssl_session_timeout: %v", err)
		}
	}
	if err = d.Set("server_side_sni", flattenServerPolicyServerPool(o["server-side-sni"], d, "server_side_sni", sv)); err != nil {
		if !fortiAPIPatch(o["server-side-sni"]) {
			return fmt.Errorf("Error reading server_side_sni: %v", err)
		}
	}
	if err = d.Set("multi_certificate", flattenServerPolicyServerPool(o["multi-certificate"], d, "multi_certificate", sv)); err != nil {
		if !fortiAPIPatch(o["multi-certificate"]) {
			return fmt.Errorf("Error reading multi_certificate: %v", err)
		}
	}
	if err = d.Set("certificate", flattenServerPolicyServerPool(o["certificate"], d, "certificate", sv)); err != nil {
		if !fortiAPIPatch(o["certificate"]) {
			return fmt.Errorf("Error reading certificate: %v", err)
		}
	}
	if err = d.Set("certificate_group", flattenServerPolicyServerPool(o["certificate-group"], d, "certificate_group", sv)); err != nil {
		if !fortiAPIPatch(o["certificate-group"]) {
			return fmt.Errorf("Error reading certificate_group: %v", err)
		}
	}
	if err = d.Set("certificate_type", flattenServerPolicyServerPool(o["certificate-type"], d, "certificate_type", sv)); err != nil {
		if !fortiAPIPatch(o["certificate-type"]) {
			return fmt.Errorf("Error reading certificate_type: %v", err)
		}
	}
	if err = d.Set("lets_certificate", flattenServerPolicyServerPool(o["lets-certificate"], d, "lets_certificate", sv)); err != nil {
		if !fortiAPIPatch(o["lets-certificate"]) {
			return fmt.Errorf("Error reading lets_certificate: %v", err)
		}
	}
	if err = d.Set("intermediate_certificate_group", flattenServerPolicyServerPool(o["intermediate-certificate-group"], d, "intermediate_certificate_group", sv)); err != nil {
		if !fortiAPIPatch(o["intermediate-certificate-group"]) {
			return fmt.Errorf("Error reading intermediate_certificate_group: %v", err)
		}
	}
	if err = d.Set("certificate_verify", flattenServerPolicyServerPool(o["certificate-verify"], d, "certificate_verify", sv)); err != nil {
		if !fortiAPIPatch(o["certificate-verify"]) {
			return fmt.Errorf("Error reading certificate_verify: %v", err)
		}
	}
	if err = d.Set("client_certificate_proxy", flattenServerPolicyServerPool(o["client-certificate-proxy"], d, "client_certificate_proxy", sv)); err != nil {
		if !fortiAPIPatch(o["client-certificate-proxy"]) {
			return fmt.Errorf("Error reading client_certificate_proxy: %v", err)
		}
	}
	if err = d.Set("client_certificate_proxy_sign_ca", flattenServerPolicyServerPool(o["client-certificate-proxy-sign-ca"], d, "client_certificate_proxy_sign_ca", sv)); err != nil {
		if !fortiAPIPatch(o["client-certificate-proxy-sign-ca"]) {
			return fmt.Errorf("Error reading client_certificate_proxy_sign_ca: %v", err)
		}
	}
	if err = d.Set("client_certificate", flattenServerPolicyServerPool(o["client-certificate"], d, "client_certificate", sv)); err != nil {
		if !fortiAPIPatch(o["client-certificate"]) {
			return fmt.Errorf("Error reading client_certificate: %v", err)
		}
	}
	if err = d.Set("server_certificate_verify", flattenServerPolicyServerPool(o["server-certificate-verify"], d, "server_certificate_verify", sv)); err != nil {
		if !fortiAPIPatch(o["server-certificate-verify"]) {
			return fmt.Errorf("Error reading server_certificate_verify: %v", err)
		}
	}
	if err = d.Set("server_certificate_verify_policy", flattenServerPolicyServerPool(o["server-certificate-verify-policy"], d, "server_certificate_verify_policy", sv)); err != nil {
		if !fortiAPIPatch(o["server-certificate-verify-policy"]) {
			return fmt.Errorf("Error reading server_certificate_verify_policy: %v", err)
		}
	}
	if err = d.Set("server_certificate_verify_action", flattenServerPolicyServerPool(o["server-certificate-verify-action"], d, "server_certificate_verify_action", sv)); err != nil {
		if !fortiAPIPatch(o["server-certificate-verify-action"]) {
			return fmt.Errorf("Error reading server_certificate_verify_action: %v", err)
		}
	}
	if err = d.Set("session_ticket_reuse", flattenServerPolicyServerPool(o["session-ticket-reuse"], d, "session_ticket_reuse", sv)); err != nil {
		if !fortiAPIPatch(o["session-ticket-reuse"]) {
			return fmt.Errorf("Error reading session_ticket_reuse: %v", err)
		}
	}
	if err = d.Set("session_id_reuse", flattenServerPolicyServerPool(o["session-id-reuse"], d, "session_id_reuse", sv)); err != nil {
		if !fortiAPIPatch(o["session-id-reuse"]) {
			return fmt.Errorf("Error reading session_id_reuse: %v", err)
		}
	}
	if err = d.Set("sni", flattenServerPolicyServerPool(o["sni"], d, "sni", sv)); err != nil {
		if !fortiAPIPatch(o["sni"]) {
			return fmt.Errorf("Error reading sni: %v", err)
		}
	}
	if err = d.Set("sni_certificate", flattenServerPolicyServerPool(o["sni-certificate"], d, "sni_certificate", sv)); err != nil {
		if !fortiAPIPatch(o["sni-certificate"]) {
			return fmt.Errorf("Error reading sni_certificate: %v", err)
		}
	}
	if err = d.Set("sni_strict", flattenServerPolicyServerPool(o["sni-strict"], d, "sni_strict", sv)); err != nil {
		if !fortiAPIPatch(o["sni-strict"]) {
			return fmt.Errorf("Error reading sni_strict: %v", err)
		}
	}
	if err = d.Set("urlcert", flattenServerPolicyServerPool(o["urlcert"], d, "urlcert", sv)); err != nil {
		if !fortiAPIPatch(o["urlcert"]) {
			return fmt.Errorf("Error reading urlcert: %v", err)
		}
	}
	if err = d.Set("urlcert_group", flattenServerPolicyServerPool(o["urlcert-group"], d, "urlcert_group", sv)); err != nil {
		if !fortiAPIPatch(o["urlcert-group"]) {
			return fmt.Errorf("Error reading urlcert_group: %v", err)
		}
	}
	if err = d.Set("urlcert_hlen", flattenServerPolicyServerPool(o["urlcert-hlen"], d, "urlcert_hlen", sv)); err != nil {
		if !fortiAPIPatch(o["urlcert-hlen"]) {
			return fmt.Errorf("Error reading urlcert_hlen: %v", err)
		}
	}
	if err = d.Set("use_ciphers_group", flattenServerPolicyServerPool(o["use-ciphers-group"], d, "use_ciphers_group", sv)); err != nil {
		if !fortiAPIPatch(o["use-ciphers-group"]) {
			return fmt.Errorf("Error reading use_ciphers_group: %v", err)
		}
	}
	if err = d.Set("ssl_ciphers_group", flattenServerPolicyServerPool(o["ssl-ciphers-group"], d, "ssl_ciphers_group", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-ciphers-group"]) {
			return fmt.Errorf("Error reading ssl_ciphers_group: %v", err)
		}
	}
	if err = d.Set("tls_v10", flattenServerPolicyServerPool(o["tls-v10"], d, "tls_v10", sv)); err != nil {
		if !fortiAPIPatch(o["tls-v10"]) {
			return fmt.Errorf("Error reading tls_v10: %v", err)
		}
	}
	if err = d.Set("tls_v11", flattenServerPolicyServerPool(o["tls-v11"], d, "tls_v11", sv)); err != nil {
		if !fortiAPIPatch(o["tls-v11"]) {
			return fmt.Errorf("Error reading tls_v11: %v", err)
		}
	}
	if err = d.Set("tls_v12", flattenServerPolicyServerPool(o["tls-v12"], d, "tls_v12", sv)); err != nil {
		if !fortiAPIPatch(o["tls-v12"]) {
			return fmt.Errorf("Error reading tls_v12: %v", err)
		}
	}
	if err = d.Set("tls_v13", flattenServerPolicyServerPool(o["tls-v13"], d, "tls_v13", sv)); err != nil {
		if !fortiAPIPatch(o["tls-v13"]) {
			return fmt.Errorf("Error reading tls_v13: %v", err)
		}
	}
	if err = d.Set("ssl_noreg", flattenServerPolicyServerPool(o["ssl-noreg"], d, "ssl_noreg", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-noreg"]) {
			return fmt.Errorf("Error reading ssl_noreg: %v", err)
		}
	}
	if err = d.Set("ssl_cipher", flattenServerPolicyServerPool(o["ssl-cipher"], d, "ssl_cipher", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-cipher"]) {
			return fmt.Errorf("Error reading ssl_cipher: %v", err)
		}
	}
	if err = d.Set("ssl_custom_cipher", flattenServerPolicyServerPool(o["ssl-custom-cipher"], d, "ssl_custom_cipher", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-custom-cipher"]) {
			return fmt.Errorf("Error reading ssl_custom_cipher: %v", err)
		}
	}
	if err = d.Set("tls13_custom_cipher", flattenServerPolicyServerPool(o["tls13-custom-cipher"], d, "tls13_custom_cipher", sv)); err != nil {
		if !fortiAPIPatch(o["tls13-custom-cipher"]) {
			return fmt.Errorf("Error reading tls13_custom_cipher: %v", err)
		}
	}
	if err = d.Set("rfc7919_comply", flattenServerPolicyServerPool(o["rfc7919-comply"], d, "rfc7919_comply", sv)); err != nil {
		if !fortiAPIPatch(o["rfc7919-comply"]) {
			return fmt.Errorf("Error reading rfc7919_comply: %v", err)
		}
	}
	if err = d.Set("supported_groups", flattenServerPolicyServerPool(o["supported-groups"], d, "supported_groups", sv)); err != nil {
		if !fortiAPIPatch(o["supported-groups"]) {
			return fmt.Errorf("Error reading supported_groups: %v", err)
		}
	}
	if err = d.Set("tls_pqc_support", flattenServerPolicyServerPool(o["tls-pqc-support"], d, "tls_pqc_support", sv)); err != nil {
		if !fortiAPIPatch(o["tls-pqc-support"]) {
			return fmt.Errorf("Error reading tls_pqc_support: %v", err)
		}
	}
	if err = d.Set("tls_pqc_groups", flattenServerPolicyServerPool(o["tls-pqc-groups"], d, "tls_pqc_groups", sv)); err != nil {
		if !fortiAPIPatch(o["tls-pqc-groups"]) {
			return fmt.Errorf("Error reading tls_pqc_groups: %v", err)
		}
	}
	if err = d.Set("hsts_header", flattenServerPolicyServerPool(o["hsts-header"], d, "hsts_header", sv)); err != nil {
		if !fortiAPIPatch(o["hsts-header"]) {
			return fmt.Errorf("Error reading hsts_header: %v", err)
		}
	}
	if err = d.Set("hsts_max_age", flattenServerPolicyServerPool(o["hsts-max-age"], d, "hsts_max_age", sv)); err != nil {
		if !fortiAPIPatch(o["hsts-max-age"]) {
			return fmt.Errorf("Error reading hsts_max_age: %v", err)
		}
	}
	if err = d.Set("hsts_include_subdomains", flattenServerPolicyServerPool(o["hsts-include-subdomains"], d, "hsts_include_subdomains", sv)); err != nil {
		if !fortiAPIPatch(o["hsts-include-subdomains"]) {
			return fmt.Errorf("Error reading hsts_include_subdomains: %v", err)
		}
	}
	if err = d.Set("hsts_preload", flattenServerPolicyServerPool(o["hsts-preload"], d, "hsts_preload", sv)); err != nil {
		if !fortiAPIPatch(o["hsts-preload"]) {
			return fmt.Errorf("Error reading hsts_preload: %v", err)
		}
	}
	if err = d.Set("hpkp_header", flattenServerPolicyServerPool(o["hpkp-header"], d, "hpkp_header", sv)); err != nil {
		if !fortiAPIPatch(o["hpkp-header"]) {
			return fmt.Errorf("Error reading hpkp_header: %v", err)
		}
	}
	if err = d.Set("client_certificate_forwarding", flattenServerPolicyServerPool(o["client-certificate-forwarding"], d, "client_certificate_forwarding", sv)); err != nil {
		if !fortiAPIPatch(o["client-certificate-forwarding"]) {
			return fmt.Errorf("Error reading client_certificate_forwarding: %v", err)
		}
	}
	if err = d.Set("client_certificate_forwarding_sub_header", flattenServerPolicyServerPool(o["client-certificate-forwarding-sub-header"], d, "client_certificate_forwarding_sub_header", sv)); err != nil {
		if !fortiAPIPatch(o["client-certificate-forwarding-sub-header"]) {
			return fmt.Errorf("Error reading client_certificate_forwarding_sub_header: %v", err)
		}
	}
	if err = d.Set("client_certificate_forwarding_cert_header", flattenServerPolicyServerPool(o["client-certificate-forwarding-cert-header"], d, "client_certificate_forwarding_cert_header", sv)); err != nil {
		if !fortiAPIPatch(o["client-certificate-forwarding-cert-header"]) {
			return fmt.Errorf("Error reading client_certificate_forwarding_cert_header: %v", err)
		}
	}
	if err = d.Set("health_check_inherit", flattenServerPolicyServerPool(o["health-check-inherit"], d, "health_check_inherit", sv)); err != nil {
		if !fortiAPIPatch(o["health-check-inherit"]) {
			return fmt.Errorf("Error reading health_check_inherit: %v", err)
		}
	}
	if err = d.Set("conn_limit", flattenServerPolicyServerPool(o["conn-limit"], d, "conn_limit", sv)); err != nil {
		if !fortiAPIPatch(o["conn-limit"]) {
			return fmt.Errorf("Error reading conn_limit: %v", err)
		}
	}
	if err = d.Set("recover", flattenServerPolicyServerPool(o["recover"], d, "recover", sv)); err != nil {
		if !fortiAPIPatch(o["recover"]) {
			return fmt.Errorf("Error reading recover: %v", err)
		}
	}
	if err = d.Set("warm_up", flattenServerPolicyServerPool(o["warm-up"], d, "warm_up", sv)); err != nil {
		if !fortiAPIPatch(o["warm-up"]) {
			return fmt.Errorf("Error reading warm_up: %v", err)
		}
	}
	if err = d.Set("warm_rate", flattenServerPolicyServerPool(o["warm-rate"], d, "warm_rate", sv)); err != nil {
		if !fortiAPIPatch(o["warm-rate"]) {
			return fmt.Errorf("Error reading warm_rate: %v", err)
		}
	}
	if err = d.Set("http2", flattenServerPolicyServerPool(o["http2"], d, "http2", sv)); err != nil {
		if !fortiAPIPatch(o["http2"]) {
			return fmt.Errorf("Error reading http2: %v", err)
		}
	}
	if err = d.Set("http2_window_size", flattenServerPolicyServerPool(o["http2-window-size"], d, "http2_window_size", sv)); err != nil {
		if !fortiAPIPatch(o["http2-window-size"]) {
			return fmt.Errorf("Error reading http2_window_size: %v", err)
		}
	}
	if err = d.Set("hlck_domain", flattenServerPolicyServerPool(o["hlck-domain"], d, "hlck_domain", sv)); err != nil {
		if !fortiAPIPatch(o["hlck-domain"]) {
			return fmt.Errorf("Error reading hlck_domain: %v", err)
		}
	}
	if err = d.Set("http_https_adaptive", flattenServerPolicyServerPool(o["http-https-adaptive"], d, "http_https_adaptive", sv)); err != nil {
		if !fortiAPIPatch(o["http-https-adaptive"]) {
			return fmt.Errorf("Error reading http_https_adaptive: %v", err)
		}
	}
	if err = d.Set("http_port", flattenServerPolicyServerPool(o["http-port"], d, "http_port", sv)); err != nil {
		if !fortiAPIPatch(o["http-port"]) {
			return fmt.Errorf("Error reading http_port: %v", err)
		}
	}
	if err = d.Set("https_port", flattenServerPolicyServerPool(o["https-port"], d, "https_port", sv)); err != nil {
		if !fortiAPIPatch(o["https-port"]) {
			return fmt.Errorf("Error reading https_port: %v", err)
		}
	}
	if err = d.Set("enforce_trust_establishment", flattenServerPolicyServerPool(o["enforce-trust-establishment"], d, "enforce_trust_establishment", sv)); err != nil {
		if !fortiAPIPatch(o["enforce-trust-establishment"]) {
			return fmt.Errorf("Error reading enforce_trust_establishment: %v", err)
		}
	}
	if err = d.Set("mkey", flattenServerPolicyServerPool(o["name"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandServerPolicyServerPool(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectServerPolicyServerPool(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("type"); ok {
		t, err := expandServerPolicyServerPool(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}
	if v, ok := d.GetOk("protocol"); ok {
		t, err := expandServerPolicyServerPool(d, v, "protocol", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}
	if v, ok := d.GetOk("server_balance"); ok {
		t, err := expandServerPolicyServerPool(d, v, "server_balance", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-balance"] = t
		}
	}
	if v, ok := d.GetOk("health"); ok {
		t, err := expandServerPolicyServerPool(d, v, "health", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["health"] = t
		}
	}
	if v, ok := d.GetOk("hlck_sip"); ok {
		t, err := expandServerPolicyServerPool(d, v, "hlck_sip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hlck-sip"] = t
		}
	}
	if v, ok := d.GetOk("hlck_sip6"); ok {
		t, err := expandServerPolicyServerPool(d, v, "hlck_sip6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hlck-sip6"] = t
		}
	}
	if v, ok := d.GetOk("lb_algo"); ok {
		t, err := expandServerPolicyServerPool(d, v, "lb_algo", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lb-algo"] = t
		}
	}
	if v, ok := d.GetOk("persistence"); ok {
		t, err := expandServerPolicyServerPool(d, v, "persistence", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["persistence"] = t
		}
	}
	if v, ok := d.GetOk("comment"); ok {
		t, err := expandServerPolicyServerPool(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}
	if v, ok := d.GetOk("flag"); ok {
		t, err := expandServerPolicyServerPool(d, v, "flag", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["flag"] = t
		}
	}
	if v, ok := d.GetOk("server_pool_id"); ok {
		t, err := expandServerPolicyServerPool(d, v, "server_pool_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-pool-id"] = t
		}
	}
	if v, ok := d.GetOk("sub_table_id"); ok {
		t, err := expandServerPolicyServerPool(d, v, "sub_table_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sub_table_id"] = t
		}
	}
	if v, ok := d.GetOk("sub_table_action"); ok {
		t, err := expandServerPolicyServerPool(d, v, "sub_table_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sub_table_action"] = t
		}
	}
	if v, ok := d.GetOk("http_reuse"); ok {
		t, err := expandServerPolicyServerPool(d, v, "http_reuse", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-reuse"] = t
		}
	}
	if v, ok := d.GetOk("reuse_conn_total_time"); ok {
		t, err := expandServerPolicyServerPool(d, v, "reuse_conn_total_time", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reuse-conn-total-time"] = t
		}
	}
	if v, ok := d.GetOk("reuse_conn_idle_time"); ok {
		t, err := expandServerPolicyServerPool(d, v, "reuse_conn_idle_time", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reuse-conn-idle-time"] = t
		}
	}
	if v, ok := d.GetOk("reuse_conn_max_request"); ok {
		t, err := expandServerPolicyServerPool(d, v, "reuse_conn_max_request", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reuse-conn-max-request"] = t
		}
	}
	if v, ok := d.GetOk("reuse_conn_max_count"); ok {
		t, err := expandServerPolicyServerPool(d, v, "reuse_conn_max_count", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reuse-conn-max-count"] = t
		}
	}
	if v, ok := d.GetOk("adfs_server_name"); ok {
		t, err := expandServerPolicyServerPool(d, v, "adfs_server_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adfs-server-name"] = t
		}
	}
	if v, ok := d.GetOk("panic_mode"); ok {
		t, err := expandServerPolicyServerPool(d, v, "panic_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["panic-mode"] = t
		}
	}
	if v, ok := d.GetOk("panic_threshold"); ok {
		t, err := expandServerPolicyServerPool(d, v, "panic_threshold", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["panic-threshold"] = t
		}
	}
	if v, ok := d.GetOk("server_type"); ok {
		t, err := expandServerPolicyServerPool(d, v, "server_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-type"] = t
		}
	}
	if v, ok := d.GetOk("ip"); ok {
		t, err := expandServerPolicyServerPool(d, v, "ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}
	if v, ok := d.GetOk("domain"); ok {
		t, err := expandServerPolicyServerPool(d, v, "domain", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain"] = t
		}
	}
	if v, ok := d.GetOk("adfs_username"); ok {
		t, err := expandServerPolicyServerPool(d, v, "adfs_username", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adfs-username"] = t
		}
	}
	if v, ok := d.GetOk("adfs_password"); ok {
		t, err := expandServerPolicyServerPool(d, v, "adfs_password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adfs-password"] = t
		}
	}
	if v, ok := d.GetOk("sdn_addr_type"); ok {
		t, err := expandServerPolicyServerPool(d, v, "sdn_addr_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdn-addr-type"] = t
		}
	}
	if v, ok := d.GetOk("sdn"); ok {
		t, err := expandServerPolicyServerPool(d, v, "sdn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdn"] = t
		}
	}
	if v, ok := d.GetOk("filter"); ok {
		t, err := expandServerPolicyServerPool(d, v, "filter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter"] = t
		}
	}
	if v, ok := d.GetOk("port"); ok {
		t, err := expandServerPolicyServerPool(d, v, "port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}
	if v, ok := d.GetOk("weight"); ok {
		t, err := expandServerPolicyServerPool(d, v, "weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight"] = t
		}
	}
	if v, ok := d.GetOk("status"); ok {
		t, err := expandServerPolicyServerPool(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}
	if v, ok := d.GetOk("server_id"); ok {
		t, err := expandServerPolicyServerPool(d, v, "server_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-id"] = t
		}
	}
	if v, ok := d.GetOk("backup_server"); ok {
		t, err := expandServerPolicyServerPool(d, v, "backup_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["backup-server"] = t
		}
	}
	if v, ok := d.GetOk("proxy_protocol"); ok {
		t, err := expandServerPolicyServerPool(d, v, "proxy_protocol", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-protocol"] = t
		}
	}
	if v, ok := d.GetOk("proxy_protocol_version"); ok {
		t, err := expandServerPolicyServerPool(d, v, "proxy_protocol_version", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-protocol-version"] = t
		}
	}
	if v, ok := d.GetOk("ssl"); ok {
		t, err := expandServerPolicyServerPool(d, v, "ssl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl"] = t
		}
	}
	if v, ok := d.GetOk("implicit_ssl"); ok {
		t, err := expandServerPolicyServerPool(d, v, "implicit_ssl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["implicit_ssl"] = t
		}
	}
	if v, ok := d.GetOk("ssl_quiet_shutdown"); ok {
		t, err := expandServerPolicyServerPool(d, v, "ssl_quiet_shutdown", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-quiet-shutdown"] = t
		}
	}
	if v, ok := d.GetOk("ssl_session_timeout"); ok {
		t, err := expandServerPolicyServerPool(d, v, "ssl_session_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-session-timeout"] = t
		}
	}
	if v, ok := d.GetOk("server_side_sni"); ok {
		t, err := expandServerPolicyServerPool(d, v, "server_side_sni", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-side-sni"] = t
		}
	}
	if v, ok := d.GetOk("multi_certificate"); ok {
		t, err := expandServerPolicyServerPool(d, v, "multi_certificate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multi-certificate"] = t
		}
	}
	if v, ok := d.GetOk("certificate"); ok {
		t, err := expandServerPolicyServerPool(d, v, "certificate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certificate"] = t
		}
	}
	if v, ok := d.GetOk("certificate_group"); ok {
		t, err := expandServerPolicyServerPool(d, v, "certificate_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certificate-group"] = t
		}
	}
	if v, ok := d.GetOk("certificate_type"); ok {
		t, err := expandServerPolicyServerPool(d, v, "certificate_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certificate-type"] = t
		}
	}
	if v, ok := d.GetOk("lets_certificate"); ok {
		t, err := expandServerPolicyServerPool(d, v, "lets_certificate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lets-certificate"] = t
		}
	}
	if v, ok := d.GetOk("intermediate_certificate_group"); ok {
		t, err := expandServerPolicyServerPool(d, v, "intermediate_certificate_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["intermediate-certificate-group"] = t
		}
	}
	if v, ok := d.GetOk("certificate_verify"); ok {
		t, err := expandServerPolicyServerPool(d, v, "certificate_verify", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certificate-verify"] = t
		}
	}
	if v, ok := d.GetOk("client_certificate_proxy"); ok {
		t, err := expandServerPolicyServerPool(d, v, "client_certificate_proxy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-certificate-proxy"] = t
		}
	}
	if v, ok := d.GetOk("client_certificate_proxy_sign_ca"); ok {
		t, err := expandServerPolicyServerPool(d, v, "client_certificate_proxy_sign_ca", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-certificate-proxy-sign-ca"] = t
		}
	}
	if v, ok := d.GetOk("client_certificate"); ok {
		t, err := expandServerPolicyServerPool(d, v, "client_certificate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-certificate"] = t
		}
	}
	if v, ok := d.GetOk("server_certificate_verify"); ok {
		t, err := expandServerPolicyServerPool(d, v, "server_certificate_verify", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-certificate-verify"] = t
		}
	}
	if v, ok := d.GetOk("server_certificate_verify_policy"); ok {
		t, err := expandServerPolicyServerPool(d, v, "server_certificate_verify_policy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-certificate-verify-policy"] = t
		}
	}
	if v, ok := d.GetOk("server_certificate_verify_action"); ok {
		t, err := expandServerPolicyServerPool(d, v, "server_certificate_verify_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-certificate-verify-action"] = t
		}
	}
	if v, ok := d.GetOk("session_ticket_reuse"); ok {
		t, err := expandServerPolicyServerPool(d, v, "session_ticket_reuse", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session-ticket-reuse"] = t
		}
	}
	if v, ok := d.GetOk("session_id_reuse"); ok {
		t, err := expandServerPolicyServerPool(d, v, "session_id_reuse", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session-id-reuse"] = t
		}
	}
	if v, ok := d.GetOk("sni"); ok {
		t, err := expandServerPolicyServerPool(d, v, "sni", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sni"] = t
		}
	}
	if v, ok := d.GetOk("sni_certificate"); ok {
		t, err := expandServerPolicyServerPool(d, v, "sni_certificate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sni-certificate"] = t
		}
	}
	if v, ok := d.GetOk("sni_strict"); ok {
		t, err := expandServerPolicyServerPool(d, v, "sni_strict", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sni-strict"] = t
		}
	}
	if v, ok := d.GetOk("urlcert"); ok {
		t, err := expandServerPolicyServerPool(d, v, "urlcert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["urlcert"] = t
		}
	}
	if v, ok := d.GetOk("urlcert_group"); ok {
		t, err := expandServerPolicyServerPool(d, v, "urlcert_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["urlcert-group"] = t
		}
	}
	if v, ok := d.GetOk("urlcert_hlen"); ok {
		t, err := expandServerPolicyServerPool(d, v, "urlcert_hlen", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["urlcert-hlen"] = t
		}
	}
	if v, ok := d.GetOk("use_ciphers_group"); ok {
		t, err := expandServerPolicyServerPool(d, v, "use_ciphers_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["use-ciphers-group"] = t
		}
	}
	if v, ok := d.GetOk("ssl_ciphers_group"); ok {
		t, err := expandServerPolicyServerPool(d, v, "ssl_ciphers_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-ciphers-group"] = t
		}
	}
	if v, ok := d.GetOk("tls_v10"); ok {
		t, err := expandServerPolicyServerPool(d, v, "tls_v10", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tls-v10"] = t
		}
	}
	if v, ok := d.GetOk("tls_v11"); ok {
		t, err := expandServerPolicyServerPool(d, v, "tls_v11", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tls-v11"] = t
		}
	}
	if v, ok := d.GetOk("tls_v12"); ok {
		t, err := expandServerPolicyServerPool(d, v, "tls_v12", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tls-v12"] = t
		}
	}
	if v, ok := d.GetOk("tls_v13"); ok {
		t, err := expandServerPolicyServerPool(d, v, "tls_v13", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tls-v13"] = t
		}
	}
	if v, ok := d.GetOk("ssl_noreg"); ok {
		t, err := expandServerPolicyServerPool(d, v, "ssl_noreg", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-noreg"] = t
		}
	}
	if v, ok := d.GetOk("ssl_cipher"); ok {
		t, err := expandServerPolicyServerPool(d, v, "ssl_cipher", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-cipher"] = t
		}
	}
	if v, ok := d.GetOk("ssl_custom_cipher"); ok {
		t, err := expandServerPolicyServerPool(d, v, "ssl_custom_cipher", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-custom-cipher"] = t
		}
	}
	if v, ok := d.GetOk("tls13_custom_cipher"); ok {
		t, err := expandServerPolicyServerPool(d, v, "tls13_custom_cipher", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tls13-custom-cipher"] = t
		}
	}
	if v, ok := d.GetOk("rfc7919_comply"); ok {
		t, err := expandServerPolicyServerPool(d, v, "rfc7919_comply", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rfc7919-comply"] = t
		}
	}
	if v, ok := d.GetOk("supported_groups"); ok {
		t, err := expandServerPolicyServerPool(d, v, "supported_groups", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["supported-groups"] = t
		}
	}
	if v, ok := d.GetOk("tls_pqc_support"); ok {
		t, err := expandServerPolicyServerPool(d, v, "tls_pqc_support", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tls-pqc-support"] = t
		}
	}
	if v, ok := d.GetOk("tls_pqc_groups"); ok {
		t, err := expandServerPolicyServerPool(d, v, "tls_pqc_groups", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tls-pqc-groups"] = t
		}
	}
	if v, ok := d.GetOk("hsts_header"); ok {
		t, err := expandServerPolicyServerPool(d, v, "hsts_header", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hsts-header"] = t
		}
	}
	if v, ok := d.GetOk("hsts_max_age"); ok {
		t, err := expandServerPolicyServerPool(d, v, "hsts_max_age", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hsts-max-age"] = t
		}
	}
	if v, ok := d.GetOk("hsts_include_subdomains"); ok {
		t, err := expandServerPolicyServerPool(d, v, "hsts_include_subdomains", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hsts-include-subdomains"] = t
		}
	}
	if v, ok := d.GetOk("hsts_preload"); ok {
		t, err := expandServerPolicyServerPool(d, v, "hsts_preload", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hsts-preload"] = t
		}
	}
	if v, ok := d.GetOk("hpkp_header"); ok {
		t, err := expandServerPolicyServerPool(d, v, "hpkp_header", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hpkp-header"] = t
		}
	}
	if v, ok := d.GetOk("client_certificate_forwarding"); ok {
		t, err := expandServerPolicyServerPool(d, v, "client_certificate_forwarding", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-certificate-forwarding"] = t
		}
	}
	if v, ok := d.GetOk("client_certificate_forwarding_sub_header"); ok {
		t, err := expandServerPolicyServerPool(d, v, "client_certificate_forwarding_sub_header", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-certificate-forwarding-sub-header"] = t
		}
	}
	if v, ok := d.GetOk("client_certificate_forwarding_cert_header"); ok {
		t, err := expandServerPolicyServerPool(d, v, "client_certificate_forwarding_cert_header", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-certificate-forwarding-cert-header"] = t
		}
	}
	if v, ok := d.GetOk("health_check_inherit"); ok {
		t, err := expandServerPolicyServerPool(d, v, "health_check_inherit", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["health-check-inherit"] = t
		}
	}
	if v, ok := d.GetOk("conn_limit"); ok {
		t, err := expandServerPolicyServerPool(d, v, "conn_limit", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["conn-limit"] = t
		}
	}
	if v, ok := d.GetOk("recover"); ok {
		t, err := expandServerPolicyServerPool(d, v, "recover", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["recover"] = t
		}
	}
	if v, ok := d.GetOk("warm_up"); ok {
		t, err := expandServerPolicyServerPool(d, v, "warm_up", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["warm-up"] = t
		}
	}
	if v, ok := d.GetOk("warm_rate"); ok {
		t, err := expandServerPolicyServerPool(d, v, "warm_rate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["warm-rate"] = t
		}
	}
	if v, ok := d.GetOk("http2"); ok {
		t, err := expandServerPolicyServerPool(d, v, "http2", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http2"] = t
		}
	}
	if v, ok := d.GetOk("http2_window_size"); ok {
		t, err := expandServerPolicyServerPool(d, v, "http2_window_size", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http2-window-size"] = t
		}
	}
	if v, ok := d.GetOk("hlck_domain"); ok {
		t, err := expandServerPolicyServerPool(d, v, "hlck_domain", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hlck-domain"] = t
		}
	}
	if v, ok := d.GetOk("http_https_adaptive"); ok {
		t, err := expandServerPolicyServerPool(d, v, "http_https_adaptive", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-https-adaptive"] = t
		}
	}
	if v, ok := d.GetOk("http_port"); ok {
		t, err := expandServerPolicyServerPool(d, v, "http_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-port"] = t
		}
	}
	if v, ok := d.GetOk("https_port"); ok {
		t, err := expandServerPolicyServerPool(d, v, "https_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["https-port"] = t
		}
	}
	if v, ok := d.GetOk("enforce_trust_establishment"); ok {
		t, err := expandServerPolicyServerPool(d, v, "enforce_trust_establishment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enforce-trust-establishment"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandServerPolicyServerPool(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
