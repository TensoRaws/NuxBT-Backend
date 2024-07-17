package torrent

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMagnet(t *testing.T) {
	mag := GetMagnet("7f3956e5a15b34b62159727c08f944c7e433ad1e", []string{})
	assert.Equal(t, "magnet:?xt=urn:btih:7f3956e5a15b34b62159727c08f944c7e433ad1e", mag)
}

func TestGetMagnet2(t *testing.T) {
	mag := GetMagnet("7f3956e5a15b34b62159727c08f944c7e433ad1e",
		[]string{"http://tracker1.com", "http://tracker2.com"})
	assert.Equal(t, "magnet:?xt=urn:btih:7f3956e5a15b34b62159727c08f944c7e433ad1e"+
		"&tr=http%3A%2F%2Ftracker1.com&tr=http%3A%2F%2Ftracker2.com", mag)
	t.Log(GetMagnet("7f3956e5a15b34b62159727c08f944c7e433ad1e", TRACKER_LIST))
}
