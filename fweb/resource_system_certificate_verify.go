// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure system/certificate.verify.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCertificateVerify() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemCertificateVerifyRead,
		Update: resourceSystemCertificateVerifyUpdate,
		Create: resourceSystemCertificateVerifyCreate,
		Delete: resourceSystemCertificateVerifyDelete,

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
			"publish_dn": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"strictly_need_cert": &schema.Schema{
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

func resourceSystemCertificateVerifyCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateVerify: type error")
	}

	obj, err := getObjectSystemCertificateVerify(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateVerify resource while getting object: %v", err)
	}

	_, err = c.CreateSystemCertificateVerify(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateVerify resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemCertificateVerifyRead(d, m)
}

func resourceSystemCertificateVerifyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemCertificateVerify(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateVerify resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemCertificateVerify(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateVerify resource: %v", err)
	}

	return resourceSystemCertificateVerifyRead(d, m)
}

func resourceSystemCertificateVerifyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSystemCertificateVerify(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCertificateVerify resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCertificateVerifyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemCertificateVerify(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateVerify resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCertificateVerify(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateVerify resource from API: %v", err)
	}

	return nil
}

func flattenSystemCertificateVerify(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemCertificateVerify(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("ca", flattenSystemCertificateVerify(o["ca"], d, "ca", sv)); err != nil {
		if !fortiAPIPatch(o["ca"]) {
			return fmt.Errorf("Error reading ca: %v", err)
		}
	}
	if err = d.Set("crl", flattenSystemCertificateVerify(o["crl"], d, "crl", sv)); err != nil {
		if !fortiAPIPatch(o["crl"]) {
			return fmt.Errorf("Error reading crl: %v", err)
		}
	}
	if err = d.Set("ocsp", flattenSystemCertificateVerify(o["ocsp"], d, "ocsp", sv)); err != nil {
		if !fortiAPIPatch(o["ocsp"]) {
			return fmt.Errorf("Error reading ocsp: %v", err)
		}
	}
	if err = d.Set("publish_dn", flattenSystemCertificateVerify(o["publish-dn"], d, "publish_dn", sv)); err != nil {
		if !fortiAPIPatch(o["publish-dn"]) {
			return fmt.Errorf("Error reading publish_dn: %v", err)
		}
	}
	if err = d.Set("strictly_need_cert", flattenSystemCertificateVerify(o["strictly-need-cert"], d, "strictly_need_cert", sv)); err != nil {
		if !fortiAPIPatch(o["strictly-need-cert"]) {
			return fmt.Errorf("Error reading strictly_need_cert: %v", err)
		}
	}
	if err = d.Set("mkey", flattenSystemCertificateVerify(o["name"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemCertificateVerify(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCertificateVerify(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ca"); ok {
		t, err := expandSystemCertificateVerify(d, v, "ca", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ca"] = t
		}
	}
	if v, ok := d.GetOk("crl"); ok {
		t, err := expandSystemCertificateVerify(d, v, "crl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["crl"] = t
		}
	}
	if v, ok := d.GetOk("ocsp"); ok {
		t, err := expandSystemCertificateVerify(d, v, "ocsp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ocsp"] = t
		}
	}
	if v, ok := d.GetOk("publish_dn"); ok {
		t, err := expandSystemCertificateVerify(d, v, "publish_dn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["publish-dn"] = t
		}
	}
	if v, ok := d.GetOk("strictly_need_cert"); ok {
		t, err := expandSystemCertificateVerify(d, v, "strictly_need_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["strictly-need-cert"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemCertificateVerify(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
