// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// Code generated by cmd/codegen. DO NOT EDIT.
//
// Post-Processing Option Group

package ytdlp

// Convert video files to audio-only files (requires ffmpeg and ffprobe)
//
// ExtractAudio maps to cli flags: -x/--extract-audio.
func (c *Command) ExtractAudio() *Command {
	c.addFlag(&Flag{
		ID:   "extractaudio",
		Flag: "--extract-audio",
		Args: nil,
	})
	return c
}

// Format to convert the audio to when -x is used. (currently supported: best
// (default), aac, alac, flac, m4a, mp3, opus, vorbis, wav). You can specify
// multiple rules using similar syntax as --remux-video
//
// AudioFormat maps to cli flags: --audio-format=FORMAT.
func (c *Command) AudioFormat(format string) *Command {
	c.addFlag(&Flag{
		ID:   "audioformat",
		Flag: "--audio-format",
		Args: []string{format},
	})
	return c
}

// Specify ffmpeg audio quality to use when converting the audio with -x. Insert a
// value between 0 (best) and 10 (worst) for VBR or a specific bitrate like 128K
// (default 5)
//
// AudioQuality maps to cli flags: --audio-quality=QUALITY.
func (c *Command) AudioQuality(quality string) *Command {
	c.addFlag(&Flag{
		ID:   "audioquality",
		Flag: "--audio-quality",
		Args: []string{quality},
	})
	return c
}

// Remux the video into another container if necessary (currently supported: avi,
// flv, gif, mkv, mov, mp4, webm, aac, aiff, alac, flac, m4a, mka, mp3, ogg, opus,
// vorbis, wav). If target container does not support the video/audio codec,
// remuxing will fail. You can specify multiple rules; e.g. "aac>m4a/mov>mp4/mkv"
// will remux aac to m4a, mov to mp4 and anything else to mkv
//
// RemuxVideo maps to cli flags: --remux-video=FORMAT.
func (c *Command) RemuxVideo(format string) *Command {
	c.addFlag(&Flag{
		ID:   "remuxvideo",
		Flag: "--remux-video",
		Args: []string{format},
	})
	return c
}

// Re-encode the video into another format if necessary. The syntax and supported
// formats are the same as --remux-video
//
// RecodeVideo maps to cli flags: --recode-video=FORMAT.
func (c *Command) RecodeVideo(format string) *Command {
	c.addFlag(&Flag{
		ID:   "recodevideo",
		Flag: "--recode-video",
		Args: []string{format},
	})
	return c
}

// Give these arguments to the postprocessors. Specify the postprocessor/executable
// name and the arguments separated by a colon ":" to give the argument to the
// specified postprocessor/executable. Supported PP are: Merger, ModifyChapters,
// SplitChapters, ExtractAudio, VideoRemuxer, VideoConvertor, Metadata,
// EmbedSubtitle, EmbedThumbnail, SubtitlesConvertor, ThumbnailsConvertor,
// FixupStretched, FixupM4a, FixupM3u8, FixupTimestamp and FixupDuration. The
// supported executables are: AtomicParsley, FFmpeg and FFprobe. You can also
// specify "PP+EXE:ARGS" to give the arguments to the specified executable only
// when being used by the specified postprocessor. Additionally, for
// ffmpeg/ffprobe, "_i"/"_o" can be appended to the prefix optionally followed by a
// number to pass the argument before the specified input/output file, e.g. --ppa
// "Merger+ffmpeg_i1:-v quiet". You can use this option multiple times to give
// different arguments to different postprocessors. (Alias: --ppa)
//
// PostprocessorArgs maps to cli flags: --postprocessor-args/--ppa=NAME:ARGS.
func (c *Command) PostprocessorArgs(nameargs string) *Command {
	c.addFlag(&Flag{
		ID:   "postprocessor_args",
		Flag: "--postprocessor-args",
		Args: []string{nameargs},
	})
	return c
}

// Keep the intermediate video file on disk after post-processing
//
// KeepVideo maps to cli flags: -k/--keep-video.
func (c *Command) KeepVideo() *Command {
	c.addFlag(&Flag{
		ID:   "keepvideo",
		Flag: "--keep-video",
		Args: nil,
	})
	return c
}

// Delete the intermediate video file after post-processing (default)
//
// NoKeepVideo maps to cli flags: --no-keep-video.
func (c *Command) NoKeepVideo() *Command {
	c.addFlag(&Flag{
		ID:   "keepvideo",
		Flag: "--no-keep-video",
		Args: nil,
	})
	return c
}

