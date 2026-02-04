// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure system/certificate.multi_local.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCertificateMultiLocal() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemCertificateMultiLocalRead,
		Update: resourceSystemCertificateMultiLocalUpdate,
		Create: resourceSystemCertificateMultiLocalCreate,
		Delete: resourceSystemCertificateMultiLocalDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"rsa_cert": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"dsa_cert": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ecc_cert": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"comment": &schema.Schema{
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

func resourceSystemCertificateMultiLocalCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateMultiLocal: type error")
	}

	obj, err := getObjectSystemCertificateMultiLocal(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateMultiLocal resource while getting object: %v", err)
	}

	_, err = c.CreateSystemCertificateMultiLocal(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateMultiLocal resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemCertificateMultiLocalRead(d, m)
}

func resourceSystemCertificateMultiLocalUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemCertificateMultiLocal(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateMultiLocal resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemCertificateMultiLocal(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateMultiLocal resource: %v", err)
	}

	return resourceSystemCertificateMultiLocalRead(d, m)
}

func resourceSystemCertificateMultiLocalDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSystemCertificateMultiLocal(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCertificateMultiLocal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCertificateMultiLocalRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemCertificateMultiLocal(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateMultiLocal resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCertificateMultiLocal(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateMultiLocal resource from API: %v", err)
	}

	return nil
}

func flattenSystemCertificateMultiLocal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemCertificateMultiLocal(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenSystemCertificateMultiLocal(o["name"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}
	if err = d.Set("rsa_cert", flattenSystemCertificateMultiLocal(o["rsa-cert"], d, "rsa_cert", sv)); err != nil {
		if !fortiAPIPatch(o["rsa-cert"]) {
			return fmt.Errorf("Error reading rsa_cert: %v", err)
		}
	}
	if err = d.Set("dsa_cert", flattenSystemCertificateMultiLocal(o["dsa-cert"], d, "dsa_cert", sv)); err != nil {
		if !fortiAPIPatch(o["dsa-cert"]) {
			return fmt.Errorf("Error reading dsa_cert: %v", err)
		}
	}
	if err = d.Set("ecc_cert", flattenSystemCertificateMultiLocal(o["ecc-cert"], d, "ecc_cert", sv)); err != nil {
		if !fortiAPIPatch(o["ecc-cert"]) {
			return fmt.Errorf("Error reading ecc_cert: %v", err)
		}
	}
	if err = d.Set("comment", flattenSystemCertificateMultiLocal(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	return nil
}

func expandSystemCertificateMultiLocal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCertificateMultiLocal(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemCertificateCaGroup(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}
	if v, ok := d.GetOk("rsa_cert"); ok {
		t, err := expandSystemCertificateCaGroup(d, v, "rsa_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsa-cert"] = t
		}
	}
	if v, ok := d.GetOk("dsa_cert"); ok {
		t, err := expandSystemCertificateCaGroup(d, v, "dsa_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dsa-cert"] = t
		}
	}
	if v, ok := d.GetOk("ecc_cert"); ok {
		t, err := expandSystemCertificateCaGroup(d, v, "ecc_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ecc-cert"] = t
		}
	}
	if v, ok := d.GetOk("comment"); ok {
		t, err := expandSystemCertificateCaGroup(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	return &obj, nil
}
