// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// Code generated by cmd/codegen. DO NOT EDIT.
//
// Workarounds Option Group

package ytdlp

import (
	"strconv"
)

// Force the specified encoding (experimental)
//
//   - See [Command.UnsetEncoding], for unsetting the flag.
//   - Encoding maps to cli flags: --encoding=ENCODING.
func (c *Command) Encoding(encoding string) *Command {
	c.addFlag(&Flag{
		ID:   "encoding",
		Flag: "--encoding",
		Args: []string{encoding},
	})
	return c
}

// UnsetEncoding unsets any flags that were previously set by
// [Command.Encoding].
func (c *Command) UnsetEncoding() *Command {
	c.removeFlagByID("encoding")
	return c
}

// Explicitly allow HTTPS connection to servers that do not support RFC 5746 secure
// renegotiation
//
//   - See [Command.UnsetLegacyServerConnect], for unsetting the flag.
//   - LegacyServerConnect maps to cli flags: --legacy-server-connect.
func (c *Command) LegacyServerConnect() *Command {
	c.addFlag(&Flag{
		ID:   "legacy_server_connect",
		Flag: "--legacy-server-connect",
		Args: nil,
	})
	return c
}

// UnsetLegacyServerConnect unsets any flags that were previously set by
// [Command.LegacyServerConnect].
func (c *Command) UnsetLegacyServerConnect() *Command {
	c.removeFlagByID("legacy_server_connect")
	return c
}

// Suppress HTTPS certificate validation
//
//   - See [Command.UnsetCheckCertificates], for unsetting the flag.
//   - NoCheckCertificates maps to cli flags: --no-check-certificates.
func (c *Command) NoCheckCertificates() *Command {
	c.addFlag(&Flag{
		ID:   "no_check_certificate",
		Flag: "--no-check-certificates",
		Args: nil,
	})
	return c
}

// Use an unencrypted connection to retrieve information about the video (Currently
// supported only for YouTube)
//
//   - See [Command.UnsetPreferInsecure], for unsetting the flag.
//   - PreferInsecure maps to cli flags: --prefer-insecure/--prefer-unsecure.
func (c *Command) PreferInsecure() *Command {
	c.addFlag(&Flag{
		ID:   "prefer_insecure",
		Flag: "--prefer-insecure",
		Args: nil,
	})
	return c
}

// UnsetPreferInsecure unsets any flags that were previously set by
// [Command.PreferInsecure].
func (c *Command) UnsetPreferInsecure() *Command {
	c.removeFlagByID("prefer_insecure")
	return c
}

// - See [Command.UnsetUserAgent], for unsetting the flag.
// - UserAgent maps to cli flags: --user-agent=UA (hidden).
func (c *Command) UserAgent(ua string) *Command {
	c.addFlag(&Flag{
		ID:   "user_agent",
		Flag: "--user-agent",
		Args: []string{ua},
	})
	return c
}

// UnsetUserAgent unsets any flags that were previously set by
// [Command.UserAgent].
func (c *Command) UnsetUserAgent() *Command {
	c.removeFlagByID("user_agent")
	return c
}

// - See [Command.UnsetReferer], for unsetting the flag.
// - Referer maps to cli flags: --referer=URL (hidden).
func (c *Command) Referer(url string) *Command {
	c.addFlag(&Flag{
		ID:   "referer",
		Flag: "--referer",
		Args: []string{url},
	})
	return c
}

// UnsetReferer unsets any flags that were previously set by
// [Command.Referer].
func (c *Command) UnsetReferer() *Command {
	c.removeFlagByID("referer")
	return c
}

// Specify a custom HTTP header and its value, separated by a colon ":". You can
// use this option multiple times
//
//   - See [Command.UnsetAddHeaders], for unsetting the flag.
//   - AddHeaders maps to cli flags: --add-headers=FIELD:VALUE.
func (c *Command) AddHeaders(fieldvalue string) *Command {
	c.addFlag(&Flag{
		ID:   "headers",
		Flag: "--add-headers",
		Args: []string{fieldvalue},
	})
	return c
}

