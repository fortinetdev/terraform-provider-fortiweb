// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure user/oauth_user.server.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceUserOauthUserServer() *schema.Resource {
	return &schema.Resource{
		Read:   resourceUserOauthUserServerRead,
		Update: resourceUserOauthUserServerUpdate,
		Create: resourceUserOauthUserServerCreate,
		Delete: resourceUserOauthUserServerDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"scope": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"token_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"oidc": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"client_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"client_secret": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"redirect_endpoint": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"authz_req": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"token_req": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"refresh_req": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"jwks_req": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"validate_req": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"validate_frequency": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"validate_interval": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"userinfo_req": &schema.Schema{
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

func resourceUserOauthUserServerCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing UserOauthUserServer: type error")
	}

	obj, err := getObjectUserOauthUserServer(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating UserOauthUserServer resource while getting object: %v", err)
	}

	_, err = c.CreateUserOauthUserServer(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating UserOauthUserServer resource: %v", err)
	}

	d.SetId(mkey)

	return resourceUserOauthUserServerRead(d, m)
}

func resourceUserOauthUserServerUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectUserOauthUserServer(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating UserOauthUserServer resource while getting object: %v", err)
	}

	_, err = c.UpdateUserOauthUserServer(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating UserOauthUserServer resource: %v", err)
	}

	return resourceUserOauthUserServerRead(d, m)
}

func resourceUserOauthUserServerDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteUserOauthUserServer(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting UserOauthUserServer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserOauthUserServerRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadUserOauthUserServer(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading UserOauthUserServer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserOauthUserServer(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading UserOauthUserServer resource from API: %v", err)
	}

	return nil
}

func flattenUserOauthUserServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectUserOauthUserServer(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mode", flattenUserOauthUserServer(o["mode"], d, "mode", sv)); err != nil {
		if !fortiAPIPatch(o["mode"]) {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}
	if err = d.Set("scope", flattenUserOauthUserServer(o["scope"], d, "scope", sv)); err != nil {
		if !fortiAPIPatch(o["scope"]) {
			return fmt.Errorf("Error reading scope: %v", err)
		}
	}
	if err = d.Set("token_timeout", flattenUserOauthUserServer(o["token-timeout"], d, "token_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["token-timeout"]) {
			return fmt.Errorf("Error reading token_timeout: %v", err)
		}
	}
	if err = d.Set("oidc", flattenUserOauthUserServer(o["oidc"], d, "oidc", sv)); err != nil {
		if !fortiAPIPatch(o["oidc"]) {
			return fmt.Errorf("Error reading oidc: %v", err)
		}
	}
	if err = d.Set("client_id", flattenUserOauthUserServer(o["client-id"], d, "client_id", sv)); err != nil {
		if !fortiAPIPatch(o["client-id"]) {
			return fmt.Errorf("Error reading client_id: %v", err)
		}
	}
	if err = d.Set("client_secret", flattenUserOauthUserServer(o["client-secret"], d, "client_secret", sv)); err != nil {
		if !fortiAPIPatch(o["client-secret"]) {
			return fmt.Errorf("Error reading client_secret: %v", err)
		}
	}
	if err = d.Set("redirect_endpoint", flattenUserOauthUserServer(o["redirect-endpoint"], d, "redirect_endpoint", sv)); err != nil {
		if !fortiAPIPatch(o["redirect-endpoint"]) {
			return fmt.Errorf("Error reading redirect_endpoint: %v", err)
		}
	}
	if err = d.Set("authz_req", flattenUserOauthUserServer(o["authz-req"], d, "authz_req", sv)); err != nil {
		if !fortiAPIPatch(o["authz-req"]) {
			return fmt.Errorf("Error reading authz_req: %v", err)
		}
	}
	if err = d.Set("token_req", flattenUserOauthUserServer(o["token-req"], d, "token_req", sv)); err != nil {
		if !fortiAPIPatch(o["token-req"]) {
			return fmt.Errorf("Error reading token_req: %v", err)
		}
	}
	if err = d.Set("refresh_req", flattenUserOauthUserServer(o["refresh-req"], d, "refresh_req", sv)); err != nil {
		if !fortiAPIPatch(o["refresh-req"]) {
			return fmt.Errorf("Error reading refresh_req: %v", err)
		}
	}
	if err = d.Set("jwks_req", flattenUserOauthUserServer(o["jwks-req"], d, "jwks_req", sv)); err != nil {
		if !fortiAPIPatch(o["jwks-req"]) {
			return fmt.Errorf("Error reading jwks_req: %v", err)
		}
	}
	if err = d.Set("validate_req", flattenUserOauthUserServer(o["validate-req"], d, "validate_req", sv)); err != nil {
		if !fortiAPIPatch(o["validate-req"]) {
			return fmt.Errorf("Error reading validate_req: %v", err)
		}
	}
	if err = d.Set("validate_frequency", flattenUserOauthUserServer(o["validate-frequency"], d, "validate_frequency", sv)); err != nil {
		if !fortiAPIPatch(o["validate-frequency"]) {
			return fmt.Errorf("Error reading validate_frequency: %v", err)
		}
	}
	if err = d.Set("validate_interval", flattenUserOauthUserServer(o["validate-interval"], d, "validate_interval", sv)); err != nil {
		if !fortiAPIPatch(o["validate-interval"]) {
			return fmt.Errorf("Error reading validate_interval: %v", err)
		}
	}
	if err = d.Set("userinfo_req", flattenUserOauthUserServer(o["userinfo-req"], d, "userinfo_req", sv)); err != nil {
		if !fortiAPIPatch(o["userinfo-req"]) {
			return fmt.Errorf("Error reading userinfo_req: %v", err)
		}
	}
	if err = d.Set("mkey", flattenUserOauthUserServer(o["name"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandUserOauthUserServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectUserOauthUserServer(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mode"); ok {
		t, err := expandUserOauthUserServer(d, v, "mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}
	if v, ok := d.GetOk("scope"); ok {
		t, err := expandUserOauthUserServer(d, v, "scope", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scope"] = t
		}
	}
	if v, ok := d.GetOk("token_timeout"); ok {
		t, err := expandUserOauthUserServer(d, v, "token_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["token-timeout"] = t
		}
	}
	if v, ok := d.GetOk("oidc"); ok {
		t, err := expandUserOauthUserServer(d, v, "oidc", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oidc"] = t
		}
	}
	if v, ok := d.GetOk("client_id"); ok {
		t, err := expandUserOauthUserServer(d, v, "client_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-id"] = t
		}
	}
	if v, ok := d.GetOk("client_secret"); ok {
		t, err := expandUserOauthUserServer(d, v, "client_secret", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-secret"] = t
		}
	}
	if v, ok := d.GetOk("redirect_endpoint"); ok {
		t, err := expandUserOauthUserServer(d, v, "redirect_endpoint", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redirect-endpoint"] = t
		}
	}
	if v, ok := d.GetOk("authz_req"); ok {
		t, err := expandUserOauthUserServer(d, v, "authz_req", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authz-req"] = t
		}
	}
	if v, ok := d.GetOk("token_req"); ok {
		t, err := expandUserOauthUserServer(d, v, "token_req", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["token-req"] = t
		}
	}
	if v, ok := d.GetOk("refresh_req"); ok {
		t, err := expandUserOauthUserServer(d, v, "refresh_req", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["refresh-req"] = t
		}
	}
	if v, ok := d.GetOk("jwks_req"); ok {
		t, err := expandUserOauthUserServer(d, v, "jwks_req", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["jwks-req"] = t
		}
	}
	if v, ok := d.GetOk("validate_req"); ok {
		t, err := expandUserOauthUserServer(d, v, "validate_req", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["validate-req"] = t
		}
	}
	if v, ok := d.GetOk("validate_frequency"); ok {
		t, err := expandUserOauthUserServer(d, v, "validate_frequency", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["validate-frequency"] = t
		}
	}
	if v, ok := d.GetOk("validate_interval"); ok {
		t, err := expandUserOauthUserServer(d, v, "validate_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["validate-interval"] = t
		}
	}
	if v, ok := d.GetOk("userinfo_req"); ok {
		t, err := expandUserOauthUserServer(d, v, "userinfo_req", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["userinfo-req"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandUserOauthUserServer(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
