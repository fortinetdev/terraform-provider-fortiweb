// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure server_policy/ssl_ciphers.custom.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceServerPolicySslCiphersCustom() *schema.Resource {
	return &schema.Resource{
		Read:   resourceServerPolicySslCiphersCustomRead,
		Update: resourceServerPolicySslCiphersCustomUpdate,
		Create: resourceServerPolicySslCiphersCustomCreate,
		Delete: resourceServerPolicySslCiphersCustomDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"tls_v10": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"tls_v11": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"tls_v12": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"tls_v13": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ssl_cipher": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ssl_custom_cipher": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"tls13_custom_cipher": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "TLS_AES_256_GCM_SHA384",
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

func resourceServerPolicySslCiphersCustomCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing ServerPolicySslCiphersCustom: type error")
	}

	obj, err := getObjectServerPolicySslCiphersCustom(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating ServerPolicySslCiphersCustom resource while getting object: %v", err)
	}

	_, err = c.CreateServerPolicySslCiphersCustom(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating ServerPolicySslCiphersCustom resource: %v", err)
	}

	d.SetId(mkey)

	return resourceServerPolicySslCiphersCustomRead(d, m)
}

func resourceServerPolicySslCiphersCustomUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectServerPolicySslCiphersCustom(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating ServerPolicySslCiphersCustom resource while getting object: %v", err)
	}

	_, err = c.UpdateServerPolicySslCiphersCustom(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating ServerPolicySslCiphersCustom resource: %v", err)
	}

	return resourceServerPolicySslCiphersCustomRead(d, m)
}

func resourceServerPolicySslCiphersCustomDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteServerPolicySslCiphersCustom(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting ServerPolicySslCiphersCustom resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceServerPolicySslCiphersCustomRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadServerPolicySslCiphersCustom(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading ServerPolicySslCiphersCustom resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectServerPolicySslCiphersCustom(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading ServerPolicySslCiphersCustom resource from API: %v", err)
	}

	return nil
}

func flattenServerPolicySslCiphersCustom(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectServerPolicySslCiphersCustom(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("tls_v10", flattenServerPolicySslCiphersCustom(o["tls-v10"], d, "tls_v10", sv)); err != nil {
		if !fortiAPIPatch(o["tls-v10"]) {
			return fmt.Errorf("Error reading tls_v10: %v", err)
		}
	}
	if err = d.Set("tls_v11", flattenServerPolicySslCiphersCustom(o["tls-v11"], d, "tls_v11", sv)); err != nil {
		if !fortiAPIPatch(o["tls-v11"]) {
			return fmt.Errorf("Error reading tls_v11: %v", err)
		}
	}
	if err = d.Set("tls_v12", flattenServerPolicySslCiphersCustom(o["tls-v12"], d, "tls_v12", sv)); err != nil {
		if !fortiAPIPatch(o["tls-v12"]) {
			return fmt.Errorf("Error reading tls_v12: %v", err)
		}
	}
	if err = d.Set("tls_v13", flattenServerPolicySslCiphersCustom(o["tls-v13"], d, "tls_v13", sv)); err != nil {
		if !fortiAPIPatch(o["tls-v13"]) {
			return fmt.Errorf("Error reading tls_v13: %v", err)
		}
	}
	if err = d.Set("ssl_cipher", flattenServerPolicySslCiphersCustom(o["ssl-cipher"], d, "ssl_cipher", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-cipher"]) {
			return fmt.Errorf("Error reading ssl_cipher: %v", err)
		}
	}
	if err = d.Set("ssl_custom_cipher", flattenServerPolicySslCiphersCustom(o["ssl-custom-cipher"], d, "ssl_custom_cipher", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-custom-cipher"]) {
			return fmt.Errorf("Error reading ssl_custom_cipher: %v", err)
		}
	}
	if err = d.Set("tls13_custom_cipher", flattenServerPolicySslCiphersCustom(o["tls13-custom-cipher"], d, "tls13_custom_cipher", sv)); err != nil {
		if !fortiAPIPatch(o["tls13-custom-cipher"]) {
			return fmt.Errorf("Error reading tls13_custom_cipher: %v", err)
		}
	}
	if err = d.Set("mkey", flattenServerPolicySslCiphersCustom(o["name"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandServerPolicySslCiphersCustom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectServerPolicySslCiphersCustom(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("tls_v10"); ok {
		t, err := expandServerPolicySslCiphersCustom(d, v, "tls_v10", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tls-v10"] = t
		}
	}
	if v, ok := d.GetOk("tls_v11"); ok {
		t, err := expandServerPolicySslCiphersCustom(d, v, "tls_v11", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tls-v11"] = t
		}
	}
	if v, ok := d.GetOk("tls_v12"); ok {
		t, err := expandServerPolicySslCiphersCustom(d, v, "tls_v12", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tls-v12"] = t
		}
	}
	if v, ok := d.GetOk("tls_v13"); ok {
		t, err := expandServerPolicySslCiphersCustom(d, v, "tls_v13", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tls-v13"] = t
		}
	}
	if v, ok := d.GetOk("ssl_cipher"); ok {
		t, err := expandServerPolicySslCiphersCustom(d, v, "ssl_cipher", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-cipher"] = t
		}
	}
	if v, ok := d.GetOk("ssl_custom_cipher"); ok {
		t, err := expandServerPolicySslCiphersCustom(d, v, "ssl_custom_cipher", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-custom-cipher"] = t
		}
	}
	if v, ok := d.GetOk("tls13_custom_cipher"); ok {
		t, err := expandServerPolicySslCiphersCustom(d, v, "tls13_custom_cipher", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tls13-custom-cipher"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandServerPolicySslCiphersCustom(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
