// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure system/maintenance.systemtime.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemMaintenanceSystemTime() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemMaintenanceSystemTimeRead,
		Update: resourceSystemMaintenanceSystemTimeUpdate,
		Create: resourceSystemMaintenanceSystemTimeUpdate,
		Delete: resourceSystemMaintenanceSystemTimeDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"daylightsaving": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"timezone": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"systemtime": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"time": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mode": &schema.Schema{
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

func resourceSystemMaintenanceSystemTimeCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemMaintenanceSystemTime: type error")
	}

	obj, err := getObjectSystemMaintenanceSystemTime(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemMaintenanceSystemTime resource while getting object: %v", err)
	}

	_, err = c.CreateSystemMaintenanceSystemTime(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemMaintenanceSystemTime resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemMaintenanceSystemTimeRead(d, m)
}

func resourceSystemMaintenanceSystemTimeUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := "SystemMaintenanceSystemTime"
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemMaintenanceSystemTime(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemMaintenanceSystemTime resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemMaintenanceSystemTime(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemMaintenanceSystemTime resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemMaintenanceSystemTimeRead(d, m)
}

func resourceSystemMaintenanceSystemTimeDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemMaintenanceSystemTime(d, c.Fv)
	(*obj)["daylightSaving"] = 0
	(*obj)["timeZone"] = 4
	if err != nil {
		return fmt.Errorf("Error updating SystemMaintenanceSystemTime resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemMaintenanceSystemTime(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemMaintenanceSystemTime resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemMaintenanceSystemTimeRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemMaintenanceSystemTime(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemMaintenanceSystemTime resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemMaintenanceSystemTime(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemMaintenanceSystemTime resource from API: %v", err)
	}

	return nil
}

func flattenSystemMaintenanceSystemTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemMaintenanceSystemTime(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("daylightsaving", flattenSystemMaintenanceSystemTime(o["daylightSaving"], d, "daylightsaving", sv)); err != nil {
		if !fortiAPIPatch(o["daylightSaving"]) {
			return fmt.Errorf("Error reading daylightsaving: %v", err)
		}
	}
	if err = d.Set("timezone", flattenSystemMaintenanceSystemTime(o["timeZone"], d, "timezone", sv)); err != nil {
		if !fortiAPIPatch(o["timeZone"]) {
			return fmt.Errorf("Error reading timezone: %v", err)
		}
	}
	if err = d.Set("systemtime", flattenSystemMaintenanceSystemTime(o["systemTime"], d, "systemtime", sv)); err != nil {
		if !fortiAPIPatch(o["systemTime"]) {
			return fmt.Errorf("Error reading systemtime: %v", err)
		}
	}
	if err = d.Set("time", flattenSystemMaintenanceSystemTime(o["time"], d, "time", sv)); err != nil {
		if !fortiAPIPatch(o["time"]) {
			return fmt.Errorf("Error reading time: %v", err)
		}
	}
	if err = d.Set("mode", flattenSystemMaintenanceSystemTime(o["mode"], d, "mode", sv)); err != nil {
		if !fortiAPIPatch(o["mode"]) {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	return nil
}

func expandSystemMaintenanceSystemTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemMaintenanceSystemTime(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("daylightsaving"); ok {
		t, err := expandSystemMaintenanceSystemTime(d, v, "daylightsaving", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["daylightSaving"] = t
		}
	}
	if v, ok := d.GetOk("timezone"); ok {
		t, err := expandSystemMaintenanceSystemTime(d, v, "timezone", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timeZone"] = t
		}
	}
	if v, ok := d.GetOk("systemtime"); ok {
		t, err := expandSystemMaintenanceSystemTime(d, v, "systemtime", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["systemTime"] = t
		}
	}
	if v, ok := d.GetOk("time"); ok {
		t, err := expandSystemMaintenanceSystemTime(d, v, "time", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["time"] = t
		}
	}
	if v, ok := d.GetOk("mode"); ok {
		t, err := expandSystemMaintenanceSystemTime(d, v, "mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	return &obj, nil
}
