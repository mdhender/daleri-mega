// daleri-mega
//
// Copyright (c) 2019 Michael D Henderson
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"log"

	"github.com/mdhender/xii"
)

// NewApp returns an initialized application.
func NewApp() (app struct {
	root string // root directory
	port int
}, err error) {
	if app.root, err = xii.AsString("DALERI_MEGA_WEBAPP_ROOT", xii.StringOpts{DefaultValue: "./webapp"}); err != nil {
		return app, err
	}
	log.Printf("[app] %-30s == %q\n", "DALERI_MEGA_WEBAPP_ROOT", app.root)

	for _, key := range []string{"DALERI_MEGA_PORT", "PORT"} {
		if app.port, err = xii.AsInt(key, xii.IntOpts{}); err == nil && app.port != 0 {
			log.Printf("[app] %-30s == %d\n", key, app.port)
			break
		}
	}
	if app.port == 0 {
		app.port = 8080
		log.Printf("[app] defaulting port to %d\n", app.port)
	}

	return app, nil
}
