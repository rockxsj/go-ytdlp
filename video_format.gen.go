// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// Code generated by cmd/codegen. DO NOT EDIT.
//
// Video Format Option Group

package ytdlp

type VideoFormatBuilder struct {
	parent *Command
}

type formatFlag struct {
	args []string
}

var _ Flag = (*formatFlag)(nil) // ensure formatFlag implements Flag interface.

func (f *formatFlag) ID() string {
	return "format"
}

func (f *formatFlag) String() string {
	return "TODO"
}

func (f *formatFlag) AsFlag() []string {
	return []string{"TODO"}
}

// Video format code, see "FORMAT SELECTION" for more details
//
// Format maps to cli flags: -f/--format=FORMAT.
func (ff *VideoFormatBuilder) Format(format string) *VideoFormatBuilder {
	ff.parent.addFlag(&formatFlag{
		args: []string{format},
	})
	return ff
}

type formatSortForceFlag struct {
}

var _ Flag = (*formatSortForceFlag)(nil) // ensure formatSortForceFlag implements Flag interface.

func (f *formatSortForceFlag) ID() string {
	return "format_sort_force"
}

func (f *formatSortForceFlag) String() string {
	return "TODO"
}

func (f *formatSortForceFlag) AsFlag() []string {
	return []string{"TODO"}
}

// Force user specified sort order to have precedence over all fields, see "Sorting
// Formats" for more details (Alias: --S-force)
//
// FormatSortForce maps to cli flags: --format-sort-force/--S-force=FORMAT.
func (ff *VideoFormatBuilder) FormatSortForce() *VideoFormatBuilder {
	ff.parent.addFlag(&formatSortForceFlag{})
	return ff
}

type noFormatSortForceFlag struct {
}

var _ Flag = (*noFormatSortForceFlag)(nil) // ensure noFormatSortForceFlag implements Flag interface.

func (f *noFormatSortForceFlag) ID() string {
	return "format_sort_force"
}

func (f *noFormatSortForceFlag) String() string {
	return "TODO"
}

func (f *noFormatSortForceFlag) AsFlag() []string {
	return []string{"TODO"}
}

// Some fields have precedence over the user specified sort order (default)
//
// NoFormatSortForce maps to cli flags: --no-format-sort-force=FORMAT.
func (ff *VideoFormatBuilder) NoFormatSortForce() *VideoFormatBuilder {
	ff.parent.addFlag(&noFormatSortForceFlag{})
	return ff
}

type videoMultistreamsFlag struct {
}

var _ Flag = (*videoMultistreamsFlag)(nil) // ensure videoMultistreamsFlag implements Flag interface.

func (f *videoMultistreamsFlag) ID() string {
	return "allow_multiple_video_streams"
}

func (f *videoMultistreamsFlag) String() string {
	return "TODO"
}

func (f *videoMultistreamsFlag) AsFlag() []string {
	return []string{"TODO"}
}

// Allow multiple video streams to be merged into a single file
//
// VideoMultistreams maps to cli flags: --video-multistreams.
func (ff *VideoFormatBuilder) VideoMultistreams() *VideoFormatBuilder {
	ff.parent.addFlag(&videoMultistreamsFlag{})
	return ff
}

type noVideoMultistreamsFlag struct {
}

var _ Flag = (*noVideoMultistreamsFlag)(nil) // ensure noVideoMultistreamsFlag implements Flag interface.

func (f *noVideoMultistreamsFlag) ID() string {
	return "allow_multiple_video_streams"
}

func (f *noVideoMultistreamsFlag) String() string {
	return "TODO"
}

func (f *noVideoMultistreamsFlag) AsFlag() []string {
	return []string{"TODO"}
}

// Only one video stream is downloaded for each output file (default)
//
// NoVideoMultistreams maps to cli flags: --no-video-multistreams.
func (ff *VideoFormatBuilder) NoVideoMultistreams() *VideoFormatBuilder {
	ff.parent.addFlag(&noVideoMultistreamsFlag{})
	return ff
}

type audioMultistreamsFlag struct {
}

var _ Flag = (*audioMultistreamsFlag)(nil) // ensure audioMultistreamsFlag implements Flag interface.

func (f *audioMultistreamsFlag) ID() string {
	return "allow_multiple_audio_streams"
}

func (f *audioMultistreamsFlag) String() string {
	return "TODO"
}

func (f *audioMultistreamsFlag) AsFlag() []string {
	return []string{"TODO"}
}

// Allow multiple audio streams to be merged into a single file
//
// AudioMultistreams maps to cli flags: --audio-multistreams.
func (ff *VideoFormatBuilder) AudioMultistreams() *VideoFormatBuilder {
	ff.parent.addFlag(&audioMultistreamsFlag{})
	return ff
}

type noAudioMultistreamsFlag struct {
}

