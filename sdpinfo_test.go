package sdp

import (
	"testing"
)

const sdpstr = "v=1\r\n" +
	"o=- 4327261771880257373 2 IN IP4 127.0.0.1\r\n" +
	"s=-\r\n" +
	"t=1 1\r\n" +
	"a=group:BUNDLE audio video\r\n" +
	"a=msid-semantic: WMS xIKmAwWv4ft4ULxNJGhkHzvPaCkc8EKo4SGj\r\n" +
	"m=audio 9 UDP/TLS/RTP/SAVPF 111 103 104 9 0 8 106 105 13 110 112 113 126\r\n" +
	"c=IN IP4 0.0.0.0\r\n" +
	"a=rtcp:9 IN IP4 0.0.0.0\r\n" +
	"a=ice-ufrag:ez5G\r\n" +
	"a=ice-pwd:1F1qS++jzWLSQi0qQDZkX/QV\r\n" +
	"a=candidate:1 1 UDP 33554431 35.188.215.104 59110 typ host\r\n" +
	"a=fingerprint:sha-256 D2:FA:0E:C3:22:59:5E:14:95:69:92:3D:13:B4:84:24:2C:C2:A2:C0:3E:FD:34:8E:5E:EA:6F:AF:52:CE:E6:0F\r\n" +
	"a=setup:actpass\r\n" +
	"a=connection:new\r\n" +
	"a=mid:audio\r\n" +
	"a=extmap:1 urn:ietf:params:rtp-hdrext:ssrc-audio-level\r\n" +
	"a=sendrecv\r\n" +
	"a=rtcp-mux\r\n" +
	"a=rtpmap:111 opus/48000/2\r\n" +
	"a=rtcp-fb:111 transport-cc\r\n" +
	"a=fmtp:111 minptime=10;useinbandfec=1\r\n" +
	"a=rtpmap:103 ISAC/16000\r\n" +
	"a=rtpmap:104 ISAC/32000\r\n" +
	"a=rtpmap:9 G722/8000\r\n" +
	"a=rtpmap:0 PCMU/8000\r\n" +
	"a=rtpmap:8 PCMA/8000\r\n" +
	"a=rtpmap:106 CN/32000\r\n" +
	"a=rtpmap:105 CN/16000\r\n" +
	"a=rtpmap:13 CN/8000\r\n" +
	"a=rtpmap:110 telephone-event/48000\r\n" +
	"a=rtpmap:112 telephone-event/32000\r\n" +
	"a=rtpmap:113 telephone-event/16000\r\n" +
	"a=rtpmap:126 telephone-event/8000\r\n" +
	"a=ssrc:3510681183 cname:loqPWNg7JMmrFUnr\r\n" +
	"a=ssrc:3510681183 msid:xIKmAwWv4ft4ULxNJGhkHzvPaCkc8EKo4SGj 7ea47500-22eb-4815-a899-c74ef321b6ee\r\n" +
	"a=ssrc:3510681183 mslabel:xIKmAwWv4ft4ULxNJGhkHzvPaCkc8EKo4SGj\r\n" +
	"a=ssrc:3510681183 label:7ea47500-22eb-4815-a899-c74ef321b6ee\r\n" +
	"m=video 9 UDP/TLS/RTP/SAVPF 96 98 100 102 127 125 97 99 101 124\r\n" +
	"c=IN IP4 0.0.0.0\r\n" +
	"a=connection:new\r\n" +
	"a=rtcp:9 IN IP4 0.0.0.0\r\n" +
	"a=ice-ufrag:ez5G\r\n" +
	"a=ice-pwd:1F1qS++jzWLSQi0qQDZkX/QV\r\n" +
	"a=candidate:1 1 UDP 33554431 35.188.215.104 59110 typ host\r\n" +
	"a=fingerprint:sha-256 D2:FA:0E:C3:22:59:5E:14:95:69:92:3D:13:B4:84:24:2C:C2:A2:C0:3E:FD:34:8E:5E:EA:6F:AF:52:CE:E6:0F\r\n" +
	"a=setup:actpass\r\n" +
	"a=mid:video\r\n" +
	"a=extmap:2 urn:ietf:params:rtp-hdrext:toffset\r\n" +
	"a=extmap:3 http://www.webrtc.org/experiments/rtp-hdrext/abs-send-time\r\n" +
	"a=extmap:4 urn:3gpp:video-orientation\r\n" +
	"a=extmap:5 http://www.ietf.org/id/draft-holmer-rmcat-transport-wide-cc-extensions-01\r\n" +
	"a=extmap:6 http://www.webrtc.org/experiments/rtp-hdrext/playout-delay\r\n" +
	"a=sendrecv\r\n" +
	"a=rtcp-mux\r\n" +
	"a=rtcp-rsize\r\n" +
	"a=rtpmap:96 VP8/90000\r\n" +
	"a=rtcp-fb:96 ccm fir\r\n" +
	"a=rtcp-fb:96 nack\r\n" +
	"a=rtcp-fb:96 nack pli\r\n" +
	"a=rtcp-fb:96 goog-remb\r\n" +
	"a=rtcp-fb:96 transport-cc\r\n" +
	"a=rtpmap:98 VP9/90000\r\n" +
	"a=rtcp-fb:98 ccm fir\r\n" +
	"a=rtcp-fb:98 nack\r\n" +
	"a=rtcp-fb:98 nack pli\r\n" +
	"a=rtcp-fb:98 goog-remb\r\n" +
	"a=rtcp-fb:98 transport-cc\r\n" +
	"a=rtpmap:100 H264/90000\r\n" +
	"a=rtcp-fb:100 ccm fir\r\n" +
	"a=rtcp-fb:100 nack\r\n" +
	"a=rtcp-fb:100 nack pli\r\n" +
	"a=rtcp-fb:100 goog-remb\r\n" +
	"a=rtcp-fb:100 transport-cc\r\n" +
	"a=fmtp:100 level-asymmetry-allowed=1;packetization-mode=1;profile-level-id=42e01f\r\n" +
	"a=rtpmap:102 red/90000\r\n" +
	"a=rtpmap:127 ulpfec/90000\r\n" +
	"a=rtpmap:125 flexfec-03/90000\r\n" +
	"a=rtcp-fb:125 ccm fir\r\n" +
	"a=rtcp-fb:125 nack\r\n" +
	"a=rtcp-fb:125 nack pli\r\n" +
	"a=rtcp-fb:125 goog-remb\r\n" +
	"a=rtcp-fb:125 transport-cc\r\n" +
	"a=fmtp:125 repair-window=10000000\r\n" +
	"a=rtpmap:97 rtx/90000\r\n" +
	"a=fmtp:97 apt=96\r\n" +
	"a=rtpmap:99 rtx/90000\r\n" +
	"a=fmtp:99 apt=98\r\n" +
	"a=rtpmap:101 rtx/90000\r\n" +
	"a=fmtp:101 apt=100\r\n" +
	"a=rtpmap:124 rtx/90000\r\n" +
	"a=fmtp:124 apt=102\r\n" +
	"a=ssrc-group:FID 3004364195 1126032854\r\n" +
	"a=ssrc-group:FEC-FR 3004364195 1080772241\r\n" +
	"a=ssrc:3004364195 cname:loqPWNg7JMmrFUnr\r\n" +
	"a=ssrc:3004364195 msid:xIKmAwWv4ft4ULxNJGhkHzvPaCkc8EKo4SGj cf093ab0-0b28-4930-8fe1-7ca8d529be25\r\n" +
	"a=ssrc:3004364195 mslabel:xIKmAwWv4ft4ULxNJGhkHzvPaCkc8EKo4SGj\r\n" +
	"a=ssrc:3004364195 label:cf093ab0-0b28-4930-8fe1-7ca8d529be25\r\n" +
	"a=ssrc:1126032854 cname:loqPWNg7JMmrFUnr\r\n" +
	"a=ssrc:1126032854 msid:xIKmAwWv4ft4ULxNJGhkHzvPaCkc8EKo4SGj cf093ab0-0b28-4930-8fe1-7ca8d529be25\r\n" +
	"a=ssrc:1126032854 mslabel:xIKmAwWv4ft4ULxNJGhkHzvPaCkc8EKo4SGj\r\n" +
	"a=ssrc:1126032854 label:cf093ab0-0b28-4930-8fe1-7ca8d529be25\r\n" +
	"a=ssrc:1080772241 cname:loqPWNg7JMmrFUnr\r\n" +
	"a=ssrc:1080772241 msid:xIKmAwWv4ft4ULxNJGhkHzvPaCkc8EKo4SGj cf093ab0-0b28-4930-8fe1-7ca8d529be25\r\n" +
	"a=ssrc:1080772241 mslabel:xIKmAwWv4ft4ULxNJGhkHzvPaCkc8EKo4SGj\r\n" +
	"a=ssrc:1080772241 label:cf093ab0-0b28-4930-8fe1-7ca8d529be25\r\n"

