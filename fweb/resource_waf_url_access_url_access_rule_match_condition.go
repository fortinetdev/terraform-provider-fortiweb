// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure waf/url_access.url_access_rule/match-condition.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWafUrlAccessUrlAccessRuleMatchCondition() *schema.Resource {
	return &schema.Resource{
		Read:   resourceWafUrlAccessUrlAccessRuleMatchConditionRead,
		Update: resourceWafUrlAccessUrlAccessRuleMatchConditionUpdate,
		Create: resourceWafUrlAccessUrlAccessRuleMatchConditionCreate,
		Delete: resourceWafUrlAccessUrlAccessRuleMatchConditionDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"sub_mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"sip_address_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"sip_address_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"sip_address_value": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"sdomain_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"sip_address_domain": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"source_domain_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"source_domain": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"reverse_dns_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"reg_exp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"url_access_parameter": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"only_method_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"only_protocol_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"only_method": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"only_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"reverse_match": &schema.Schema{
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

func resourceWafUrlAccessUrlAccessRuleMatchConditionCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafUrlAccessUrlAccessRuleMatchCondition: type error")
	}

	obj, err := getObjectWafUrlAccessUrlAccessRuleMatchCondition(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WafUrlAccessUrlAccessRuleMatchCondition resource while getting object: %v", err)
	}

	_, err = c.CreateWafUrlAccessUrlAccessRuleMatchCondition(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error creating WafUrlAccessUrlAccessRuleMatchCondition resource: %v", err)
	}

	d.SetId(mkey)

	return resourceWafUrlAccessUrlAccessRuleMatchConditionRead(d, m)
}

func resourceWafUrlAccessUrlAccessRuleMatchConditionUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	sub_mkey := "1"
	t := d.Get("sub_mkey")
	if v, ok := t.(string); ok {
		sub_mkey = v
	} else {
		return fmt.Errorf("Error describing WafUrlAccessUrlAccessRuleMatchCondition: type error")
	}

	obj, err := getObjectWafUrlAccessUrlAccessRuleMatchCondition(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WafUrlAccessUrlAccessRuleMatchCondition resource while getting object: %v", err)
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	_, err = c.UpdateWafUrlAccessUrlAccessRuleMatchCondition(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating WafUrlAccessUrlAccessRuleMatchCondition resource: %v", err)
	}

	return resourceWafUrlAccessUrlAccessRuleMatchConditionRead(d, m)
}