// Overwrite post-processed files (default)
//
// PostOverwrites maps to cli flags: --post-overwrites.
func (c *Command) PostOverwrites() *Command {
	c.addFlag(&Flag{
		ID:   "nopostoverwrites",
		Flag: "--post-overwrites",
		Args: nil,
	})
	return c
}

// Do not overwrite post-processed files
//
// NoPostOverwrites maps to cli flags: --no-post-overwrites.
func (c *Command) NoPostOverwrites() *Command {
	c.addFlag(&Flag{
		ID:   "nopostoverwrites",
		Flag: "--no-post-overwrites",
		Args: nil,
	})
	return c
}

// Embed subtitles in the video (only for mp4, webm and mkv videos)
//
// EmbedSubs maps to cli flags: --embed-subs.
func (c *Command) EmbedSubs() *Command {
	c.addFlag(&Flag{
		ID:   "embedsubtitles",
		Flag: "--embed-subs",
		Args: nil,
	})
	return c
}

// Do not embed subtitles (default)
//
// NoEmbedSubs maps to cli flags: --no-embed-subs.
func (c *Command) NoEmbedSubs() *Command {
	c.addFlag(&Flag{
		ID:   "embedsubtitles",
		Flag: "--no-embed-subs",
		Args: nil,
	})
	return c
}

// Embed thumbnail in the video as cover art
//
// EmbedThumbnail maps to cli flags: --embed-thumbnail.
func (c *Command) EmbedThumbnail() *Command {
	c.addFlag(&Flag{
		ID:   "embedthumbnail",
		Flag: "--embed-thumbnail",
		Args: nil,
	})
	return c
}

// Do not embed thumbnail (default)
//
// NoEmbedThumbnail maps to cli flags: --no-embed-thumbnail.
func (c *Command) NoEmbedThumbnail() *Command {
	c.addFlag(&Flag{
		ID:   "embedthumbnail",
		Flag: "--no-embed-thumbnail",
		Args: nil,
	})
	return c
}

// Embed metadata to the video file. Also embeds chapters/infojson if present
// unless --no-embed-chapters/--no-embed-info-json are used (Alias: --add-metadata)
//
// EmbedMetadata maps to cli flags: --embed-metadata/--add-metadata.
func (c *Command) EmbedMetadata() *Command {
	c.addFlag(&Flag{
		ID:   "addmetadata",
		Flag: "--embed-metadata",
		Args: nil,
	})
	return c
}

// Do not add metadata to file (default) (Alias: --no-add-metadata)
//
// NoEmbedMetadata maps to cli flags: --no-embed-metadata/--no-add-metadata.
func (c *Command) NoEmbedMetadata() *Command {
	c.addFlag(&Flag{
		ID:   "addmetadata",
		Flag: "--no-embed-metadata",
		Args: nil,
	})
	return c
}

// Add chapter markers to the video file (Alias: --add-chapters)
//
// EmbedChapters maps to cli flags: --embed-chapters/--add-chapters.
func (c *Command) EmbedChapters() *Command {
	c.addFlag(&Flag{
		ID:   "addchapters",
		Flag: "--embed-chapters",
		Args: nil,
	})
	return c
}

// Do not add chapter markers (default) (Alias: --no-add-chapters)
//
// NoEmbedChapters maps to cli flags: --no-embed-chapters/--no-add-chapters.
func (c *Command) NoEmbedChapters() *Command {
	c.addFlag(&Flag{
		ID:   "addchapters",
		Flag: "--no-embed-chapters",
		Args: nil,
	})
	return c
}

// Embed the infojson as an attachment to mkv/mka video files
//
// EmbedInfoJson maps to cli flags: --embed-info-json.
func (c *Command) EmbedInfoJson() *Command {
	c.addFlag(&Flag{
		ID:   "embed_infojson",
		Flag: "--embed-info-json",
		Args: nil,
	})
	return c
}

// Do not embed the infojson as an attachment to the video file
//
// NoEmbedInfoJson maps to cli flags: --no-embed-info-json.
func (c *Command) NoEmbedInfoJson() *Command {
	c.addFlag(&Flag{
		ID:   "embed_infojson",
		Flag: "--no-embed-info-json",
		Args: nil,
	})
	return c
}

// MetadataFromTitle maps to cli flags: --metadata-from-title=FORMAT.
func (c *Command) MetadataFromTitle(format string) *Command {
	c.addFlag(&Flag{
		ID:   "metafromtitle",
		Flag: "--metadata-from-title",
		Args: []string{format},
	})
	return c
}

