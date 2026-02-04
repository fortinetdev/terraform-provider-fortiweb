// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure system/certificate.ocsp_signing_certs.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCertificateOcspSigningCerts() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemCertificateOcspSigningCertsRead,
		Update: resourceSystemCertificateOcspSigningCertsUpdate,
		Create: resourceSystemCertificateOcspSigningCertsCreate,
		Delete: resourceSystemCertificateOcspSigningCertsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"srcfile": &schema.Schema{
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

func resourceSystemCertificateOcspSigningCertsCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateOcspSigningCerts: type error")
	}

	obj, err := getObjectSystemCertificateOcspSigningCerts(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateOcspSigningCerts resource while getting object: %v", err)
	}

	_, err = c.CreateSystemCertificateOcspSigningCerts(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateOcspSigningCerts resource: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func resourceSystemCertificateOcspSigningCertsUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceSystemCertificateOcspSigningCertsRead(d, m)
}

func resourceSystemCertificateOcspSigningCertsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSystemCertificateOcspSigningCerts(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCertificateOcspSigningCerts resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCertificateOcspSigningCertsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemCertificateOcspSigningCerts(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateOcspSigningCerts resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCertificateOcspSigningCerts(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateOcspSigningCerts resource from API: %v", err)
	}

	return nil
}

func flattenSystemCertificateOcspSigningCerts(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemCertificateOcspSigningCerts(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenSystemCertificateOcspSigningCerts(o["name"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemCertificateOcspSigningCerts(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCertificateOcspSigningCerts(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("srcfile"); ok {
		t, err := expandSystemCertificateOcspSigningCerts(d, v, "srcfile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcfile"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemCertificateOcspSigningCerts(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
