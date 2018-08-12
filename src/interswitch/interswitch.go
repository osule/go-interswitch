package interswitch

import (
	"net/http"
	"time"
	"encoding/json"
	"io/ioutil"
)

type Conf struct {
	Endpoint string
	SchemeName string
	ClientId string
	SharedSecret string
	SignatureMethod string
	Channel string
	Version string
}

type interswitch struct {
	conf *Conf
}

func Client(c *Conf) (i *interswitch){
	return &interswitch {
		conf: c,
	}
}

func (t *interswitch) PaymentMethods() (output interface{}, err error) {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	
	resp, err := httpClient.Get(t.conf.Endpoint)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err := json.Unmarshal(body, &output); err != nil {
		return nil, err
	}

	return output, nil
}
