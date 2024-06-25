package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.47

import (
	"context"
	"fmt"

	model "github.com/iota-agency/iota-erp/graph/gqlmodels"
)

// CreateExpense is the resolver for the createExpense field.
func (r *mutationResolver) CreateExpense(ctx context.Context, input model.CreateExpense) (*model.Expense, error) {
	panic(fmt.Errorf("not implemented: CreateExpense - createExpense"))
}

// UpdateExpense is the resolver for the updateExpense field.
func (r *mutationResolver) UpdateExpense(ctx context.Context, id int64, input model.UpdateExpense) (*model.Expense, error) {
	panic(fmt.Errorf("not implemented: UpdateExpense - updateExpense"))
}

// DeleteExpense is the resolver for the deleteExpense field.
func (r *mutationResolver) DeleteExpense(ctx context.Context, id int64) (*model.Expense, error) {
	panic(fmt.Errorf("not implemented: DeleteExpense - deleteExpense"))
}

// Data is the resolver for the data field.
func (r *paginatedExpensesResolver) Data(ctx context.Context, obj *model.PaginatedExpenses) ([]*model.Expense, error) {
	panic(fmt.Errorf("not implemented: Data - data"))
}

// Total is the resolver for the total field.
func (r *paginatedExpensesResolver) Total(ctx context.Context, obj *model.PaginatedExpenses) (int64, error) {
	panic(fmt.Errorf("not implemented: Total - total"))
}

// Expense is the resolver for the expense field.
func (r *queryResolver) Expense(ctx context.Context, id int64) (*model.Expense, error) {
	panic(fmt.Errorf("not implemented: Expense - expense"))
}

// Expenses is the resolver for the expenses field.
func (r *queryResolver) Expenses(ctx context.Context, offset int, limit int, sortBy []string) (*model.PaginatedExpenses, error) {
	panic(fmt.Errorf("not implemented: Expenses - expenses"))
}

// ExpenseCreated is the resolver for the expenseCreated field.
func (r *subscriptionResolver) ExpenseCreated(ctx context.Context) (<-chan *model.Expense, error) {
	panic(fmt.Errorf("not implemented: ExpenseCreated - expenseCreated"))
}

// ExpenseUpdated is the resolver for the expenseUpdated field.
func (r *subscriptionResolver) ExpenseUpdated(ctx context.Context) (<-chan *model.Expense, error) {
	panic(fmt.Errorf("not implemented: ExpenseUpdated - expenseUpdated"))
}

// ExpenseDeleted is the resolver for the expenseDeleted field.
func (r *subscriptionResolver) ExpenseDeleted(ctx context.Context) (<-chan *model.Expense, error) {
	panic(fmt.Errorf("not implemented: ExpenseDeleted - expenseDeleted"))
}

// PaginatedExpenses returns PaginatedExpensesResolver implementation.
func (r *Resolver) PaginatedExpenses() PaginatedExpensesResolver {
	return &paginatedExpensesResolver{r}
}

type paginatedExpensesResolver struct{ *Resolver }
