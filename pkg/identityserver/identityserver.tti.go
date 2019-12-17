// Copyright © 2019 The Things Industries B.V.

package identityserver

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/jinzhu/gorm"
	"go.thethings.network/lorawan-stack/pkg/identityserver/store"
	"go.thethings.network/lorawan-stack/pkg/log"
	"go.thethings.network/lorawan-stack/pkg/random"
	"go.thethings.network/lorawan-stack/pkg/tenant"
)

func (conf Config) Apply(ctx context.Context) Config {
	deriv := conf
	deriv.Email.Config = conf.Email.Config.Apply(ctx)
	tenantID := tenant.FromContext(ctx)
	if tenantID.TenantID == "" {
		return deriv
	}
	if tenantFetcher, ok := tenant.FetcherFromContext(ctx); ok {
		if tenant, err := tenantFetcher.FetchTenant(ctx, &tenantID, "name", "configuration"); err == nil {
			userRegistration := tenant.GetConfiguration().GetDefaultCluster().GetIS().GetUserRegistration()
			if required := userRegistration.GetInvitation().GetRequired(); required != nil {
				deriv.UserRegistration.Invitation.Required = required.Value
			}
			if tokenTTL := userRegistration.GetInvitation().GetTokenTTL(); tokenTTL != nil {
				deriv.UserRegistration.Invitation.TokenTTL = *tokenTTL
			}
			if required := userRegistration.GetContactInfoValidation().GetRequired(); required != nil {
				deriv.UserRegistration.ContactInfoValidation.Required = required.Value
			}
			if required := userRegistration.GetAdminApproval().GetRequired(); required != nil {
				deriv.UserRegistration.AdminApproval.Required = required.Value
			}
			if w := userRegistration.GetPasswordRequirements().GetMinLength(); w != nil {
				deriv.UserRegistration.PasswordRequirements.MinLength = int(w.Value)
			}
			if w := userRegistration.GetPasswordRequirements().GetMaxLength(); w != nil {
				deriv.UserRegistration.PasswordRequirements.MaxLength = int(w.Value)
			}
			if w := userRegistration.GetPasswordRequirements().GetMinUppercase(); w != nil {
				deriv.UserRegistration.PasswordRequirements.MinUppercase = int(w.Value)
			}
			if w := userRegistration.GetPasswordRequirements().GetMinDigits(); w != nil {
				deriv.UserRegistration.PasswordRequirements.MinDigits = int(w.Value)
			}
			if w := userRegistration.GetPasswordRequirements().GetMinSpecial(); w != nil {
				deriv.UserRegistration.PasswordRequirements.MinSpecial = int(w.Value)
			}
		}
	}
	return deriv
}

// TenancyConfig is the configuration for tenancy in the Identity Server.
type TenancyConfig struct {
	AdminKeys []string `name:"admin-keys" description:"Keys that can be used for tenant administration"`

	decodedAdminKeys [][]byte
}

func (c *TenancyConfig) decodeAdminKeys(ctx context.Context) error {
	for i, key := range c.AdminKeys {
		decodedKey, err := hex.DecodeString(key)
		if err != nil {
			return errInvalidTenantAdminKey.WithCause(err)
		}
		switch len(decodedKey) {
		case 16, 24, 32:
		default:
			return errInvalidTenantAdminKey.WithCause(fmt.Errorf("invalid length for key %d: must be 16, 24 or 32 bytes", i))
		}
		c.decodedAdminKeys = append(c.decodedAdminKeys, decodedKey)
	}
	if c.decodedAdminKeys == nil {
		c.decodedAdminKeys = [][]byte{random.Bytes(32)}
		log.FromContext(ctx).WithField("key", hex.EncodeToString(c.decodedAdminKeys[0])).Warn("No tenant admin key configured, generated a random one")
	}
	return nil
}

func (is *IdentityServer) withReadDatabase(ctx context.Context, f func(*gorm.DB) error) error {
	db := is.db
	if is.readDB != nil {
		db = is.readDB
	}
	return store.Transact(ctx, db, func(db *gorm.DB) error {
		if err := db.Exec("SET TRANSACTION READ ONLY").Error; err != nil {
			return err
		}
		return f(db)
	})
}
