// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure system/ntp.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemNtp() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemNtpRead,
		Update: resourceSystemNtpUpdate,
		Create: resourceSystemNtpUpdate,
		Delete: resourceSystemNtpDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"ntpsync": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"syncinterval": &schema.Schema{
				Type:     schema.TypeInt,
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

func resourceSystemNtpCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemNtp: type error")
	}

	obj, err := getObjectSystemNtp(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemNtp resource while getting object: %v", err)
	}

	_, err = c.CreateSystemNtp(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemNtp resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemNtpRead(d, m)
}

func resourceSystemNtpUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := "SystemNtp"
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemNtp(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemNtp resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemNtp(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemNtp resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemNtpRead(d, m)
}

func resourceSystemNtpDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemNtp(d, c.Fv)
	(*obj)["ntpsync"] = "enable"
	(*obj)["syncinterval"] = 60
	if err != nil {
		return fmt.Errorf("Error updating SystemNtp resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemNtp(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemNtp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemNtpRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemNtp(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemNtp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemNtp(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemNtp resource from API: %v", err)
	}

	return nil
}

func flattenSystemNtp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemNtp(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("ntpsync", flattenSystemNtp(o["ntpsync"], d, "ntpsync", sv)); err != nil {
		if !fortiAPIPatch(o["ntpsync"]) {
			return fmt.Errorf("Error reading ntpsync: %v", err)
		}
	}
	if err = d.Set("syncinterval", flattenSystemNtp(o["syncinterval"], d, "syncinterval", sv)); err != nil {
		if !fortiAPIPatch(o["syncinterval"]) {
			return fmt.Errorf("Error reading syncinterval: %v", err)
		}
	}

	return nil
}

func expandSystemNtp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemNtp(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ntpsync"); ok {
		t, err := expandSystemNtp(d, v, "ntpsync", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ntpsync"] = t
		}
	}
	if v, ok := d.GetOk("syncinterval"); ok {
		t, err := expandSystemNtp(d, v, "syncinterval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["syncinterval"] = t
		}
	}

	return &obj, nil
}
