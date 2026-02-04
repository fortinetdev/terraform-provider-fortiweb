// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure system/admin.settings.

package fortiweb

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemAdminSettings() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemAdminSettingsRead,
		Update: resourceSystemAdminSettingsUpdate,
		Create: resourceSystemAdminSettingsUpdate,
		Delete: resourceSystemAdminSettingsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"http": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"https": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"httpsservercertificate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"admin_tls_v12": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"admin_tls_v13": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"confsync_enable": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"configsync": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"idletimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"language": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"passwordpolicy": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"min_length_option": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"min_length": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"single_admin_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"character_requirements": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"min_lower_case_letter": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"min_upper_case_letter": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"min_number": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"min_non_alphanumeric": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"forbid_password_reuse": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"history_password_number": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"expire_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"expire_day": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSystemAdminSettingsCreate(d *schema.ResourceData, m interface{}) error {
	mkey := ""
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemAdminSettings(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemAdminSettings resource while getting object: %v", err)
	}

	_, err = c.CreateSystemAdminSettings(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemAdminSettings resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemAdminSettingsRead(d, m)
}

func resourceSystemAdminSettingsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := "SystemAdminSettings"
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemAdminSettings(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemAdminSettings resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemAdminSettings(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemAdminSettings resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemAdminSettingsRead(d, m)
}

func resourceSystemAdminSettingsDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")

	return nil
}

func resourceSystemAdminSettingsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemAdminSettings(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemAdminSettings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAdminSettings(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemAdminSettings resource from API: %v", err)
	}

	return nil
}

func flattenSystemAdminSettings(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if pre == "passwordpolicy" {
		if vMap, ok := v.(map[string]interface{}); ok {
			convertedMap := make(map[string]interface{})
			for key, value := range vMap {
				schemaKey := strings.Replace(key, "-", "_", 4)
				convertedMap[schemaKey] = value
			}
			return []interface{}{convertedMap}
		}
		return v
	}
	return v
}

func refreshObjectSystemAdminSettings(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("hostname", flattenSystemAdminSettings(o["hostname"], d, "hostname", sv)); err != nil {
		if !fortiAPIPatch(o["hostname"]) {
			return fmt.Errorf("Error reading hostname: %v", err)
		}
	}
	if err = d.Set("http", flattenSystemAdminSettings(o["http"], d, "http", sv)); err != nil {
		if !fortiAPIPatch(o["http"]) {
			return fmt.Errorf("Error reading http: %v", err)
		}
	}
	if err = d.Set("https", flattenSystemAdminSettings(o["https"], d, "https", sv)); err != nil {
		if !fortiAPIPatch(o["https"]) {
			return fmt.Errorf("Error reading https: %v", err)
		}
	}
	if err = d.Set("httpsservercertificate", flattenSystemAdminSettings(o["httpsServerCertificate"], d, "httpsservercertificate", sv)); err != nil {
		if !fortiAPIPatch(o["httpsServerCertificate"]) {
			return fmt.Errorf("Error reading httpsservercertificate: %v", err)
		}
	}
	if err = d.Set("admin_tls_v12", flattenSystemAdminSettings(o["admin-tls-v12"], d, "admin_tls_v12", sv)); err != nil {
		if !fortiAPIPatch(o["admin-tls-v12"]) {
			return fmt.Errorf("Error reading admin_tls_v12: %v", err)
		}
	}
	if err = d.Set("admin_tls_v13", flattenSystemAdminSettings(o["admin-tls-v13"], d, "admin_tls_v13", sv)); err != nil {
		if !fortiAPIPatch(o["admin-tls-v13"]) {
			return fmt.Errorf("Error reading admin_tls_v13: %v", err)
		}
	}
	if err = d.Set("confsync_enable", flattenSystemAdminSettings(o["confsync-enable"], d, "confsync_enable", sv)); err != nil {
		if !fortiAPIPatch(o["confsync-enable"]) {
			return fmt.Errorf("Error reading confsync_enable: %v", err)
		}
	}
	if err = d.Set("configsync", flattenSystemAdminSettings(o["configSync"], d, "configsync", sv)); err != nil {
		if !fortiAPIPatch(o["configSync"]) {
			return fmt.Errorf("Error reading configsync: %v", err)
		}
	}
	if err = d.Set("idletimeout", flattenSystemAdminSettings(o["idleTimeout"], d, "idletimeout", sv)); err != nil {
		if !fortiAPIPatch(o["idleTimeout"]) {
			return fmt.Errorf("Error reading idletimeout: %v", err)
		}
	}
	if err = d.Set("language", flattenSystemAdminSettings(o["language"], d, "language", sv)); err != nil {
		if !fortiAPIPatch(o["language"]) {
			return fmt.Errorf("Error reading language: %v", err)
		}
	}
	if err = d.Set("passwordpolicy", flattenSystemAdminSettings(o["passwordPolicy"], d, "passwordpolicy", sv)); err != nil {
		if !fortiAPIPatch(o["passwordPolicy"]) {
			return fmt.Errorf("Error reading passwordpolicy: %v", err)
		}
	}
	return nil
}

func expandSystemAdminSettings(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	if pre == "passwordpolicy" {
		if vSlice, ok := v.([]interface{}); ok && len(vSlice) > 0 {
			if vMap, ok := vSlice[0].(map[string]interface{}); ok {
				convertedMap := make(map[string]interface{})
				for key, value := range vMap {
					schemaKey := strings.Replace(key, "_", "-", 4)
					convertedMap[schemaKey] = value
				}
				return convertedMap, nil
			}
		}
		return v, nil
	}
	return v, nil
}

func getObjectSystemAdminSettings(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("hostname"); ok {
		t, err := expandSystemAdminSettings(d, v, "hostname", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hostname"] = t
		}
	}
	if v, ok := d.GetOk("http"); ok {
		t, err := expandSystemAdminSettings(d, v, "http", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http"] = t
		}
	}
	if v, ok := d.GetOk("https"); ok {
		t, err := expandSystemAdminSettings(d, v, "https", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["https"] = t
		}
	}
	if v, ok := d.GetOk("httpsservercertificate"); ok {
		t, err := expandSystemAdminSettings(d, v, "httpsservercertificate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["httpsServerCertificate"] = t
		}
	}
	if v, ok := d.GetOk("admin_tls_v12"); ok {
		t, err := expandSystemAdminSettings(d, v, "admin_tls_v12", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-tls-v12"] = t
		}
	}
	if v, ok := d.GetOk("admin_tls_v13"); ok {
		t, err := expandSystemAdminSettings(d, v, "admin_tls_v13", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-tls-v13"] = t
		}
	}
	if v, ok := d.GetOk("confsync_enable"); ok {
		t, err := expandSystemAdminSettings(d, v, "confsync_enable", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["confsync-enable"] = t
		}
	}
	if v, ok := d.GetOk("configsync"); ok {
		t, err := expandSystemAdminSettings(d, v, "configsync", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["configSync"] = t
		}
	}
	if v, ok := d.GetOk("idletimeout"); ok {
		t, err := expandSystemAdminSettings(d, v, "idletimeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idleTimeout"] = t
		}
	}
	if v, ok := d.GetOk("language"); ok {
		t, err := expandSystemAdminSettings(d, v, "language", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["language"] = t
		}
	}
	if v, ok := d.GetOk("passwordpolicy"); ok {
		t, err := expandSystemAdminSettings(d, v, "passwordpolicy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["passwordPolicy"] = t
		}
	}

	return &obj, nil
}
