package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/suborbital/reactr/rwasm/runtime"
)

const (
	methodGet int32 = iota
	methodHead
	methodOptions
	methodPost
	methodPut
	methodPatch
	methodDelete
)

const (
	contentTypeJSON        = "application/json"
	contentTypeTextPlain   = "text/plain"
	contentTypeOctetStream = "application/octet-stream"
)

var methodValToMethod = map[int32]string{
	methodGet:     http.MethodGet,
	methodHead:    http.MethodHead,
	methodOptions: http.MethodOptions,
	methodPost:    http.MethodPost,
	methodPut:     http.MethodPut,
	methodPatch:   http.MethodPatch,
	methodDelete:  http.MethodDelete,
}

func FetchURLHandler() runtime.HostFn {
	fn := func(args ...interface{}) (interface{}, error) {
		method := args[0].(int32)
		urlPointer := args[1].(int32)
		urlSize := args[2].(int32)
		bodyPointer := args[3].(int32)
		bodySize := args[4].(int32)
		ident := args[5].(int32)

		ret := fetch_url(method, urlPointer, urlSize, bodyPointer, bodySize, ident)

		return ret, nil
	}

	return runtime.NewHostFn("fetch_url", 6, true, fn)
}

func fetch_url(method int32, urlPointer int32, urlSize int32, bodyPointer int32, bodySize int32, identifier int32) int32 {
	// fetch makes a network request on bahalf of the wasm runner.
	// fetch writes the http response body into memory starting at returnBodyPointer, and the return value is a pointer to that memory
	inst, err := runtime.InstanceForIdentifier(identifier, true)
	if err != nil {
		runtime.InternalLogger().Error(errors.Wrap(err, "[rwasm] alert: failed to InstanceForIdentifier"))
		return -1
	}

	httpMethod, exists := methodValToMethod[method]
	if !exists {
		runtime.InternalLogger().ErrorString("invalid method provided: ", method)
		return -2
	}

	urlBytes := inst.ReadMemory(urlPointer, urlSize)

	// the URL is encoded with headers added on the end, each seperated by ::
	// eg. https://google.com/somepage::authorization:bearer qdouwrnvgoquwnrg::anotherheader:nicetomeetyou
	urlParts := strings.Split(string(urlBytes), "::")
	urlString := urlParts[0]

	headers, err := parseHTTPHeaders(urlParts)
	if err != nil {
		runtime.InternalLogger().Error(errors.Wrap(err, "could not parse URL headers"))
		return -2
	}

	body := inst.ReadMemory(bodyPointer, bodySize)

	if len(body) > 0 {
		if headers.Get("Content-Type") == "" {
			headers.Add("Content-Type", contentTypeOctetStream)
		}
	}

	// wrap everything in a function so any errors get collected
	resp, err := func() ([]byte, error) {
		// filter the request through the capabilities
		resp, err := inst.Ctx().HTTPClient.Do(inst.Ctx().Auth, httpMethod, urlString, body, *headers)
		if err != nil {
			runtime.InternalLogger().Error(errors.Wrap(err, "failed to Do request"))
			return nil, err
		}

		defer resp.Body.Close()
		respBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			runtime.InternalLogger().Error(errors.Wrap(err, "failed to Read response body"))
		}

		if resp.StatusCode > 299 {
			runtime.InternalLogger().Debug("runnable's http request returned non-200 response:", resp.StatusCode)
			return nil, fmt.Errorf("%d: %s", resp.StatusCode, string(respBytes))
		}

		return respBytes, nil
	}()

	result, err := inst.Ctx().SetFFIResult(resp, err)
	if err != nil {
		runtime.InternalLogger().ErrorString("[rwasm] failed to SetFFIResult", err.Error())
		return -1
	}

	return result.FFISize()
}

func parseHTTPHeaders(urlParts []string) (*http.Header, error) {
	headers := &http.Header{}

	if len(urlParts) > 1 {
		for _, p := range urlParts[1:] {
			headerParts := strings.Split(p, ":")
			if len(headerParts) != 2 {
				return nil, errors.New("header was not formatted correctly")
			}

			headers.Add(headerParts[0], headerParts[1])
		}
	}

	return headers, nil
}
