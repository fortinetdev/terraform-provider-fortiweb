// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure system/certificate.hpkp.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCertificateHpkp() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemCertificateHpkpRead,
		Update: resourceSystemCertificateHpkpUpdate,
		Create: resourceSystemCertificateHpkpCreate,
		Delete: resourceSystemCertificateHpkpDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"pin_sha256": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_age": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"subdomains": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"report_uri": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"report_only": &schema.Schema{
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

func resourceSystemCertificateHpkpCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateHpkp: type error")
	}

	obj, err := getObjectSystemCertificateHpkp(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateHpkp resource while getting object: %v", err)
	}

	_, err = c.CreateSystemCertificateHpkp(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateHpkp resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemCertificateHpkpRead(d, m)
}

func resourceSystemCertificateHpkpUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemCertificateHpkp(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateHpkp resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemCertificateHpkp(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateHpkp resource: %v", err)
	}

	return resourceSystemCertificateHpkpRead(d, m)
}

func resourceSystemCertificateHpkpDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSystemCertificateHpkp(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCertificateHpkp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCertificateHpkpRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemCertificateHpkp(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateHpkp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCertificateHpkp(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateHpkp resource from API: %v", err)
	}

	return nil
}

func flattenSystemCertificateHpkp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemCertificateHpkp(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("pin_sha256", flattenSystemCertificateHpkp(o["pin-sha256"], d, "pin_sha256", sv)); err != nil {
		if !fortiAPIPatch(o["pin-sha256"]) {
			return fmt.Errorf("Error reading pin_sha256: %v", err)
		}
	}
	if err = d.Set("max_age", flattenSystemCertificateHpkp(o["max-age"], d, "max_age", sv)); err != nil {
		if !fortiAPIPatch(o["max-age"]) {
			return fmt.Errorf("Error reading max_age: %v", err)
		}
	}
	if err = d.Set("subdomains", flattenSystemCertificateHpkp(o["subdomains"], d, "subdomains", sv)); err != nil {
		if !fortiAPIPatch(o["subdomains"]) {
			return fmt.Errorf("Error reading subdomains: %v", err)
		}
	}
	if err = d.Set("report_uri", flattenSystemCertificateHpkp(o["report-uri"], d, "report_uri", sv)); err != nil {
		if !fortiAPIPatch(o["report-uri"]) {
			return fmt.Errorf("Error reading report_uri: %v", err)
		}
	}
	if err = d.Set("report_only", flattenSystemCertificateHpkp(o["report-only"], d, "report_only", sv)); err != nil {
		if !fortiAPIPatch(o["report-only"]) {
			return fmt.Errorf("Error reading report_only: %v", err)
		}
	}
	if err = d.Set("mkey", flattenSystemCertificateHpkp(o["name"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemCertificateHpkp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCertificateHpkp(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("pin_sha256"); ok {
		t, err := expandSystemCertificateHpkp(d, v, "pin_sha256", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pin-sha256"] = t
		}
	}
	if v, ok := d.GetOk("max_age"); ok {
		t, err := expandSystemCertificateHpkp(d, v, "max_age", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-age"] = t
		}
	}
	if v, ok := d.GetOk("subdomains"); ok {
		t, err := expandSystemCertificateHpkp(d, v, "subdomains", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subdomains"] = t
		}
	}
	if v, ok := d.GetOk("report_uri"); ok {
		t, err := expandSystemCertificateHpkp(d, v, "report_uri", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["report-uri"] = t
		}
	}
	if v, ok := d.GetOk("report_only"); ok {
		t, err := expandSystemCertificateHpkp(d, v, "report_only", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["report-only"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemCertificateHpkp(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
