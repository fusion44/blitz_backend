package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/fusion44/raspiblitz-backend/domain"
	"github.com/fusion44/raspiblitz-backend/graph/generated"
	"github.com/fusion44/raspiblitz-backend/graph/model"
	"github.com/fusion44/raspiblitz-backend/utils"
)

func (r *mutationResolver) Register(ctx context.Context, input model.RegisterInput) (*model.AuthResponse, error) {
	isValid := validation(ctx, input)

	if !isValid {
		return nil, domain.ErrInvalidInput
	}

	return r.Domain.Register(ctx, input)
}

func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (*model.AuthResponse, error) {
	isValid := validation(ctx, input)

	if !isValid {
		return nil, domain.ErrInvalidInput
	}

	return r.Domain.Login(ctx, input)
}

func (r *mutationResolver) PushUpdatedDeviceInfo(ctx context.Context, input model.UpdatedDeviceInfo) (*model.DeviceInfo, error) {
	r.Domain.PushUpdatedDeviceInfo(&input)
	return &model.DeviceInfo{State: input.State, Message: input.Message}, nil
}

func (r *queryResolver) DeviceInfo(ctx context.Context) (*model.DeviceInfo, error) {
	return r.Domain.InfoRepo.GetInfo()
}

func (r *subscriptionResolver) DeviceInfo(ctx context.Context) (<-chan *model.DeviceInfo, error) {
	// Get a random ID for the observer
	id := utils.RandString(8)

	go func() {
		// Delete the observer once the client disconnects
		<-ctx.Done()
		r.Domain.SetupRepo.DeleteDeviceInfoObserver(id)
	}()

	channel := r.Domain.SetupRepo.AddDeviceInfoObserver(id)

	return channel, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
