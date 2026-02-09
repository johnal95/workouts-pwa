package testutil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/johnal95/workouts-pwa/internal/httpx"
	"github.com/stretchr/testify/require"
)

func DoRequest(a *TestApp, method, path, token string, body, out any) *http.Response {
	buf := &bytes.Buffer{}
	if body != nil {
		data, err := json.Marshal(body)
		require.NoError(a.T, err)
		buf = bytes.NewBuffer(data)
	}

	req, err := http.NewRequest(method, a.Server.URL+path, buf)
	require.NoError(a.T, err)
	req.Header.Set("Cookie", fmt.Sprintf("%s=%s", httpx.CookieSessionToken, token))

	resp, err := http.DefaultClient.Do(req)
	require.NoError(a.T, err)
	if out != nil {
		defer resp.Body.Close()
		err = json.NewDecoder(resp.Body).Decode(out)
		require.NoError(a.T, err)
	}
	return resp
}
