package publicip

import (
	"testing"

	"github.com/miekg/dns"
)

func TestGetAddress(t *testing.T) {
	_, err := GetIP()
	if err != nil {
		t.Errorf("Error: %s\n", err)
	}
}

func TestGenerateClientError(t *testing.T) {
	config := dns.ClientConfig{Servers: []string{"0.0.0.0"}, Port: "53"}
	dnsClient := new(dns.Client)
	message := new(dns.Msg)
	message.SetQuestion("myip.opendns.com.", dns.TypeA)
	message.RecursionDesired = false
	_, err := doDNSLookup(config, dnsClient, message)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestGenerateLookupError(t *testing.T) {
	config := dns.ClientConfig{Servers: []string{"208.67.220.220", "208.67.222.222"}, Port: "53"}
	dnsClient := new(dns.Client)
	message := new(dns.Msg)
	message.SetQuestion("notmyip.opendns.com.", dns.TypeA)
	message.RecursionDesired = false
	_, err := doDNSLookup(config, dnsClient, message)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
