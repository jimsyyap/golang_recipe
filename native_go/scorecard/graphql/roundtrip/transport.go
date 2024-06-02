// Copyright 2021 Security Scorecard Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rt

import (
	"context"
	"fmt"
	"net/http"
)

func NewTransport(t string) http.RoundTripper {
	return &githubTransport{
		innerTransport: http.DefaultTransport,
		token:          t,
	}
}

// githubTransport handles authorization using GitHub personal access token (PATs) during HTTP requests.
type githubTransport struct {
	innerTransport http.RoundTripper
	token          string
}

func (gt *githubTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	*r = *r.WithContext(context.Background())

	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", gt.token))
	resp, err := gt.innerTransport.RoundTrip(r)
	if err != nil {
		return nil, fmt.Errorf("error in HTTP: %w", err)
	}

	return resp, nil
}

/*
* summary
* purpose - the code sets up a custom http ttransport that adds a githuub personal access token to the authorization header of every request. This is crucial for making authenticated requests to githubs api
* components:
* newtransport function - initializes the custom transport with the provided token
* githubtransport struct - holds the inner http transport and the token
* roundtrip method - modifies each http request to include the token in the authorization header and then performs the request.
* this code ensures that any http request made using this custom transport will be authenticated with the provided github token, allowing ccess to github api endpoints that require authentication.
 */