var Capabilities = map[string]*Capability{
	"audio": &Capability{
		Codecs: []string{"opus"},
	},
	"video": &Capability{
		Codecs: []string{"vp8"},
		Rtx:    true,
		Rtcpfbs: []*RtcpFeedback{
			&RtcpFeedback{
				ID: "goog-remb",
			},
			&RtcpFeedback{
				ID: "transport-cc",
			},
			&RtcpFeedback{
				ID:     "ccm",
				Params: []string{"fir"},
			},
			&RtcpFeedback{
				ID:     "nack",
				Params: []string{"pli"},
			},
		},
		Extensions: []string{
			"urn:3gpp:video-orientation",
			"http://www.ietf.org/id/draft-holmer-rmcat-transport-wide-cc-extensions-01",
			"http://www.webrtc.org/experiments/rtp-hdrext/abs-send-time",
			"urn:ietf:params:rtp-hdrext:toffse",
			"urn:ietf:params:rtp-hdrext:crypto:rtp-stream-id",
			"urn:ietf:params:rtp-hdrext:crypto:mid",
		},
	},
}

func Test_Parse(t *testing.T) {

	offer, err := Parse(sdpstr)
	if err != nil {
		t.Error(err)
	}

	candidate := offer.GetCandidates()[0]

	if candidate.GetAddress() != "35.188.215.104" {
		t.Error("can not find candidate")
	}

	mediaInfo := offer.GetMedia("audio")

	codecInfo := mediaInfo.GetCodecForType(111)

	if codecInfo.GetCodec() != "opus" {
		t.Error("can not find opus codec")
	}

	if len(offer.GetMedias()) != 2 {
		t.Error("parse media error")
	}

	dtls := offer.GetDTLS()

	if dtls.GetHash() != "sha-256" {
		t.Error("parse dtls error")
	}

	stream := offer.GetStream("xIKmAwWv4ft4ULxNJGhkHzvPaCkc8EKo4SGj")

	if stream == nil {
		t.Error("stream parse error")
	}

	track := stream.GetTrack("7ea47500-22eb-4815-a899-c74ef321b6ee")

	if track == nil {
		t.Error("track parse error")
	}

	if track.GetMediaType() != "audio" {
		t.Error("track can not get media")
	}

	trackClone := track.Clone()

	if trackClone.GetID() != track.GetID() {
		t.Error("track clone error")
	}
}

