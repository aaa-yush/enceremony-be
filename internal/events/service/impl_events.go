package service

import (
	"context"
	"enceremony-be/internal/database/mysql/models"
	"enceremony-be/pkg/events"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"time"
)

func (i *impl) InsertEvent(ctx context.Context, insertData *events.EventDetails) error {

	repoReq := models.Event{
		Id:          uuid.NewString(),
		Name:        insertData.Name,
		Description: insertData.EventDescription,
		CreatedAt:   insertData.CAt,
		UpdatedAt:   time.Now().UTC(),
		EventDate:   insertData.EventDate,
	}

	err := i.repo.InsertEvent(ctx, &repoReq)
	if err != nil {
		i.logger.Errorw("InsertEvent", zap.Error(err))
		return err
	}

	return nil
}

func (i *impl) GetAllEventsByUserId(ctx context.Context, userId string) (*events.EventList, error) {
	res := events.EventList{Events: []events.EventListItem{}}

	repoRes, err := i.repo.GetAllEventsByUserId(ctx, userId)
	if err != nil {
		i.logger.Errorw("GetAllEventsByUserIdRepo", zap.Error(err))
		return nil, err
	}

	for _, re := range repoRes {
		res.Events = append(res.Events, events.EventListItem{
			events.EventDetailCommons{
				Id: re.Id,

				Creator: &events.Creator{
					Id:   re.UserId,
					Name: "",
				},
				CAt:       re.CreatedAt,
				UAt:       re.UpdatedAt,
				EventDate: re.EventDate,
				Name:      re.Name,
				//ShareUrl:  "",
			},
		})
	}

	return &res, nil
}

func (i *impl) UpdateEvent(ctx context.Context, updateEvent *events.EventDetails) (*events.EventDetails, error) {

	repoReq := &models.Event{
		Id:          updateEvent.Id,
		Name:        updateEvent.Name,
		Description: updateEvent.EventDescription,
		CreatedAt:   updateEvent.CAt,
		UpdatedAt:   time.Now().UTC(),
		EventDate:   updateEvent.EventDate,
	}

	res, err := i.repo.UpdateEvent(ctx, repoReq)
	if err != nil {
		i.logger.Errorw("UpdateEvent", zap.Error(err))
		return nil, err
	}

	return &events.EventDetails{
		EventDetailCommons: events.EventDetailCommons{
			Id:        res.Id,
			Creator:   &events.Creator{Id: res.UserId},
			CAt:       res.CreatedAt,
			UAt:       res.UpdatedAt,
			EventDate: res.EventDate,
			Name:      res.Name,
			//ShareUrl:  "",
		},
		//PhotoUrl:         res.PhotoUrl,
		Products:         nil,
		EventDescription: res.Description,
	}, nil

}

func (i *impl) DeleteEvent(ctx context.Context, eventId string) error {
	return i.repo.DeleteEvent(ctx, eventId)
}