func resourceWafUrlAccessUrlAccessRuleMatchConditionDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	sub_mkey := "1"
	t := d.Get("sub_mkey")
	if v, ok := t.(string); ok {
		sub_mkey = v
	} else {
		return fmt.Errorf("Error describing WafUrlAccessUrlAccessRuleMatchCondition: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	err := c.DeleteWafUrlAccessUrlAccessRuleMatchCondition(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting WafUrlAccessUrlAccessRuleMatchCondition resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWafUrlAccessUrlAccessRuleMatchConditionRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	sub_mkey := "1"
	t := d.Get("sub_mkey")
	if v, ok := t.(string); ok {
		sub_mkey = v
	} else {
		return fmt.Errorf("Error describing WafUrlAccessUrlAccessRuleMatchCondition: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	o, err := c.ReadWafUrlAccessUrlAccessRuleMatchCondition(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading WafUrlAccessUrlAccessRuleMatchCondition resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWafUrlAccessUrlAccessRuleMatchCondition(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WafUrlAccessUrlAccessRuleMatchCondition resource from API: %v", err)
	}

	return nil
}

func flattenWafUrlAccessUrlAccessRuleMatchCondition(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWafUrlAccessUrlAccessRuleMatchCondition(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("sub_mkey", flattenWafUrlAccessUrlAccessRuleMatchCondition(o["id"], d, "sub_mkey", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading sub_mkey: %v", err)
		}
	}
	if err = d.Set("reg_exp", flattenWafUrlAccessUrlAccessRuleMatchCondition(o["reg-exp"], d, "reg_exp", sv)); err != nil {
		if !fortiAPIPatch(o["reg-exp"]) {
			return fmt.Errorf("Error reading reg_exp: %v", err)
		}
	}
	if err = d.Set("type", flattenWafUrlAccessUrlAccessRuleMatchCondition(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}
	if err = d.Set("reverse_match", flattenWafUrlAccessUrlAccessRuleMatchCondition(o["reverse-match"], d, "reverse_match", sv)); err != nil {
		if !fortiAPIPatch(o["reverse-match"]) {
			return fmt.Errorf("Error reading reverse_match: %v", err)
		}
	}
	if err = d.Set("sip_address_check", flattenWafUrlAccessUrlAccessRuleMatchCondition(o["sip-address-check"], d, "sip_address_check", sv)); err != nil {
		if !fortiAPIPatch(o["sip-address-check"]) {
			return fmt.Errorf("Error reading sip_address_check: %v", err)
		}
	}
	if err = d.Set("sip_address_type", flattenWafUrlAccessUrlAccessRuleMatchCondition(o["sip-address-type"], d, "sip_address_type", sv)); err != nil {
		if !fortiAPIPatch(o["sip-address-type"]) {
			return fmt.Errorf("Error reading sip_address_type: %v", err)
		}
	}
	if err = d.Set("reverse_dns_timeout", flattenWafUrlAccessUrlAccessRuleMatchCondition(o["reverse-dns-timeout"], d, "reverse_dns_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["reverse-dns-timeout"]) {
			return fmt.Errorf("Error reading reverse_dns_timeout: %v", err)
		}
	}
	if err = d.Set("sip_address_value", flattenWafUrlAccessUrlAccessRuleMatchCondition(o["sip-address-value"], d, "sip_address_value", sv)); err != nil {
		if !fortiAPIPatch(o["sip-address-value"]) {
			return fmt.Errorf("Error reading sip_address_value: %v", err)
		}
	}
	if err = d.Set("sip_address_domain", flattenWafUrlAccessUrlAccessRuleMatchCondition(o["sip-address-domain"], d, "sip_address_domain", sv)); err != nil {
		if !fortiAPIPatch(o["sip-address-domain"]) {
			return fmt.Errorf("Error reading sip_address_domain: %v", err)
		}
	}
	if err = d.Set("sdomain_type", flattenWafUrlAccessUrlAccessRuleMatchCondition(o["sdomain-type"], d, "sdomain_type", sv)); err != nil {
		if !fortiAPIPatch(o["sdomain-type"]) {
			return fmt.Errorf("Error reading sdomain_type: %v", err)
		}
	}
	if err = d.Set("source_domain", flattenWafUrlAccessUrlAccessRuleMatchCondition(o["source-domain"], d, "source_domain", sv)); err != nil {
		if !fortiAPIPatch(o["source-domain"]) {
			return fmt.Errorf("Error reading source_domain: %v", err)
		}
	}
	if err = d.Set("source_domain_type", flattenWafUrlAccessUrlAccessRuleMatchCondition(o["source-domain-type"], d, "source_domain_type", sv)); err != nil {
		if !fortiAPIPatch(o["source-domain-type"]) {
			return fmt.Errorf("Error reading source_domain_type: %v", err)
		}
	}
	if err = d.Set("url_access_parameter", flattenWafUrlAccessUrlAccessRuleMatchCondition(o["url-access-parameter"], d, "url_access_parameter", sv)); err != nil {
		if !fortiAPIPatch(o["url-access-parameter"]) {
			return fmt.Errorf("Error reading url_access_parameter: %v", err)
		}
	}
	if err = d.Set("only_method_check", flattenWafUrlAccessUrlAccessRuleMatchCondition(o["only-method-check"], d, "only_method_check", sv)); err != nil {
		if !fortiAPIPatch(o["only-method-check"]) {
			return fmt.Errorf("Error reading only_method_check: %v", err)
		}
	}
	if err = d.Set("only_protocol_check", flattenWafUrlAccessUrlAccessRuleMatchCondition(o["only-protocol-check"], d, "only_protocol_check", sv)); err != nil {
		if !fortiAPIPatch(o["only-protocol-check"]) {
			return fmt.Errorf("Error reading only_protocol_check: %v", err)
		}
	}
	if err = d.Set("only_method", flattenWafUrlAccessUrlAccessRuleMatchCondition(o["only-method"], d, "only_method", sv)); err != nil {
		if !fortiAPIPatch(o["only-method"]) {
			return fmt.Errorf("Error reading only_method: %v", err)
		}
	}
	if err = d.Set("only_protocol", flattenWafUrlAccessUrlAccessRuleMatchCondition(o["only-protocol"], d, "only_protocol", sv)); err != nil {
		if !fortiAPIPatch(o["only-protocol"]) {
			return fmt.Errorf("Error reading only_protocol: %v", err)
		}
	}
	if err = d.Set("mkey", flattenWafUrlAccessUrlAccessRuleMatchCondition(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandWafUrlAccessUrlAccessRuleMatchCondition(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWafUrlAccessUrlAccessRuleMatchCondition(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("sub_mkey"); ok {
		t, err := expandWafUrlAccessUrlAccessRuleMatchCondition(d, v, "sub_mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}
	if v, ok := d.GetOk("reg_exp"); ok {
		t, err := expandWafUrlAccessUrlAccessRuleMatchCondition(d, v, "reg_exp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reg-exp"] = t
		}
	}
	if v, ok := d.GetOk("type"); ok {
		t, err := expandWafUrlAccessUrlAccessRuleMatchCondition(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}
	if v, ok := d.GetOk("reverse_match"); ok {
		t, err := expandWafUrlAccessUrlAccessRuleMatchCondition(d, v, "reverse_match", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reverse-match"] = t
		}
	}
	if v, ok := d.GetOk("sip_address_check"); ok {
		t, err := expandWafUrlAccessUrlAccessRuleMatchCondition(d, v, "sip_address_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sip-address-check"] = t
		}
	}
	if v, ok := d.GetOk("sip_address_type"); ok {
		t, err := expandWafUrlAccessUrlAccessRuleMatchCondition(d, v, "sip_address_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sip-address-type"] = t
		}
	}
	if v, ok := d.GetOk("reverse_dns_timeout"); ok {
		t, err := expandWafUrlAccessUrlAccessRuleMatchCondition(d, v, "reverse_dns_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reverse-dns-timeout"] = t
		}
	}
	if v, ok := d.GetOk("sip_address_value"); ok {
		t, err := expandWafUrlAccessUrlAccessRuleMatchCondition(d, v, "sip_address_value", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sip-address-value"] = t
		}
	}
	if v, ok := d.GetOk("sip_address_domain"); ok {
		t, err := expandWafUrlAccessUrlAccessRuleMatchCondition(d, v, "sip_address_domain", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sip-address-domain"] = t
		}
	}
	if v, ok := d.GetOk("sdomain_type"); ok {
		t, err := expandWafUrlAccessUrlAccessRuleMatchCondition(d, v, "sdomain_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdomain-type"] = t
		}
	}
	if v, ok := d.GetOk("source_domain"); ok {
		t, err := expandWafUrlAccessUrlAccessRuleMatchCondition(d, v, "source_domain", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-domain"] = t
		}
	}
	if v, ok := d.GetOk("source_domain_type"); ok {
		t, err := expandWafUrlAccessUrlAccessRuleMatchCondition(d, v, "source_domain_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-domain-type"] = t
		}
	}
	if v, ok := d.GetOk("url_access_parameter"); ok {
		t, err := expandWafUrlAccessUrlAccessRuleMatchCondition(d, v, "url_access_parameter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-access-parameter"] = t
		}
	}
	if v, ok := d.GetOk("only_method_check"); ok {
		t, err := expandWafUrlAccessUrlAccessRuleMatchCondition(d, v, "only_method_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["only-method-check"] = t
		}
	}
	if v, ok := d.GetOk("only_protocol_check"); ok {
		t, err := expandWafUrlAccessUrlAccessRuleMatchCondition(d, v, "only_protocol_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["only-protocol-check"] = t
		}
	}
	if v, ok := d.GetOk("only_method"); ok {
		t, err := expandWafUrlAccessUrlAccessRuleMatchCondition(d, v, "only_method", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["only-method"] = t
		}
	}
	if v, ok := d.GetOk("only_protocol"); ok {
		t, err := expandWafUrlAccessUrlAccessRuleMatchCondition(d, v, "only_protocol", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["only-protocol"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandWafUrlAccessUrlAccessRuleMatchCondition(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
