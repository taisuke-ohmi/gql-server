package gql_server

import (
	"context"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Billing(ctx context.Context, id string) (Billing, error) {
	return MFKInvoice{
		ID:         "1",
		SenderName: "invoice",
		InvoiceID:  "10",
	}, nil
}
