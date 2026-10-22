package notificationprobe

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"
	"sync"
	"time"
)

// These types are test-only. The double does not implement or validate Ably's
// protocol/SDK. Its signed capability and notice envelope are synthetic.
type notification struct {
	GameID   string `json:"gameId"`
	Revision int64  `json:"revision"`
}
type snapshot struct {
	Revision int64 `json:"revision"`
	Effects  int64 `json:"effects"`
	Receipts int64 `json:"receipts"`
}
type publication struct {
	Notice    notification `json:"notice"`
	Committed snapshot     `json:"committedAtPublish"`
}
type capability struct {
	Kind     string `json:"kind"`
	Audience string `json:"audience"`
	Channel  string `json:"channel"`
	Action   string `json:"action"`
	Expires  int64  `json:"expires"`
}
type broker struct {
	clientKey, publisherKey []byte
	mu                      sync.Mutex
	mode                    string
	messages                map[string][]notification
	publications            []publication
	readCommitted           func(context.Context, notification) (snapshot, error)
	cancelled               chan struct{}
}

func newBroker() *broker {
	return &broker{clientKey: []byte("synthetic-client-key-32-bytes-long"), publisherKey: []byte("synthetic-game-key-32-bytes-long!!"), messages: map[string][]notification{}, cancelled: make(chan struct{}, 1)}
}
func (b *broker) token(kind, channel, action string, until time.Time) string {
	raw, _ := json.Marshal(capability{kind, "local-notification-double", channel, action, until.Unix()})
	p := base64.RawURLEncoding.EncodeToString(raw)
	key := b.clientKey
	if kind == "game" {
		key = b.publisherKey
	}
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(p))
	return p + "." + base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
}
func (b *broker) authorize(r *http.Request, channel string) int {
	h := r.Header.Get("Authorization")
	if !strings.HasPrefix(h, "Bearer ") || len(h) > 2048 {
		return 401
	}
	parts := strings.Split(strings.TrimPrefix(h, "Bearer "), ".")
	if len(parts) != 2 {
		return 401
	}
	raw, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return 401
	}
	var c capability
	decoder := json.NewDecoder(strings.NewReader(string(raw)))
	decoder.DisallowUnknownFields()
	if decoder.Decode(&c) != nil {
		return 401
	}
	key := b.clientKey
	if c.Kind == "game" {
		key = b.publisherKey
	} else if c.Kind != "client" {
		return 401
	}
	sig, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return 401
	}
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(parts[0]))
	now := time.Now().Unix()
	if !hmac.Equal(sig, mac.Sum(nil)) || c.Audience != "local-notification-double" || c.Expires <= now || c.Expires > now+300 {
		return 401
	}
	if c.Channel != channel {
		return 403
	}
	if r.Method == "POST" && (c.Kind != "game" || c.Action != "publish") {
		return 403
	}
	if r.Method == "GET" && (c.Kind != "client" || c.Action != "subscribe") {
		return 403
	}
	return 0
}
func (b *broker) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) != 3 || parts[0] != "channels" || parts[2] != "messages" {
		http.NotFound(w, r)
		return
	}
	if r.Method != "POST" && r.Method != "GET" {
		w.WriteHeader(405)
		return
	}
	channel := parts[1]
	if status := b.authorize(r, channel); status != 0 {
		w.WriteHeader(status)
		return
	}
	if r.Method == "GET" {
		b.mu.Lock()
		messages := append([]notification{}, b.messages[channel]...)
		delete(b.messages, channel)
		b.mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(messages)
		return
	}
	var n notification
	if json.NewDecoder(http.MaxBytesReader(w, r.Body, 2048)).Decode(&n) != nil || n.GameID != channel || n.Revision < 1 {
		w.WriteHeader(400)
		return
	}
	var observed snapshot
	if b.readCommitted != nil {
		var err error
		observed, err = b.readCommitted(r.Context(), n)
		if err != nil || observed.Revision != n.Revision || observed.Effects != n.Revision || observed.Receipts != n.Revision {
			w.WriteHeader(500)
			return
		}
	}
	b.mu.Lock()
	b.publications = append(b.publications, publication{n, observed})
	mode := b.mode
	switch mode {
	case "loss", "unresponsive":
		// Deliberate transport failure: committed Game state remains separate.
	case "duplicate":
		b.messages[channel] = append(b.messages[channel], n, n)
	case "reverse":
		b.messages[channel] = append([]notification{n}, b.messages[channel]...)
	default:
		b.messages[channel] = append(b.messages[channel], n)
	}
	b.mu.Unlock()
	if mode == "unresponsive" {
		select {
		case <-r.Context().Done():
			select {
			case b.cancelled <- struct{}{}:
			default:
			}
		case <-time.After(2 * time.Second):
		}
		return
	}
	w.WriteHeader(202)
}
func (b *broker) audit() []publication {
	b.mu.Lock()
	defer b.mu.Unlock()
	return append([]publication{}, b.publications...)
}
