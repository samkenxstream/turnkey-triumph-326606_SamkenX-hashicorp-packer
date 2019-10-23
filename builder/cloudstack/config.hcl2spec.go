// Code generated by "mapstructure-to-hcl2 -type Config"; DO NOT EDIT.
package cloudstack

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName           *string           `mapstructure:"packer_build_name" cty:"packer_build_name"`
	PackerBuilderType         *string           `mapstructure:"packer_builder_type" cty:"packer_builder_type"`
	PackerDebug               *bool             `mapstructure:"packer_debug" cty:"packer_debug"`
	PackerForce               *bool             `mapstructure:"packer_force" cty:"packer_force"`
	PackerOnError             *string           `mapstructure:"packer_on_error" cty:"packer_on_error"`
	PackerUserVars            map[string]string `mapstructure:"packer_user_variables" cty:"packer_user_variables"`
	PackerSensitiveVars       []string          `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables"`
	HTTPDir                   *string           `mapstructure:"http_directory" cty:"http_directory"`
	HTTPPortMin               *int              `mapstructure:"http_port_min" cty:"http_port_min"`
	HTTPPortMax               *int              `mapstructure:"http_port_max" cty:"http_port_max"`
	Type                      *string           `mapstructure:"communicator" cty:"communicator"`
	PauseBeforeConnect        *string           `mapstructure:"pause_before_connecting" cty:"pause_before_connecting"`
	SSHHost                   *string           `mapstructure:"ssh_host" cty:"ssh_host"`
	SSHPort                   *int              `mapstructure:"ssh_port" cty:"ssh_port"`
	SSHUsername               *string           `mapstructure:"ssh_username" cty:"ssh_username"`
	SSHPassword               *string           `mapstructure:"ssh_password" cty:"ssh_password"`
	SSHKeyPairName            *string           `mapstructure:"ssh_keypair_name" cty:"ssh_keypair_name"`
	SSHTemporaryKeyPairName   *string           `mapstructure:"temporary_key_pair_name" cty:"temporary_key_pair_name"`
	SSHClearAuthorizedKeys    *bool             `mapstructure:"ssh_clear_authorized_keys" cty:"ssh_clear_authorized_keys"`
	SSHPrivateKeyFile         *string           `mapstructure:"ssh_private_key_file" cty:"ssh_private_key_file"`
	SSHPty                    *bool             `mapstructure:"ssh_pty" cty:"ssh_pty"`
	SSHTimeout                *string           `mapstructure:"ssh_timeout" cty:"ssh_timeout"`
	SSHAgentAuth              *bool             `mapstructure:"ssh_agent_auth" cty:"ssh_agent_auth"`
	SSHDisableAgentForwarding *bool             `mapstructure:"ssh_disable_agent_forwarding" cty:"ssh_disable_agent_forwarding"`
	SSHHandshakeAttempts      *int              `mapstructure:"ssh_handshake_attempts" cty:"ssh_handshake_attempts"`
	SSHBastionHost            *string           `mapstructure:"ssh_bastion_host" cty:"ssh_bastion_host"`
	SSHBastionPort            *int              `mapstructure:"ssh_bastion_port" cty:"ssh_bastion_port"`
	SSHBastionAgentAuth       *bool             `mapstructure:"ssh_bastion_agent_auth" cty:"ssh_bastion_agent_auth"`
	SSHBastionUsername        *string           `mapstructure:"ssh_bastion_username" cty:"ssh_bastion_username"`
	SSHBastionPassword        *string           `mapstructure:"ssh_bastion_password" cty:"ssh_bastion_password"`
	SSHBastionPrivateKeyFile  *string           `mapstructure:"ssh_bastion_private_key_file" cty:"ssh_bastion_private_key_file"`
	SSHFileTransferMethod     *string           `mapstructure:"ssh_file_transfer_method" cty:"ssh_file_transfer_method"`
	SSHProxyHost              *string           `mapstructure:"ssh_proxy_host" cty:"ssh_proxy_host"`
	SSHProxyPort              *int              `mapstructure:"ssh_proxy_port" cty:"ssh_proxy_port"`
	SSHProxyUsername          *string           `mapstructure:"ssh_proxy_username" cty:"ssh_proxy_username"`
	SSHProxyPassword          *string           `mapstructure:"ssh_proxy_password" cty:"ssh_proxy_password"`
	SSHKeepAliveInterval      *string           `mapstructure:"ssh_keep_alive_interval" cty:"ssh_keep_alive_interval"`
	SSHReadWriteTimeout       *string           `mapstructure:"ssh_read_write_timeout" cty:"ssh_read_write_timeout"`
	SSHRemoteTunnels          []string          `mapstructure:"ssh_remote_tunnels" cty:"ssh_remote_tunnels"`
	SSHLocalTunnels           []string          `mapstructure:"ssh_local_tunnels" cty:"ssh_local_tunnels"`
	SSHPublicKey              []byte            `cty:"ssh_public_key"`
	SSHPrivateKey             []byte            `cty:"ssh_private_key"`
	WinRMUser                 *string           `mapstructure:"winrm_username" cty:"winrm_username"`
	WinRMPassword             *string           `mapstructure:"winrm_password" cty:"winrm_password"`
	WinRMHost                 *string           `mapstructure:"winrm_host" cty:"winrm_host"`
	WinRMPort                 *int              `mapstructure:"winrm_port" cty:"winrm_port"`
	WinRMTimeout              *string           `mapstructure:"winrm_timeout" cty:"winrm_timeout"`
	WinRMUseSSL               *bool             `mapstructure:"winrm_use_ssl" cty:"winrm_use_ssl"`
	WinRMInsecure             *bool             `mapstructure:"winrm_insecure" cty:"winrm_insecure"`
	WinRMUseNTLM              *bool             `mapstructure:"winrm_use_ntlm" cty:"winrm_use_ntlm"`
	APIURL                    *string           `mapstructure:"api_url" required:"true" cty:"api_url"`
	APIKey                    *string           `mapstructure:"api_key" required:"true" cty:"api_key"`
	SecretKey                 *string           `mapstructure:"secret_key" required:"true" cty:"secret_key"`
	AsyncTimeout              *string           `mapstructure:"async_timeout" required:"false" cty:"async_timeout"`
	HTTPGetOnly               *bool             `mapstructure:"http_get_only" required:"false" cty:"http_get_only"`
	SSLNoVerify               *bool             `mapstructure:"ssl_no_verify" required:"false" cty:"ssl_no_verify"`
	CIDRList                  []string          `mapstructure:"cidr_list" required:"false" cty:"cidr_list"`
	CreateSecurityGroup       *bool             `mapstructure:"create_security_group" required:"false" cty:"create_security_group"`
	DiskOffering              *string           `mapstructure:"disk_offering" required:"false" cty:"disk_offering"`
	DiskSize                  *int64            `mapstructure:"disk_size" required:"false" cty:"disk_size"`
	EjectISO                  *bool             `mapstructure:"eject_iso" cty:"eject_iso"`
	EjectISODelay             *string           `mapstructure:"eject_iso_delay" cty:"eject_iso_delay"`
	Expunge                   *bool             `mapstructure:"expunge" required:"false" cty:"expunge"`
	Hypervisor                *string           `mapstructure:"hypervisor" required:"false" cty:"hypervisor"`
	InstanceName              *string           `mapstructure:"instance_name" required:"false" cty:"instance_name"`
	Network                   *string           `mapstructure:"network" required:"true" cty:"network"`
	Project                   *string           `mapstructure:"project" required:"false" cty:"project"`
	PublicIPAddress           *string           `mapstructure:"public_ip_address" required:"false" cty:"public_ip_address"`
	PublicPort                *int              `mapstructure:"public_port" required:"false" cty:"public_port"`
	SecurityGroups            []string          `mapstructure:"security_groups" required:"false" cty:"security_groups"`
	ServiceOffering           *string           `mapstructure:"service_offering" required:"true" cty:"service_offering"`
	PreventFirewallChanges    *bool             `mapstructure:"prevent_firewall_changes" required:"false" cty:"prevent_firewall_changes"`
	SourceISO                 *string           `mapstructure:"source_iso" required:"true" cty:"source_iso"`
	SourceTemplate            *string           `mapstructure:"source_template" required:"true" cty:"source_template"`
	TemporaryKeypairName      *string           `mapstructure:"temporary_keypair_name" required:"false" cty:"temporary_keypair_name"`
	UseLocalIPAddress         *bool             `mapstructure:"use_local_ip_address" required:"false" cty:"use_local_ip_address"`
	UserData                  *string           `mapstructure:"user_data" required:"false" cty:"user_data"`
	UserDataFile              *string           `mapstructure:"user_data_file" required:"false" cty:"user_data_file"`
	Zone                      *string           `mapstructure:"zone" required:"true" cty:"zone"`
	TemplateName              *string           `mapstructure:"template_name" required:"false" cty:"template_name"`
	TemplateDisplayText       *string           `mapstructure:"template_display_text" required:"false" cty:"template_display_text"`
	TemplateOS                *string           `mapstructure:"template_os" required:"true" cty:"template_os"`
	TemplateFeatured          *bool             `mapstructure:"template_featured" required:"false" cty:"template_featured"`
	TemplatePublic            *bool             `mapstructure:"template_public" required:"false" cty:"template_public"`
	TemplatePasswordEnabled   *bool             `mapstructure:"template_password_enabled" required:"false" cty:"template_password_enabled"`
	TemplateRequiresHVM       *bool             `mapstructure:"template_requires_hvm" required:"false" cty:"template_requires_hvm"`
	TemplateScalable          *bool             `mapstructure:"template_scalable" required:"false" cty:"template_scalable"`
	TemplateTag               *string           `mapstructure:"template_tag" cty:"template_tag"`
	Tags                      map[string]string `mapstructure:"tags" cty:"tags"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{} { return new(FlatConfig) }

// HCL2Spec returns the hcldec.Spec of a FlatConfig.
// This spec is used by HCL to read the fields of FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":            &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":          &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_debug":                 &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":                 &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":              &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":        &hcldec.BlockAttrsSpec{TypeName: "packer_user_variables", ElementType: cty.String, Required: false},
		"packer_sensitive_variables":   &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"http_directory":               &hcldec.AttrSpec{Name: "http_directory", Type: cty.String, Required: false},
		"http_port_min":                &hcldec.AttrSpec{Name: "http_port_min", Type: cty.Number, Required: false},
		"http_port_max":                &hcldec.AttrSpec{Name: "http_port_max", Type: cty.Number, Required: false},
		"communicator":                 &hcldec.AttrSpec{Name: "communicator", Type: cty.String, Required: false},
		"pause_before_connecting":      &hcldec.AttrSpec{Name: "pause_before_connecting", Type: cty.String, Required: false},
		"ssh_host":                     &hcldec.AttrSpec{Name: "ssh_host", Type: cty.String, Required: false},
		"ssh_port":                     &hcldec.AttrSpec{Name: "ssh_port", Type: cty.Number, Required: false},
		"ssh_username":                 &hcldec.AttrSpec{Name: "ssh_username", Type: cty.String, Required: false},
		"ssh_password":                 &hcldec.AttrSpec{Name: "ssh_password", Type: cty.String, Required: false},
		"ssh_keypair_name":             &hcldec.AttrSpec{Name: "ssh_keypair_name", Type: cty.String, Required: false},
		"temporary_key_pair_name":      &hcldec.AttrSpec{Name: "temporary_key_pair_name", Type: cty.String, Required: false},
		"ssh_clear_authorized_keys":    &hcldec.AttrSpec{Name: "ssh_clear_authorized_keys", Type: cty.Bool, Required: false},
		"ssh_private_key_file":         &hcldec.AttrSpec{Name: "ssh_private_key_file", Type: cty.String, Required: false},
		"ssh_pty":                      &hcldec.AttrSpec{Name: "ssh_pty", Type: cty.Bool, Required: false},
		"ssh_timeout":                  &hcldec.AttrSpec{Name: "ssh_timeout", Type: cty.String, Required: false},
		"ssh_agent_auth":               &hcldec.AttrSpec{Name: "ssh_agent_auth", Type: cty.Bool, Required: false},
		"ssh_disable_agent_forwarding": &hcldec.AttrSpec{Name: "ssh_disable_agent_forwarding", Type: cty.Bool, Required: false},
		"ssh_handshake_attempts":       &hcldec.AttrSpec{Name: "ssh_handshake_attempts", Type: cty.Number, Required: false},
		"ssh_bastion_host":             &hcldec.AttrSpec{Name: "ssh_bastion_host", Type: cty.String, Required: false},
		"ssh_bastion_port":             &hcldec.AttrSpec{Name: "ssh_bastion_port", Type: cty.Number, Required: false},
		"ssh_bastion_agent_auth":       &hcldec.AttrSpec{Name: "ssh_bastion_agent_auth", Type: cty.Bool, Required: false},
		"ssh_bastion_username":         &hcldec.AttrSpec{Name: "ssh_bastion_username", Type: cty.String, Required: false},
		"ssh_bastion_password":         &hcldec.AttrSpec{Name: "ssh_bastion_password", Type: cty.String, Required: false},
		"ssh_bastion_private_key_file": &hcldec.AttrSpec{Name: "ssh_bastion_private_key_file", Type: cty.String, Required: false},
		"ssh_file_transfer_method":     &hcldec.AttrSpec{Name: "ssh_file_transfer_method", Type: cty.String, Required: false},
		"ssh_proxy_host":               &hcldec.AttrSpec{Name: "ssh_proxy_host", Type: cty.String, Required: false},
		"ssh_proxy_port":               &hcldec.AttrSpec{Name: "ssh_proxy_port", Type: cty.Number, Required: false},
		"ssh_proxy_username":           &hcldec.AttrSpec{Name: "ssh_proxy_username", Type: cty.String, Required: false},
		"ssh_proxy_password":           &hcldec.AttrSpec{Name: "ssh_proxy_password", Type: cty.String, Required: false},
		"ssh_keep_alive_interval":      &hcldec.AttrSpec{Name: "ssh_keep_alive_interval", Type: cty.String, Required: false},
		"ssh_read_write_timeout":       &hcldec.AttrSpec{Name: "ssh_read_write_timeout", Type: cty.String, Required: false},
		"ssh_remote_tunnels":           &hcldec.AttrSpec{Name: "ssh_remote_tunnels", Type: cty.List(cty.String), Required: false},
		"ssh_local_tunnels":            &hcldec.AttrSpec{Name: "ssh_local_tunnels", Type: cty.List(cty.String), Required: false},
		"ssh_public_key":               &hcldec.AttrSpec{Name: "ssh_public_key", Type: cty.List(cty.Number), Required: false},
		"ssh_private_key":              &hcldec.AttrSpec{Name: "ssh_private_key", Type: cty.List(cty.Number), Required: false},
		"winrm_username":               &hcldec.AttrSpec{Name: "winrm_username", Type: cty.String, Required: false},
		"winrm_password":               &hcldec.AttrSpec{Name: "winrm_password", Type: cty.String, Required: false},
		"winrm_host":                   &hcldec.AttrSpec{Name: "winrm_host", Type: cty.String, Required: false},
		"winrm_port":                   &hcldec.AttrSpec{Name: "winrm_port", Type: cty.Number, Required: false},
		"winrm_timeout":                &hcldec.AttrSpec{Name: "winrm_timeout", Type: cty.String, Required: false},
		"winrm_use_ssl":                &hcldec.AttrSpec{Name: "winrm_use_ssl", Type: cty.Bool, Required: false},
		"winrm_insecure":               &hcldec.AttrSpec{Name: "winrm_insecure", Type: cty.Bool, Required: false},
		"winrm_use_ntlm":               &hcldec.AttrSpec{Name: "winrm_use_ntlm", Type: cty.Bool, Required: false},
		"api_url":                      &hcldec.AttrSpec{Name: "api_url", Type: cty.String, Required: false},
		"api_key":                      &hcldec.AttrSpec{Name: "api_key", Type: cty.String, Required: false},
		"secret_key":                   &hcldec.AttrSpec{Name: "secret_key", Type: cty.String, Required: false},
		"async_timeout":                &hcldec.AttrSpec{Name: "async_timeout", Type: cty.String, Required: false},
		"http_get_only":                &hcldec.AttrSpec{Name: "http_get_only", Type: cty.Bool, Required: false},
		"ssl_no_verify":                &hcldec.AttrSpec{Name: "ssl_no_verify", Type: cty.Bool, Required: false},
		"cidr_list":                    &hcldec.AttrSpec{Name: "cidr_list", Type: cty.List(cty.String), Required: false},
		"create_security_group":        &hcldec.AttrSpec{Name: "create_security_group", Type: cty.Bool, Required: false},
		"disk_offering":                &hcldec.AttrSpec{Name: "disk_offering", Type: cty.String, Required: false},
		"disk_size":                    &hcldec.AttrSpec{Name: "disk_size", Type: cty.Number, Required: false},
		"eject_iso":                    &hcldec.AttrSpec{Name: "eject_iso", Type: cty.Bool, Required: false},
		"eject_iso_delay":              &hcldec.AttrSpec{Name: "eject_iso_delay", Type: cty.String, Required: false},
		"expunge":                      &hcldec.AttrSpec{Name: "expunge", Type: cty.Bool, Required: false},
		"hypervisor":                   &hcldec.AttrSpec{Name: "hypervisor", Type: cty.String, Required: false},
		"instance_name":                &hcldec.AttrSpec{Name: "instance_name", Type: cty.String, Required: false},
		"network":                      &hcldec.AttrSpec{Name: "network", Type: cty.String, Required: false},
		"project":                      &hcldec.AttrSpec{Name: "project", Type: cty.String, Required: false},
		"public_ip_address":            &hcldec.AttrSpec{Name: "public_ip_address", Type: cty.String, Required: false},
		"public_port":                  &hcldec.AttrSpec{Name: "public_port", Type: cty.Number, Required: false},
		"security_groups":              &hcldec.AttrSpec{Name: "security_groups", Type: cty.List(cty.String), Required: false},
		"service_offering":             &hcldec.AttrSpec{Name: "service_offering", Type: cty.String, Required: false},
		"prevent_firewall_changes":     &hcldec.AttrSpec{Name: "prevent_firewall_changes", Type: cty.Bool, Required: false},
		"source_iso":                   &hcldec.AttrSpec{Name: "source_iso", Type: cty.String, Required: false},
		"source_template":              &hcldec.AttrSpec{Name: "source_template", Type: cty.String, Required: false},
		"temporary_keypair_name":       &hcldec.AttrSpec{Name: "temporary_keypair_name", Type: cty.String, Required: false},
		"use_local_ip_address":         &hcldec.AttrSpec{Name: "use_local_ip_address", Type: cty.Bool, Required: false},
		"user_data":                    &hcldec.AttrSpec{Name: "user_data", Type: cty.String, Required: false},
		"user_data_file":               &hcldec.AttrSpec{Name: "user_data_file", Type: cty.String, Required: false},
		"zone":                         &hcldec.AttrSpec{Name: "zone", Type: cty.String, Required: false},
		"template_name":                &hcldec.AttrSpec{Name: "template_name", Type: cty.String, Required: false},
		"template_display_text":        &hcldec.AttrSpec{Name: "template_display_text", Type: cty.String, Required: false},
		"template_os":                  &hcldec.AttrSpec{Name: "template_os", Type: cty.String, Required: false},
		"template_featured":            &hcldec.AttrSpec{Name: "template_featured", Type: cty.Bool, Required: false},
		"template_public":              &hcldec.AttrSpec{Name: "template_public", Type: cty.Bool, Required: false},
		"template_password_enabled":    &hcldec.AttrSpec{Name: "template_password_enabled", Type: cty.Bool, Required: false},
		"template_requires_hvm":        &hcldec.AttrSpec{Name: "template_requires_hvm", Type: cty.Bool, Required: false},
		"template_scalable":            &hcldec.AttrSpec{Name: "template_scalable", Type: cty.Bool, Required: false},
		"template_tag":                 &hcldec.AttrSpec{Name: "template_tag", Type: cty.String, Required: false},
		"tags":                         &hcldec.BlockAttrsSpec{TypeName: "tags", ElementType: cty.String, Required: false},
	}
	return s
}
