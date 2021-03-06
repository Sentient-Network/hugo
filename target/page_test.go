// Copyright 2015 The Hugo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package target

import (
	"path/filepath"
	"testing"

	"github.com/spf13/hugo/hugofs"
)

func TestPageTranslator(t *testing.T) {
	fs := hugofs.NewMem()

	tests := []struct {
		content  string
		expected string
	}{
		{"/", "index.html"},
		{"index.html", "index.html"},
		{"bar/index.html", "bar/index.html"},
		{"foo", "foo/index.html"},
		{"foo.html", "foo/index.html"},
		{"foo.xhtml", "foo/index.xhtml"},
		{"section", "section/index.html"},
		{"section/", "section/index.html"},
		{"section/foo", "section/foo/index.html"},
		{"section/foo.html", "section/foo/index.html"},
		{"section/foo.rss", "section/foo/index.rss"},
	}

	for _, test := range tests {
		f := &PagePub{Fs: fs}
		dest, err := f.Translate(filepath.FromSlash(test.content))
		expected := filepath.FromSlash(test.expected)
		if err != nil {
			t.Fatalf("Translate returned and unexpected err: %s", err)
		}

		if dest != expected {
			t.Errorf("Translate expected return: %s, got: %s", expected, dest)
		}
	}
}

func TestPageTranslatorBase(t *testing.T) {
	tests := []struct {
		content  string
		expected string
	}{
		{"/", "a/base/index.html"},
	}

	for _, test := range tests {
		f := &PagePub{PublishDir: "a/base"}
		fts := &PagePub{PublishDir: "a/base/"}

		for _, fs := range []*PagePub{f, fts} {
			dest, err := fs.Translate(test.content)
			if err != nil {
				t.Fatalf("Translated returned and err: %s", err)
			}

			if dest != filepath.FromSlash(test.expected) {
				t.Errorf("Translate expected: %s, got: %s", test.expected, dest)
			}
		}
	}
}

func TestTranslateUglyURLs(t *testing.T) {
	tests := []struct {
		content  string
		expected string
	}{
		{"foo.html", "foo.html"},
		{"/", "index.html"},
		{"section", "section.html"},
		{"index.html", "index.html"},
	}

	for _, test := range tests {
		f := &PagePub{UglyURLs: true}
		dest, err := f.Translate(filepath.FromSlash(test.content))
		if err != nil {
			t.Fatalf("Translate returned an unexpected err: %s", err)
		}

		if dest != test.expected {
			t.Errorf("Translate expected return: %s, got: %s", test.expected, dest)
		}
	}
}

func TestTranslateDefaultExtension(t *testing.T) {
	f := &PagePub{DefaultExtension: ".foobar"}
	dest, _ := f.Translate("baz")
	if dest != filepath.FromSlash("baz/index.foobar") {
		t.Errorf("Translate expected return: %s, got %s", "baz/index.foobar", dest)
	}
}
