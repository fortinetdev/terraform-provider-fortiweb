---
subcategory: "FortiWEB WAF"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_waf_file_upload_restriction_policy"
description: |-
  Configure FortiWEB waf file upload restriction policy configuration.
---

# fortiweb_waf_file_upload_restriction_policy
Configure FortiWEB waf file upload restriction policy configuration.

## Example Usage
```hcl
resource "fortiweb_waf_file_upload_restriction_policy" "test" {
	mkey = "test"
	action = "alert"
	severity = "Medium"
	trigger = "test"
	av_scan = "enable"
	signature_check = "enable"
	fortisandbox_check = "enable"
	hold_session_while_scanning_file = "enable"
	exchange_mail_detection = "enable"
	owa_protocol = "enable"
	activesync_protocol = "enable"
	mapi_protocol = "enable"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the waf file upload restriction policy entry.
* `action` - Enter the action you want FortiWeb to perform when the policy is violated. Valid values: alert, alert_deny, block-period, deny_no_log.
* `severity` - Select the severity level to use in logs and reports generated when a violation of the rule occurs. Valid values: High, Medium, Low, Info.
* `trigger` - Enter the name of the trigger to apply when this policy is violated. 
* `av_scan` - Enter enable to scan for viruses, malware, and greyware.
* `signature_check` - Enter enable to check for signature.
* `fortisandbox_check` - Enter enable to send matching files to FortiSandbox for evaluation.
* `hold_session_while_scanning_file` - Enable it, and FortiWeb waits for up to 30 minutes. If FortiWeb holds the session for over 30 minutes while FortiSandbox scans the file in the request, FortiWeb will forward the session without taking any other actions.
* `exchange_mail_detection` - Enter enable so that FortiWeb will scan email attachments in applications using OWA or ActiveSync protocols. If enabled, FortiWeb will perform Trojan detection, an antivirus scan, and will send the attachments to FortiSandbox.
* `owa_protocol` - Available only when exchange-mail-detection {enable | disable} is set to enable. If enabled, FortiWeb will scan attachments in Exchange Email sent and received via a web browser login.
* `activesync_protocol` - Available only when exchange-mail-detection {enable | disable} is set to enable. If enabled, FortiWeb will scan attachments in Exchange Email sent and received via a mobile phone login.
* `mapi_protocol` - FortiWeb will scan attachments in Email sent and received via the Messaging Application Programming Interface (MAPI), a new transport protocol implemented in Microsoft Exchange Server 2013 Service Pack 1 (SP1).

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Waf File Upload Restriction Policy can be imported using any of these accepted formats:
```
$ terraform import fortiweb_waf_file_upload_restriction_policy.labelname {mkey}
```
