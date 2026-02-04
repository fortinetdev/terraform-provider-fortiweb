// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure system/certificate.tsl_ca.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCertificateTslCa() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemCertificateTslCaRead,
		Update: resourceSystemCertificateTslCaUpdate,
		Create: resourceSystemCertificateTslCaCreate,
		Delete: resourceSystemCertificateTslCaDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"srcfile": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
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

func resourceSystemCertificateTslCaCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateTslCa: type error")
	}

	obj, err := getObjectSystemCertificateTslCa(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateTslCa resource while getting object: %v", err)
	}

	_, err = c.CreateSystemCertificateTslCa(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateTslCa resource: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func resourceSystemCertificateTslCaUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceSystemCertificateTslCaRead(d, m)
}

func resourceSystemCertificateTslCaDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSystemCertificateTslCa(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCertificateTslCa resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCertificateTslCaRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemCertificateTslCa(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateTslCa resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCertificateTslCa(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateTslCa resource from API: %v", err)
	}

	return nil
}

func flattenSystemCertificateTslCa(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemCertificateTslCa(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenSystemCertificateTslCa(o["name"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemCertificateTslCa(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCertificateTslCa(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("url"); ok {
		t, err := expandSystemCertificateTslCa(d, v, "url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url"] = t
		}
	}
	if v, ok := d.GetOk("srcfile"); ok {
		t, err := expandSystemCertificateTslCa(d, v, "srcfile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcfile"] = t
		}
	}
	if v, ok := d.GetOk("type"); ok {
		t, err := expandSystemCertificateTslCa(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemCertificateTslCa(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
