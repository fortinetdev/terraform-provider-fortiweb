// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure system/certificate.offline_sni/members.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCertificateOfflineSniMembers() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemCertificateOfflineSniMembersRead,
		Update: resourceSystemCertificateOfflineSniMembersUpdate,
		Create: resourceSystemCertificateOfflineSniMembersCreate,
		Delete: resourceSystemCertificateOfflineSniMembersDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"sub_mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"domain_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"local_cert": &schema.Schema{
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

func resourceSystemCertificateOfflineSniMembersCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateOfflineSniMembers: type error")
	}

	obj, err := getObjectSystemCertificateOfflineSniMembers(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateOfflineSniMembers resource while getting object: %v", err)
	}

	_, err = c.CreateSystemCertificateOfflineSniMembers(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateOfflineSniMembers resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemCertificateOfflineSniMembersRead(d, m)
}

func resourceSystemCertificateOfflineSniMembersUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateOfflineSniMembers: type error")
	}

	obj, err := getObjectSystemCertificateOfflineSniMembers(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateOfflineSniMembers resource while getting object: %v", err)
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	_, err = c.UpdateSystemCertificateOfflineSniMembers(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateOfflineSniMembers resource: %v", err)
	}

	return resourceSystemCertificateOfflineSniMembersRead(d, m)
}

func resourceSystemCertificateOfflineSniMembersDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateOfflineSniMembers: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	err := c.DeleteSystemCertificateOfflineSniMembers(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCertificateOfflineSniMembers resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCertificateOfflineSniMembersRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateOfflineSniMembers: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	o, err := c.ReadSystemCertificateOfflineSniMembers(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateOfflineSniMembers resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCertificateOfflineSniMembers(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateOfflineSniMembers resource from API: %v", err)
	}

	return nil
}

func flattenSystemCertificateOfflineSniMembers(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemCertificateOfflineSniMembers(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("sub_mkey", flattenSystemCertificateOfflineSniMembers(o["id"], d, "sub_mkey", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading sub_mkey: %v", err)
		}
	}
	if err = d.Set("domain_type", flattenSystemCertificateOfflineSniMembers(o["domain-type"], d, "domain_type", sv)); err != nil {
		if !fortiAPIPatch(o["domain-type"]) {
			return fmt.Errorf("Error reading domain_type: %v", err)
		}
	}
	if err = d.Set("domain", flattenSystemCertificateOfflineSniMembers(o["domain"], d, "domain", sv)); err != nil {
		if !fortiAPIPatch(o["domain"]) {
			return fmt.Errorf("Error reading domain: %v", err)
		}
	}
	if err = d.Set("local_cert", flattenSystemCertificateOfflineSniMembers(o["local-cert"], d, "local_cert", sv)); err != nil {
		if !fortiAPIPatch(o["local-cert"]) {
			return fmt.Errorf("Error reading local_cert: %v", err)
		}
	}
	if err = d.Set("mkey", flattenSystemCertificateOfflineSniMembers(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemCertificateOfflineSniMembers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCertificateOfflineSniMembers(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("sub_mkey"); ok {
		t, err := expandSystemCertificateOfflineSniMembers(d, v, "sub_mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}
	if v, ok := d.GetOk("domain_type"); ok {
		t, err := expandSystemCertificateOfflineSniMembers(d, v, "domain_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain-type"] = t
		}
	}
	if v, ok := d.GetOk("domain"); ok {
		t, err := expandSystemCertificateOfflineSniMembers(d, v, "domain", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain"] = t
		}
	}
	if v, ok := d.GetOk("local_cert"); ok {
		t, err := expandSystemCertificateOfflineSniMembers(d, v, "local_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-cert"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemCertificateOfflineSniMembers(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
