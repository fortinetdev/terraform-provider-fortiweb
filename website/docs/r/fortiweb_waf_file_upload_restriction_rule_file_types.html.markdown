---
subcategory: "FortiWEB WAF"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_waf_file_upload_restriction_rule_file_types"
description: |-
  Configure FortiWEB waf file upload restriction rule file types configuration.
---

# fortiweb_waf_file_upload_restriction_rule_file_types
Configure FortiWEB waf file upload restriction rule file types configuration.

## Example Usage
```hcl
resource "fortiweb_waf_file_upload_restriction_rule_file_types" "test" {
	mkey = "test"
	sub_mkey = "1"
	file_type_name = "3GPP"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the waf file upload restriction rule.
* `sub_mkey` - The id of the waf file upload restriction rule file types entry.
* `file_type_name` - Select the numeric name that corresponds to the file type. Valid values: GIF, JPG, PDF, XML, MP3, MIDI, WAVE, FLV for a Macromedia Flash Video, RAR, ZIP, BMP, RM for RealMedia, MPEG for MPEG v, 3GPP, MSI, BAT.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Waf File Upload Restriction Rule File Types can be imported using any of these accepted formats:
```
$ terraform import fortiweb_waf_file_upload_restriction_rule_file_types.labelname {mkey}
```
