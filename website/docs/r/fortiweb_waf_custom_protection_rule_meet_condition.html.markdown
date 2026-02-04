---
subcategory: "FortiWEB WAF"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_waf_custom_protection_rule_meet_condition"
description: |-
  Configure FortiWEB waf custom protection rule meet condition configuration.
---

# fortiweb_waf_custom_protection_rule_meet_condition
Configure FortiWEB waf custom protection rule meet condition configuration.

## Example Usage
```hcl
resource "fortiweb_waf_custom_protection_rule_meet_condition" "test" {
	mkey = "test"
	sub_mkey = "1"
	operator = "RE"
	case_sensitive = "enable"
	expression = "test"
	request_target = "REQUEST_FILENAME REQUEST_HEADERS_NAMES"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the waf custom protection rule.
* `sub_mkey` - The id of the waf custom protection rule meet condition entry.
* `operator` - The signature matches when specified target has the different condition between the value of threshold.  Valid values: RE, GT, LT, NE, EQ.
* `case_sensitive` - Enable to differentiate upper case and lower case letters when evaluating the web server’s response for data leaks according to expression.
* `expression` - When operator is RE, type a regular expression that matches either an attack from a client or a data leak from the server.
* `request_target` - The name of one or more locations in the HTTP request to scan for a signature match. Valid values: REQUEST_FILENAME, REQUEST_URI, REQUEST_HEADERS_NAMES, REQUEST_HEADERS, REQUEST_COOKIES_NAMES, REQUEST_COOKIES, ARGS_NAMES, ARGS_VALUE, REQUEST_RAW_URI, REQUEST_BODY, CONTENT_LENGTH, HEADER_LENGTH, BODY_LENGTH, COOKIE_NUMBER, ARGS_NUMBER, HTTP_METHOD, HTTP_METHOD.
* `response_target` - The name of one or more locations in the HTTP response to scan for a signature match. Valid values: RESPONSE_BODY, RESPONSE_HEADER, CONTENT_LENGTH, HEADER_LENGTH, BODY_LENGTH, RESPONSE_CODE.
* `threshold` - The value that FortiWeb compares to the target value to determine if a request or response matches.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Waf Custom Protection Rule Meet Condition can be imported using any of these accepted formats:
```
$ terraform import fortiweb_waf_custom_protection_rule_meet_condition.labelname {mkey}
```
