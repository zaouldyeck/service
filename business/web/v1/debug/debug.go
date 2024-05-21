// Package debug provides handler support for the debugging endpoints.
package debug

import (
	"expvar"
	"net/http"
	"net/http/pprof"
)

// Mux registers all the debug routes from the stdlib into a new mux
// bypassing the use of DefaultServeMux. Using the DefaultServeMux would
// be a security risk since it uses an init func.
func Mux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	mux.Handle("/debug/vars", expvar.Handler())

	return mux
}
