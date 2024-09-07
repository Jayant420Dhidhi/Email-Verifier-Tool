package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("%-30s %-6s %-6s %-50s %-6s %-50s\n", "Domain", "MX", "SPF", "SPF Record", "DMARC", "DMARC Record")
	fmt.Println(strings.Repeat("-", 150))

	for scanner.Scan() {
		domain := extractDomain(scanner.Text())
		if domain != "" {
			checkDomain(domain)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error: could not read from input: %v", err)
	}
}

// extractDomain extracts the domain from an email address or returns the input if it's already a domain
func extractDomain(input string) string {
	// Check if the input contains an '@', indicating it's likely an email address
	atIndex := strings.LastIndex(input, "@")
	if atIndex != -1 && atIndex+1 < len(input) {
		return input[atIndex+1:] // Extract the domain part
	}
	return input // Return the input as-is if it's not an email address
}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	// Check MX records
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error: failed to lookup MX records for domain %s: %v", domain, err)
	} else if len(mxRecords) > 0 {
		hasMX = true
	}

	// Check TXT records for SPF
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error: failed to lookup TXT records for domain %s: %v", domain, err)
	} else {
		for _, record := range txtRecords {
			if strings.HasPrefix(record, "v=spf1") {
				hasSPF = true
				spfRecord = record
				break
			}
		}
	}

	// Check TXT records for DMARC
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error: failed to lookup DMARC records for domain %s: %v", domain, err)
	} else {
		for _, record := range dmarcRecords {
			if strings.HasPrefix(record, "v=DMARC1") {
				hasDMARC = true
				dmarcRecord = record
				break
			}
		}
	}

	fmt.Printf("%-30s %-6t %-6t %-50s %-6t %-50s\n", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}