func Test_SDPTOString(t *testing.T) {

	sdpInfo, err := Parse(sdpstr)
	if err != nil {
		t.Error(err)
	}

	sdpstring := sdpInfo.String()

	sdpInfo2, err := Parse(sdpstring)

	if err != nil {
		t.Log(sdpInfo2)
		t.Error(err)
	}

	if sdpInfo2.GetFirstStream().GetID() != sdpInfo.GetFirstStream().GetID() {
		t.Error("stream does not equal")
	}

	if len(sdpInfo2.GetMedias()) != len(sdpInfo.GetMedias()) {
		t.Error("media count does not equal")
	}

	if sdpInfo2.ice.GetUfrag() != sdpInfo.ice.GetUfrag() {
		t.Error("ufrag does not equal")
	}

	if sdpInfo2.dtls.GetHash() != sdpInfo.dtls.GetHash() {
		t.Error("dtls hash does not match")
	}
}

func Test_SDPAnswer(t *testing.T) {

	offer, err := Parse(sdpstr)

	if err != nil {
		t.Error(err)
	}

	iceInfo := ICEInfoGenerate(true)
	dtlsInfo := NewDTLSInfo(SETUPACTPASS, "sha-256", "F2:AA:0E:C3:22:59:5E:14:95:69:92:3D:13:B4:84:24:2C:C2:A2:C0:3E:FD:34:8E:5E:EA:6F:AF:52:CE:E6:0F")
	candidate := NewCandidateInfo("1", 1, "UDP", 33554431, "127.0.0.1", 10000, "host", "", 0)

	answer := offer.Answer(iceInfo, dtlsInfo, []*CandidateInfo{candidate}, Capabilities)

	if answer.GetMedia("audio") == nil {
		t.Error("answer audio error")
	}

	if answer.GetMedia("video") == nil {
		t.Error("answer video error")
	}
}

func BenchmarkParse(b *testing.B) {

	for n := 0; n < b.N; n++ {
		_, err := Parse(sdpstr)
		if err != nil {
			b.Error("parse error")
		}
	}
}
