package apis

import (
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http"
)

var Client HttpInterface = &HttpClient{}

type HttpInterface interface {
	Get(string) ([]byte, error)
}

type HttpClient struct {}

func (c *HttpClient) Get(url string) ([]byte, error) {
	return httpGet(url)
}

func httpGet(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	return ioutil.ReadAll(resp.Body)
}

// mock Test
type MockHttpClient struct {
	mock.Mock
}

func (m *MockHttpClient) Get(url string) ([]byte, error) {
	called := m.Called(url)
	return called.Get(0).([]byte), called.Error(1)
}

var MockGet func(string) ([]byte, error)