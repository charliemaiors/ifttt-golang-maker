package maker

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/juju/errors"
)

const baseURL = "https://maker.ifttt.com/trigger/"

type IFTTTClient struct {
	requestKey string
	httpClient *http.Client
}

func NewClient(reqKey string) (*IFTTTClient, error) {

	if reqKey == "" {
		return nil, errors.New("IFTTT Request string must not be null")
	}

	return &IFTTTClient{
		httpClient: http.DefaultClient,
		requestKey: reqKey,
	}, nil
}

func (client *IFTTTClient) Do(event string, values Values) (string, error) {
	url := baseURL + event + "/with/key/" + client.requestKey
	req, err := http.NewRequest("POST", url, strings.NewReader(values.String()))

	if err != nil {
		fmt.Printf("Got error creating request %v", err)
		return "", err
	}

	resp, err := client.httpClient.Do(req)

	if err != nil {
		fmt.Printf("Got error performing request %v", err)
		return "", err
	}

	responseBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("Error reading response body %v", err)
		return "", err
	}

	return string(responseBody), nil
}
