package seed

import (
	"context"
	"github.com/iota-agency/iota-erp/elxolding/constants"
	"github.com/iota-agency/iota-erp/internal/application"
	"github.com/iota-agency/iota-erp/internal/domain/aggregates/role"
	"github.com/iota-agency/iota-erp/internal/domain/aggregates/user"
	"github.com/iota-agency/iota-erp/internal/infrastructure/persistence"
)

func SeedUser(ctx context.Context, app *application.Application) error {
	roleRepository := persistence.NewRoleRepository()

	for _, r := range constants.Roles {
		if err := roleRepository.CreateOrUpdate(ctx, &r); err != nil {
			return err
		}
	}
	userRepository := persistence.NewUserRepository()
	usr := &user.User{
		//nolint:exhaustruct
		ID:         1,
		FirstName:  "Admin",
		LastName:   "User",
		Email:      "test@gmail.com",
		UILanguage: user.UILanguageRU,
		Roles: []*role.Role{
			&constants.CEO,
		},
	}
	if err := usr.SetPassword("TestPass123!"); err != nil {
		return err
	}
	return userRepository.CreateOrUpdate(ctx, usr)
}
