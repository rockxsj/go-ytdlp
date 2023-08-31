// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// Code generated by cmd/codegen. DO NOT EDIT.
//
// Network Option Group

package ytdlp

type NetworkBuilder struct {
	parent *Command
}

type proxyFlag struct {
	args []string
}

var _ Flag = (*proxyFlag)(nil) // ensure proxyFlag implements Flag interface.

func (f *proxyFlag) ID() string {
	return "proxy"
}

func (f *proxyFlag) String() string {
	return "TODO"
}

func (f *proxyFlag) AsFlag() []string {
	return []string{"TODO"}
}

// Use the specified HTTP/HTTPS/SOCKS proxy. To enable SOCKS proxy, specify a
// proper scheme, e.g. socks5://user:pass@127.0.0.1:1080/. Pass in an empty string
// (--proxy "") for direct connection
//
// Proxy maps to cli flags: --proxy=URL.
func (ff *NetworkBuilder) Proxy(url string) *NetworkBuilder {
	ff.parent.addFlag(&proxyFlag{
		args: []string{url},
	})
	return ff
}

type socketTimeoutFlag struct {
	args []float64
}

var _ Flag = (*socketTimeoutFlag)(nil) // ensure socketTimeoutFlag implements Flag interface.

func (f *socketTimeoutFlag) ID() string {
	return "socket_timeout"
}

func (f *socketTimeoutFlag) String() string {
	return "TODO"
}

func (f *socketTimeoutFlag) AsFlag() []string {
	return []string{"TODO"}
}

// Time to wait before giving up, in seconds
//
// SocketTimeout maps to cli flags: --socket-timeout=SECONDS.
func (ff *NetworkBuilder) SocketTimeout(seconds float64) *NetworkBuilder {
	ff.parent.addFlag(&socketTimeoutFlag{
		args: []float64{seconds},
	})
	return ff
}

type sourceAddressFlag struct {
	args []string
}

var _ Flag = (*sourceAddressFlag)(nil) // ensure sourceAddressFlag implements Flag interface.

func (f *sourceAddressFlag) ID() string {
	return "source_address"
}

func (f *sourceAddressFlag) String() string {
	return "TODO"
}

func (f *sourceAddressFlag) AsFlag() []string {
	return []string{"TODO"}
}

// Client-side IP address to bind to
//
// SourceAddress maps to cli flags: --source-address=IP.
func (ff *NetworkBuilder) SourceAddress(ip string) *NetworkBuilder {
	ff.parent.addFlag(&sourceAddressFlag{
		args: []string{ip},
	})
	return ff
}

type forceIpv4Flag struct {
}

var _ Flag = (*forceIpv4Flag)(nil) // ensure forceIpv4Flag implements Flag interface.

func (f *forceIpv4Flag) ID() string {
	return "source_address"
}

func (f *forceIpv4Flag) String() string {
	return "TODO"
}

func (f *forceIpv4Flag) AsFlag() []string {
	return []string{"TODO"}
}

type forceIpv6Flag struct {
}

var _ Flag = (*forceIpv6Flag)(nil) // ensure forceIpv6Flag implements Flag interface.

func (f *forceIpv6Flag) ID() string {
	return "source_address"
}

func (f *forceIpv6Flag) String() string {
	return "TODO"
}

func (f *forceIpv6Flag) AsFlag() []string {
	return []string{"TODO"}
}

type enableFileUrlsFlag struct {
}

var _ Flag = (*enableFileUrlsFlag)(nil) // ensure enableFileUrlsFlag implements Flag interface.

func (f *enableFileUrlsFlag) ID() string {
	return "enable_file_urls"
}

func (f *enableFileUrlsFlag) String() string {
	return "TODO"
}

func (f *enableFileUrlsFlag) AsFlag() []string {
	return []string{"TODO"}
}

// Enable file:// URLs. This is disabled by default for security reasons.
//
// EnableFileUrls maps to cli flags: --enable-file-urls.
func (ff *NetworkBuilder) EnableFileUrls() *NetworkBuilder {
	ff.parent.addFlag(&enableFileUrlsFlag{})
	return ff
}
