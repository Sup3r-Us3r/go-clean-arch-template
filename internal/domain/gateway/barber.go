package gateway

import (
	"context"

	"github.com/Sup3r-Us3r/barber-server/internal/domain/apperr"
	"github.com/Sup3r-Us3r/barber-server/internal/domain/entity"
)

type BarberGatewayInterface interface {
	CreateBarber(ctx context.Context, barber *entity.Barber) *apperr.AppErr
	FetchBarbers(ctx context.Context) []entity.Barber
	GetBarberById(ctx context.Context, barberId string) (entity.Barber, *apperr.AppErr)
	GetBarberByEmail(ctx context.Context, email string) (entity.Barber, *apperr.AppErr)
}
