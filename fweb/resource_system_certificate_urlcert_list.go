// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure system/certificate.urlcert/list.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCertificateUrlcertList() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemCertificateUrlcertListRead,
		Update: resourceSystemCertificateUrlcertListUpdate,
		Create: resourceSystemCertificateUrlcertListCreate,
		Delete: resourceSystemCertificateUrlcertListDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"sub_mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"require": &schema.Schema{
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

func resourceSystemCertificateUrlcertListCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateUrlcertList: type error")
	}

	obj, err := getObjectSystemCertificateUrlcertList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateUrlcertList resource while getting object: %v", err)
	}

	_, err = c.CreateSystemCertificateUrlcertList(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateUrlcertList resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemCertificateUrlcertListRead(d, m)
}

func resourceSystemCertificateUrlcertListUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	sub_mkey := "1"
	t := d.Get("sub_mkey")
	if v, ok := t.(string); ok {
		sub_mkey = v
	} else {
		return fmt.Errorf("Error describing SystemCertificateUrlcertList: type error")
	}

	obj, err := getObjectSystemCertificateUrlcertList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateUrlcertList resource while getting object: %v", err)
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	_, err = c.UpdateSystemCertificateUrlcertList(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateUrlcertList resource: %v", err)
	}

	return resourceSystemCertificateUrlcertListRead(d, m)
}

func resourceSystemCertificateUrlcertListDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	sub_mkey := "1"
	t := d.Get("sub_mkey")
	if v, ok := t.(string); ok {
		sub_mkey = v
	} else {
		return fmt.Errorf("Error describing SystemCertificateUrlcertList: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	err := c.DeleteSystemCertificateUrlcertList(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCertificateUrlcertList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCertificateUrlcertListRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	sub_mkey := "1"
	t := d.Get("sub_mkey")
	if v, ok := t.(string); ok {
		sub_mkey = v
	} else {
		return fmt.Errorf("Error describing SystemCertificateUrlcertList: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	o, err := c.ReadSystemCertificateUrlcertList(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateUrlcertList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCertificateUrlcertList(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateUrlcertList resource from API: %v", err)
	}

	return nil
}

func flattenSystemCertificateUrlcertList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemCertificateUrlcertList(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("sub_mkey", flattenSystemCertificateUrlcertList(o["id"], d, "sub_mkey", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading sub_mkey: %v", err)
		}
	}
	if err = d.Set("url", flattenSystemCertificateUrlcertList(o["url"], d, "url", sv)); err != nil {
		if !fortiAPIPatch(o["url"]) {
			return fmt.Errorf("Error reading url: %v", err)
		}
	}
	if err = d.Set("require", flattenSystemCertificateUrlcertList(o["require"], d, "require", sv)); err != nil {
		if !fortiAPIPatch(o["require"]) {
			return fmt.Errorf("Error reading require: %v", err)
		}
	}
	if err = d.Set("mkey", flattenSystemCertificateUrlcertList(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemCertificateUrlcertList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCertificateUrlcertList(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("sub_mkey"); ok {
		t, err := expandSystemCertificateUrlcertList(d, v, "sub_mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}
	if v, ok := d.GetOk("url"); ok {
		t, err := expandSystemCertificateUrlcertList(d, v, "url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url"] = t
		}
	}
	if v, ok := d.GetOk("require"); ok {
		t, err := expandSystemCertificateUrlcertList(d, v, "require", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["require"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemCertificateUrlcertList(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
