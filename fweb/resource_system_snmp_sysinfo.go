// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure system/snmp.sysinfo.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemSnmpSysinfo() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemSnmpSysinfoRead,
		Update: resourceSystemSnmpSysinfoUpdate,
		Create: resourceSystemSnmpSysinfoUpdate,
		Delete: resourceSystemSnmpSysinfoDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"engine_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"contact_info": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"location": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSystemSnmpSysinfoCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemSnmpSysinfo: type error")
	}

	obj, err := getObjectSystemSnmpSysinfo(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemSnmpSysinfo resource while getting object: %v", err)
	}

	_, err = c.CreateSystemSnmpSysinfo(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemSnmpSysinfo resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemSnmpSysinfoRead(d, m)
}

func resourceSystemSnmpSysinfoUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := "SnmpSysinfo"
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemSnmpSysinfo(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemSnmpSysinfo resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSnmpSysinfo(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemSnmpSysinfo resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemSnmpSysinfoRead(d, m)
}

func resourceSystemSnmpSysinfoDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemSnmpSysinfo(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemSnmpSysinfo resource while getting object: %v", err)
	}

	(*obj)["status"] = "disable"
	(*obj)["description"] = ""
	(*obj)["location"] = ""
	(*obj)["contact-info"] = ""

	_, err = c.UpdateSystemSnmpSysinfo(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemSnmpSysinfo resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSnmpSysinfoRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemSnmpSysinfo(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemSnmpSysinfo resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSnmpSysinfo(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemSnmpSysinfo resource from API: %v", err)
	}

	return nil
}

func flattenSystemSnmpSysinfo(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemSnmpSysinfo(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenSystemSnmpSysinfo(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}
	if err = d.Set("engine_id", flattenSystemSnmpSysinfo(o["engine-id"], d, "engine_id", sv)); err != nil {
		if !fortiAPIPatch(o["engine-id"]) {
			return fmt.Errorf("Error reading engine_id: %v", err)
		}
	}
	if err = d.Set("description", flattenSystemSnmpSysinfo(o["description"], d, "description", sv)); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}
	if err = d.Set("contact_info", flattenSystemSnmpSysinfo(o["contact-info"], d, "contact_info", sv)); err != nil {
		if !fortiAPIPatch(o["contact-info"]) {
			return fmt.Errorf("Error reading contact_info: %v", err)
		}
	}
	if err = d.Set("location", flattenSystemSnmpSysinfo(o["location"], d, "location", sv)); err != nil {
		if !fortiAPIPatch(o["location"]) {
			return fmt.Errorf("Error reading location: %v", err)
		}
	}
	if err = d.Set("mkey", flattenSystemSnmpSysinfo(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemSnmpSysinfo(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSnmpSysinfo(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		t, err := expandSystemSnmpSysinfo(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}
	if v, ok := d.GetOk("engine_id"); ok {
		t, err := expandSystemSnmpSysinfo(d, v, "engine_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["engine-id"] = t
		}
	}
	if v, ok := d.GetOk("description"); ok {
		t, err := expandSystemSnmpSysinfo(d, v, "description", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}
	if v, ok := d.GetOk("contact_info"); ok {
		t, err := expandSystemSnmpSysinfo(d, v, "contact_info", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["contact-info"] = t
		}
	}
	if v, ok := d.GetOk("location"); ok {
		t, err := expandSystemSnmpSysinfo(d, v, "location", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["location"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemSnmpSysinfo(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
