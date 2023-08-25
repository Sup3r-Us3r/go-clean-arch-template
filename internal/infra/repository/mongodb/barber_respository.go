package mongodb

import (
	"context"

	"github.com/Sup3r-Us3r/barber-server/internal/domain/apperr"
	"github.com/Sup3r-Us3r/barber-server/internal/domain/entity"
	"github.com/Sup3r-Us3r/barber-server/internal/mapper"
	"github.com/Sup3r-Us3r/barber-server/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BarberRepositoryMongo struct {
	DB *mongo.Database
}

func NewBarberRepositoryMongo(database *mongo.Database) *BarberRepositoryMongo {
	return &BarberRepositoryMongo{
		DB: database,
	}
}

func (brm BarberRepositoryMongo) GetBarberById(ctx context.Context, barberId string) (entity.Barber, *apperr.AppErr) {
	collection := brm.DB.Collection("barber")

	var barberData mapper.BarberMongo
	filter := bson.D{{Key: "_id", Value: barberId}}
	err := collection.FindOne(ctx, filter).Decode(&barberData)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return entity.Barber{}, apperr.ErrBarberNotFound
		}
	}

	result := mapper.MapBarberMongoToEntity(barberData)

	return *result, nil
}

func (brm BarberRepositoryMongo) GetBarberByEmail(ctx context.Context, email string) (entity.Barber, *apperr.AppErr) {
	collection := brm.DB.Collection("barber")

	var barberData mapper.BarberMongo
	filter := bson.D{{Key: "email", Value: email}}
	err := collection.FindOne(ctx, filter).Decode(&barberData)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return entity.Barber{}, apperr.ErrBarberNotFound
		}
	}

	result := mapper.MapBarberMongoToEntity(barberData)

	return *result, nil
}

func (brm BarberRepositoryMongo) FetchBarbers(ctx context.Context) []entity.Barber {
	collection := brm.DB.Collection("barber")

	filter := bson.D{}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return []entity.Barber{}
	}
	defer cursor.Close(ctx)

	var barbers []entity.Barber
	for cursor.Next(ctx) {
		var barber mapper.BarberMongo

		err := cursor.Decode(&barber)
		if err != nil {
			log.Error(err.Error())
			barbers = []entity.Barber{}
			break
		}

		barbers = append(barbers, *mapper.MapBarberMongoToEntity(barber))
	}

	return barbers
}

func (brm *BarberRepositoryMongo) CreateBarber(ctx context.Context, barber *entity.Barber) *apperr.AppErr {
	collection := brm.DB.Collection("barber")

	var barberAlreadyExists mapper.BarberMongo
	filter := bson.D{{Key: "email", Value: barber.Email}}
	collection.FindOne(ctx, filter).Decode(&barberAlreadyExists)

	if barberAlreadyExists.ID != "" {
		return apperr.ErrBarberAlreadyExists
	}

	_, err := collection.InsertOne(ctx, mapper.MapBarberEntityToMongo(barber))
	if err != nil {
		log.Error(err.Error())
		return apperr.NewInternalServerError("unable to register barber")
	}

	return nil
}

func (brm *BarberRepositoryMongo) UpdateBarber(ctx context.Context, id string, updateData *entity.Barber) *apperr.AppErr {
	collection := brm.DB.Collection("barber")

	var barberAlreadyExists mapper.BarberMongo
	filter := bson.D{{Key: "_id", Value: id}}
	collection.FindOne(ctx, filter).Decode(&barberAlreadyExists)

	if barberAlreadyExists.ID == "" {
		return apperr.ErrBarberNotFound
	}

	barberUpdateData := barberAlreadyExists

	if updateData.Name != "" {
		barberUpdateData.Name = updateData.Name
	}

	if updateData.Email != "" {
		barberUpdateData.Email = updateData.Email
	}

	if updateData.Phone != "" {
		barberUpdateData.Phone = updateData.Phone
	}

	_, err := collection.UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: barberUpdateData}})
	if err != nil {
		log.Error(err.Error())
		return apperr.NewInternalServerError("unable to update barber")
	}

	return nil
}

func (brm *BarberRepositoryMongo) DeleteBarber(ctx context.Context, id string) *apperr.AppErr {
	collection := brm.DB.Collection("barber")

	var barberAlreadyExists mapper.BarberMongo
	filter := bson.D{{Key: "_id", Value: id}}
	collection.FindOne(ctx, filter).Decode(&barberAlreadyExists)

	if barberAlreadyExists.ID == "" {
		return apperr.ErrBarberNotFound
	}

	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Error(err.Error())
		return apperr.NewInternalServerError("unable to delete barber")
	}

	return nil
}