// Parse additional metadata like title/artist from other fields; see "MODIFYING
// METADATA" for details. Supported values of "WHEN" are the same as that of
// --use-postprocessor (default: pre_process)
//
// ParseMetadata maps to cli flags: --parse-metadata=[WHEN:]FROM:TO.
func (c *Command) ParseMetadata(fromto string) *Command {
	c.addFlag(&Flag{
		ID:   "parse_metadata",
		Flag: "--parse-metadata",
		Args: []string{fromto},
	})
	return c
}

// Replace text in a metadata field using the given regex. This option can be used
// multiple times. Supported values of "WHEN" are the same as that of
// --use-postprocessor (default: pre_process)
//
// ReplaceInMetadata maps to cli flags: --replace-in-metadata=[WHEN:]FIELDS REGEX REPLACE.
func (c *Command) ReplaceInMetadata(fields, regex, replace string) *Command {
	c.addFlag(&Flag{
		ID:   "parse_metadata",
		Flag: "--replace-in-metadata",
		Args: []string{fields, regex, replace},
	})
	return c
}

// Write metadata to the video file's xattrs (using dublin core and xdg standards)
//
// Xattrs maps to cli flags: --xattrs/--xattr.
func (c *Command) Xattrs() *Command {
	c.addFlag(&Flag{
		ID:   "xattrs",
		Flag: "--xattrs",
		Args: nil,
	})
	return c
}

var concatPlaylistChoices = []string{
	"never",
	"always",
	"multi_video",
}

// Concatenate videos in a playlist. One of "never", "always", or "multi_video"
// (default; only when the videos form a single show). All the video files must
// have same codecs and number of streams to be concatable. The "pl_video:" prefix
// can be used with "--paths" and "--output" to set the output filename for the
// concatenated files. See "OUTPUT TEMPLATE" for details
//
// ConcatPlaylist maps to cli flags: --concat-playlist=POLICY.
func (c *Command) ConcatPlaylist(policy string) *Command {
	c.addFlag(&Flag{
		ID:   "concat_playlist",
		Flag: "--concat-playlist",
		Args: []string{policy},
	})
	return c
}

var fixupChoices = []string{
	"never",
	"ignore",
	"warn",
	"detect_or_warn",
	"force",
}

// Automatically correct known faults of the file. One of never (do nothing), warn
// (only emit a warning), detect_or_warn (the default; fix file if we can, warn
// otherwise), force (try fixing even if file already exists)
//
// Fixup maps to cli flags: --fixup=POLICY.
func (c *Command) Fixup(policy string) *Command {
	c.addFlag(&Flag{
		ID:   "fixup",
		Flag: "--fixup",
		Args: []string{policy},
	})
	return c
}

// PreferAvconv sets the "prefer-avconv" flag (no description specified).
//
// PreferAvconv maps to cli flags: --prefer-avconv/--no-prefer-ffmpeg.
func (c *Command) PreferAvconv() *Command {
	c.addFlag(&Flag{
		ID:   "prefer_ffmpeg",
		Flag: "--prefer-avconv",
		Args: nil,
	})
	return c
}

// PreferFfmpeg sets the "prefer-ffmpeg" flag (no description specified).
//
// PreferFfmpeg maps to cli flags: --prefer-ffmpeg/--no-prefer-avconv.
func (c *Command) PreferFfmpeg() *Command {
	c.addFlag(&Flag{
		ID:   "prefer_ffmpeg",
		Flag: "--prefer-ffmpeg",
		Args: nil,
	})
	return c
}

// Location of the ffmpeg binary; either the path to the binary or its containing
// directory
//
// FfmpegLocation maps to cli flags: --ffmpeg-location/--avconv-location=PATH.
func (c *Command) FfmpegLocation(path string) *Command {
	c.addFlag(&Flag{
		ID:   "ffmpeg_location",
		Flag: "--ffmpeg-location",
		Args: []string{path},
	})
	return c
}

// Execute a command, optionally prefixed with when to execute it, separated by a
// ":". Supported values of "WHEN" are the same as that of --use-postprocessor
// (default: after_move). Same syntax as the output template can be used to pass
// any field as arguments to the command. If no fields are passed,
// %(filepath,_filename|)q is appended to the end of the command. This option can
// be used multiple times
//
// Exec maps to cli flags: --exec=[WHEN:]CMD.
func (c *Command) Exec(cmd string) *Command {
	c.addFlag(&Flag{
		ID:   "exec_cmd",
		Flag: "--exec",
		Args: []string{cmd},
	})
	return c
}

