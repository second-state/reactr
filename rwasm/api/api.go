package api

import "github.com/suborbital/reactr/rwasm/runtime"

// API returns the full Runnable API as runtime Host Functions
func API() []runtime.HostFn {

	api := []runtime.HostFn{
		ReturnResultHandler(),
		ReturnErrorHandler(),
		GetFFIResultHandler(),
		AddFFIVariableHandler(),
		FetchURLHandler(),
		GraphQLQueryHandler(),
		CacheSetHandler(),
		CacheGetHandler(),
		LogMsgHandler(),
		RequestGetFieldHandler(),
		RequestSetFieldHandler(),
		RespSetHeaderHandler(),
		GetStaticFileHandler(),
		DBExecHandler(),
		AbortHandler(),
	}

	return api
}
