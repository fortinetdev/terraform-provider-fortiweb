---
subcategory: "FortiWEB WAF"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_waf_url_access_url_access_parameter_url_access_parameter_list"
description: |-
  Configure FortiWEB waf url access url access parameter url access parameter list configuration.
---

# fortiweb_waf_url_access_url_access_parameter_url_access_parameter_list
Configure FortiWEB waf url access url access parameter url access parameter list configuration.

## Example Usage
```hcl
resource "fortiweb_waf_url_access_url_access_parameter_url_access_parameter_list" "test" {
	mkey = "test"
	sub_mkey = "1"
	argument_name = "test"
	argument_name_type = "regular"
	optional = "enable"
	type_checked = "enable"
	argument_type = "data-type"
	data_type = "String"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the waf url access url access parameter entry.
* `sub_mkey` - The id of the waf url access url access parameter url access parameter list entry.
* `argument_name_type` - Select the type of argument name. Valid values: plain, regular.
* `argument_name` - The name of argument.
* `argument_type` - Select the type of argument. Valid values: data-type, regular-expression, custom-data-type.
* `optional` - Enable the optional check. Valid values: enable, disable.
* `type_checked` - Enable the type check. Valid values: enable, disable.
* `data_type` - Select the name of Data-Type if argument_type is data-type.
* `argument_expression` - Select the name of Argument-Expression if argument_type is regular-expression.
* `custom_data_type` - Select the name of Custom-Data-Type if argument_type is custom-data-type.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Waf Url Access Url Access Parameter Url Access Parameter List can be imported using any of these accepted formats:
```
$ terraform import fortiweb_waf_url_access_url_access_parameter_url_access_parameter_list.labelname {mkey}
```
