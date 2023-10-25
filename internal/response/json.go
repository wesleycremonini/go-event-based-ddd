package response

import (
	"encoding/json"
	"net/http"
)

type Envelope map[string]any

func JSONWithHeaders(w http.ResponseWriter, status int, data any, headers http.Header) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	js = append(js, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

func Success(w http.ResponseWriter, status int, msg string, data any) error {
	envelope := &Envelope{"status": "success", "message": msg, "data": data}

	return JSONWithHeaders(w, status, envelope, nil)
}

func Error(w http.ResponseWriter, status int, msg string, data any) error {
	envelope := &Envelope{"status": "error", "message": msg, "data": data}

	return JSONWithHeaders(w, status, envelope, nil)
}

func SuccessWithHeaders(w http.ResponseWriter, status int, msg string, data any, headers http.Header) error {
	envelope := &Envelope{"status": "success", "message": msg, "data": data}

	return JSONWithHeaders(w, status, envelope, headers)
}

func ErrorWithHeaders(w http.ResponseWriter, status int, msg string, data any, headers http.Header) error {
	envelope := &Envelope{"status": "error", "message": msg, "data": data}

	return JSONWithHeaders(w, status, envelope, headers)
}
