package main

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/miekg/dns"
	"github.com/weppos/publicsuffix-go/publicsuffix"
)

// stringify DomainName to string
func stringify(domain *publicsuffix.DomainName) string {
	return domain.TRD + "." + domain.SLD + "." + domain.TLD
}

// getResolver get a random trusted resolver
func getResolver() string {
	resolvers := []string{"8.8.8.8:53", "8.8.4.4:53"}
	return resolvers[rand.Intn(len(resolvers))]
}

// hasWildcard checks if has wildcard record, checks * and random string, assumes same level
func hasWildcard(root string) (bool, error) {
	asterisk, err := doesResolve("*." + root)
	if err != nil {
		return false, err
	}
	random, err := doesResolve(randStringBytes(rand.Intn(6)+4) + "." + root)
	if err != nil {
		return false, err
	}
	if asterisk == "NOERROR" || random == "NOERROR" {
		return true, nil
	}
	if asterisk == "NXDOMAIN" && random == "NXDOMAIN" {
		return false, nil
	}
	return false, errors.New(fmt.Sprintf("weird return codes (%s) and (%s)", asterisk, random))
}

// TODO: see what happens if there is a non A record, does this thing fails
// doesResolve checks if A resolves to anything
func doesResolve(domain string) (string, error) {
	var msg dns.Msg
	msg.SetQuestion(dns.Fqdn(domain), dns.TypeA)
	resp, err := dns.Exchange(&msg, getResolver())
	if err != nil {
		return dns.RcodeToString[resp.Rcode], err
	}
	if len(resp.Answer) < 1 {
		return dns.RcodeToString[resp.Rcode], nil
	}
	fmt.Println(resp)
	return dns.RcodeToString[resp.Rcode], nil
}
