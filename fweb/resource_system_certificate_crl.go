// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure system/certificate.crl.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCertificateCrl() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemCertificateCrlRead,
		Update: resourceSystemCertificateCrlUpdate,
		Create: resourceSystemCertificateCrlCreate,
		Delete: resourceSystemCertificateCrlDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"httpurl": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"scepurl": &schema.Schema{
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

func resourceSystemCertificateCrlCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateCrl: type error")
	}

	obj, err := getObjectSystemCertificateCrl(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateCrl resource while getting object: %v", err)
	}

	_, err = c.CreateSystemCertificateCrl(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateCrl resource: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func resourceSystemCertificateCrlUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceSystemCertificateCrlRead(d, m)
}

func resourceSystemCertificateCrlDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSystemCertificateCrl(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCertificateCrl resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCertificateCrlRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemCertificateCrl(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateCrl resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCertificateCrl(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateCrl resource from API: %v", err)
	}

	return nil
}

func flattenSystemCertificateCrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemCertificateCrl(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenSystemCertificateCrl(o["name"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemCertificateCrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCertificateCrl(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("httpurl"); ok {
		t, err := expandSystemCertificateCrl(d, v, "httpurl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["httpUrl"] = t
		}
	}
	if v, ok := d.GetOk("scepurl"); ok {
		t, err := expandSystemCertificateCrl(d, v, "scepurl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scepUrl"] = t
		}
	}
	if v, ok := d.GetOk("type"); ok {
		t, err := expandSystemCertificateCrl(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}
	if v, ok := d.GetOk("srcfile"); ok {
		t, err := expandSystemCertificateCrl(d, v, "srcfile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uploadedFile"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemCertificateCrl(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
