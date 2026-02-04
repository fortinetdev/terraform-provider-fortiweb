// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure system/accprofile.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemAccprofile() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemAccprofileRead,
		Update: resourceSystemAccprofileUpdate,
		Create: resourceSystemAccprofileCreate,
		Delete: resourceSystemAccprofileDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"mntgrp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"admingrp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"sysgrp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"netgrp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"loggrp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"authusergrp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"traroutegrp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"wafgrp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"wadgrp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"wvsgrp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mlgrp": &schema.Schema{
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

func resourceSystemAccprofileCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemAccprofile: type error")
	}

	obj, err := getObjectSystemAccprofile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemAccprofile resource while getting object: %v", err)
	}

	_, err = c.CreateSystemAccprofile(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemAccprofile resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemAccprofileRead(d, m)
}

func resourceSystemAccprofileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemAccprofile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemAccprofile resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemAccprofile(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemAccprofile resource: %v", err)
	}

	return resourceSystemAccprofileRead(d, m)
}

func resourceSystemAccprofileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSystemAccprofile(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAccprofile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAccprofileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemAccprofile(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemAccprofile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAccprofile(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemAccprofile resource from API: %v", err)
	}

	return nil
}

func flattenSystemAccprofile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemAccprofile(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mntgrp", flattenSystemAccprofile(o["mntgrp"], d, "mntgrp", sv)); err != nil {
		if !fortiAPIPatch(o["mntgrp"]) {
			return fmt.Errorf("Error reading mntgrp: %v", err)
		}
	}
	if err = d.Set("admingrp", flattenSystemAccprofile(o["admingrp"], d, "admingrp", sv)); err != nil {
		if !fortiAPIPatch(o["admingrp"]) {
			return fmt.Errorf("Error reading admingrp: %v", err)
		}
	}
	if err = d.Set("sysgrp", flattenSystemAccprofile(o["sysgrp"], d, "sysgrp", sv)); err != nil {
		if !fortiAPIPatch(o["sysgrp"]) {
			return fmt.Errorf("Error reading sysgrp: %v", err)
		}
	}
	if err = d.Set("netgrp", flattenSystemAccprofile(o["netgrp"], d, "netgrp", sv)); err != nil {
		if !fortiAPIPatch(o["netgrp"]) {
			return fmt.Errorf("Error reading netgrp: %v", err)
		}
	}
	if err = d.Set("loggrp", flattenSystemAccprofile(o["loggrp"], d, "loggrp", sv)); err != nil {
		if !fortiAPIPatch(o["loggrp"]) {
			return fmt.Errorf("Error reading loggrp: %v", err)
		}
	}
	if err = d.Set("authusergrp", flattenSystemAccprofile(o["authusergrp"], d, "authusergrp", sv)); err != nil {
		if !fortiAPIPatch(o["authusergrp"]) {
			return fmt.Errorf("Error reading authusergrp: %v", err)
		}
	}
	if err = d.Set("traroutegrp", flattenSystemAccprofile(o["traroutegrp"], d, "traroutegrp", sv)); err != nil {
		if !fortiAPIPatch(o["traroutegrp"]) {
			return fmt.Errorf("Error reading traroutegrp: %v", err)
		}
	}
	if err = d.Set("wafgrp", flattenSystemAccprofile(o["wafgrp"], d, "wafgrp", sv)); err != nil {
		if !fortiAPIPatch(o["wafgrp"]) {
			return fmt.Errorf("Error reading wafgrp: %v", err)
		}
	}
	if err = d.Set("wadgrp", flattenSystemAccprofile(o["wadgrp"], d, "wadgrp", sv)); err != nil {
		if !fortiAPIPatch(o["wadgrp"]) {
			return fmt.Errorf("Error reading wadgrp: %v", err)
		}
	}
	if err = d.Set("wvsgrp", flattenSystemAccprofile(o["wvsgrp"], d, "wvsgrp", sv)); err != nil {
		if !fortiAPIPatch(o["wvsgrp"]) {
			return fmt.Errorf("Error reading wvsgrp: %v", err)
		}
	}
	if err = d.Set("mlgrp", flattenSystemAccprofile(o["mlgrp"], d, "mlgrp", sv)); err != nil {
		if !fortiAPIPatch(o["mlgrp"]) {
			return fmt.Errorf("Error reading mlgrp: %v", err)
		}
	}
	if err = d.Set("mkey", flattenSystemAccprofile(o["name"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemAccprofile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAccprofile(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mntgrp"); ok {
		t, err := expandSystemAccprofile(d, v, "mntgrp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mntgrp"] = t
		}
	}
	if v, ok := d.GetOk("admingrp"); ok {
		t, err := expandSystemAccprofile(d, v, "admingrp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admingrp"] = t
		}
	}
	if v, ok := d.GetOk("sysgrp"); ok {
		t, err := expandSystemAccprofile(d, v, "sysgrp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sysgrp"] = t
		}
	}
	if v, ok := d.GetOk("netgrp"); ok {
		t, err := expandSystemAccprofile(d, v, "netgrp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["netgrp"] = t
		}
	}
	if v, ok := d.GetOk("loggrp"); ok {
		t, err := expandSystemAccprofile(d, v, "loggrp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["loggrp"] = t
		}
	}
	if v, ok := d.GetOk("authusergrp"); ok {
		t, err := expandSystemAccprofile(d, v, "authusergrp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authusergrp"] = t
		}
	}
	if v, ok := d.GetOk("traroutegrp"); ok {
		t, err := expandSystemAccprofile(d, v, "traroutegrp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traroutegrp"] = t
		}
	}
	if v, ok := d.GetOk("wafgrp"); ok {
		t, err := expandSystemAccprofile(d, v, "wafgrp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wafgrp"] = t
		}
	}
	if v, ok := d.GetOk("wadgrp"); ok {
		t, err := expandSystemAccprofile(d, v, "wadgrp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wadgrp"] = t
		}
	}
	if v, ok := d.GetOk("wvsgrp"); ok {
		t, err := expandSystemAccprofile(d, v, "wvsgrp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wvsgrp"] = t
		}
	}
	if v, ok := d.GetOk("mlgrp"); ok {
		t, err := expandSystemAccprofile(d, v, "mlgrp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mlgrp"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemAccprofile(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
