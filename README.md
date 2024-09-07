# Email Verifier Tool

This Go application checks DNS records for a given list of domains to identify the presence of MX, SPF, and DMARC records. It reads domains from standard input, performs DNS lookups, and outputs the results in a tabular format.

## Features

- **MX Records**: Checks for the presence of MX (Mail Exchange) records for the domain.
- **SPF Records**: Checks for the presence of SPF (Sender Policy Framework) records in TXT records.
- **DMARC Records**: Checks for the presence of DMARC (Domain-based Message Authentication, Reporting, and Conformance) records in TXT records.

## Usage

1. **Run the Program**: Execute the program from the command line. You can provide a list of email addresses or domain names via standard input.

    ```sh
    go run main.go
    ```

2. **Input**: Enter email addresses or domain names, one per line. Press `Ctrl+D` (or `Ctrl+Z` on Windows) to end input and trigger the checks.

3. **Output**: The program will output a table with columns for domain, MX record presence, SPF record presence, SPF record content, DMARC record presence, and DMARC record content.


## Functions

- `extractDomain(input string) string`: Extracts the domain from an email address or returns the input if it's already a domain.
- `checkDomain(domain string)`: Checks MX, SPF, and DMARC records for the given domain and prints the results.

## Error Handling

The program logs errors encountered during DNS lookups to standard error but continues to process subsequent domains.

## Dependencies

- Go standard library

## EXAMPLE
<img width="916" alt="Screenshot 2024-09-07 at 7 37 38 PM" src="https://github.com/user-attachments/assets/d952eb89-36ec-4591-8509-4c6432e3e728">
