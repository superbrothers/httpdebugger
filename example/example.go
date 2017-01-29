/*
Copyright (c) 2017 Kazuki Suda.

For the full copyright and license information, please view the LICENSE
file that was distributed with this source code.
*/
package main

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/superbrothers/httpdebugger"
)

var log = logrus.New()

func main() {
	w := log.Writer()
	defer w.Close()

	client := &http.Client{
		Transport: httpdebugger.NewDebuggingRoundTripper(
			&http.Transport{},
			w,
			httpdebugger.JustURL,
			httpdebugger.URLTiming,
			httpdebugger.CurlCommand,
			httpdebugger.RequestHeaders,
			httpdebugger.ResponseStatus,
			httpdebugger.ResponseHeaders,
		),
	}

	if _, err := client.Get("https://kubernetes.io/"); err != nil {
		log.Fatal(err)
	}
}
