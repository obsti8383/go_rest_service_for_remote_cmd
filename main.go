package main

import (
	//	"context"
	//	"crypto/tls"
	//	"fmt"
	"net/http"
	"os"
	//	"golang.org/x/crypto/acme/autocert"
)

func main() {
	a := &Router{
		RemoteAccessHandler: AuthenticationHandler(LoggingHandler(os.Stdout, new(RemoteAccessHandler))),
	}
	http.ListenAndServeTLS(":8000", "cert.pem", "key.pem", a)

	///// Autocert / Lets Encrypt code
	//	hostPolicy := func(ctx context.Context, host string) error {
	//		// Note: change to your real domain
	//		allowedHost := "www.blablub.de"
	//		if host == allowedHost {
	//			return nil
	//		}
	//		return fmt.Errorf("acme/autocert: only %s host is allowed", allowedHost)
	//	}

	//	certManager := autocert.Manager{
	//		Prompt:     autocert.AcceptTOS,
	//		HostPolicy: hostPolicy,
	//		Cache:      autocert.DirCache("certs"),
	//	}

	//	server := &http.Server{
	//		Addr:    ":443",
	//		Handler: a,
	//		TLSConfig: &tls.Config{
	//			GetCertificate: certManager.GetCertificate,
	//		},
	//	}

	//	go http.ListenAndServe(":80", certManager.HTTPHandler(nil))
	//	server.ListenAndServeTLS("", "")
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}
