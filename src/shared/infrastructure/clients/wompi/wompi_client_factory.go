package wompi_client

func WompiClientFactory() *WompiClient {
	config := NewConfigClient()
	return NewWompiClient(config)
}
