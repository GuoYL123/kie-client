/*
 * Copyright 2020 Huawei Technologies Co., Ltd
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package kie

import "strconv"

const (
	defaultProject = "default"
)

//GetOption is the functional option of client func
type GetOption func(*GetOptions)

//OpOption is the functional option of client func
type OpOption func(*OpOptions)

//GetOptions is the options of client func
type GetOptions struct {
	Labels   []map[string]string
	Project  string
	Key      string
	Wait     string
	Exact    bool
	Revision string
}

//OpOptions is the options of client func
type OpOptions struct {
	Project string
}

//WithLabels query kv by labels
func WithLabels(l ...map[string]string) GetOption {
	return func(options *GetOptions) {
		options.Labels = append(options.Labels, l...)
	}
}

//WithGetProject query keys with certain project
func WithGetProject(project string) GetOption {
	return func(options *GetOptions) {
		options.Project = project
	}
}

//WithExact means label exact match
func WithExact() GetOption {
	return func(options *GetOptions) {
		options.Exact = true
	}
}

//WithWait is for long polling,format is 1s,2m
func WithWait(duration string) GetOption {
	return func(options *GetOptions) {
		options.Wait = duration
	}
}

//WithKey query keys with certain key
func WithKey(k string) GetOption {
	return func(options *GetOptions) {
		options.Key = k
	}
}

//WithRevision query keys with certain revision
func WithRevision(revision int) GetOption {
	return func(options *GetOptions) {
		options.Revision = strconv.Itoa(revision)
	}
}

//WithProject set project to param
func WithProject(project string) OpOption {
	return func(options *OpOptions) {
		options.Project = project
	}
}
