package domainutil

import (
	"errors"
	"strings"

	"golang.org/x/net/publicsuffix"
)

//Split  split www.abc.com =>  host:www  sld:abc  tld:com err:nil
func Split(s string) (host, sld, tld string, err error) {
	s, err = Punycode(s)
	if err != nil {
		return
	}
	var ok bool
	tld, ok = publicsuffix.PublicSuffix(s)
	if !ok {
		err = errors.New("publicsuffix")
		return
	}
	var domain string
	domain, err = publicsuffix.EffectiveTLDPlusOne(s)
	if err != nil {
		return
	}
	host = strings.TrimSuffix(strings.TrimSuffix(s, domain), ".")
	host, err = Unicode(host)
	if err != nil {
		return
	}
	sld = strings.TrimSuffix(strings.TrimSuffix(domain, tld), ".")
	sld, err = Unicode(sld)
	if err != nil {
		return
	}
	tld, err = Unicode(tld)
	return
}
