package gateway

import (
	"context"

	"github.com/Sup3r-Us3r/barber-server/internal/domain/apperr"
	"github.com/Sup3r-Us3r/barber-server/internal/domain/entity"
)

type BarberGatewayInterface interface {
	GetBarberById(ctx context.Context, id string) (entity.Barber, *apperr.AppErr)
	GetBarberByEmail(ctx context.Context, email string) (entity.Barber, *apperr.AppErr)
	FetchBarbers(ctx context.Context) []entity.Barber
	CreateBarber(ctx context.Context, barber *entity.Barber) *apperr.AppErr
	UpdateBarber(ctx context.Context, id string, updateData *entity.Barber) *apperr.AppErr
	DeleteBarber(ctx context.Context, id string) *apperr.AppErr
}
