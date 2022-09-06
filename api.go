package defectdojo

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/google/go-querystring/query"
)

type APIConfig struct {
	Host     string       // DefectDojo Server, https://example.org
	APIToken string       // DefectDojo V2 API Token
	Client   *http.Client // Optional, can provide a custom HTTP Client, defaults to http.DefaultClient
	Verbose  bool         // Prints stack traces API request errors
}

type DefectDojoAPI struct {
	host     string
	apiToken string
	client   *http.Client
	verbose  bool
}

type RequestOptions struct {
	Offset int
	Limit  int
}

var DefaultRequestOptions = &RequestOptions{
	Offset: 0,
	Limit:  100,
}

var ErrorResponseNon200 = errors.New("api returned a non-200 response code")
var ErrorNoResultsFound = errors.New("results returned no items, please check your search string")
var ErrorInvalidOptions = errors.New("request options is invalid")

func New(config APIConfig) *DefectDojoAPI {
	d := DefectDojoAPI{
		apiToken: config.APIToken,
		verbose:  config.Verbose,
	}

	if config.Client == nil {
		d.client = http.DefaultClient
	}

	if strings.HasSuffix(config.Host, "/") {
		d.host = strings.TrimSuffix(config.Host, "/")
	} else {
		d.host = config.Host
	}
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	return &d
}

func (d *DefectDojoAPI) get(ctx context.Context, path string, options *RequestOptions, in interface{}, out interface{}) error {
	if options.Limit == 0 {
		return ErrorInvalidOptions
	}

	u, err := url.Parse(path)
	if err != nil {
		return err
	}
	q := u.Query()
	q.Add("offset", strconv.Itoa(options.Offset))
	q.Add("limit", strconv.Itoa(options.Limit))
	u.RawQuery = q.Encode()

	res, err := d.request(ctx, http.MethodGet, u.String(), in)
	if err != nil {
		d.errorWithStackTrace(err)
		return err
	}

	if res.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(res.Body)
		err := errors.New("Unexpected HTTP Response. URL: " + res.Request.URL.String() +
			" Expected: " + strconv.Itoa(http.StatusOK) +
			" Got: " + strconv.Itoa(res.StatusCode) +
			" Response: " + string(b))
		d.errorWithStackTrace(err)
		return err
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, out)
	if err != nil {
		return err
	}
	return nil
}

func (d *DefectDojoAPI) post(ctx context.Context, path string, in interface{}, out interface{}) error {
	res, err := d.request(ctx, http.MethodPost, path, in)
	if err != nil {
		d.errorWithStackTrace(err)
		return err
	}

	if res.StatusCode != http.StatusCreated {
		b, _ := io.ReadAll(res.Body)
		err := errors.New("unexpected HTTP Response. URL: " + res.Request.URL.String() +
			" Expected: " + strconv.Itoa(http.StatusCreated) +
			" Got: " + strconv.Itoa(res.StatusCode) +
			" Response: " + string(b))
		d.errorWithStackTrace(err)
		return err
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, out)
	if err != nil {
		return err
	}
	return nil
}

func (d *DefectDojoAPI) patch(ctx context.Context, path string, in interface{}, out interface{}) error {
	res, err := d.request(ctx, http.MethodPatch, path, in)
	if err != nil {
		d.errorWithStackTrace(err)
		return err
	}

	if res.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(res.Body)
		err := errors.New("unexpected HTTP Response. URL: " + res.Request.URL.String() +
			" Expected: " + strconv.Itoa(http.StatusCreated) +
			" Got: " + strconv.Itoa(res.StatusCode) +
			" Response: " + string(b))
		d.errorWithStackTrace(err)
		return err
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, out)
	if err != nil {
		return err
	}
	return nil
}

func (d *DefectDojoAPI) request(ctx context.Context, method string, path string, body interface{}) (*http.Response, error) {
	var req *http.Request
	if body != nil {
		if method == http.MethodPost || method == http.MethodPatch {
			b, err := json.Marshal(body)
			if err != nil {
				return nil, err
			}
			req, err = http.NewRequestWithContext(ctx, method, d.host+path, bytes.NewReader(b))
			if err != nil {
				return nil, err
			}
			req.Header.Add("content-type", "application/json")
		} else {
			values, err := query.Values(body)
			if err != nil {
				return nil, err
			}

			separator := "?"
			if strings.Contains(path, "?") {
				separator = "&"
			}
			u, err := url.Parse(d.host + path + separator + values.Encode())
			if err != nil {
				return nil, err
			}

			req, err = http.NewRequestWithContext(ctx, method, u.String(), http.NoBody)
			if err != nil {
				return nil, err
			}
		}
	} else {
		var err error
		req, err = http.NewRequestWithContext(ctx, method, d.host+path, http.NoBody)
		if err != nil {
			return nil, err
		}
	}

	req.Header.Add("Authorization", "Token "+d.apiToken)
	res, err := d.client.Do(req)
	req.Header.Set("Authorization", "Token <redacted>")

	if d.verbose {
		log.Printf("DefectDojo api request %s %s %v %v\n", req.Method, req.URL.String(), res.StatusCode, err)
	}
	return res, err
}
