// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure system/certificate.ocsp_stapling.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCertificateOcspStapling() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemCertificateOcspStaplingRead,
		Update: resourceSystemCertificateOcspStaplingUpdate,
		Create: resourceSystemCertificateOcspStaplingCreate,
		Delete: resourceSystemCertificateOcspStaplingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"certificate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"local_cert": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ocsp_url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"comment": &schema.Schema{
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

func resourceSystemCertificateOcspStaplingCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateOcspStapling: type error")
	}

	obj, err := getObjectSystemCertificateOcspStapling(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateOcspStapling resource while getting object: %v", err)
	}

	_, err = c.CreateSystemCertificateOcspStapling(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateOcspStapling resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemCertificateOcspStaplingRead(d, m)
}

func resourceSystemCertificateOcspStaplingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemCertificateOcspStapling(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateOcspStapling resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemCertificateOcspStapling(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateOcspStapling resource: %v", err)
	}

	return resourceSystemCertificateOcspStaplingRead(d, m)
}

func resourceSystemCertificateOcspStaplingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSystemCertificateOcspStapling(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCertificateOcspStapling resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCertificateOcspStaplingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemCertificateOcspStapling(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateOcspStapling resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCertificateOcspStapling(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateOcspStapling resource from API: %v", err)
	}

	return nil
}

func flattenSystemCertificateOcspStapling(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemCertificateOcspStapling(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("certificate", flattenSystemCertificateOcspStapling(o["certificate"], d, "certificate", sv)); err != nil {
		if !fortiAPIPatch(o["certificate"]) {
			return fmt.Errorf("Error reading certificate: %v", err)
		}
	}
	if err = d.Set("local_cert", flattenSystemCertificateOcspStapling(o["local-cert"], d, "local_cert", sv)); err != nil {
		if !fortiAPIPatch(o["local-cert"]) {
			return fmt.Errorf("Error reading local_cert: %v", err)
		}
	}
	if err = d.Set("ocsp_url", flattenSystemCertificateOcspStapling(o["ocsp_url"], d, "ocsp_url", sv)); err != nil {
		if !fortiAPIPatch(o["ocsp_url"]) {
			return fmt.Errorf("Error reading ocsp_url: %v", err)
		}
	}
	if err = d.Set("comment", flattenSystemCertificateOcspStapling(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}
	if err = d.Set("mkey", flattenSystemCertificateOcspStapling(o["name"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemCertificateOcspStapling(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCertificateOcspStapling(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("certificate"); ok {
		t, err := expandSystemCertificateOcspStapling(d, v, "certificate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certificate"] = t
		}
	}
	if v, ok := d.GetOk("local_cert"); ok {
		t, err := expandSystemCertificateOcspStapling(d, v, "local_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-cert"] = t
		}
	}
	if v, ok := d.GetOk("ocsp_url"); ok {
		t, err := expandSystemCertificateOcspStapling(d, v, "ocsp_url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ocsp_url"] = t
		}
	}
	if v, ok := d.GetOk("comment"); ok {
		t, err := expandSystemCertificateOcspStapling(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemCertificateOcspStapling(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
