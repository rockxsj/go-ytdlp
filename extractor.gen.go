// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// Code generated by cmd/codegen. DO NOT EDIT.
//
// Extractor Option Group

package ytdlp

type ExtractorBuilder struct {
	parent *Command
}

type extractorRetriesFlag struct {
	args []string
}

var _ Flag = (*extractorRetriesFlag)(nil) // ensure extractorRetriesFlag implements Flag interface.

func (f *extractorRetriesFlag) ID() string {
	return "extractor_retries"
}

func (f *extractorRetriesFlag) String() string {
	return "TODO"
}

func (f *extractorRetriesFlag) AsFlag() []string {
	return []string{"TODO"}
}

// Number of retries for known extractor errors (default is %default), or
// "infinite"
//
// ExtractorRetries maps to cli flags: --extractor-retries=RETRIES.
func (ff *ExtractorBuilder) ExtractorRetries(retries string) *ExtractorBuilder {
	ff.parent.addFlag(&extractorRetriesFlag{
		args: []string{retries},
	})
	return ff
}

type allowDynamicMpdFlag struct {
}

var _ Flag = (*allowDynamicMpdFlag)(nil) // ensure allowDynamicMpdFlag implements Flag interface.

func (f *allowDynamicMpdFlag) ID() string {
	return "dynamic_mpd"
}

func (f *allowDynamicMpdFlag) String() string {
	return "TODO"
}

func (f *allowDynamicMpdFlag) AsFlag() []string {
	return []string{"TODO"}
}

// Process dynamic DASH manifests (default) (Alias: --no-ignore-dynamic-mpd)
//
// AllowDynamicMpd maps to cli flags: --allow-dynamic-mpd/--no-ignore-dynamic-mpd.
func (ff *ExtractorBuilder) AllowDynamicMpd() *ExtractorBuilder {
	ff.parent.addFlag(&allowDynamicMpdFlag{})
	return ff
}

type ignoreDynamicMpdFlag struct {
}

var _ Flag = (*ignoreDynamicMpdFlag)(nil) // ensure ignoreDynamicMpdFlag implements Flag interface.

func (f *ignoreDynamicMpdFlag) ID() string {
	return "dynamic_mpd"
}

func (f *ignoreDynamicMpdFlag) String() string {
	return "TODO"
}

func (f *ignoreDynamicMpdFlag) AsFlag() []string {
	return []string{"TODO"}
}

// Do not process dynamic DASH manifests (Alias: --no-allow-dynamic-mpd)
//
// IgnoreDynamicMpd maps to cli flags: --ignore-dynamic-mpd/--no-allow-dynamic-mpd.
func (ff *ExtractorBuilder) IgnoreDynamicMpd() *ExtractorBuilder {
	ff.parent.addFlag(&ignoreDynamicMpdFlag{})
	return ff
}

type hlsSplitDiscontinuityFlag struct {
}

var _ Flag = (*hlsSplitDiscontinuityFlag)(nil) // ensure hlsSplitDiscontinuityFlag implements Flag interface.

func (f *hlsSplitDiscontinuityFlag) ID() string {
	return "hls_split_discontinuity"
}

func (f *hlsSplitDiscontinuityFlag) String() string {
	return "TODO"
}

func (f *hlsSplitDiscontinuityFlag) AsFlag() []string {
	return []string{"TODO"}
}

// Split HLS playlists to different formats at discontinuities such as ad breaks
//
// HlsSplitDiscontinuity maps to cli flags: --hls-split-discontinuity.
func (ff *ExtractorBuilder) HlsSplitDiscontinuity() *ExtractorBuilder {
	ff.parent.addFlag(&hlsSplitDiscontinuityFlag{})
	return ff
}

type noHlsSplitDiscontinuityFlag struct {
}

var _ Flag = (*noHlsSplitDiscontinuityFlag)(nil) // ensure noHlsSplitDiscontinuityFlag implements Flag interface.

func (f *noHlsSplitDiscontinuityFlag) ID() string {
	return "hls_split_discontinuity"
}

func (f *noHlsSplitDiscontinuityFlag) String() string {
	return "TODO"
}

func (f *noHlsSplitDiscontinuityFlag) AsFlag() []string {
	return []string{"TODO"}
}

