package gdrive

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os/exec"
	"strings"
	"time"

	"github.com/brunoga/go-gdrivefs/settings"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	drive "google.golang.org/api/drive/v2"
)

type Auth struct {
	oAuthClient *http.Client
}

func NewAuth(s *settings.Settings) (*Auth, error) {
	oAuthClient, err := getOAuthClient(s)
	if err != nil {
		return nil, err
	}

	return &Auth{
		oAuthClient,
	}, nil
}

func (a *Auth) Client() *http.Client {
	return a.oAuthClient
}

func getOAuthClient(s *settings.Settings) (*http.Client, error) {
	oAuthClientId, err := s.Get("oAuthClientId")
	if err != nil {
		return nil, fmt.Errorf("no oAuthClientId in settings")
	}

	oAuthClientSecret, err := s.Get("oAuthClientSecret")
	if err != nil {
		return nil, fmt.Errorf("no oAuthClientSecret in settings")
	}

	config := &oauth2.Config{
		ClientID:     oAuthClientId,
		ClientSecret: oAuthClientSecret,
		Endpoint:     google.Endpoint,
		Scopes: []string{
			drive.DriveScope,
		},
	}

	ctx := context.Background()

	var t *oauth2.Token

	oAuthClientToken, err := s.Get("oAuthClientToken")
	if err != nil {
		oobAuth, err := s.Get("oobAuth")
		if err != nil {
			oobAuth = "false"
		}

		t, err = getAuthTokenFromWeb(config, ctx, oobAuth == "true")
		if err != nil {
			return nil, err
		}
		b := new(bytes.Buffer)
		gob.NewEncoder(base64.NewEncoder(base64.StdEncoding,
			b)).Encode(t)

		s.Set("oAuthClientToken", b.String())

		err = s.Save()
		if err != nil {
			log.Printf("Could not save settings file (%q). Using "+
				"in-memory.", err)
		}
	} else {
		t := new(oauth2.Token)
		err = gob.NewDecoder(base64.NewDecoder(base64.StdEncoding,
			strings.NewReader(oAuthClientToken))).Decode(t)
		if err != nil {
			return nil, err
		}
	}

	c := config.Client(ctx, t)

	return c, nil
}

func getAuthTokenFromWeb(config *oauth2.Config, ctx context.Context,
	oobAuth bool) (*oauth2.Token, error) {
	randState := fmt.Sprintf("st%d", time.Now().UnixNano())

	var code string

	if oobAuth {
		config.RedirectURL = "urn:ietf:wg:oauth:2.0:oob"

		authURL := config.AuthCodeURL(randState)
		log.Printf("Authorize this app at: %s", authURL)

		log.Printf("Enter verification code: ")

		fmt.Scanln(&code)
	} else {
		ch := make(chan string)
		ts := httptest.NewServer(http.HandlerFunc(
			func(rw http.ResponseWriter, req *http.Request) {
				if req.URL.Path == "/favicon.ico" {
					http.Error(rw, "", 404)
					return
				}
				if req.FormValue("state") != randState {
					log.Printf("State doesn't match: req = %#v", req)
					http.Error(rw, "", 500)
					return
				}
				if code := req.FormValue("code"); code != "" {
					fmt.Fprintf(rw, "<h1>Success</h1>Authorized.")
					rw.(http.Flusher).Flush()
					ch <- code
					return
				}
				log.Printf("no code")
				http.Error(rw, "", 500)
			},
		))
		defer ts.Close()

		config.RedirectURL = ts.URL
		authURL := config.AuthCodeURL(randState)
		go openURL(authURL)
		log.Printf("Authorize this app at: %s", authURL)
		code = <-ch
	}

	t, err := config.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}

	return t, err
}

func openURL(url string) {
	try := []string{"xdg-open", "google-chrome", "open"}
	for _, bin := range try {
		err := exec.Command(bin, url).Run()
		if err == nil {
			return
		}
	}
	log.Printf("Error opening URL in browser.")
}