var _ Flag = (*noAudioMultistreamsFlag)(nil) // ensure noAudioMultistreamsFlag implements Flag interface.

func (f *noAudioMultistreamsFlag) ID() string {
	return "allow_multiple_audio_streams"
}

func (f *noAudioMultistreamsFlag) String() string {
	return "TODO"
}

func (f *noAudioMultistreamsFlag) AsFlag() []string {
	return []string{"TODO"}
}

// Only one audio stream is downloaded for each output file (default)
//
// NoAudioMultistreams maps to cli flags: --no-audio-multistreams.
func (ff *VideoFormatBuilder) NoAudioMultistreams() *VideoFormatBuilder {
	ff.parent.addFlag(&noAudioMultistreamsFlag{})
	return ff
}

type allFormatsFlag struct {
}

var _ Flag = (*allFormatsFlag)(nil) // ensure allFormatsFlag implements Flag interface.

func (f *allFormatsFlag) ID() string {
	return "format"
}

func (f *allFormatsFlag) String() string {
	return "TODO"
}

func (f *allFormatsFlag) AsFlag() []string {
	return []string{"TODO"}
}

type preferFreeFormatsFlag struct {
}

var _ Flag = (*preferFreeFormatsFlag)(nil) // ensure preferFreeFormatsFlag implements Flag interface.

func (f *preferFreeFormatsFlag) ID() string {
	return "prefer_free_formats"
}

func (f *preferFreeFormatsFlag) String() string {
	return "TODO"
}

func (f *preferFreeFormatsFlag) AsFlag() []string {
	return []string{"TODO"}
}

// Prefer video formats with free containers over non-free ones of same quality.
// Use with "-S ext" to strictly prefer free containers irrespective of quality
//
// PreferFreeFormats maps to cli flags: --prefer-free-formats.
func (ff *VideoFormatBuilder) PreferFreeFormats() *VideoFormatBuilder {
	ff.parent.addFlag(&preferFreeFormatsFlag{})
	return ff
}

type noPreferFreeFormatsFlag struct {
}

var _ Flag = (*noPreferFreeFormatsFlag)(nil) // ensure noPreferFreeFormatsFlag implements Flag interface.

func (f *noPreferFreeFormatsFlag) ID() string {
	return "prefer_free_formats"
}

func (f *noPreferFreeFormatsFlag) String() string {
	return "TODO"
}

func (f *noPreferFreeFormatsFlag) AsFlag() []string {
	return []string{"TODO"}
}

// Don't give any special preference to free containers (default)
//
// NoPreferFreeFormats maps to cli flags: --no-prefer-free-formats.
func (ff *VideoFormatBuilder) NoPreferFreeFormats() *VideoFormatBuilder {
	ff.parent.addFlag(&noPreferFreeFormatsFlag{})
	return ff
}

type checkFormatsFlag struct {
}

var _ Flag = (*checkFormatsFlag)(nil) // ensure checkFormatsFlag implements Flag interface.

func (f *checkFormatsFlag) ID() string {
	return "check_formats"
}

func (f *checkFormatsFlag) String() string {
	return "TODO"
}

func (f *checkFormatsFlag) AsFlag() []string {
	return []string{"TODO"}
}

type checkAllFormatsFlag struct {
}

var _ Flag = (*checkAllFormatsFlag)(nil) // ensure checkAllFormatsFlag implements Flag interface.

func (f *checkAllFormatsFlag) ID() string {
	return "check_formats"
}

func (f *checkAllFormatsFlag) String() string {
	return "TODO"
}

func (f *checkAllFormatsFlag) AsFlag() []string {
	return []string{"TODO"}
}

// Check all formats for whether they are actually downloadable
//
// CheckAllFormats maps to cli flags: --check-all-formats.
func (ff *VideoFormatBuilder) CheckAllFormats() *VideoFormatBuilder {
	ff.parent.addFlag(&checkAllFormatsFlag{})
	return ff
}

type noCheckFormatsFlag struct {
}

var _ Flag = (*noCheckFormatsFlag)(nil) // ensure noCheckFormatsFlag implements Flag interface.

func (f *noCheckFormatsFlag) ID() string {
	return "check_formats"
}

func (f *noCheckFormatsFlag) String() string {
	return "TODO"
}

func (f *noCheckFormatsFlag) AsFlag() []string {
	return []string{"TODO"}
}

// Do not check that the formats are actually downloadable
//
// NoCheckFormats maps to cli flags: --no-check-formats.
func (ff *VideoFormatBuilder) NoCheckFormats() *VideoFormatBuilder {
	ff.parent.addFlag(&noCheckFormatsFlag{})
	return ff
}

type listFormatsFlag struct {
}

var _ Flag = (*listFormatsFlag)(nil) // ensure listFormatsFlag implements Flag interface.

