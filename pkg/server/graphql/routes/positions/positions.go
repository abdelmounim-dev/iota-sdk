package positions

import (
	"github.com/graphql-go/graphql"
	"github.com/iota-agency/iota-erp/models"
	"github.com/iota-agency/iota-erp/pkg/server/graphql/adapters"
	"github.com/jmoiron/sqlx"
)

func Queries(db *sqlx.DB) []*graphql.Field {
	return adapters.DefaultQueries(db, &models.Position{}, "position", "positions")
}

func Mutations(db *sqlx.DB) []*graphql.Field {
	return adapters.DefaultMutations(db, &models.Position{}, "position")
}
