package instrument_serializer

import (
	"time"

	"github.com/google/uuid"
	instrument_domain "github.com/juanfer2/whorship_band/src/instruments/domain"
)

type InstrumentSerializer struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewInstrumentSerializerFromtInstrument(
	i instrument_domain.Instrument,
) *InstrumentSerializer {
	return &InstrumentSerializer{
		ID:   i.ID,
		Name: i.Name,
	}
}

func NewInstrumentSerializerInstrumentList(
	instruments []instrument_domain.Instrument,
) []InstrumentSerializer {
	serializerList := make([]InstrumentSerializer, 0, len(instruments))

	for _, instrument := range instruments {
		serializerList = append(serializerList, *NewInstrumentSerializerFromtInstrument(instrument))
	}

	return serializerList
}
