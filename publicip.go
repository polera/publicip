package publicip

import (
	"errors"
	"fmt"
	"log"

	"github.com/miekg/dns"
)

/*
GetIP returns the public facing IPv4 address of the requesting client by querying servers
at OpenDNS.
*/
func GetIP() (string, error) {
	config := dns.ClientConfig{Servers: []string{"208.67.220.220", "208.67.222.222"}, Port: "53"}
	dnsClient := new(dns.Client)
	message := new(dns.Msg)
	message.SetQuestion("myip.opendns.com.", dns.TypeA)
	message.RecursionDesired = false
	return doDNSLookup(config, dnsClient, message)
}

func doDNSLookup(config dns.ClientConfig, client *dns.Client, message *dns.Msg) (string, error) {
	var err error
	for _, server := range config.Servers {
		serverAddr := fmt.Sprintf("%s:%s", server, config.Port)
		response, _, err := client.Exchange(message, serverAddr)
		if err != nil {
			log.Printf("Error on DNS lookup: %s", err)
			return "", err
		}
		if response.Rcode != dns.RcodeSuccess {
			errMsg := fmt.Sprintf("DNS call not successful.  Response code: %d", response.Rcode)
			log.Printf(errMsg)
			return "", errors.New(errMsg)
		}
		for _, answer := range response.Answer {
			if aRecord, ok := answer.(*dns.A); ok {
				return aRecord.A.String(), nil
			}
		}
	}
	return "", err
}
