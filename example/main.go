package main

import (
	"fmt"
	"github.com/visitor-analytics/visa-3as-go-sdk/sdk"
	"log"
)

const pKey = `-----BEGIN PRIVATE KEY-----
MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQDIJUr96Qdc/AkL
xuXoKIpuvzY3FeJOlukVcZEvwQkPpKK9IHWGgNJ4hHXz5egEjWI92/QXjpKQ4A2C
1M9eSWsUqKqmM69S1897yaIUN/RjXCLmHJNEKT9mCgX66U/mfP7FmCN7ApbPrLg4
OBGkGaafH+mgqaImDfaIy1xC3nu5lLm3yD9ziqJHYTXQZs3i8ZaLuuiomu6EP8nR
OCC/mUHofktSGk0PEJjpkLf1lZKi9/JKix3FSkbcrIkPStDwm82C4ABfKPe/oaVl
hgCTGo8DrrW81acGvK+I09Gx32yWNd6MFv+mJmH4LYZ7oApZ1Vt7cM/OcAvisGul
UCy5Hxx5AgMBAAECggEANw5rLqWnrOB37JFuNQrR6d0XoNeYRV4mCAwRkD1drLvx
OG1xZuqQ+y4U9F7OEFKEL9gNsV063DeF//Nih+FSX4B2UPnuxK8xGDBDMaSiyaJS
x1IdAKcIFZ20P21m2VSREPHk8LWpnr84fd6Om8GioCs7LUL8o9M7ei3W1140Urwm
Kn94gkeqv0SA6vs8LdT3O0UPfQRJWSxS9dnjO7qeAgTk+8YJyerHcp2pCpQP/Bkt
ZST1K+Wojo0q+0oj9MvSB7ak4jlAGP6xYCLOdb2ZGMJAC37rtVMJ6G2zWbSsdYxC
GjrjjteIP8FCYBxnW0VxbszhHmxCH1Qt7iCorqUYAQKBgQDrIQk9Q1T0mIxQElEK
7VS73bryascKydoT7s3nMNu4GbebHt8esB7uvIjvnLZs3psmw+fD6R2O4O7BOwl/
I0b40CwmDvIwczCX4Q2nuwlwqRZw3lHEUaZRvkDqbnn8EeuQu+Sc1Z12JkPVgdJY
OqasYCRw9P5yYE3O2dsC4Rn+uQKBgQDZ6U+Ue0woKPaiYmgiXoHNVAYOouEIoLMg
sgJDs7D18mfYdlOVS0hvw+B7Q/QSDhK0t0hhZwMFSTFZD8QkUegwJ03xs+znpJz7
lcqrw9hVpOYwHLn4knTzPbbN8Ev2TqkNXbrLpM+1IZUkUXbMbNE8tPe9Je88FAu/
IFZzhiUrwQKBgEYar3qInMfgw9UL4QX1BRKOZbLpizb4QAE5bkLEGn6ljEy/w56O
vGpJ5Dos62dCZ9gDCRMsahezkPwj8gzqI7sDtmYShrtTXOWrwDqGVaY1g+9bGd4C
yigNJaXAErbJUQbyPpNUTYJwnkEGWATeV9uFPtg+865+cDSAWABfxTRhAoGAfyzS
7O0ofSerCQo1jBlr19F514DnpIllAWfiOnDcji1yvboQ/ch59gBzOn1mLENaV23A
KFheQu98hWXWKvxCbhgCPVWspWRE2e+J4MTjtNgQH3QkdRXEe1FBJt9e9djigJJ7
Oe5t6mA3EoMYuiWn164mB5XkEUQBtwHAcpuPhgECgYB1Wi8BJUO9jAAWhnh0gfcx
HzXFZZnQVZpt1qxslZuqriW1QMC57SCj3bMXrIGg1PO3XX7y6oW5rq6znyaYulml
eI8aL18e6FCA4YUMnBmIW3fS8yW9fkE62fdlOCVN1+ikF6HbP+OXoms1udKefe5R
TTkcmhGs+3xNxuppNxse9w==
-----END PRIVATE KEY-----`

func main() {
	twipla, err := sdk.NewTwipla(sdk.TwiplaArgs{
		Intp: sdk.Intp{
			ID:   "abb7fa69-4915-4702-80b0-bad2e8cb856c",
			PKey: pKey,
		},
		Env: sdk.TwiplaProduction,
	})
	if err != nil {
		log.Fatal(err)
	}

	//p, err := twipla.Package.GetByID("1c7da0ae-44a7-4221-8946-951e4a47dd33")
	//if err != nil {
	//	log.Fatal(err)
	//}

	//p, err := twipla.Package.List()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(p)

	//p, err := twipla.Website.List(sdk.PagArgs{
	//	Page:     0,
	//	PageSize: 100,
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(p)

	p, err := twipla.Website.GetByID("ad8e90dc-483b-4522-a1c6-b40c4c39503e")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*p.TrackingCodeMaxPrivacy, *p.TrackingCode, p.PackageID)

	//c, err := twipla.Customer.List(sdk.PagArgs{
	//	Page:     0,
	//	PageSize: 10,
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(c)

	//err = twipla.Website.New("customer-1", sdk.NewWebsiteArgs{
	//	ExtID:     "ad8e90dc-483b-4522-a1c6-b40c4c39503e",
	//	Domain:    "example.io",
	//	PackageID: "5e89e24b-72f4-4167-b31b-bbdf3922f0a1",
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}

	//p, err := twipla.Customer.List(sdk.PagArgs{
	//	Page:     0,
	//	PageSize: 10,
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(p)

	//p, err := twipla.Customer.GetByID("customer-1")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(p)

	//err := twipla.Customer.NewINTPToken(sdk.NewCustomerArgs{
	//	ExtID: "",
	//	Email: "",
	//	Website: sdk.NewWebsiteArgs{
	//		ExtID:     "",
	//		Domain:    "",
	//		PackageID: "",
	//	},
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}

	//err = twipla.Website.NewINTPToken("abc", sdk.NewWebsiteArgs{
	//	ExtID:     "",
	//	Domain:    "",
	//	PackageID: "",
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}

	//err = twipla.Website.NewINTPToken("abc", sdk.NewWebsiteArgs{
	//	ExtID:     "",
	//	Domain:    "",
	//	PackageID: "",
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}
}
