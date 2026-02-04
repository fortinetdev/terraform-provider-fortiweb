// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure system/certificate.ocsp_responder.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCertificateOcspResponder() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemCertificateOcspResponderRead,
		Update: resourceSystemCertificateOcspResponderUpdate,
		Create: resourceSystemCertificateOcspResponderCreate,
		Delete: resourceSystemCertificateOcspResponderDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"ocsp_url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ocsp_signing_certs": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"caching": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"caching_ttl": &schema.Schema{
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

func resourceSystemCertificateOcspResponderCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateOcspResponder: type error")
	}

	obj, err := getObjectSystemCertificateOcspResponder(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateOcspResponder resource while getting object: %v", err)
	}

	_, err = c.CreateSystemCertificateOcspResponder(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateOcspResponder resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemCertificateOcspResponderRead(d, m)
}

func resourceSystemCertificateOcspResponderUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemCertificateOcspResponder(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateOcspResponder resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemCertificateOcspResponder(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateOcspResponder resource: %v", err)
	}

	return resourceSystemCertificateOcspResponderRead(d, m)
}

func resourceSystemCertificateOcspResponderDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSystemCertificateOcspResponder(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCertificateOcspResponder resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCertificateOcspResponderRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemCertificateOcspResponder(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateOcspResponder resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCertificateOcspResponder(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateOcspResponder resource from API: %v", err)
	}

	return nil
}

func flattenSystemCertificateOcspResponder(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemCertificateOcspResponder(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("ocsp_url", flattenSystemCertificateOcspResponder(o["ocsp-url"], d, "ocsp_url", sv)); err != nil {
		if !fortiAPIPatch(o["ocsp-url"]) {
			return fmt.Errorf("Error reading ocsp_url: %v", err)
		}
	}
	if err = d.Set("ocsp_signing_certs", flattenSystemCertificateOcspResponder(o["ocsp-signing-certs"], d, "ocsp_signing_certs", sv)); err != nil {
		if !fortiAPIPatch(o["ocsp-signing-certs"]) {
			return fmt.Errorf("Error reading ocsp_signing_certs: %v", err)
		}
	}
	if err = d.Set("timeout", flattenSystemCertificateOcspResponder(o["timeout"], d, "timeout", sv)); err != nil {
		if !fortiAPIPatch(o["timeout"]) {
			return fmt.Errorf("Error reading timeout: %v", err)
		}
	}
	if err = d.Set("comment", flattenSystemCertificateOcspResponder(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}
	if err = d.Set("caching", flattenSystemCertificateOcspResponder(o["caching"], d, "caching", sv)); err != nil {
		if !fortiAPIPatch(o["caching"]) {
			return fmt.Errorf("Error reading caching: %v", err)
		}
	}
	if err = d.Set("caching_ttl", flattenSystemCertificateOcspResponder(o["caching-ttl"], d, "caching_ttl", sv)); err != nil {
		if !fortiAPIPatch(o["caching-ttl"]) {
			return fmt.Errorf("Error reading caching_ttl: %v", err)
		}
	}
	if err = d.Set("mkey", flattenSystemCertificateOcspResponder(o["name"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemCertificateOcspResponder(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCertificateOcspResponder(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ocsp_url"); ok {
		t, err := expandSystemCertificateOcspResponder(d, v, "ocsp_url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ocsp-url"] = t
		}
	}
	if v, ok := d.GetOk("ocsp_signing_certs"); ok {
		t, err := expandSystemCertificateOcspResponder(d, v, "ocsp_signing_certs", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ocsp-signing-certs"] = t
		}
	}
	if v, ok := d.GetOk("timeout"); ok {
		t, err := expandSystemCertificateOcspResponder(d, v, "timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timeout"] = t
		}
	}
	if v, ok := d.GetOk("comment"); ok {
		t, err := expandSystemCertificateOcspResponder(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}
	if v, ok := d.GetOk("caching"); ok {
		t, err := expandSystemCertificateOcspResponder(d, v, "caching", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["caching"] = t
		}
	}
	if v, ok := d.GetOk("caching_ttl"); ok {
		t, err := expandSystemCertificateOcspResponder(d, v, "caching_ttl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["caching-ttl"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemCertificateOcspResponder(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
