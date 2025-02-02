package role

import (
	"github.com/iota-uz/iota-sdk/modules/core/domain/entities/permission"
	"time"
)

func NewWithID(
	id uint,
	name string,
	description string,
	permissions []*permission.Permission,
	createdAt time.Time,
	updatedAt time.Time,
) (Role, error) {
	return &role{
		id:          id,
		name:        name,
		description: description,
		permissions: permissions,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
	}, nil
}

func New(
	name string,
	description string,
	permissions []*permission.Permission,
) (Role, error) {
	return &role{
		id:          0,
		name:        name,
		description: description,
		permissions: permissions,
		createdAt:   time.Now(),
		updatedAt:   time.Now(),
	}, nil
}

type role struct {
	id          uint
	name        string
	description string
	permissions []*permission.Permission
	createdAt   time.Time
	updatedAt   time.Time
}

func (r *role) ID() uint {
	return r.id
}

func (r *role) Name() string {
	return r.name
}

func (r *role) Description() string {
	return r.description
}

func (r *role) Permissions() []*permission.Permission {
	return r.permissions
}

func (r *role) CreatedAt() time.Time {
	return r.createdAt
}

func (r *role) UpdatedAt() time.Time {
	return r.updatedAt
}

func (r *role) SetName(name string) Role {
	return &role{
		id:          r.id,
		name:        name,
		description: r.description,
		permissions: r.permissions,
		createdAt:   r.createdAt,
		updatedAt:   time.Now(),
	}
}

func (r *role) SetDescription(description string) Role {
	return &role{
		id:          r.id,
		name:        r.name,
		description: description,
		permissions: r.permissions,
		createdAt:   r.createdAt,
		updatedAt:   time.Now(),
	}
}

func (r *role) AddPermission(p *permission.Permission) Role {
	return &role{
		id:          r.id,
		name:        r.name,
		description: r.description,
		permissions: append(r.permissions, p),
		createdAt:   r.createdAt,
		updatedAt:   time.Now(),
	}
}

func (r *role) SetPermissions(permissions []*permission.Permission) Role {
	return &role{
		id:          r.id,
		name:        r.name,
		description: r.description,
		permissions: permissions,
		createdAt:   r.createdAt,
		updatedAt:   time.Now(),
	}
}

func (r *role) Can(perm *permission.Permission) bool {
	for _, p := range r.permissions {
		if p.Equals(*perm) {
			return true
		}
	}
	return false
}
