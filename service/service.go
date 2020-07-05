package service

import (
	"net/http"
	"net/url"

	"github.com/mmzou/geektime-dl/requester"
)

var (
	GeekBangCommURL = &url.URL{
		Scheme: "https",
		Host:   "geekbang.org",
	}
)

//Service geek time service
type Service struct {
	client *requester.HTTPClient
}

//NewService new service
func NewService(gcid, gcess, serviceID string) *Service {
	client := requester.NewHTTPClient()
	client.ResetCookieJar()
	cookies := []*http.Cookie{}
	cookies = append(cookies, &http.Cookie{
		Name:   "GCID",
		Value:  gcid,
		Domain: "." + GeekBangCommURL.Host,
	})
	cookies = append(cookies, &http.Cookie{
		Name:   "GCESS",
		Value:  gcess,
		Domain: "." + GeekBangCommURL.Host,
	})
	cookies = append(cookies, &http.Cookie{
		Name:   "SERVERID",
		Value:  serviceID,
		Domain: "." + GeekBangCommURL.Host,
	})
	client.Jar.SetCookies(GeekBangCommURL, cookies)

	return &Service{client: client}
}

//Cookies get cookies string
func (s *Service) Cookies() map[string]string {
	cookies := s.client.Jar.Cookies(GeekBangCommURL)

	cstr := map[string]string{}

	for _, cookie := range cookies {
		cstr[cookie.Name] = cookie.Value
	}

	return cstr
}
