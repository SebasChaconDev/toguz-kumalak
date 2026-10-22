package notificationprobe

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// The break this catches is a simulated broker accepting a client credential for
// publishing, another resource, or outside its validity period. It proves only
// the local double's capability contract, never Ably's implementation.
func TestSimulatedCapabilityBoundary(t *testing.T) {
	b := newBroker()
	s := httptest.NewServer(b)
	defer s.Close()
	channel := "notification-capability-owned"
	client := b.token("client", channel, "subscribe", time.Now().Add(time.Minute))
	publisher := b.token("game", channel, "publish", time.Now().Add(time.Minute))
	if client == publisher || bytes.Equal(b.clientKey, b.publisherKey) {
		t.Fatal("client and Game publisher credentials must be distinct")
	}
	cases := []struct {
		id, method, channel, credential string
		want                            int
	}{
		{"CAP-01-exact-subscribe", "GET", channel, client, 200},
		{"CAP-02-client-official-publish-denied", "POST", channel, client, 403},
		{"CAP-03-other-resource-denied", "GET", "notification-capability-foreign", client, 403},
		{"CAP-04-expired-denied", "GET", channel, b.token("client", channel, "subscribe", time.Now().Add(-time.Second)), 401},
		{"CAP-05-invalid-signature-denied", "GET", channel, client + "x", 401},
		{"CAP-06-publisher-publish", "POST", channel, publisher, 202},
		{"CAP-07-publisher-subscribe-denied", "GET", channel, publisher, 403},
		{"CAP-08-missing-denied", "GET", channel, "", 401},
	}
	for _, tc := range cases {
		t.Run(tc.id, func(t *testing.T) {
			r, _ := http.NewRequest(tc.method, s.URL+"/channels/"+tc.channel+"/messages", bytes.NewBufferString(`{"gameId":"notification-capability-owned","revision":1}`))
			if tc.credential != "" {
				r.Header.Set("Authorization", "Bearer "+tc.credential)
			}
			resp, err := s.Client().Do(r)
			if err != nil {
				t.Fatal("local capability request failed")
			}
			io.Copy(io.Discard, resp.Body)
			resp.Body.Close()
			if resp.StatusCode != tc.want {
				t.Fatalf("status=%d want=%d", resp.StatusCode, tc.want)
			}
			v, _ := json.Marshal(map[string]any{"caseId": tc.id, "status": "PASS", "httpStatus": resp.StatusCode, "syntheticCapability": true, "ablyRealValidated": false})
			t.Log("CASE " + string(v))
		})
	}
}
