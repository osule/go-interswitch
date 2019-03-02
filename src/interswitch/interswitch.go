package interswitch

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"time"
)

// Conf is the configuration for the client.
type Conf struct {
	Endpoint        string
	SchemeName      string
	ClientID        string
	SharedSecret    string
	SignatureMethod string
	Channel         string
	Version         string
}

// Client interface for callers.
type Client interface {
	PaymentMethods() (interface{}, error)
}

type client struct {
	conf     *Conf
	endpoint *endpoint
	http     *http.Client
}

type endpoint struct {
	URL *url.URL
}

func (e *endpoint) String() string {
	return e.URL.String()
}
func validEndpoint(v string) (*endpoint, error) {
	var ve endpoint
	u, err := url.Parse(string(v))
	if err != nil {
		return &ve, err
	}
	ve.URL = u
	return &ve, err
}

func (e endpoint) Path(p string) string {
	u := e.URL
	u.Path = path.Join(u.Path, p)
	return u.String()
}

func request(req func(string) (*http.Response, error), ep string) (interface{}, error) {
	var output interface{}
	res, err := req(ep)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return body, err
	}

	err = json.Unmarshal(body, &output)
	return output, err
}

// ClientUsingConf creates Client and configures it.
func ClientUsingConf(c *Conf) (Client, error) {
	e, err := validEndpoint(c.Endpoint)
	http := &http.Client{Timeout: time.Second * 10}
	return client{conf: c, endpoint: e, http: http}, err
}

// PaymentMethods lists available payment methods.
func (c client) PaymentMethods() (interface{}, error) {
	return request(c.http.Get, c.endpoint.Path("paymethods"))
}
