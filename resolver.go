package main

import (
	"math/rand"
	"strings"

	"github.com/miekg/dns"
	"github.com/weppos/publicsuffix-go/publicsuffix"
)

// getResolver get a random trusted resolver
func getResolver() string {
	resolvers := []string{"8.8.8.8:53", "8.8.4.4:53"}
	return resolvers[rand.Intn(len(resolvers))]
}

// hasWildcard checks if has wildcard record, checks * and random string, assumes same level
func hasWildcard(domains []*publicsuffix.DomainName) (bool, error) {
	domain := stringify(domains[rand.Intn(len(domains))])
	root := strings.Join(strings.Split(domain, ".")[1:], ".")

	wildcard, err := doesResolve("*." + root)
	if err != nil {
		return false, err
	}
	if wildcard {
		return true, nil
	}

	random, err := doesResolve("asdascwecw." + root)
	if err != nil {
		return false, err
	}
	if random {
		return true, nil
	}
	return false, nil
}

// TODO: see what happens if there is a non A record, does this thing fails
// doesResolve checks if A resolves to anything
func doesResolve(domain string) (bool, error) {
	var msg dns.Msg
	msg.SetQuestion(dns.Fqdn(domain), dns.TypeA)
	resp, err := dns.Exchange(&msg, getResolver())
	if err != nil {
		return false, err
	}
	if len(resp.Answer) < 1 {
		return false, nil
	}
	return true, nil
}
