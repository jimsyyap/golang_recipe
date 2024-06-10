package main

import "github.com/miekg/dns"

func main() {
	var msg dns.Msg
	fqdn := dns.Fqdn("example.com")
	msg.SetQuestion(fqdn, dns.TypeA)
	dns.Exchange(&msg, "8.8.8.8:53")
}
