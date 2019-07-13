package worldtime

import "testing"

func TestClient_Timezones(t *testing.T) {
	c := NewDefaultClient()
	tz, _ := c.Timezone("Europe/London")
	t.Log(tz)
}
