// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttipb

import fmt "fmt"

func (dst *Configuration) SetFields(src *Configuration, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "default_cluster":
			if len(subs) > 0 {
				var newDst, newSrc *Configuration_Cluster
				if (src == nil || src.DefaultCluster == nil) && dst.DefaultCluster == nil {
					continue
				}
				if src != nil {
					newSrc = src.DefaultCluster
				}
				if dst.DefaultCluster != nil {
					newDst = dst.DefaultCluster
				} else {
					newDst = &Configuration_Cluster{}
					dst.DefaultCluster = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.DefaultCluster = src.DefaultCluster
				} else {
					dst.DefaultCluster = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *Configuration_UI) SetFields(src *Configuration_UI, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "branding_base_url":
			if len(subs) > 0 {
				return fmt.Errorf("'branding_base_url' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.BrandingBaseURL = src.BrandingBaseURL
			} else {
				var zero string
				dst.BrandingBaseURL = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *Configuration_Cluster) SetFields(src *Configuration_Cluster, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "ui":
			if len(subs) > 0 {
				var newDst, newSrc *Configuration_UI
				if (src == nil || src.UI == nil) && dst.UI == nil {
					continue
				}
				if src != nil {
					newSrc = src.UI
				}
				if dst.UI != nil {
					newDst = dst.UI
				} else {
					newDst = &Configuration_UI{}
					dst.UI = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.UI = src.UI
				} else {
					dst.UI = nil
				}
			}
		case "is":
			if len(subs) > 0 {
				var newDst, newSrc *Configuration_Cluster_IdentityServer
				if (src == nil || src.IS == nil) && dst.IS == nil {
					continue
				}
				if src != nil {
					newSrc = src.IS
				}
				if dst.IS != nil {
					newDst = dst.IS
				} else {
					newDst = &Configuration_Cluster_IdentityServer{}
					dst.IS = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.IS = src.IS
				} else {
					dst.IS = nil
				}
			}
		case "ns":
			if len(subs) > 0 {
				var newDst, newSrc *Configuration_Cluster_NetworkServer
				if (src == nil || src.NS == nil) && dst.NS == nil {
					continue
				}
				if src != nil {
					newSrc = src.NS
				}
				if dst.NS != nil {
					newDst = dst.NS
				} else {
					newDst = &Configuration_Cluster_NetworkServer{}
					dst.NS = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.NS = src.NS
				} else {
					dst.NS = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *Configuration_Cluster_IdentityServer) SetFields(src *Configuration_Cluster_IdentityServer, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "user_registration":
			if len(subs) > 0 {
				var newDst, newSrc *Configuration_Cluster_IdentityServer_UserRegistration
				if (src == nil || src.UserRegistration == nil) && dst.UserRegistration == nil {
					continue
				}
				if src != nil {
					newSrc = src.UserRegistration
				}
				if dst.UserRegistration != nil {
					newDst = dst.UserRegistration
				} else {
					newDst = &Configuration_Cluster_IdentityServer_UserRegistration{}
					dst.UserRegistration = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.UserRegistration = src.UserRegistration
				} else {
					dst.UserRegistration = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *Configuration_Cluster_NetworkServer) SetFields(src *Configuration_Cluster_NetworkServer, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "dev_addr_prefixes":
			if len(subs) > 0 {
				return fmt.Errorf("'dev_addr_prefixes' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.DevAddrPrefixes = src.DevAddrPrefixes
			} else {
				dst.DevAddrPrefixes = nil
			}
		case "deduplication_window":
			if len(subs) > 0 {
				return fmt.Errorf("'deduplication_window' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.DeduplicationWindow = src.DeduplicationWindow
			} else {
				dst.DeduplicationWindow = nil
			}
		case "cooldown_window":
			if len(subs) > 0 {
				return fmt.Errorf("'cooldown_window' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.CooldownWindow = src.CooldownWindow
			} else {
				dst.CooldownWindow = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *Configuration_Cluster_IdentityServer_UserRegistration) SetFields(src *Configuration_Cluster_IdentityServer_UserRegistration, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "invitation":
			if len(subs) > 0 {
				var newDst, newSrc *Configuration_Cluster_IdentityServer_UserRegistration_Invitation
				if (src == nil || src.Invitation == nil) && dst.Invitation == nil {
					continue
				}
				if src != nil {
					newSrc = src.Invitation
				}
				if dst.Invitation != nil {
					newDst = dst.Invitation
				} else {
					newDst = &Configuration_Cluster_IdentityServer_UserRegistration_Invitation{}
					dst.Invitation = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Invitation = src.Invitation
				} else {
					dst.Invitation = nil
				}
			}
		case "contact_info_validation":
			if len(subs) > 0 {
				var newDst, newSrc *Configuration_Cluster_IdentityServer_UserRegistration_ContactInfoValidation
				if (src == nil || src.ContactInfoValidation == nil) && dst.ContactInfoValidation == nil {
					continue
				}
				if src != nil {
					newSrc = src.ContactInfoValidation
				}
				if dst.ContactInfoValidation != nil {
					newDst = dst.ContactInfoValidation
				} else {
					newDst = &Configuration_Cluster_IdentityServer_UserRegistration_ContactInfoValidation{}
					dst.ContactInfoValidation = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ContactInfoValidation = src.ContactInfoValidation
				} else {
					dst.ContactInfoValidation = nil
				}
			}
		case "admin_approval":
			if len(subs) > 0 {
				var newDst, newSrc *Configuration_Cluster_IdentityServer_UserRegistration_AdminApproval
				if (src == nil || src.AdminApproval == nil) && dst.AdminApproval == nil {
					continue
				}
				if src != nil {
					newSrc = src.AdminApproval
				}
				if dst.AdminApproval != nil {
					newDst = dst.AdminApproval
				} else {
					newDst = &Configuration_Cluster_IdentityServer_UserRegistration_AdminApproval{}
					dst.AdminApproval = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.AdminApproval = src.AdminApproval
				} else {
					dst.AdminApproval = nil
				}
			}
		case "password_requirements":
			if len(subs) > 0 {
				var newDst, newSrc *Configuration_Cluster_IdentityServer_UserRegistration_PasswordRequirements
				if (src == nil || src.PasswordRequirements == nil) && dst.PasswordRequirements == nil {
					continue
				}
				if src != nil {
					newSrc = src.PasswordRequirements
				}
				if dst.PasswordRequirements != nil {
					newDst = dst.PasswordRequirements
				} else {
					newDst = &Configuration_Cluster_IdentityServer_UserRegistration_PasswordRequirements{}
					dst.PasswordRequirements = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.PasswordRequirements = src.PasswordRequirements
				} else {
					dst.PasswordRequirements = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *Configuration_Cluster_IdentityServer_UserRegistration_Invitation) SetFields(src *Configuration_Cluster_IdentityServer_UserRegistration_Invitation, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "required":
			if len(subs) > 0 {
				return fmt.Errorf("'required' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Required = src.Required
			} else {
				dst.Required = nil
			}
		case "token_ttl":
			if len(subs) > 0 {
				return fmt.Errorf("'token_ttl' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.TokenTTL = src.TokenTTL
			} else {
				dst.TokenTTL = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *Configuration_Cluster_IdentityServer_UserRegistration_ContactInfoValidation) SetFields(src *Configuration_Cluster_IdentityServer_UserRegistration_ContactInfoValidation, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "required":
			if len(subs) > 0 {
				return fmt.Errorf("'required' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Required = src.Required
			} else {
				dst.Required = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *Configuration_Cluster_IdentityServer_UserRegistration_AdminApproval) SetFields(src *Configuration_Cluster_IdentityServer_UserRegistration_AdminApproval, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "required":
			if len(subs) > 0 {
				return fmt.Errorf("'required' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Required = src.Required
			} else {
				dst.Required = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *Configuration_Cluster_IdentityServer_UserRegistration_PasswordRequirements) SetFields(src *Configuration_Cluster_IdentityServer_UserRegistration_PasswordRequirements, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "min_length":
			if len(subs) > 0 {
				return fmt.Errorf("'min_length' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.MinLength = src.MinLength
			} else {
				dst.MinLength = nil
			}
		case "max_length":
			if len(subs) > 0 {
				return fmt.Errorf("'max_length' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.MaxLength = src.MaxLength
			} else {
				dst.MaxLength = nil
			}
		case "min_uppercase":
			if len(subs) > 0 {
				return fmt.Errorf("'min_uppercase' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.MinUppercase = src.MinUppercase
			} else {
				dst.MinUppercase = nil
			}
		case "min_digits":
			if len(subs) > 0 {
				return fmt.Errorf("'min_digits' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.MinDigits = src.MinDigits
			} else {
				dst.MinDigits = nil
			}
		case "min_special":
			if len(subs) > 0 {
				return fmt.Errorf("'min_special' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.MinSpecial = src.MinSpecial
			} else {
				dst.MinSpecial = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}
