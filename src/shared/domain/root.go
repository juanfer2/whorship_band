package domain

type FetchParams struct {
	Url     string
	Method  string
	Body    any
	Headers any
}

type File struct {
	Filename string `json:"Filename"`
	Header   any
	/*Header   struct {
		ContentDisposition []string `json:"Content-Disposition"`
		ContentType        []string `json:"Content-Type"`
	} `json:"Header"`
	*/
	Size int `json:"Size"`
}
