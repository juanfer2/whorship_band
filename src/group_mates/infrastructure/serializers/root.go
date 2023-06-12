package group_mates_serializers

import (
	"time"

	"github.com/google/uuid"
	group_mate_domain "github.com/juanfer2/whorship_band/src/group_mates/domain"
)

type GroupMateSerializer struct {
	ID          uuid.UUID              `json:"id"`
	Name        string                 `json:"name"`
	URLImage    string                 `json:"urlImage"`
	Instruments []InstrumentSerializer `json:"instruments"`

	CreatedAt time.Time `json:"createdAt"`
}

type InstrumentSerializer struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`

	CreatedAt time.Time `json:"createdAt"`
}

func NewGroupMateSerializerFromGroupMate(gr group_mate_domain.GroupMate) *GroupMateSerializer {
	return &GroupMateSerializer{
		ID:       gr.ID,
		Name:     gr.Name,
		URLImage: gr.URLImage,
		//Instruments: gr.Instruments,
		Instruments: NewInstrumentSerializerFromGroupMateInstruments(gr.Instruments),
		CreatedAt:   gr.CreatedAt,
	}
}

func NewGroupMateSerializerFromGroupMateList(
	grs []group_mate_domain.GroupMate,
) []GroupMateSerializer {
	groupMateSerializerList := make([]GroupMateSerializer, 0, len(grs))

	for _, gr := range grs {
		groupMateSerializerList = append(
			groupMateSerializerList,
			*NewGroupMateSerializerFromGroupMate(gr),
		)
	}

	return groupMateSerializerList
}

func NewInstrumentSerializerFromGroupMateInstrument(
	i group_mate_domain.GroupMateInstrument,
) *InstrumentSerializer {
	return &InstrumentSerializer{
		ID:   i.Instrument.ID,
		Name: i.Instrument.Name,
	}
}

func NewInstrumentSerializerFromGroupMateInstruments(
	i []group_mate_domain.GroupMateInstrument,
) []InstrumentSerializer {
	instrumentslist := make([]InstrumentSerializer, 0, len(i))

	for _, value := range i {
		instrumentslist = append(
			instrumentslist,
			*NewInstrumentSerializerFromGroupMateInstrument(value),
		)
	}

	return instrumentslist
}
