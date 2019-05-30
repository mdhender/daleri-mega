// qd8pb: a user friendly content management system
//
// Copyright (C) 2019  Michael D Henderson
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// * Redistributions of source code must retain the above copyright notice, this
//   list of conditions and the following disclaimer.
//
// * Redistributions in binary form must reproduce the above copyright notice,
//   this list of conditions and the following disclaimer in the documentation
//   and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

// Package static serves files.
package static

import (
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

// Handler returns a new static file handler.
func Handler(prefix string, root string, spaRouting bool) http.Handler {
	log.Println("[static] initializing")
	defer log.Println("[static] initialized")

	log.Printf("[static] strip: %q\n", prefix)
	log.Printf("[static]  root: %q\n", root)

	var indexFile string
	if spaRouting {
		indexFile = root + "/index.html"
		log.Printf("[static] index: %q\n", indexFile)
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}

		file := root + path.Clean("/"+strings.TrimPrefix(r.URL.Path, prefix))
		log.Printf("[static] %q\n", file)

		stat, err := os.Stat(file)
		if err != nil {
			if spaRouting {
				// try serving index file for SPA routing instead
				if stat, err := os.Stat(indexFile); err == nil {
					if rdr, err := os.Open(indexFile); err == nil {
						defer rdr.Close()
						http.ServeContent(w, r, file, stat.ModTime(), rdr)
						return
					}
				}
			}
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		// we never want to give a directory listing, so change raw directory request to fetch the index.html instead.
		if stat.IsDir() {
			file = file + "/index.html"
			stat, err = os.Stat(file)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
				return
			}
		}

		// only serve regular files (this avoids serving a directory named index.html)
		if !stat.Mode().IsRegular() {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		// pretty sure that we have a regular file at this point.
		rdr, err := os.Open(file)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		defer rdr.Close()

		http.ServeContent(w, r, file, stat.ModTime(), rdr)
	})
}
