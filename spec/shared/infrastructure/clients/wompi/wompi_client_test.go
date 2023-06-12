package wompi_client_test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	wompi_client "github.com/juanfer2/whorship_band/src/shared/infrastructure/clients/wompi"
	"github.com/juanfer2/whorship_band/src/shared/infrastructure/config"
	"github.com/stretchr/testify/assert"
)

func TestDoStuffWithTestServer(t *testing.T) {
	config.LoadEnv()

	assert := assert.New(t)
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		assert.Equal(t, req.URL.String(), "some/path")
		fmt.Println("req.URL.String()")
		fmt.Println(req.URL.String())
		// Send response to be tested
		rw.Write([]byte(`OK`))
	}))
	// Close the server when test finishes
	defer server.Close()

	// Use Client & URL from our local test server
	api := wompi_client.WompiClientFactory()
	_, body := api.Fetch(wompi_client.FetchParams{
		Method:        "POST",
		Url:           "/some/path",
		UsePrivateKey: false,
	})
	fmt.Println("Yeahh")
	// fmt.Println(err.Error())
	// assert.OK(t, err)
	rs := bytes.NewBuffer(body).String()
	fmt.Println(rs)

	assert.Equal(t, []byte("OK"), body)

}