// UnsetAddHeaders unsets any flags that were previously set by
// [Command.AddHeaders].
func (c *Command) UnsetAddHeaders() *Command {
	c.removeFlagByID("headers")
	return c
}

// Work around terminals that lack bidirectional text support. Requires bidiv or
// fribidi executable in PATH
//
//   - See [Command.UnsetBidiWorkaround], for unsetting the flag.
//   - BidiWorkaround maps to cli flags: --bidi-workaround.
func (c *Command) BidiWorkaround() *Command {
	c.addFlag(&Flag{
		ID:   "bidi_workaround",
		Flag: "--bidi-workaround",
		Args: nil,
	})
	return c
}

// UnsetBidiWorkaround unsets any flags that were previously set by
// [Command.BidiWorkaround].
func (c *Command) UnsetBidiWorkaround() *Command {
	c.removeFlagByID("bidi_workaround")
	return c
}

// Number of seconds to sleep between requests during data extraction
//
//   - See [Command.UnsetSleepRequests], for unsetting the flag.
//   - SleepRequests maps to cli flags: --sleep-requests=SECONDS.
func (c *Command) SleepRequests(seconds float64) *Command {
	c.addFlag(&Flag{
		ID:   "sleep_interval_requests",
		Flag: "--sleep-requests",
		Args: []string{
			strconv.FormatFloat(seconds, 'g', -1, 64),
		},
	})
	return c
}

// UnsetSleepRequests unsets any flags that were previously set by
// [Command.SleepRequests].
func (c *Command) UnsetSleepRequests() *Command {
	c.removeFlagByID("sleep_interval_requests")
	return c
}

// Number of seconds to sleep before each download. This is the minimum time to
// sleep when used along with --max-sleep-interval
//
//   - See [Command.UnsetSleepInterval], for unsetting the flag.
//   - SleepInterval maps to cli flags: --sleep-interval/--min-sleep-interval=SECONDS.
func (c *Command) SleepInterval(seconds float64) *Command {
	c.addFlag(&Flag{
		ID:   "sleep_interval",
		Flag: "--sleep-interval",
		Args: []string{
			strconv.FormatFloat(seconds, 'g', -1, 64),
		},
	})
	return c
}

// UnsetSleepInterval unsets any flags that were previously set by
// [Command.SleepInterval].
func (c *Command) UnsetSleepInterval() *Command {
	c.removeFlagByID("sleep_interval")
	return c
}

// Maximum number of seconds to sleep. Can only be used along with
// --min-sleep-interval
//
//   - See [Command.UnsetMaxSleepInterval], for unsetting the flag.
//   - MaxSleepInterval maps to cli flags: --max-sleep-interval=SECONDS.
func (c *Command) MaxSleepInterval(seconds float64) *Command {
	c.addFlag(&Flag{
		ID:   "max_sleep_interval",
		Flag: "--max-sleep-interval",
		Args: []string{
			strconv.FormatFloat(seconds, 'g', -1, 64),
		},
	})
	return c
}

// UnsetMaxSleepInterval unsets any flags that were previously set by
// [Command.MaxSleepInterval].
func (c *Command) UnsetMaxSleepInterval() *Command {
	c.removeFlagByID("max_sleep_interval")
	return c
}

// Number of seconds to sleep before each subtitle download
//
//   - See [Command.UnsetSleepSubtitles], for unsetting the flag.
//   - SleepSubtitles maps to cli flags: --sleep-subtitles=SECONDS.
func (c *Command) SleepSubtitles(seconds int) *Command {
	c.addFlag(&Flag{
		ID:   "sleep_interval_subtitles",
		Flag: "--sleep-subtitles",
		Args: []string{
			strconv.Itoa(seconds),
		},
	})
	return c
}

// UnsetSleepSubtitles unsets any flags that were previously set by
// [Command.SleepSubtitles].
func (c *Command) UnsetSleepSubtitles() *Command {
	c.removeFlagByID("sleep_interval_subtitles")
	return c
}
