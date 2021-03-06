/*
 * Copyright (c) 2016 Giovanni Liboni
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"),
 * to deal in the Software without restriction, including without limitation
 * the rights to use, copy, modify, merge, publish, distribute, sublicense,
 * and/or sell copies of the Software, and to permit persons to whom
 * the Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
 * EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
 * MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
 * IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
 * DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
 * TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
 * OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */
package web

import (
	"net/http"
	"strings"

	"github.com/gorilla/schema"
)

// FillValuesFromForm serve per riempire la struttura passata con i campi relativi presenti nel form
func FillValuesFromForm(i interface{}, r *http.Request) error {
	decoder := schema.NewDecoder()
	// Parse the body depending on the content type.
	contentType := strings.ToLower(strings.TrimSpace(strings.Split(r.Header.Get("Content-Type"), ";")[0]))
	switch contentType {
	case "application/x-www-form-urlencoded":
		// Typical form.
		if err := r.ParseForm(); err != nil {
			return err
		}
		decoder.Decode(i, r.Form)

	case "multipart/form-data":
		// Multipart form.
		// TODO: Extract the multipart form param so app can set it.
		if err := r.ParseMultipartForm(32 << 20 /* 32 MB */); err != nil {
			return err
		}
		decoder.Decode(i, r.MultipartForm.Value)
	}

	return nil
}
