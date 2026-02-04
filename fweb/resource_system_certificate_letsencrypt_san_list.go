// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure system/certificate.letsencrypt/san-list.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCertificateLetsencryptSanList() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemCertificateLetsencryptSanListRead,
		Update: resourceSystemCertificateLetsencryptSanListUpdate,
		Create: resourceSystemCertificateLetsencryptSanListCreate,
		Delete: resourceSystemCertificateLetsencryptSanListDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"sub_mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"san_dns": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"type": &schema.Schema{
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

func resourceSystemCertificateLetsencryptSanListCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateLetsencryptSanList: type error")
	}

	obj, err := getObjectSystemCertificateLetsencryptSanList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateLetsencryptSanList resource while getting object: %v", err)
	}

	_, err = c.CreateSystemCertificateLetsencryptSanList(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateLetsencryptSanList resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemCertificateLetsencryptSanListRead(d, m)
}

func resourceSystemCertificateLetsencryptSanListUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateLetsencryptSanList: type error")
	}

	obj, err := getObjectSystemCertificateLetsencryptSanList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateLetsencryptSanList resource while getting object: %v", err)
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	_, err = c.UpdateSystemCertificateLetsencryptSanList(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateLetsencryptSanList resource: %v", err)
	}

	return resourceSystemCertificateLetsencryptSanListRead(d, m)
}

func resourceSystemCertificateLetsencryptSanListDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateLetsencryptSanList: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	err := c.DeleteSystemCertificateLetsencryptSanList(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCertificateLetsencryptSanList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCertificateLetsencryptSanListRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateLetsencryptSanList: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	o, err := c.ReadSystemCertificateLetsencryptSanList(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateLetsencryptSanList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCertificateLetsencryptSanList(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateLetsencryptSanList resource from API: %v", err)
	}

	return nil
}

func flattenSystemCertificateLetsencryptSanList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemCertificateLetsencryptSanList(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("sub_mkey", flattenSystemCertificateLetsencryptSanList(o["id"], d, "sub_mkey", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading sub_mkey: %v", err)
		}
	}
	if err = d.Set("san_dns", flattenSystemCertificateLetsencryptSanList(o["san-dns"], d, "san_dns", sv)); err != nil {
		if !fortiAPIPatch(o["san-dns"]) {
			return fmt.Errorf("Error reading san_dns: %v", err)
		}
	}
	if err = d.Set("mkey", flattenSystemCertificateLetsencryptSanList(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}
	if err = d.Set("type", flattenSystemCertificateLetsencryptSanList(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func expandSystemCertificateLetsencryptSanList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCertificateLetsencryptSanList(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("sub_mkey"); ok {
		t, err := expandSystemCertificateLetsencryptSanList(d, v, "sub_mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}
	if v, ok := d.GetOk("san_dns"); ok {
		t, err := expandSystemCertificateLetsencryptSanList(d, v, "san_dns", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["san-dns"] = t
		}
	}
	if v, ok := d.GetOk("type"); ok {
		t, err := expandSystemCertificateLetsencryptSanList(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemCertificateLetsencryptSanList(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
