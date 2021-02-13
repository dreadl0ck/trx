/*
 * MALTEGO - Go package that provides datastructures for interacting with the Maltego graphical link analysis tool.
 * Copyright (c) 2021 Philipp Mieden <dreadl0ck [at] protonmail [dot] ch>
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package main

import (
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"log"
	"net/http"
)

func setupOauthServer() {
	// OAuth manager
	manager := manage.NewDefaultManager()

	// token memory store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	var domain string
	if *flagTLS {
		domain += "https://"
	} else {
		domain += "https://"
	}
	if *flagDomain != "" {
		domain += *flagDomain
	} else {
		domain += "localhost"
	}

	// client memory store
	clientStore := store.NewClientStore()
	err := clientStore.Set("000000", &models.Client{
		ID:     "000000",
		Secret: "999999", // TODO: pass via env?
		Domain: domain,
	})
	if err != nil {
		log.Println("failed to init client store:", err)
	}
	manager.MapClientStorage(clientStore)

	s := server.NewDefaultServer(manager)
	s.SetAllowGetAccessRequest(true)
	s.SetClientInfoHandler(server.ClientFormHandler)

	s.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	s.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	http.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) {
		err := s.HandleAuthorizeRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		err := s.HandleTokenRequest(w, r)
		if err != nil {
			log.Println("failed to handle token request", err)
		}
	})
}
