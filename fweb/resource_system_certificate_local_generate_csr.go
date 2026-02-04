// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure system/certificate.local.generate_csr

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCertificateLocalGenerateCsr() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemCertificateLocalGenerateCsrRead,
		Update: resourceSystemCertificateLocalGenerateCsrUpdate,
		Create: resourceSystemCertificateLocalGenerateCsrCreate,
		Delete: resourceSystemCertificateLocalGenerateCsrDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			//hostIp, domainName, email
			"idtype": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"domainname": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"subjectemail": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"organizationunit_1": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"organizationunit_2": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"organization": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"localitycity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"stateprovince": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"countryregion": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			//US
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			//ip, domainname, email
			"subject_alternative_type_1": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"subject_alternative_text_1": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"subject_alternative_type_2": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"subject_alternative_text_2": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"subject_alternative_type_3": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"subject_alternative_text_3": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			//1024, 1536, 2048, 4096
			"keysize": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			//SHA1, SHA256
			"digestalgorithm": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			//file, scep
			"enrollmentmethod": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"caserverurl": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"challengepassword": &schema.Schema{
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

func resourceSystemCertificateLocalGenerateCsrCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateLocalGenerateCsr: type error")
	}

	obj, err := getObjectSystemCertificateLocalGenerateCsr(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateLocalGenerateCsrresource while getting object: %v", err)
	}

	_, err = c.CreateSystemCertificateLocalGenerateCsr(obj, vdom)

	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateLocalGenerateCsr resource: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func resourceSystemCertificateLocalGenerateCsrUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceSystemCertificateLocalGenerateCsrRead(d, m)
}

func resourceSystemCertificateLocalGenerateCsrDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSystemCertificateLocalGenerateCsr(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCertificateLocalGenerateCsr resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCertificateLocalGenerateCsrRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemCertificateLocalGenerateCsr(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateGenerateCsr resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCertificateLocalGenerateCsr(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateLocalGenerateCsr resource from API: %v", err)
	}

	return nil
}

func flattenSystemCertificateLocalGenerateCsr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemCertificateLocalGenerateCsr(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenSystemCertificateLocalGenerateCsr(o["name"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemCertificateLocalGenerateCsr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCertificateLocalGenerateCsr(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("idtype"); ok {
		t, err := expandSystemCertificateLocalGenerateCsr(d, v, "idtype", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idType"] = t
		}
	}
	if v, ok := d.GetOk("ip"); ok {
		t, err := expandSystemCertificateLocalGenerateCsr(d, v, "ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}
	if v, ok := d.GetOk("domainname"); ok {
		t, err := expandSystemCertificateLocalGenerateCsr(d, v, "domainname", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domainName"] = t
		}
	}
	if v, ok := d.GetOk("subjectemail"); ok {
		t, err := expandSystemCertificateLocalGenerateCsr(d, v, "subjectemail", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subjectEmail"] = t
		}
	}
	if v, ok := d.GetOk("organizationunit_1"); ok {
		t, err := expandSystemCertificateLocalGenerateCsr(d, v, "organizationunit_1", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["organizationUnit_1"] = t
		}
	}
	if v, ok := d.GetOk("organizationunit_2"); ok {
		t, err := expandSystemCertificateLocalGenerateCsr(d, v, "organizationunit_2", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["organizationUnit_2"] = t
		}
	}
	if v, ok := d.GetOk("organization"); ok {
		t, err := expandSystemCertificateLocalGenerateCsr(d, v, "organization", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["organization"] = t
		}
	}
	if v, ok := d.GetOk("localitycity"); ok {
		t, err := expandSystemCertificateLocalGenerateCsr(d, v, "localitycity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["localityCity"] = t
		}
	}
	if v, ok := d.GetOk("stateprovince"); ok {
		t, err := expandSystemCertificateLocalGenerateCsr(d, v, "stateprovince", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["stateProvince"] = t
		}
	}
	if v, ok := d.GetOk("countryregion"); ok {
		t, err := expandSystemCertificateLocalGenerateCsr(d, v, "countryregion", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["countryRegion"] = t
		}
	}
	if v, ok := d.GetOk("email"); ok {
		t, err := expandSystemCertificateLocalGenerateCsr(d, v, "email", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["email"] = t
		}
	}
	if v, ok := d.GetOk("subject_alternative_type_1"); ok {
		t, err := expandSystemCertificateLocalGenerateCsr(d, v, "subject_alternative_type_1", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			switch t {
			case "ip":
				obj["alt_name_type_1"] = "7"
			case "domainname":
				obj["alt_name_type_1"] = "2"
			case "email":
				obj["alt_name_type_1"] = "1"
			default:
				obj["alt_name_type_1"] = "7"
			}
		}
	}
	if v, ok := d.GetOk("subject_alternative_txt_1"); ok {
		t, err := expandSystemCertificateLocalGenerateCsr(d, v, "subject_alternative_text_1", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alt_name_text_1"] = t
		}
	}
	if v, ok := d.GetOk("subject_alternative_type_2"); ok {
		t, err := expandSystemCertificateLocalGenerateCsr(d, v, "subject_alternative_type_2", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			switch t {
			case "ip":
				obj["alt_name_type_2"] = "7"
			case "domainname":
				obj["alt_name_type_2"] = "2"
			case "email":
				obj["alt_name_type_2"] = "1"
			default:
				obj["alt_name_type_2"] = "7"
			}
		}
	}
	if v, ok := d.GetOk("subject_alternative_text_2"); ok {
		t, err := expandSystemCertificateLocalGenerateCsr(d, v, "subject_alternative_text_2", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alt_name_text_2"] = t
		}
	}
	if v, ok := d.GetOk("subject_alternative_type_2"); ok {
		t, err := expandSystemCertificateLocalGenerateCsr(d, v, "subject_alternative_type_2", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			switch t {
			case "ip":
				obj["alt_name_type_3"] = "7"
			case "domainname":
				obj["alt_name_type_3"] = "2"
			case "email":
				obj["alt_name_type_3"] = "1"
			default:
				obj["alt_name_type_3"] = "7"
			}
		}
	}
	if v, ok := d.GetOk("subject_alternative_text_3"); ok {
		t, err := expandSystemCertificateLocalGenerateCsr(d, v, "subject_alternative_text_3", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alt_name_text_3"] = t
		}
	}
	if v, ok := d.GetOk("keysize"); ok {
		t, err := expandSystemCertificateLocalGenerateCsr(d, v, "keysize", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keySize"] = t
		}
	}
	if v, ok := d.GetOk("digestalgorithm"); ok {
		t, err := expandSystemCertificateLocalGenerateCsr(d, v, "digestalgorithm", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["digestAlgorithm"] = t
		}
	}
	if v, ok := d.GetOk("enrollmentmethod"); ok {
		t, err := expandSystemCertificateLocalGenerateCsr(d, v, "enrollmentmethod", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enrollmentMethod"] = t
		}
	}
	if v, ok := d.GetOk("caserverurl"); ok {
		t, err := expandSystemCertificateLocalGenerateCsr(d, v, "caserverurl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["caServerURL"] = t
		}
	}
	if v, ok := d.GetOk("challengepassword"); ok {
		t, err := expandSystemCertificateLocalGenerateCsr(d, v, "challengepassword", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["challengePassword"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemCertificateLocalGenerateCsr(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	obj["alt_name_type"] = "2"

	return &obj, nil
}
