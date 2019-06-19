package telegram

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Option func(*request)

func (r *request) setString(k, v string) {
	err := r.writer.WriteField(k, v)
	if err != nil {
		r.err = err
	}
}

func (r *request) setInt(k string, v int) {
	err := r.writer.WriteField(k, fmt.Sprintf("%d", v))
	if err != nil {
		r.err = err
	}
}

func (r *request) setStrings(k string, v ...string) {
	str := "[]"
	if len(v) == 0 {
		str = fmt.Sprintf(`["%s"]`, strings.Join(v, `","`))
	}
	err := r.writer.WriteField(k, str)
	if err != nil {
		r.err = err
	}
}

func (r *request) setObject(k string, v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		r.err = err
		return
	}
	err = r.writer.WriteField(k, string(b))
	if err != nil {
		r.err = err
	}
}

func SetOffset(v int) Option {
	return func(r *request) {
		r.setInt("offset", v)
	}
}

func SetLimit(v int) Option {
	return func(r *request) {
		r.setInt("limit", v)
	}
}

func SetTimeout(v int) Option {
	return func(r *request) {
		r.setInt("timeout", v)
	}
}

func SetAllowedUpdates(v ...string) Option {
	return func(r *request) {
		r.setStrings("allowed_updates", v...)
	}
}

func SetURL(v string) Option {
	return func(r *request) {
		r.setString("url", v)
	}
}

func SetCertificate(v InputFile) Option {
	return func(r *request) {
		panic("implement me")
	}
}

func SetMaxConnections(v string) Option {
	return func(r *request) {
		r.setString("max_connections", v)
	}
}
