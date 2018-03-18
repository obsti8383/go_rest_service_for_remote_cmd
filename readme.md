REST Web Service in Golang (Go) that allows access to predefined commands on the operating system

Currently only static password in code

DON'T USE FOR PRODUCTION USE!

DON'T USE ON COMPUTERS DIRECTLY CONNECTED TO THE INTERNET!

Several important things are missing:
* Real Authentication Backend
* Configuration in config file
* Command line parameters
* Encoding of responses

Since this uses TLS you need a certificate chain (cert.pem) and a Key (key.pem) for your domain in your program directory
