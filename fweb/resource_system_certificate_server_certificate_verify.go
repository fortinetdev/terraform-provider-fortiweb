// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure system/certificate.server_certificate_verify.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCertificateServerCertificateVerify() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemCertificateServerCertificateVerifyRead,
		Update: resourceSystemCertificateServerCertificateVerifyUpdate,
		Create: resourceSystemCertificateServerCertificateVerifyCreate,
		Delete: resourceSystemCertificateServerCertificateVerifyDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"ca": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"crl": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ocsp": &schema.Schema{
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

func resourceSystemCertificateServerCertificateVerifyCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateServerCertificateVerify: type error")
	}

	obj, err := getObjectSystemCertificateServerCertificateVerify(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateServerCertificateVerify resource while getting object: %v", err)
	}

	_, err = c.CreateSystemCertificateServerCertificateVerify(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateServerCertificateVerify resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemCertificateServerCertificateVerifyRead(d, m)
}

func resourceSystemCertificateServerCertificateVerifyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemCertificateServerCertificateVerify(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateServerCertificateVerify resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemCertificateServerCertificateVerify(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateServerCertificateVerify resource: %v", err)
	}

	return resourceSystemCertificateServerCertificateVerifyRead(d, m)
}

func resourceSystemCertificateServerCertificateVerifyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSystemCertificateServerCertificateVerify(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCertificateServerCertificateVerify resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCertificateServerCertificateVerifyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemCertificateServerCertificateVerify(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateServerCertificateVerify resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCertificateServerCertificateVerify(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateServerCertificateVerify resource from API: %v", err)
	}

	return nil
}

func flattenSystemCertificateServerCertificateVerify(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemCertificateServerCertificateVerify(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("ca", flattenSystemCertificateServerCertificateVerify(o["ca"], d, "ca", sv)); err != nil {
		if !fortiAPIPatch(o["ca"]) {
			return fmt.Errorf("Error reading ca: %v", err)
		}
	}
	if err = d.Set("crl", flattenSystemCertificateServerCertificateVerify(o["crl"], d, "crl", sv)); err != nil {
		if !fortiAPIPatch(o["crl"]) {
			return fmt.Errorf("Error reading crl: %v", err)
		}
	}
	if err = d.Set("ocsp", flattenSystemCertificateServerCertificateVerify(o["ocsp"], d, "ocsp", sv)); err != nil {
		if !fortiAPIPatch(o["ocsp"]) {
			return fmt.Errorf("Error reading ocsp: %v", err)
		}
	}
	if err = d.Set("mkey", flattenSystemCertificateServerCertificateVerify(o["name"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemCertificateServerCertificateVerify(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCertificateServerCertificateVerify(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ca"); ok {
		t, err := expandSystemCertificateServerCertificateVerify(d, v, "ca", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ca"] = t
		}
	}
	if v, ok := d.GetOk("crl"); ok {
		t, err := expandSystemCertificateServerCertificateVerify(d, v, "crl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["crl"] = t
		}
	}
	if v, ok := d.GetOk("ocsp"); ok {
		t, err := expandSystemCertificateServerCertificateVerify(d, v, "ocsp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ocsp"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemCertificateServerCertificateVerify(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
