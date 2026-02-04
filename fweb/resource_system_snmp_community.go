// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure system/snmp.community.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemSnmpCommunity() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemSnmpCommunityRead,
		Update: resourceSystemSnmpCommunityUpdate,
		Create: resourceSystemSnmpCommunityCreate,
		Delete: resourceSystemSnmpCommunityDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"query_v1_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"query_v1_port": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"query_v2c_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"query_v2c_port": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"trap_v1_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"trap_v1_lport": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"trap_v1_rport": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"trap_v2c_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"trap_v2c_lport": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"trap_v2c_rport": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"events": &schema.Schema{
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

func resourceSystemSnmpCommunityCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemSnmpCommunity: type error")
	}

	obj, err := getObjectSystemSnmpCommunity(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemSnmpCommunity resource while getting object: %v", err)
	}

	_, err = c.CreateSystemSnmpCommunity(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemSnmpCommunity resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemSnmpCommunityRead(d, m)
}

func resourceSystemSnmpCommunityUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemSnmpCommunity(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemSnmpCommunity resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSnmpCommunity(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemSnmpCommunity resource: %v", err)
	}

	return resourceSystemSnmpCommunityRead(d, m)
}

func resourceSystemSnmpCommunityDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSystemSnmpCommunity(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSnmpCommunity resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSnmpCommunityRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemSnmpCommunity(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemSnmpCommunity resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSnmpCommunity(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemSnmpCommunity resource from API: %v", err)
	}

	return nil
}

func flattenSystemSnmpCommunity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemSnmpCommunity(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenSystemSnmpCommunity(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}
	if err = d.Set("status", flattenSystemSnmpCommunity(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}
	if err = d.Set("query_v1_status", flattenSystemSnmpCommunity(o["query-v1-status"], d, "query_v1_status", sv)); err != nil {
		if !fortiAPIPatch(o["query-v1-status"]) {
			return fmt.Errorf("Error reading query_v1_status: %v", err)
		}
	}
	if err = d.Set("query_v1_port", flattenSystemSnmpCommunity(o["query-v1-port"], d, "query_v1_port", sv)); err != nil {
		if !fortiAPIPatch(o["query-v1-port"]) {
			return fmt.Errorf("Error reading query_v1_port: %v", err)
		}
	}
	if err = d.Set("query_v2c_status", flattenSystemSnmpCommunity(o["query-v2c-status"], d, "query_v2c_status", sv)); err != nil {
		if !fortiAPIPatch(o["query-v2c-status"]) {
			return fmt.Errorf("Error reading query_v2c_status: %v", err)
		}
	}
	if err = d.Set("query_v2c_port", flattenSystemSnmpCommunity(o["query-v2c-port"], d, "query_v2c_port", sv)); err != nil {
		if !fortiAPIPatch(o["query-v2c-port"]) {
			return fmt.Errorf("Error reading query_v2c_port: %v", err)
		}
	}
	if err = d.Set("trap_v1_status", flattenSystemSnmpCommunity(o["trap-v1-status"], d, "trap_v1_status", sv)); err != nil {
		if !fortiAPIPatch(o["trap-v1-status"]) {
			return fmt.Errorf("Error reading trap_v1_status: %v", err)
		}
	}
	if err = d.Set("trap_v1_lport", flattenSystemSnmpCommunity(o["trap-v1-lport"], d, "trap_v1_lport", sv)); err != nil {
		if !fortiAPIPatch(o["trap-v1-lport"]) {
			return fmt.Errorf("Error reading trap_v1_lport: %v", err)
		}
	}
	if err = d.Set("trap_v1_rport", flattenSystemSnmpCommunity(o["trap-v1-rport"], d, "trap_v1_rport", sv)); err != nil {
		if !fortiAPIPatch(o["trap-v1-rport"]) {
			return fmt.Errorf("Error reading trap_v1_rport: %v", err)
		}
	}
	if err = d.Set("trap_v2c_status", flattenSystemSnmpCommunity(o["trap-v2c-status"], d, "trap_v2c_status", sv)); err != nil {
		if !fortiAPIPatch(o["trap-v2c-status"]) {
			return fmt.Errorf("Error reading trap_v2c_status: %v", err)
		}
	}
	if err = d.Set("trap_v2c_lport", flattenSystemSnmpCommunity(o["trap-v2c-lport"], d, "trap_v2c_lport", sv)); err != nil {
		if !fortiAPIPatch(o["trap-v2c-lport"]) {
			return fmt.Errorf("Error reading trap_v2c_lport: %v", err)
		}
	}
	if err = d.Set("trap_v2c_rport", flattenSystemSnmpCommunity(o["trap-v2c-rport"], d, "trap_v2c_rport", sv)); err != nil {
		if !fortiAPIPatch(o["trap-v2c-rport"]) {
			return fmt.Errorf("Error reading trap_v2c_rport: %v", err)
		}
	}
	if err = d.Set("events", flattenSystemSnmpCommunity(o["events"], d, "events", sv)); err != nil {
		if !fortiAPIPatch(o["events"]) {
			return fmt.Errorf("Error reading events: %v", err)
		}
	}
	if err = d.Set("mkey", flattenSystemSnmpCommunity(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemSnmpCommunity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSnmpCommunity(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemSnmpCommunity(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}
	if v, ok := d.GetOk("status"); ok {
		t, err := expandSystemSnmpCommunity(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}
	if v, ok := d.GetOk("query_v1_status"); ok {
		t, err := expandSystemSnmpCommunity(d, v, "query_v1_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["query-v1-status"] = t
		}
	}
	if v, ok := d.GetOk("query_v1_port"); ok {
		t, err := expandSystemSnmpCommunity(d, v, "query_v1_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["query-v1-port"] = t
		}
	}
	if v, ok := d.GetOk("query_v2c_status"); ok {
		t, err := expandSystemSnmpCommunity(d, v, "query_v2c_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["query-v2c-status"] = t
		}
	}
	if v, ok := d.GetOk("query_v2c_port"); ok {
		t, err := expandSystemSnmpCommunity(d, v, "query_v2c_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["query-v2c-port"] = t
		}
	}
	if v, ok := d.GetOk("trap_v1_status"); ok {
		t, err := expandSystemSnmpCommunity(d, v, "trap_v1_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-v1-status"] = t
		}
	}
	if v, ok := d.GetOk("trap_v1_lport"); ok {
		t, err := expandSystemSnmpCommunity(d, v, "trap_v1_lport", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-v1-lport"] = t
		}
	}
	if v, ok := d.GetOk("trap_v1_rport"); ok {
		t, err := expandSystemSnmpCommunity(d, v, "trap_v1_rport", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-v1-rport"] = t
		}
	}
	if v, ok := d.GetOk("trap_v2c_status"); ok {
		t, err := expandSystemSnmpCommunity(d, v, "trap_v2c_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-v2c-status"] = t
		}
	}
	if v, ok := d.GetOk("trap_v2c_lport"); ok {
		t, err := expandSystemSnmpCommunity(d, v, "trap_v2c_lport", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-v2c-lport"] = t
		}
	}
	if v, ok := d.GetOk("trap_v2c_rport"); ok {
		t, err := expandSystemSnmpCommunity(d, v, "trap_v2c_rport", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-v2c-rport"] = t
		}
	}
	if v, ok := d.GetOk("events"); ok {
		t, err := expandSystemSnmpCommunity(d, v, "events", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["events"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemSnmpCommunity(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
