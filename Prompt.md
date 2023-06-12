Tu tarea es generar código en golang dada las siguientes instrucciones

generar diferentes archivos basados en un struct

-> genera un struct que se llame Instrument con los siguientes campos
---
ID,
Name
---

-> generar una interface con los siguientes campos sin ningula implementación

type <NombreDelStruct>Repository interface {
	FindBy(id int) GroupMate
	FindByUuid(id string) (GroupMate, error)
	Create(groupMate GroupMate) (GroupMate, error)
	Update(groupMate GroupMate) (GroupMate, error)
	UpdateByUuid(id string, updateRecord GroupMate) (GroupMate, error)
	CreateInBatches(groupMates []GroupMate)
	All() []GroupMate
	Delete(groupMate GroupMate) GroupMate
	WhereBy(query any, arg ...any) GroupMate
}


-> generar un struct llamada repository con el siguiente template

package <nombre_del_dominio>_repositories

type <NombreDelStruct>PGRepository struct {
	postgres.PostgresRepository[group_mate_domain.NombreDelStruct]
}

func New<NombreDelStruct>PGRepository() group_mate_domain.NombreDelStructRepository {
	repository := NombreDelStructPGRepository{}
	repository.PostgresClient = postgres.CreateClientFactory()

	return &repository
}

func (p *<NombreDelStruct>PGRepository) Where(query any, arg ...any) (NombreDelStruct []group_mate_domain.NombreDelStruct) {
	return
}

-> Genera un struct llamado use_case con el siguiente template

type <Name>UseCase struct {
	repository group_mate_domain.<Name>Repository
}

func New<Name>UseCase(
	Repository group_mate_domain.<Name>Repository,
) *<Name>UseCase {
	return &<Name>UseCase{repository: Repository}
}

func (uc *<Name>UseCase) FindById(id int) group_mate_domain.<Name> {
	return uc.repository.FindBy(id)
}

func (uc *<Name>UseCase) Create(
	<Name> group_mate_domain.<Name>,
) (group_mate_domain.<Name>, error) {
	new<Name>, err := uc.repository.Create(<Name>)

	return new<Name>, err
}

func (uc *<name>UseCase) CreateInBatches(<name>s []group_mate_domain.<name>) {
	uc.repository.CreateInBatches(<name>s)
}

func (uc *<name>UseCase) GetAll() []group_mate_domain.<name> {
	return uc.repository.All()
}

func (uc *<name>UseCase) Delete(id string) (group_mate_domain.<name>, error) {
	<name>, err := uc.repository.FindByUuid(id)

	if err != nil {
		return <name>, err
	}

	uc.repository.Delete(<name>)

	return <name>, nil
}

func (uc *<name>UseCase) Update(
	id string, <name> group_mate_domain.<name>,
) (group_mate_domain.<name>, error) {
	updateGm, err := uc.repository.UpdateByUuid(id, <name>)

	if err != nil {
		return <name>, err
	}

	return updateGm, nil
}

-> Genera un controller con el siguiente template

type <name>Controller struct {
	<name>UseCase *<name>_application.<name>UseCase
}

func New<name>Controller(
	<name>UseCase *<name>_application.<name>UseCase,
) *<name>Controller {
	return &<name>Controller{<name>UseCase: <name>UseCase}
}

func (<name>Controller *<name>Controller) All<name>(c *fiber.Ctx) error {
	<name>s := <name>Controller.<name>UseCase.GetAll()

	return c.JSON(<name>_serializers.New<name>SerializerFrom<name>List(<name>s))
}

func (<name>Controller *<name>Controller) Create<name>(c *fiber.Ctx) error {
	<name>Params, err := getParams<name>(c)
	if err != nil {
		errorHandler := middlewares.ErrorHandlers(err)
		return c.Status(403).JSON(errorHandler)
	}

	<name>, err := <name>Controller.<name>UseCase.Create(<name>Params)
	if err != nil {
		errorHandler := middlewares.ErrorHandlers(err)
		return c.Status(errorHandler.Status).JSON(errorHandler)
	}

	<name>_serializers.New<name>SerializerFrom<name>(<name>)

	return c.JSON(<name>_serializers.New<name>SerializerFrom<name>(<name>))
}

func (<name>Controller *<name>Controller) Update<name>(c *fiber.Ctx) error {
	id := c.Params("id")
	<name>Params, err := getParams<name>(c)
	if err != nil {
		errorHandler := middlewares.ErrorHandlers(err)
		return c.Status(403).JSON(errorHandler)
	}

	<name>, err := <name>Controller.<name>UseCase.Update(id, <name>Params)

	if err != nil {
		errorHandler := middlewares.ErrorHandlers(err)
		return c.Status(errorHandler.Status).JSON(errorHandler)
	}

	return c.JSON(<name>_serializers.New<name>SerializerFrom<name>(<name>))
}

func (<name>Controller *<name>Controller) Delete<name>(c *fiber.Ctx) error {
	id := c.Params("id")
	<name>, err := <name>Controller.<name>UseCase.Delete(id)
	if err != nil {
		errorHandler := middlewares.ErrorHandlers(err)
		return c.Status(errorHandler.Status).JSON(errorHandler)
	}

	return c.JSON(<name>_serializers.New<name>SerializerFrom<name>(<name>))
}

-> genera un struct que se llame Instrument con los siguientes campos
---
ID,
Name
---

-> Genera un struct con el siguiente template con y los campos


type NameStructSerializer struct {
	ID          uuid.UUID              `json:"id"`
	...
	CreatedAt time.Time `json:"createdAt"`
}

func NewNameStructSerializerName(gr group_mate_domain.NameStruct) *NameStructSerializer {
	return &NameStructSerializer{
		ID:       gr.ID,
		Name:     gr.Name,
	}
}

func NewNameStructSerializerNameList(
	grs []group_mate_domain.NameStruct,
) []NameStructSerializer {
	NameStructSerializerList := make([]NameStructSerializer, 0, len(grs))

	for _, gr := range grs {
		NameStructSerializerList = append(
			NameStructSerializerList,
			*NewNameStructSerializerName(gr),
		)
	}

	return nameStructSerializerList
}


-> generar un struct y una función con el siguiente template

type GroupMateParams struct {
	...
	Name        string                                  `json:"name"`
}

func getParamsNameStruct(c *fiber.Ctx) (group_mate_domain.NameStruct, error) {
	var NameStructParams NameStructParams
	if err := c.BodyParser(&NameStructParams); err != nil {
		return group_mate_domain.NameStruct{}, err
	}

	utilities.PrintJson(NameStructParams)

	fmt.Println(NameStructParams)

	return group_mate_domain.NameStcut{
		Name:        NameStructParams.Name,
		URLImage:    NameStructParams.URLImage,
		Instruments: NameStructParams.Instruments,
	}, nil
}

donde dice name struct por el nombre instrument
