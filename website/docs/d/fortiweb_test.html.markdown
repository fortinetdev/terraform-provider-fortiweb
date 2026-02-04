---
subcategory: "FortiWEB Config"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_test"
description: |-
  Get information on an fortiweb test
---

# Data Source: fortiweb_test
Use this data source to get information on an fortiweb config sync list

## Example Usage

```hcl
 data "fortiweb_test" sample1 {
  mkey = "1"
}

output output1 {
  value = data.fortiweb_test.sample1
}
```

## Argument Reference
* `mkey` - (Required) Specify the mkey of the desired  config sync list.
* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `mkey` - config test name.

