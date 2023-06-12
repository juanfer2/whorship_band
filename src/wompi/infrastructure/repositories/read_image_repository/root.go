package read_image_repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/juanfer2/whorship_band/src/shared/domain"
	"github.com/juanfer2/whorship_band/src/shared/utilities"
	wompi_entities "github.com/juanfer2/whorship_band/src/wompi/domain/entities"
	wompi_repository "github.com/juanfer2/whorship_band/src/wompi/domain/repository"
)

type ReadImageRepository struct {
	Url string
}

func NewReadImageRepository() wompi_repository.ReadImageRepository {
	return &ReadImageRepository{Url: "http://127.0.0.1:5000"}
}

func (readImageRepository *ReadImageRepository) ExtractText(file domain.File) wompi_entities.CreditCard {
	err, response := readImageRepository.Fetch(domain.FetchParams{
		Url:    "api/v1/images/convert",
		Method: "POST",
	}, file)
	if err != nil {
		fmt.Print("Eerrrror")
		fmt.Print(err.Error())
	}

	var responseBody ImageRedResponse
	errf := json.Unmarshal(response, &responseBody)
	if errf != nil {
		fmt.Println("Error")
		utilities.PrintJson(responseBody)
		log.Fatalln(errf)
	}

	utilities.PrintJson(responseBody)
	return readImageRepository.GetCreditCard(responseBody)
}

func (readImageRepository *ReadImageRepository) GetCreditCard(
	creditCardResponse ImageRedResponse,
) (creditCard wompi_entities.CreditCard) {
	extractCreditCardNumber(&creditCard, creditCardResponse.Text)
	extractCreditCardDateExpiration(&creditCard, creditCardResponse.Text)
	setType(&creditCard)

	utilities.PrintJson(creditCard)

	return
}

func setType(creditCard *wompi_entities.CreditCard) *wompi_entities.CreditCard {
	number := fmt.Sprintf("%s", creditCard.Number)

	if creditCard.Number == "" || creditCard.Number == " " {
		return creditCard
	}

	if number[0:1] == "4" {
		creditCard.Type = "visa"
	}

	if number[0:1] == "5" {
		creditCard.Type = "mastercard"
	}

	return creditCard
}

func extractCreditCardNumber(
	creditCard *wompi_entities.CreditCard, text string,
) *wompi_entities.CreditCard {
	re := regexp.MustCompile(`\d{4}[-\s]?\d{4}[-\s]?\d{4}[-\s]?\d{4}`)

	if re.MatchString(text) {
		creditCardNumber := re.FindString(text)
		creditCard.Number = strings.ReplaceAll(creditCardNumber, " ", "")
	}

	return creditCard
}

func extractCreditCardDateExpiration(
	creditCard *wompi_entities.CreditCard, text string,
) *wompi_entities.CreditCard {
	re := regexp.MustCompile(`\b(0[1-9]|1[0-2])/([0-9]{2})\b`)
	fmt.Println(re.MatchString(text))

	if re.MatchString(text) {
		dateExpirationPattner := re.FindString(text)
		dateExpiration := strings.Split(dateExpirationPattner, "/")
		creditCard.ExpYear = dateExpiration[1]
		creditCard.ExpMonth = dateExpiration[0]
	}

	return creditCard
}

func (readImageRepository *ReadImageRepository) Fetch(params domain.FetchParams, file domain.File) (error, []byte) {
	fmt.Println("-------------Start")
	client := http.Client{}

	//////
	// Abre el archivo que se va a enviar
	newFile, err := os.Open(file.Filename)
	if err != nil {
		fmt.Println("-------------Error when open file")
		fmt.Print(err.Error())
		return err, nil
	}
	defer newFile.Close()
	fmt.Println("-------------Abre el archivo que se va a enviar")

	// Lee el contenido del archivo en memoria
	fileContents, err := ioutil.ReadAll(newFile)
	if err != nil {
		fmt.Print(err.Error())
		return err, nil
	}

	fmt.Println("-------------Lee el contenido del archivo en memoria")

	// Crea un objeto multipart para adjuntar el archivo a la petición
	requestBody := &bytes.Buffer{}
	requestWriter := multipart.NewWriter(requestBody)

	fileWriter, err := requestWriter.CreateFormFile("image", file.Filename)
	if err != nil {
		fmt.Print(err.Error())
		return err, nil
	}

	_, err = io.Copy(fileWriter, bytes.NewReader(fileContents))
	if err != nil {
		fmt.Print(err.Error())
		return err, nil
	}

	// Agrega el parámetro adicional a la petición
	fieldWriter, err := requestWriter.CreateFormField("lang")
	if err != nil {
		fmt.Print(err.Error())
		return err, nil
	}
	_, err = fieldWriter.Write([]byte("es"))
	if err != nil {
		fmt.Print(err.Error())
		return err, nil
	}

	requestWriter.Close()

	fmt.Println("http://127.0.0.1:5000" + "/" + params.Url)
	// Crea una petición POST para enviar el archivo al endpoint de destino
	request, err := http.NewRequest(params.Method, "http://localhost:5000/api/v1/images/convert",
		requestBody)
	if err != nil {
		fmt.Print(err.Error())
		return err, nil
		//os.Exit(1)
	}

	request.Header.Add("Content-Type", requestWriter.FormDataContentType())

	/////////

	response, err := client.Do(request)
	if err != nil {
		fmt.Print(err.Error())
		return err, nil
		//os.Exit(1)
	}

	fmt.Println("-------------response")
	fmt.Println(response)

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return err, nil
	}

	response.Body.Close()

	return err, responseData
}
