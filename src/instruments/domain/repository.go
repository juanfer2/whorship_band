package intrument_domain

type InstrumentRepository interface {
	FindBy(id int) Instrument
	FindByUuid(id string) (Instrument, error)
	Create(instrument Instrument) (Instrument, error)
	Update(instrument Instrument) (Instrument, error)
	UpdateByUuid(id string, updateRecord Instrument) (Instrument, error)
	CreateInBatches(instruments []Instrument)
	All() []Instrument
	Delete(instrument Instrument) Instrument
	WhereBy(query interface{}, arg ...interface{}) Instrument
}
