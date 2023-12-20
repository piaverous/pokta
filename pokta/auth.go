package pokta

import (
	b64 "encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"crypto/x509"

	"github.com/google/uuid"
	"github.com/kataras/jwt"
	"github.com/piaverous/pokta/pokta/types"
)

func (app *App) AuthenticatePKJWT(signOnly bool) (string, error) {
	var selfSignedToken string
	pKeyBytes, err := b64.StdEncoding.DecodeString(app.Config.Okta.PrivateKey)
	if err != nil {
		return "", err
	}

	privateKey, err := x509.ParsePKCS8PrivateKey(pKeyBytes)
	if err != nil {
		return "", err
	}

	// Generate a selfSignedToken:
	now := time.Now()
	selfSignedTokenBytes, err := jwt.Sign(jwt.RS256, privateKey, map[string]interface{}{
		"iat": now.Unix(),
		"exp": now.Add(5 * time.Minute).Unix(),
		"aud": app.Config.Okta.Aud,
		"kid": app.Config.Okta.Kid,
		"iss": app.Config.Okta.ClientId,
		"sub": app.Config.Okta.ClientId,
		"jid": uuid.NewString(),
	})
	if err != nil {
		return "", err
	}
	selfSignedToken = string(selfSignedTokenBytes)

	if signOnly {
		return selfSignedToken, nil
	}

	// Build HTTP request to Okta using self signed token
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("scope", app.Config.Okta.Scope)
	data.Set("client_assertion_type", "urn:ietf:params:oauth:client-assertion-type:jwt-bearer")
	data.Set("client_assertion", selfSignedToken)

	u, err := url.ParseRequestURI(app.Config.Okta.Aud)
	if err != nil {
		return selfSignedToken, err
	}

	client := &http.Client{}
	r, err := http.NewRequest(http.MethodPost, u.String(), strings.NewReader(data.Encode())) // URL-encoded payload
	if err != nil {
		return selfSignedToken, err
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Send HTTP Request
	resp, err := client.Do(r)
	if err != nil {
		return selfSignedToken, err
	}

	// Parse HTTP Response
	var parsedResponse types.OktaAuthResponse
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return selfSignedToken, err
	}

	if err := json.Unmarshal(body, &parsedResponse); err != nil {
		return selfSignedToken, err
	}

	return parsedResponse.AccessToken, nil
}
