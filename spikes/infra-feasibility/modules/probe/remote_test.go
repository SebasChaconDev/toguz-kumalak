package probe

import "testing"

func TestRemoteOriginRejectsProxyAndCredentialTargets(t *testing.T) {
	if _, e := RemoteOrigin("https://tk-spike01-ai-ultra-seneca.vercel.app"); e != nil {
		t.Fatal(e)
	}
	for _, target := range []string{
		"https://example.com", "https://tk-spike01-ai-ultra-seneca.vercel.app.evil.test",
		"http://tk-spike01-ai-ultra-seneca.vercel.app", "https://u:p@tk-spike01-ai-ultra-seneca.vercel.app",
		"https://tk-spike01-ai-ultra-seneca.vercel.app/private", "https://tk-spike01-ai-ultra-seneca.vercel.app?secret=x",
		"https://tk-spike01-ai-ultra-seneca.vercel.app:443", "https://tk-spike01-ai-ultra-seneca.vercel.app#fragment",
	} {
		if _, e := RemoteOrigin(target); e == nil {
			t.Fatal("unsafe target accepted")
		}
	}
}