// Do not split HLS playlists to different formats at discontinuities such as ad
// breaks (default)
//
// NoHlsSplitDiscontinuity maps to cli flags: --no-hls-split-discontinuity.
func (ff *ExtractorBuilder) NoHlsSplitDiscontinuity() *ExtractorBuilder {
	ff.parent.addFlag(&noHlsSplitDiscontinuityFlag{})
	return ff
}

type youtubeIncludeDashManifestFlag struct {
}

var _ Flag = (*youtubeIncludeDashManifestFlag)(nil) // ensure youtubeIncludeDashManifestFlag implements Flag interface.

func (f *youtubeIncludeDashManifestFlag) ID() string {
	return "youtube_include_dash_manifest"
}

func (f *youtubeIncludeDashManifestFlag) String() string {
	return "TODO"
}

func (f *youtubeIncludeDashManifestFlag) AsFlag() []string {
	return []string{"TODO"}
}

// YoutubeIncludeDashManifest sets the "youtube-include-dash-manifest" flag to "true".
//
// YoutubeIncludeDashManifest maps to cli flags: --youtube-include-dash-manifest/--no-youtube-skip-dash-manifest.
func (ff *ExtractorBuilder) YoutubeIncludeDashManifest() *ExtractorBuilder {
	ff.parent.addFlag(&youtubeIncludeDashManifestFlag{})
	return ff
}

type youtubeSkipDashManifestFlag struct {
}

var _ Flag = (*youtubeSkipDashManifestFlag)(nil) // ensure youtubeSkipDashManifestFlag implements Flag interface.

func (f *youtubeSkipDashManifestFlag) ID() string {
	return "youtube_include_dash_manifest"
}

func (f *youtubeSkipDashManifestFlag) String() string {
	return "TODO"
}

func (f *youtubeSkipDashManifestFlag) AsFlag() []string {
	return []string{"TODO"}
}

// YoutubeSkipDashManifest sets the "youtube-skip-dash-manifest" flag to "false".
//
// YoutubeSkipDashManifest maps to cli flags: --youtube-skip-dash-manifest/--no-youtube-include-dash-manifest.
func (ff *ExtractorBuilder) YoutubeSkipDashManifest() *ExtractorBuilder {
	ff.parent.addFlag(&youtubeSkipDashManifestFlag{})
	return ff
}

type youtubeIncludeHlsManifestFlag struct {
}

var _ Flag = (*youtubeIncludeHlsManifestFlag)(nil) // ensure youtubeIncludeHlsManifestFlag implements Flag interface.

func (f *youtubeIncludeHlsManifestFlag) ID() string {
	return "youtube_include_hls_manifest"
}

func (f *youtubeIncludeHlsManifestFlag) String() string {
	return "TODO"
}

func (f *youtubeIncludeHlsManifestFlag) AsFlag() []string {
	return []string{"TODO"}
}

// YoutubeIncludeHlsManifest sets the "youtube-include-hls-manifest" flag to "true".
//
// YoutubeIncludeHlsManifest maps to cli flags: --youtube-include-hls-manifest/--no-youtube-skip-hls-manifest.
func (ff *ExtractorBuilder) YoutubeIncludeHlsManifest() *ExtractorBuilder {
	ff.parent.addFlag(&youtubeIncludeHlsManifestFlag{})
	return ff
}

type youtubeSkipHlsManifestFlag struct {
}

var _ Flag = (*youtubeSkipHlsManifestFlag)(nil) // ensure youtubeSkipHlsManifestFlag implements Flag interface.

func (f *youtubeSkipHlsManifestFlag) ID() string {
	return "youtube_include_hls_manifest"
}

func (f *youtubeSkipHlsManifestFlag) String() string {
	return "TODO"
}

func (f *youtubeSkipHlsManifestFlag) AsFlag() []string {
	return []string{"TODO"}
}

// YoutubeSkipHlsManifest sets the "youtube-skip-hls-manifest" flag to "false".
//
// YoutubeSkipHlsManifest maps to cli flags: --youtube-skip-hls-manifest/--no-youtube-include-hls-manifest.
func (ff *ExtractorBuilder) YoutubeSkipHlsManifest() *ExtractorBuilder {
	ff.parent.addFlag(&youtubeSkipHlsManifestFlag{})
	return ff
}
