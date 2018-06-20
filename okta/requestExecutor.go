/*
 * Copyright 2018 - Present Okta, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package okta

import (
	"io"
	"io/ioutil"
	"net/http"
)

type RequestExecutor struct {
	httpClient *http.Client
	config *Config
}

func NewRequestExecutor(httpClient *http.Client, config *Config) *RequestExecutor {
	re := RequestExecutor{}
	re.httpClient = httpClient
	re.config = config

	if (httpClient == nil) {
		re.httpClient = &http.Client{}
	}

	return &re
}

func (re *RequestExecutor) Get(url string) ([]byte, error){
	return re.doRequest("GET", url, nil)
}

func (re *RequestExecutor) Post(url string, body  io.Reader) ([]byte, error) {
	return re.doRequest("POST", url, body)
}

func (re *RequestExecutor) Put(url string, body  io.Reader) ([]byte, error) {
	return re.doRequest("PUT", url, body)
}

func (re *RequestExecutor) Delete(url string) ([]byte, error) {
	return re.doRequest("DELETE", url, nil)
}

func (re *RequestExecutor) doRequest(method string, url string, body io.Reader) ([]byte, error) {
	url = re.config.Okta.Client.OrgUrl + "api/v1" + url

	req, err := http.NewRequest(method, url, body)
	if (err != nil ) {
		return nil, err
	}

	req.Header.Add("Authorization", "SSWS " + re.config.Okta.Client.Token)
	req.Header.Add("User-Agent", "okta-sdk-golang/0.0.0-development") //TODO: create a User-Agent object to fill and create this on build.

	resp, err := re.httpClient.Do(req)
	if (err != nil ) {
		return nil, err
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if (err != nil ) {
		return nil, err
	}

	return bodyBytes, nil
}