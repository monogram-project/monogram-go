//go:build !withweb
// +build !withweb

package main

import "github.com/monogram-project/monogram-go/mg"

// Controlled by a build tag (withweb) to include or exclude web server functionality.
var WithWeb bool = false

// startTestServer starts an HTTP listener on the specified port and opens the browser.
func startTestServer(port string, _ bool, options *mg.FormatOptions) {
	// A stub
}
