# security.txt scanning

Scan websites for security.txt files

## Usage

`go build sectxtscanner.go`

`./sectxtscanner.go <input domains file>`

## Why

This project aims to implement a simple way to scan a list of domains for security.txt files (as per [RFC 9116: A File Format to Aid in Security Vulnerability Disclosure](https://www.rfc-editor.org/rfc/rfc9116.html), and act as a central repository to check for security.txt files for Australian domains (but can be used with any domains).

It's important to have a security.txt file so that responsible disclosure of vulnerabilities belonging to your organization can go to the right place, like when you are vulnerable to CVEs being exploited by ransomware that need urgent patching.

A very small percentage organizations actually implement this, although it's very easy to do. See [securitytxt.org](https://securitytxt.org/) for how to do it.

