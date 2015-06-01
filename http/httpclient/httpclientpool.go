package httpclient

import (
	"github.com/niean/gotools/container/nmap"
	"net/http"
	"time"
)

// TODO unit test
// TODO support timeout in GetXXX()

var (
	httpClientMap = nmap.NewSafeMap()
)

func Stop() {
	CloseHttpClients()
}

func GetHttpClient(name string) *http.Client {
	hci, found := httpClientMap.Get(name)
	if !found {
		hci = newHttpClient(5, 20)
		httpClientMap.Put(name, hci)
	}

	return hci.(*http.Client)
}

func CloseHttpClient(name string) {
	if client, found := httpClientMap.Get(name); found {
		if client.(*http.Client).Transport != nil {
			client.(*http.Client).Transport.(*Transport).Close()
		}
	}
}

func CloseHttpClients() {
	for _, key := range httpClientMap.Keys() {
		clienti, found := httpClientMap.Get(key)
		if found {
			client := clienti.(*http.Client)
			if client.Transport != nil {
				client.Transport.(*Transport).Close()
			}
		}
	}
}

// internal
func newHttpClient(connTimeout int, reqTimeout int) *http.Client {
	transport := &Transport{
		ConnectTimeout: time.Duration(connTimeout) * time.Second,
		RequestTimeout: time.Duration(reqTimeout) * time.Second,
	}
	return &http.Client{Transport: transport}
}
