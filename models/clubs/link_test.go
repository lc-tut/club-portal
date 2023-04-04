package clubs

import (
	"github.com/google/go-cmp/cmp"
	"github.com/lc-tut/club-portal/consts"
	"testing"
)

var links = []ClubLink{
	{1, consts.DummyUUID, "Twitter", "https://twitter.com/TwitterJP"},
	{2, consts.DummyUUID, "Instagram", "https://www.instagram.com/"},
	{3, consts.DummyUUID, "HP", "https://example.com"},
}

func TestLinks_ToLinkResponse(t *testing.T) {
	res := []LinkResponse{
		{"Twitter", "https://twitter.com/TwitterJP"},
		{"Instagram", "https://www.instagram.com/"},
		{"HP", "https://example.com"},
	}
	tests := []struct {
		name string
		l    Links
		want []LinkResponse
	}{
		{"link", links, res},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.l.ToLinkResponse()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("ToLinkResponse() = %v, want %v\n%v", got, tt.want, diff)
			}
		})
	}
}

func TestLinks_ToRestrictedLinkResponse(t *testing.T) {
	res := []LinkResponse{
		{"Twitter", "https://twitter.com/TwitterJP"},
		{"Instagram", "https://www.instagram.com/"},
	}
	tests := []struct {
		name string
		l    Links
		want []LinkResponse
	}{
		{"link_restricted", links, res},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.l.ToRestrictedLinkResponse()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("ToRestrictedLinkResponse() = %v, want %v\n%v", got, tt.want, diff)
			}
		})
	}
}
