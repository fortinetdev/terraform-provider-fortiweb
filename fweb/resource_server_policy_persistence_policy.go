// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure server_policy/persistence_policy.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceServerPolicyPersistencePolicy() *schema.Resource {
	return &schema.Resource{
		Read:   resourceServerPolicyPersistencePolicyRead,
		Update: resourceServerPolicyPersistencePolicyUpdate,
		Create: resourceServerPolicyPersistencePolicyCreate,
		Delete: resourceServerPolicyPersistencePolicyDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"cookie_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ipv4_netmask": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ipv6_mask_length": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"http_header": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"url_parameter": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"cookie_path": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"cookie_domain": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"secure_cookie": &schema.Schema{
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

func resourceServerPolicyPersistencePolicyCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing ServerPolicyPersistencePolicy: type error")
	}

	obj, err := getObjectServerPolicyPersistencePolicy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating ServerPolicyPersistencePolicy resource while getting object: %v", err)
	}

	_, err = c.CreateServerPolicyPersistencePolicy(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating ServerPolicyPersistencePolicy resource: %v", err)
	}

	d.SetId(mkey)

	return resourceServerPolicyPersistencePolicyRead(d, m)
}

func resourceServerPolicyPersistencePolicyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectServerPolicyPersistencePolicy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating ServerPolicyPersistencePolicy resource while getting object: %v", err)
	}

	_, err = c.UpdateServerPolicyPersistencePolicy(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating ServerPolicyPersistencePolicy resource: %v", err)
	}

	return resourceServerPolicyPersistencePolicyRead(d, m)
}

func resourceServerPolicyPersistencePolicyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteServerPolicyPersistencePolicy(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting ServerPolicyPersistencePolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceServerPolicyPersistencePolicyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadServerPolicyPersistencePolicy(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading ServerPolicyPersistencePolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectServerPolicyPersistencePolicy(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading ServerPolicyPersistencePolicy resource from API: %v", err)
	}

	return nil
}

func flattenServerPolicyPersistencePolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectServerPolicyPersistencePolicy(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("type", flattenServerPolicyPersistencePolicy(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}
	if err = d.Set("cookie_name", flattenServerPolicyPersistencePolicy(o["cookie-name"], d, "cookie_name", sv)); err != nil {
		if !fortiAPIPatch(o["cookie-name"]) {
			return fmt.Errorf("Error reading cookie_name: %v", err)
		}
	}
	if err = d.Set("timeout", flattenServerPolicyPersistencePolicy(o["timeout"], d, "timeout", sv)); err != nil {
		if !fortiAPIPatch(o["timeout"]) {
			return fmt.Errorf("Error reading timeout: %v", err)
		}
	}
	if err = d.Set("uuid", flattenServerPolicyPersistencePolicy(o["uuid"], d, "uuid", sv)); err != nil {
		if !fortiAPIPatch(o["uuid"]) {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}
	if err = d.Set("ipv4_netmask", flattenServerPolicyPersistencePolicy(o["ipv4-netmask"], d, "ipv4_netmask", sv)); err != nil {
		if !fortiAPIPatch(o["ipv4-netmask"]) {
			return fmt.Errorf("Error reading ipv4_netmask: %v", err)
		}
	}
	if err = d.Set("ipv6_mask_length", flattenServerPolicyPersistencePolicy(o["ipv6-mask-length"], d, "ipv6_mask_length", sv)); err != nil {
		if !fortiAPIPatch(o["ipv6-mask-length"]) {
			return fmt.Errorf("Error reading ipv6_mask_length: %v", err)
		}
	}
	if err = d.Set("http_header", flattenServerPolicyPersistencePolicy(o["http-header"], d, "http_header", sv)); err != nil {
		if !fortiAPIPatch(o["http-header"]) {
			return fmt.Errorf("Error reading http_header: %v", err)
		}
	}
	if err = d.Set("url_parameter", flattenServerPolicyPersistencePolicy(o["url-parameter"], d, "url_parameter", sv)); err != nil {
		if !fortiAPIPatch(o["url-parameter"]) {
			return fmt.Errorf("Error reading url_parameter: %v", err)
		}
	}
	if err = d.Set("cookie_path", flattenServerPolicyPersistencePolicy(o["cookie-path"], d, "cookie_path", sv)); err != nil {
		if !fortiAPIPatch(o["cookie-path"]) {
			return fmt.Errorf("Error reading cookie_path: %v", err)
		}
	}
	if err = d.Set("cookie_domain", flattenServerPolicyPersistencePolicy(o["cookie-domain"], d, "cookie_domain", sv)); err != nil {
		if !fortiAPIPatch(o["cookie-domain"]) {
			return fmt.Errorf("Error reading cookie_domain: %v", err)
		}
	}
	if err = d.Set("secure_cookie", flattenServerPolicyPersistencePolicy(o["secure-cookie"], d, "secure_cookie", sv)); err != nil {
		if !fortiAPIPatch(o["secure-cookie"]) {
			return fmt.Errorf("Error reading secure_cookie: %v", err)
		}
	}
	if err = d.Set("mkey", flattenServerPolicyPersistencePolicy(o["name"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandServerPolicyPersistencePolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectServerPolicyPersistencePolicy(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("type"); ok {
		t, err := expandServerPolicyPersistencePolicy(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}
	if v, ok := d.GetOk("cookie_name"); ok {
		t, err := expandServerPolicyPersistencePolicy(d, v, "cookie_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cookie-name"] = t
		}
	}
	if v, ok := d.GetOk("timeout"); ok {
		t, err := expandServerPolicyPersistencePolicy(d, v, "timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timeout"] = t
		}
	}
	if v, ok := d.GetOk("uuid"); ok {
		t, err := expandServerPolicyPersistencePolicy(d, v, "uuid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}
	if v, ok := d.GetOk("ipv4_netmask"); ok {
		t, err := expandServerPolicyPersistencePolicy(d, v, "ipv4_netmask", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-netmask"] = t
		}
	}
	if v, ok := d.GetOk("ipv6_mask_length"); ok {
		t, err := expandServerPolicyPersistencePolicy(d, v, "ipv6_mask_length", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-mask-length"] = t
		}
	}
	if v, ok := d.GetOk("http_header"); ok {
		t, err := expandServerPolicyPersistencePolicy(d, v, "http_header", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-header"] = t
		}
	}
	if v, ok := d.GetOk("url_parameter"); ok {
		t, err := expandServerPolicyPersistencePolicy(d, v, "url_parameter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-parameter"] = t
		}
	}
	if v, ok := d.GetOk("cookie_path"); ok {
		t, err := expandServerPolicyPersistencePolicy(d, v, "cookie_path", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cookie-path"] = t
		}
	}
	if v, ok := d.GetOk("cookie_domain"); ok {
		t, err := expandServerPolicyPersistencePolicy(d, v, "cookie_domain", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cookie-domain"] = t
		}
	}
	if v, ok := d.GetOk("secure_cookie"); ok {
		t, err := expandServerPolicyPersistencePolicy(d, v, "secure_cookie", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secure-cookie"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandServerPolicyPersistencePolicy(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
