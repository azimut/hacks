package main

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/miekg/dns"
)

// getResolver get a random trusted resolver
func getResolver() string {
	resolvers := []string{"8.8.8.8:53", "8.8.4.4:53", "1.1.1.1:53", "9.9.9.9:53"}
	return resolvers[rand.Intn(len(resolvers))]
}

// TODO: handle non A, like CNAME
func getAnswers(msg *dns.Msg) []string {
	var ips = make([]string, 0)
	for _, answer := range msg.Answer {
		if a, ok := answer.(*dns.A); ok {
			ips = append(ips, a.A.String())
		}
	}
	return ips
}

// doesResolve checks if A resolves to anything
func doesResolve(domain string) (*dns.Msg, error) {
	var msg dns.Msg
	msg.SetQuestion(dns.Fqdn(domain), dns.TypeA)
	resp, err := dns.Exchange(&msg, getResolver())
	if err != nil {
		return resp, err
	}
	if rcode(resp) != "NOERROR" && rcode(resp) != "NXDOMAIN" {
		return resp, errors.New(fmt.Sprintf("strange return code(%s) for simple A query to domain(%s)", rcode(resp), domain))
	}
	return resp, nil
}

func append_unique(a []string, b []string) []string {
	ret := make([]string, 0)
	ret = append(ret, a...)
	for _, elb := range b {
		found := false
		for _, ela := range a {
			if ela == elb {
				found = true
				break
			}
		}
		if found == false {
			ret = append(ret, elb)
		}
	}
	return ret
}

func rcode(msg *dns.Msg) string {
	return dns.RcodeToString[msg.Rcode]
}

// hasWildcard checks if has wildcard record, checks * and random string, assumes same level
func hasWildcard(root string) (bool, []string, error) {
	asterisk, err := doesResolve("*." + root)
	if err != nil {
		return false, nil, err
	}
	random, err := doesResolve(randStringBytes(rand.Intn(6)+4) + "." + root)
	if err != nil {
		return false, nil, err
	}
	if rcode(asterisk) == "NOERROR" || rcode(random) == "NOERROR" {
		return true, append_unique(getAnswers(asterisk), getAnswers(random)), nil
	}
	if rcode(asterisk) == "NXDOMAIN" && rcode(random) == "NXDOMAIN" {
		return false, nil, nil
	}
	return false, nil, errors.New(fmt.Sprintf("weird return codes (%d) and (%d)", asterisk.Rcode, random.Rcode))
}
