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
