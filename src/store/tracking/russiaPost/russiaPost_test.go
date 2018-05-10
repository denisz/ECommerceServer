package russiaPost

import (
	"testing"
	"fmt"
)

func TestTracking_RussiaPost(t *testing.T) {
	track := "RA644000001RU"
	data, err := TrackRussiaPost(track)

	if err != nil {
		t.Error(err)
		return
	}

	fmt.Printf("%v", data)
}
