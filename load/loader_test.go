package load

import (
	"github.com/TakuSemba/go-hls-hosting/parse"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var FakeTsMasterPlayList = parse.MasterPlaylist{
	Path: "testMasterPlaylist.m3u8",
	Tags: []string{
		"#EXTM3U",
		"#EXT-X-VERSION:4",
		"#EXT-X-STREAM-INF:AVERAGE-BANDWIDTH=2231539,BANDWIDTH=2984657,CODECS=\"avc1.64001F,mp4a.40.2\",RESOLUTION=1280x720",
	},
	MediaPlaylists: []parse.MediaPlaylist{
		{
			Path: "testMediaPlaylist.m3u8",
			Tags: []string{
				"#EXTM3U",
				"#EXT-X-VERSION:4",
				"#EXT-X-PLAYLIST-TYPE:VOD",
				"#EXT-X-INDEPENDENT-SEGMENTS",
				"#EXT-X-TARGETDURATION:5",
				"#EXT-X-MEDIA-SEQUENCE:0",
				"#EXT-X-DISCONTINUITY-SEQUENCE:0",
				"#EXTINF:3,",
				"#EXTINF:4,",
				"#EXTINF:5,",
				"#EXT-X-ENDLIST",
			},
			Segments: []parse.Segment{
				{Path: "segment-0.ts", DurationMs: 3000, DiscontinuitySequence: 0, FileExtension: ".ts", ContainerFormat: parse.Ts, RequestType: parse.SegmentBySegment},
				{Path: "segment-1.ts", DurationMs: 4000, DiscontinuitySequence: 0, FileExtension: ".ts", ContainerFormat: parse.Ts, RequestType: parse.SegmentBySegment},
				{Path: "segment-2.ts", DurationMs: 5000, DiscontinuitySequence: 0, FileExtension: ".ts", ContainerFormat: parse.Ts, RequestType: parse.SegmentBySegment},
			},
			TotalDurationMs:         12 * 1000,
			TotalDiscontinuityCount: 0,
		},
	},
}

var FakeFmp4MasterPlayList = parse.MasterPlaylist{
	Path: "testMasterPlaylist.m3u8",
	Tags: []string{
		"#EXTM3U",
		"#EXT-X-VERSION:4",
		"#EXT-X-STREAM-INF:AVERAGE-BANDWIDTH=2231539,BANDWIDTH=2984657,CODECS=\"avc1.64001F,mp4a.40.2\",RESOLUTION=1280x720",
	},
	MediaPlaylists: []parse.MediaPlaylist{
		{
			Path: "testMediaPlaylist.m3u8",
			Tags: []string{
				"#EXTM3U",
				"#EXT-X-VERSION:4",
				"#EXT-X-PLAYLIST-TYPE:VOD",
				"#EXT-X-INDEPENDENT-SEGMENTS",
				"#EXT-X-TARGETDURATION:5",
				"#EXT-X-MEDIA-SEQUENCE:0",
				"#EXT-X-DISCONTINUITY-SEQUENCE:0",
				"#EXTINF:3,",
				"#EXTINF:4,",
				"#EXTINF:5,",
				"#EXT-X-ENDLIST",
			},
			Segments: []parse.Segment{
				{Path: "segment-0.mp4", DurationMs: 3000, DiscontinuitySequence: 0, FileExtension: ".mp4", ContainerFormat: parse.Fmp4, RequestType: parse.SegmentBySegment},
				{Path: "segment-1.mp4", DurationMs: 4000, DiscontinuitySequence: 0, FileExtension: ".mp4", ContainerFormat: parse.Fmp4, RequestType: parse.SegmentBySegment},
				{Path: "segment-2.mp4", DurationMs: 5000, DiscontinuitySequence: 0, FileExtension: ".mp4", ContainerFormat: parse.Fmp4, RequestType: parse.SegmentBySegment},
			},
			TotalDurationMs:         12 * 1000,
			TotalDiscontinuityCount: 0,
		},
	},
}

var FakeByteRangeMasterPlayList = parse.MasterPlaylist{
	Path: "testMasterPlaylist.m3u8",
	Tags: []string{
		"#EXTM3U",
		"#EXT-X-VERSION:4",
		"#EXT-X-STREAM-INF:AVERAGE-BANDWIDTH=2231539,BANDWIDTH=2984657,CODECS=\"avc1.64001F,mp4a.40.2\",RESOLUTION=1280x720",
	},
	MediaPlaylists: []parse.MediaPlaylist{
		{
			Path: "testMediaPlaylist.m3u8",
			Tags: []string{
				"#EXTM3U",
				"#EXT-X-VERSION:4",
				"#EXT-X-PLAYLIST-TYPE:VOD",
				"#EXT-X-INDEPENDENT-SEGMENTS",
				"#EXT-X-TARGETDURATION:5",
				"#EXT-X-MEDIA-SEQUENCE:0",
				"#EXT-X-DISCONTINUITY-SEQUENCE:0",
				"#EXTINF:3,",
				"#EXT-X-BYTERANGE:300000@100",
				"#EXTINF:4,",
				"#EXT-X-BYTERANGE:400000@300100",
				"#EXTINF:5,",
				"#EXT-X-BYTERANGE:500000@700100",
				"#EXT-X-ENDLIST",
			},
			Segments: []parse.Segment{
				{Path: "segment-0.mp4", DurationMs: 3000, DiscontinuitySequence: 0, FileExtension: ".mp4", ContainerFormat: parse.Fmp4, RequestType: parse.ByteRange},
				{Path: "segment-0.mp4", DurationMs: 4000, DiscontinuitySequence: 0, FileExtension: ".mp4", ContainerFormat: parse.Fmp4, RequestType: parse.ByteRange},
				{Path: "segment-0.mp4", DurationMs: 5000, DiscontinuitySequence: 0, FileExtension: ".mp4", ContainerFormat: parse.Fmp4, RequestType: parse.ByteRange},
			},
			TotalDurationMs:         36 * 1000,
			TotalDiscontinuityCount: 0,
		},
	},
}

func TestLoadTsMasterPlaylist(t *testing.T) {
	loader := NewDefaultLoader(FakeTsMasterPlayList)
	actual, err := loader.LoadMasterPlaylist()

	masterPlaylist := `
		#EXTM3U
		#EXT-X-VERSION:4
		#EXT-X-STREAM-INF:AVERAGE-BANDWIDTH=2231539,BANDWIDTH=2984657,CODECS="avc1.64001F,mp4a.40.2",RESOLUTION=1280x720
		0/playlist.m3u8
	`
	if assert.NoError(t, err) {
		assert.Equal(t, trimIndent(t, masterPlaylist), string(actual))

	}
}

// trim indent from multiple-string expression.
func trimIndent(t *testing.T, s string) string {
	t.Helper()
	// remove first indent.
	a := strings.Replace(s, "\n\t\t", "", 1)
	// remove mid indents.
	b := strings.Replace(a, "\n\t\t", "\n", -1)
	// remove last indent.
	c := strings.Replace(b, "\t", "", -1)
	return c
}