func (f *listFormatsFlag) ID() string {
	return "listformats"
}

func (f *listFormatsFlag) String() string {
	return "TODO"
}

func (f *listFormatsFlag) AsFlag() []string {
	return []string{"TODO"}
}

// List available formats of each video. Simulate unless --no-simulate is used
//
// ListFormats maps to cli flags: -F/--list-formats.
func (ff *VideoFormatBuilder) ListFormats() *VideoFormatBuilder {
	ff.parent.addFlag(&listFormatsFlag{})
	return ff
}

type listFormatsAsTableFlag struct {
}

var _ Flag = (*listFormatsAsTableFlag)(nil) // ensure listFormatsAsTableFlag implements Flag interface.

func (f *listFormatsAsTableFlag) ID() string {
	return "listformats_table"
}

func (f *listFormatsAsTableFlag) String() string {
	return "TODO"
}

func (f *listFormatsAsTableFlag) AsFlag() []string {
	return []string{"TODO"}
}

// ListFormatsAsTable sets the "list-formats-as-table" flag to "true".
//
// ListFormatsAsTable maps to cli flags: --list-formats-as-table.
func (ff *VideoFormatBuilder) ListFormatsAsTable() *VideoFormatBuilder {
	ff.parent.addFlag(&listFormatsAsTableFlag{})
	return ff
}

type listFormatsOldFlag struct {
}

var _ Flag = (*listFormatsOldFlag)(nil) // ensure listFormatsOldFlag implements Flag interface.

func (f *listFormatsOldFlag) ID() string {
	return "listformats_table"
}

func (f *listFormatsOldFlag) String() string {
	return "TODO"
}

func (f *listFormatsOldFlag) AsFlag() []string {
	return []string{"TODO"}
}

// ListFormatsOld sets the "list-formats-old" flag to "false".
//
// ListFormatsOld maps to cli flags: --list-formats-old/--no-list-formats-as-table.
func (ff *VideoFormatBuilder) ListFormatsOld() *VideoFormatBuilder {
	ff.parent.addFlag(&listFormatsOldFlag{})
	return ff
}

type mergeOutputFormatFlag struct {
	args []string
}

var _ Flag = (*mergeOutputFormatFlag)(nil) // ensure mergeOutputFormatFlag implements Flag interface.

func (f *mergeOutputFormatFlag) ID() string {
	return "merge_output_format"
}

func (f *mergeOutputFormatFlag) String() string {
	return "TODO"
}

func (f *mergeOutputFormatFlag) AsFlag() []string {
	return []string{"TODO"}
}

// Containers that may be used when merging formats, separated by "/", e.g.
// "mp4/mkv". Ignored if no merge is required. (currently supported: avi, flv, mkv,
// mov, mp4, webm)
//
// MergeOutputFormat maps to cli flags: --merge-output-format=FORMAT.
func (ff *VideoFormatBuilder) MergeOutputFormat(format string) *VideoFormatBuilder {
	ff.parent.addFlag(&mergeOutputFormatFlag{
		args: []string{format},
	})
	return ff
}

type allowUnplayableFormatsFlag struct {
}

var _ Flag = (*allowUnplayableFormatsFlag)(nil) // ensure allowUnplayableFormatsFlag implements Flag interface.

func (f *allowUnplayableFormatsFlag) ID() string {
	return "allow_unplayable_formats"
}

func (f *allowUnplayableFormatsFlag) String() string {
	return "TODO"
}

func (f *allowUnplayableFormatsFlag) AsFlag() []string {
	return []string{"TODO"}
}

// AllowUnplayableFormats sets the "allow-unplayable-formats" flag to "true".
//
// AllowUnplayableFormats maps to cli flags: --allow-unplayable-formats.
func (ff *VideoFormatBuilder) AllowUnplayableFormats() *VideoFormatBuilder {
	ff.parent.addFlag(&allowUnplayableFormatsFlag{})
	return ff
}

type noAllowUnplayableFormatsFlag struct {
}

var _ Flag = (*noAllowUnplayableFormatsFlag)(nil) // ensure noAllowUnplayableFormatsFlag implements Flag interface.

func (f *noAllowUnplayableFormatsFlag) ID() string {
	return "allow_unplayable_formats"
}

func (f *noAllowUnplayableFormatsFlag) String() string {
	return "TODO"
}

func (f *noAllowUnplayableFormatsFlag) AsFlag() []string {
	return []string{"TODO"}
}

// NoAllowUnplayableFormats sets the "no-allow-unplayable-formats" flag to "false".
//
// NoAllowUnplayableFormats maps to cli flags: --no-allow-unplayable-formats.
func (ff *VideoFormatBuilder) NoAllowUnplayableFormats() *VideoFormatBuilder {
	ff.parent.addFlag(&noAllowUnplayableFormatsFlag{})
	return ff
}
