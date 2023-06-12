package utilities

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

func DataByteToStruct() {

}

func GenerateRandomString() string {
	rand.Seed(time.Now().Unix())

	str := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

	shuff := []rune(str)

	// Shuffling the string
	rand.Shuffle(len(shuff), func(i, j int) {
		shuff[i], shuff[j] = shuff[j], shuff[i]
	})

	valueString := string(shuff)
	// Displaying the random string
	fmt.Println(valueString)

	return valueString
}

func TypeConverter[R any](data any) (*R, error) {
	var result R
	b, err := json.Marshal(&data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, err
	}
	return &result, err
}

func StringToDate(stringDate string) (time.Time, error) {
	date, dateErr := time.Parse("2006-01-02", stringDate)

	if dateErr != nil {
		return date, dateErr
	}

	return date, nil
}
