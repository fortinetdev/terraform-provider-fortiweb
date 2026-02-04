// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure system/certificate.sni/members.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCertificateSniMembers() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemCertificateSniMembersRead,
		Update: resourceSystemCertificateSniMembersUpdate,
		Create: resourceSystemCertificateSniMembersCreate,
		Delete: resourceSystemCertificateSniMembersDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"sub_mkey": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"domain_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"multi_local_cert": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"multi_local_cert_group": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"local_cert": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"certificate_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"lets_certificate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"inter_group": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"verify": &schema.Schema{
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

func resourceSystemCertificateSniMembersCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateSniMembers: type error")
	}

	obj, err := getObjectSystemCertificateSniMembers(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateSniMembers resource while getting object: %v", err)
	}

	_, err = c.CreateSystemCertificateSniMembers(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateSniMembers resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemCertificateSniMembersRead(d, m)
}

func resourceSystemCertificateSniMembersUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateSniMembers: type error")
	}

	obj, err := getObjectSystemCertificateSniMembers(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateSniMembers resource while getting object: %v", err)
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	_, err = c.UpdateSystemCertificateSniMembers(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateSniMembers resource: %v", err)
	}

	return resourceSystemCertificateSniMembersRead(d, m)
}

func resourceSystemCertificateSniMembersDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateSniMembers: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	err := c.DeleteSystemCertificateSniMembers(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCertificateSniMembers resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCertificateSniMembersRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateSniMembers: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	o, err := c.ReadSystemCertificateSniMembers(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateSniMembers resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCertificateSniMembers(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateSniMembers resource from API: %v", err)
	}

	return nil
}

func flattenSystemCertificateSniMembers(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemCertificateSniMembers(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("sub_mkey", flattenSystemCertificateSniMembers(o["id"], d, "sub_mkey", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading sub_mkey: %v", err)
		}
	}
	if err = d.Set("domain", flattenSystemCertificateSniMembers(o["domain"], d, "domain", sv)); err != nil {
		if !fortiAPIPatch(o["domain"]) {
			return fmt.Errorf("Error reading domain: %v", err)
		}
	}
	if err = d.Set("domain_type", flattenSystemCertificateSniMembers(o["domain-type"], d, "domain_type", sv)); err != nil {
		if !fortiAPIPatch(o["domain-type"]) {
			return fmt.Errorf("Error reading domain_type: %v", err)
		}
	}
	if err = d.Set("multi_local_cert", flattenSystemCertificateSniMembers(o["multi-local-cert"], d, "multi_local_cert", sv)); err != nil {
		if !fortiAPIPatch(o["multi-local-cert"]) {
			return fmt.Errorf("Error reading multi_local_cert: %v", err)
		}
	}
	if err = d.Set("multi_local_cert_group", flattenSystemCertificateSniMembers(o["multi-local-cert-group"], d, "multi_local_cert_group", sv)); err != nil {
		if !fortiAPIPatch(o["multi-local-cert-group"]) {
			return fmt.Errorf("Error reading multi_local_cert_group: %v", err)
		}
	}
	if err = d.Set("local_cert", flattenSystemCertificateSniMembers(o["local-cert"], d, "local_cert", sv)); err != nil {
		if !fortiAPIPatch(o["local-cert"]) {
			return fmt.Errorf("Error reading local_cert: %v", err)
		}
	}
	if err = d.Set("certificate_type", flattenSystemCertificateSniMembers(o["certificate-type"], d, "certificate_type", sv)); err != nil {
		if !fortiAPIPatch(o["certificate-type"]) {
			return fmt.Errorf("Error reading certificate_type: %v", err)
		}
	}
	if err = d.Set("lets_certificate", flattenSystemCertificateSniMembers(o["lets-certificate"], d, "lets_certificate", sv)); err != nil {
		if !fortiAPIPatch(o["lets-certificate"]) {
			return fmt.Errorf("Error reading lets_certificate: %v", err)
		}
	}
	if err = d.Set("inter_group", flattenSystemCertificateSniMembers(o["inter-group"], d, "inter_group", sv)); err != nil {
		if !fortiAPIPatch(o["inter-group"]) {
			return fmt.Errorf("Error reading inter_group: %v", err)
		}
	}
	if err = d.Set("verify", flattenSystemCertificateSniMembers(o["verify"], d, "verify", sv)); err != nil {
		if !fortiAPIPatch(o["verify"]) {
			return fmt.Errorf("Error reading verify: %v", err)
		}
	}
	if err = d.Set("mkey", flattenSystemCertificateSniMembers(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemCertificateSniMembers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCertificateSniMembers(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("sub_mkey"); ok {
		t, err := expandSystemCertificateSniMembers(d, v, "sub_mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}
	if v, ok := d.GetOk("domain"); ok {
		t, err := expandSystemCertificateSniMembers(d, v, "domain", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain"] = t
		}
	}
	if v, ok := d.GetOk("domain_type"); ok {
		t, err := expandSystemCertificateSniMembers(d, v, "domain_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain-type"] = t
		}
	}
	if v, ok := d.GetOk("multi_local_cert"); ok {
		t, err := expandSystemCertificateSniMembers(d, v, "multi_local_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multi-local-cert"] = t
		}
	}
	if v, ok := d.GetOk("multi_local_cert_group"); ok {
		t, err := expandSystemCertificateSniMembers(d, v, "multi_local_cert_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multi-local-cert-group"] = t
		}
	}
	if v, ok := d.GetOk("local_cert"); ok {
		t, err := expandSystemCertificateSniMembers(d, v, "local_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-cert"] = t
		}
	}
	if v, ok := d.GetOk("certificate_type"); ok {
		t, err := expandSystemCertificateSniMembers(d, v, "certificate_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certificate-type"] = t
		}
	}
	if v, ok := d.GetOk("lets_certificate"); ok {
		t, err := expandSystemCertificateSniMembers(d, v, "lets_certificate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lets-certificate"] = t
		}
	}
	if v, ok := d.GetOk("inter_group"); ok {
		t, err := expandSystemCertificateSniMembers(d, v, "inter_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inter-group"] = t
		}
	}
	if v, ok := d.GetOk("verify"); ok {
		t, err := expandSystemCertificateSniMembers(d, v, "verify", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["verify"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemCertificateSniMembers(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