// ExecBeforeDownload maps to cli flags: --exec-before-download=CMD.
func (c *Command) ExecBeforeDownload(cmd string) *Command {
	c.addFlag(&Flag{
		ID:   "exec_before_dl_cmd",
		Flag: "--exec-before-download",
		Args: []string{cmd},
	})
	return c
}

// Convert the subtitles to another format (currently supported: ass, lrc, srt,
// vtt) (Alias: --convert-subtitles)
//
// ConvertSubs maps to cli flags: --convert-subs/--convert-sub/--convert-subtitles=FORMAT.
func (c *Command) ConvertSubs(format string) *Command {
	c.addFlag(&Flag{
		ID:   "convertsubtitles",
		Flag: "--convert-subs",
		Args: []string{format},
	})
	return c
}

// Convert the thumbnails to another format (currently supported: jpg, png, webp).
// You can specify multiple rules using similar syntax as --remux-video
//
// ConvertThumbnails maps to cli flags: --convert-thumbnails=FORMAT.
func (c *Command) ConvertThumbnails(format string) *Command {
	c.addFlag(&Flag{
		ID:   "convertthumbnails",
		Flag: "--convert-thumbnails",
		Args: []string{format},
	})
	return c
}

// Split video into multiple files based on internal chapters. The "chapter:"
// prefix can be used with "--paths" and "--output" to set the output filename for
// the split files. See "OUTPUT TEMPLATE" for details
//
// SplitChapters maps to cli flags: --split-chapters/--split-tracks.
func (c *Command) SplitChapters() *Command {
	c.addFlag(&Flag{
		ID:   "split_chapters",
		Flag: "--split-chapters",
		Args: nil,
	})
	return c
}

// Do not split video based on chapters (default)
//
// NoSplitChapters maps to cli flags: --no-split-chapters/--no-split-tracks.
func (c *Command) NoSplitChapters() *Command {
	c.addFlag(&Flag{
		ID:   "split_chapters",
		Flag: "--no-split-chapters",
		Args: nil,
	})
	return c
}

// Remove chapters whose title matches the given regular expression. The syntax is
// the same as --download-sections. This option can be used multiple times
//
// RemoveChapters maps to cli flags: --remove-chapters=REGEX.
func (c *Command) RemoveChapters(regex string) *Command {
	c.addFlag(&Flag{
		ID:   "remove_chapters",
		Flag: "--remove-chapters",
		Args: []string{regex},
	})
	return c
}

// Force keyframes at cuts when downloading/splitting/removing sections. This is
// slow due to needing a re-encode, but the resulting video may have fewer
// artifacts around the cuts
//
// ForceKeyframesAtCuts maps to cli flags: --force-keyframes-at-cuts.
func (c *Command) ForceKeyframesAtCuts() *Command {
	c.addFlag(&Flag{
		ID:   "force_keyframes_at_cuts",
		Flag: "--force-keyframes-at-cuts",
		Args: nil,
	})
	return c
}

// Do not force keyframes around the chapters when cutting/splitting (default)
//
// NoForceKeyframesAtCuts maps to cli flags: --no-force-keyframes-at-cuts.
func (c *Command) NoForceKeyframesAtCuts() *Command {
	c.addFlag(&Flag{
		ID:   "force_keyframes_at_cuts",
		Flag: "--no-force-keyframes-at-cuts",
		Args: nil,
	})
	return c
}

// The (case sensitive) name of plugin postprocessors to be enabled, and
// (optionally) arguments to be passed to it, separated by a colon ":". ARGS are a
// semicolon ";" delimited list of NAME=VALUE. The "when" argument determines when
// the postprocessor is invoked. It can be one of "pre_process" (after video
// extraction), "after_filter" (after video passes filter), "video" (after
// --format; before --print/--output), "before_dl" (before each video download),
// "post_process" (after each video download; default), "after_move" (after moving
// video file to it's final locations), "after_video" (after downloading and
// processing all formats of a video), or "playlist" (at end of playlist). This
// option can be used multiple times to add different postprocessors
//
// UsePostprocessor maps to cli flags: --use-postprocessor=NAME[:ARGS].
func (c *Command) UsePostprocessor(name string) *Command {
	c.addFlag(&Flag{
		ID:   "add_postprocessors",
		Flag: "--use-postprocessor",
		Args: []string{name},
	})
	return c
}
