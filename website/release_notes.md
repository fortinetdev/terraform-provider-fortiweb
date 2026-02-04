Release Notes
==============================

Release Version 1.0.0
--------------------

Release Targets
---------------

FortiWEB Terraform 1.0.0

Features
---------------

- Add new modules:

  - fortiweb_router_policy
  - fortiweb_router_static
  - fortiweb_server_policy_persistence_policy
  - fortiweb_server_policy_server_pool
  - fortiweb_server_policy_server_pool_pserver_list
  - fortiweb_server_policy_service_custom
  - fortiweb_server_policy_ssl_ciphers_custom
  - fortiweb_server_policy_vserver
  - fortiweb_server_policy_vserver_vip_list
  - fortiweb_system_accprofile
  - fortiweb_system_admin
  - fortiweb_system_admin_settings
  - fortiweb_system_certificate_ca
  - fortiweb_system_certificate_ca_group
  - fortiweb_system_certificate_ca_group_members
  - fortiweb_system_certificate_crl
  - fortiweb_system_certificate_crl_group
  - fortiweb_system_certificate_crl_group_members
  - fortiweb_system_certificate_hpkp
  - fortiweb_system_certificate_intermediate_certificate
  - fortiweb_system_certificate_intermediate_certificate_group
  - fortiweb_system_certificate_intermediate_certificate_group_members
  - fortiweb_system_certificate_letsencrypt
  - fortiweb_system_certificate_letsencrypt_san_list
  - fortiweb_system_certificate_local_generate_csr
  - fortiweb_system_certificate_local_import_certificate
  - fortiweb_system_certificate_multi_local
  - fortiweb_system_certificate_ocsp_responder
  - fortiweb_system_certificate_ocsp_signing_certs
  - fortiweb_system_certificate_ocsp_stapling
  - fortiweb_system_certificate_offline_sni
  - fortiweb_system_certificate_offline_sni_members
  - fortiweb_system_certificate_server_certificate_verify
  - fortiweb_system_certificate_sign_ca
  - fortiweb_system_certificate_sni
  - fortiweb_system_certificate_sni_members
  - fortiweb_system_certificate_tsl_ca
  - fortiweb_system_certificate_urlcert
  - fortiweb_system_certificate_urlcert_list
  - fortiweb_system_certificate_verify
  - fortiweb_system_certificate_xml_client_certificate
  - fortiweb_system_certificate_xml_client_certificate_group
  - fortiweb_system_certificate_xml_client_certificate_group_members
  - fortiweb_system_certificate_xml_server_certificate
  - fortiweb_system_dns
  - fortiweb_system_ha
  - fortiweb_system_maintenance_systemtime
  - fortiweb_system_ntp
  - fortiweb_system_ntp_ntpserver
  - fortiweb_system_snmp_community
  - fortiweb_system_snmp_community_hosts
  - fortiweb_system_snmp_sysinfo
  - fortiweb_system_snmp_user
  - fortiweb_system_snmp_user_hosts
  - fortiweb_system_sso_admin
  - fortiweb_system_vip
  - fortiweb_user_admin_usergrp
  - fortiweb_user_admin_usergrp_members
  - fortiweb_user_kerberos_user
  - fortiweb_user_kerberos_user_server_members
  - fortiweb_user_ldap_user
  - fortiweb_user_ldap_user_extracted_attritubes
  - fortiweb_user_ntlm_user
  - fortiweb_user_oauth_user_server
  - fortiweb_user_radius_user
  - fortiweb_user_recaptcha_user
  - fortiweb_user_tacacs_user
  - fortiweb_waf_allow_method_exceptions
  - fortiweb_waf_allow_method_exceptions_allow_method_exception_list
  - fortiweb_waf_allow_method_policy
  - fortiweb_waf_allowed_origins
  - fortiweb_waf_allowed_origins_origin_list
  - fortiweb_waf_cors_protection_rule
  - fortiweb_waf_cors_protection_rule_allowed_headers_list
  - fortiweb_waf_cors_protection_rule_allowed_methods_list
  - fortiweb_waf_cors_protection_rule_exposed_headers_list
  - fortiweb_waf_custom_protection_group
  - fortiweb_waf_custom_protection_group_type_list
  - fortiweb_waf_custom_protection_rule
  - fortiweb_waf_custom_protection_rule_meet_condition
  - fortiweb_waf_file_upload_restriction_policy
  - fortiweb_waf_file_upload_restriction_policy_rule
  - fortiweb_waf_file_upload_restriction_rule
  - fortiweb_waf_file_upload_restriction_rule_file_types
  - fortiweb_waf_http_constraints_exceptions
  - fortiweb_waf_http_constraints_exceptions_http_constraints_exception_list
  - fortiweb_waf_http_header_security
  - fortiweb_waf_http_header_security_exception
  - fortiweb_waf_http_header_security_exception_list
  - fortiweb_waf_http_header_security_http_header_security_list
  - fortiweb_waf_http_protocol_parameter_restriction
  - fortiweb_waf_signature
  - fortiweb_waf_url_access_url_access_parameter
  - fortiweb_waf_url_access_url_access_parameter_url_access_parameter_list
  - fortiweb_waf_url_access_url_access_policy
  - fortiweb_waf_url_access_url_access_policy_rule
  - fortiweb_waf_url_access_url_access_rule
  - fortiweb_waf_url_access_url_access_rule_match_condition
  - fortiweb_waf_url_encryption_url_encryption_policy
  - fortiweb_waf_url_encryption_url_encryption_policy_rule_list
  - fortiweb_waf_url_encryption_url_encryption_rule
  - fortiweb_waf_url_encryption_url_encryption_rule_exceptions
  - fortiweb_waf_url_encryption_url_encryption_rule_url_list


- FortiWEB version: v8.0

Notice
---------------

- For detailed configurations of the new modules, please refer to the internal examples within each module.
