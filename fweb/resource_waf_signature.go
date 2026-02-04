// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure waf/signature.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWafSignature() *schema.Resource {
	return &schema.Resource{
		Read:   resourceWafSignatureRead,
		Update: resourceWafSignatureUpdate,
		Create: resourceWafSignatureCreate,
		Delete: resourceWafSignatureDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"credit_card_detection_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"custom_signature_group": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"sensitivity_level": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"main_class_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"block_period": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"main_class_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"trigger": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fpm_cap": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"fpm_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
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

func resourceWafSignatureCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafSignature: type error")
	}

	obj, err := getObjectWafSignature(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WafSignature resource while getting object: %v", err)
	}
	fmt.Println("resourceWafSignatureCreate final obj: ", obj)

	_, err = c.CreateWafSignature(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating WafSignature resource: %v", err)
	}

	d.SetId(mkey)

	return resourceWafSignatureRead(d, m)
}

func resourceWafSignatureUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectWafSignature(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WafSignature resource while getting object: %v", err)
	}
	fmt.Println("resourceWafSignatureUpdate final obj: ", obj)

	_, err = c.UpdateWafSignature(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating WafSignature resource: %v", err)
	}

	return resourceWafSignatureRead(d, m)
}

func resourceWafSignatureDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteWafSignature(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting WafSignature resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWafSignatureRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadWafSignature(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading WafSignature resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWafSignature(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WafSignature resource from API: %v", err)
	}

	return nil
}

func flattenWafSignature(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if pre == "main_class_list" {
		if vSlice, ok := v.([]interface{}); ok {
			result := make([]interface{}, 0, len(vSlice))
			for _, item := range vSlice {
				if itemMap, ok := item.(map[string]interface{}); ok {
					convertedMap := make(map[string]interface{})
					for key, value := range itemMap {
						if key != "main_class_id" && key != "fpm_cap" {
							schemaKey := strings.Replace(key, "-", "_", -1)
							convertedMap[schemaKey] = value
						} else {
							convertedMap[key] = value
						}

					}
					result = append(result, convertedMap)
				}
			}
			return result
		} /*else if vMap, ok := v.(map[string]interface{}); ok {
			convertedMap := make(map[string]interface{})
			for key, value := range itemMap {
				if key != "main_class_id" && key != "fpm_cap" {
					schemaKey := strings.Replace(key, "-", "_", -1)
					convertedMap[schemaKey] = value
				} else {
					convertedMap[key] = value
				}
			}
			result := []interface{}{convertedMap}
			return result
		}*/
		return v
	}
	return v
}

func refreshObjectWafSignature(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("credit_card_detection_threshold", flattenWafSignature(o["creditCardDetectionThreshold"], d, "credit_card_detection_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["creditCardDetectionThreshold"]) {
			return fmt.Errorf("Error reading credit_card_detection_threshold: %v", err)
		}
	}
	if err = d.Set("comment", flattenWafSignature(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}
	if err = d.Set("custom_signature_group", flattenWafSignature(o["customSignatureGroup"], d, "custom_signature_group", sv)); err != nil {
		if !fortiAPIPatch(o["customSignatureGroup"]) {
			return fmt.Errorf("Error reading custom_signature_group: %v", err)
		}
	}
	if err = d.Set("sensitivity_level", flattenWafSignature(o["sensitivity_level"], d, "sensitivity_level", sv)); err != nil {
		if !fortiAPIPatch(o["sensitivity_level"]) {
			return fmt.Errorf("Error reading sensitivity_level: %v", err)
		}
	}
	mainClassListData := o["main_class_list"]
	flattenedData := flattenWafSignature(mainClassListData, d, "main_class_list", sv)
	if err = d.Set("main_class_list", flattenedData); err != nil {
		if !fortiAPIPatch(mainClassListData) {
			return fmt.Errorf("Error reading main_class_list: %v", err)
		}
	}
	if err = d.Set("mkey", flattenWafSignature(o["name"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandWafSignature(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	if pre == "main_class_list" {
		if vSlice, ok := v.([]interface{}); ok {
			result := make([]interface{}, 0, len(vSlice))
			for _, item := range vSlice {
				if itemMap, ok := item.(map[string]interface{}); ok {
					convertedMap := make(map[string]interface{})
					for key, value := range itemMap {
						if key == "main_class_id" || key == "fpm_cap" {
							convertedMap[key] = value
						} else {
							apiKey := strings.Replace(key, "_", "-", -1)
							convertedMap[apiKey] = value
						}
					}
					result = append(result, convertedMap)
				}
			}
			return result, nil
		}
		return v, nil
	}
	return v, nil
}

func getObjectWafSignature(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("credit_card_detection_threshold"); ok {
		t, err := expandWafSignature(d, v, "credit_card_detection_threshold", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["creditCardDetectionThreshold"] = t
		}
	}
	if v, ok := d.GetOk("comment"); ok {
		t, err := expandWafSignature(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}
	if v, ok := d.GetOk("custom_signature_group"); ok {
		t, err := expandWafSignature(d, v, "custom_signature_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["customSignatureGroup"] = t
		}
	}
	if v, ok := d.GetOk("sensitivity_level"); ok {
		t, err := expandWafSignature(d, v, "sensitivity_level", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sensitivity_level"] = t
		}
	}
	if v, ok := d.GetOk("main_class_list"); ok {
		t, err := expandWafSignature(d, v, "main_class_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["main_class_list"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandWafSignature(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
