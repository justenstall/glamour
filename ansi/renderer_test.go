package ansi

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/muesli/termenv"
	"github.com/yuin/goldmark"

	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

const (
	generateExamples = false
	generateIssues   = false
	examplesDir      = "../styles/examples/"
	issuesDir        = "../testdata/issues/"
)

func TestRenderer(t *testing.T) {
	files, err := filepath.Glob(examplesDir + "*.md")
	if err != nil {
		t.Fatal(err)
	}

	for _, f := range files {
		bn := strings.TrimSuffix(filepath.Base(f), ".md")
		sn := filepath.Join(examplesDir, bn+".style")
		tn := filepath.Join("../testdata", bn+".test")

		in, err := os.ReadFile(f)
		if err != nil {
			t.Fatal(err)
		}
		b, err := os.ReadFile(sn)
		if err != nil {
			t.Fatal(err)
		}

		options := Options{
			WordWrap:     80,
			ColorProfile: termenv.TrueColor,
		}
		err = json.Unmarshal(b, &options.Styles)
		if err != nil {
			t.Fatal(err)
		}

		md := goldmark.New(
			goldmark.WithExtensions(
				extension.GFM,
				extension.DefinitionList,
				// emoji.Emoji,
			),
			goldmark.WithParserOptions(
				parser.WithAutoHeadingID(),
			),
		)

		ar := NewRenderer(options)
		md.SetRenderer(
			renderer.NewRenderer(
				renderer.WithNodeRenderers(util.Prioritized(ar, 1000))))

		var buf bytes.Buffer
		err = md.Convert(in, &buf)
		if err != nil {
			t.Error(err)
		}

		// generate
		if generateExamples {
			err = os.WriteFile(tn, buf.Bytes(), 0o644)
			if err != nil {
				t.Fatal(err)
			}
			continue
		}

		// verify
		td, err := os.ReadFile(tn)
		if err != nil {
			t.Fatal(err)
		}

		if !bytes.Equal(td, buf.Bytes()) {
			t.Errorf("Rendered output for %s doesn't match!\nExpected: `\n%s`\nGot: `\n%s`\n",
				bn, string(td), buf.String())
		}
	}
}

func TestRendererIssues(t *testing.T) {
	files, err := filepath.Glob(issuesDir + "*.md")
	if err != nil {
		t.Fatal(err)
	}

	for _, f := range files {
		bn := strings.TrimSuffix(filepath.Base(f), ".md")
		t.Run(bn, func(t *testing.T) {
			tn := filepath.Join(issuesDir, bn+".test")

			in, err := os.ReadFile(f)
			if err != nil {
				t.Fatal(err)
			}
			b, err := os.ReadFile("../styles/dark.json")
			if err != nil {
				t.Fatal(err)
			}

			options := Options{
				WordWrap:     80,
				ColorProfile: termenv.TrueColor,
			}
			err = json.Unmarshal(b, &options.Styles)
			if err != nil {
				t.Fatal(err)
			}

			md := goldmark.New(
				goldmark.WithExtensions(
					extension.GFM,
					extension.DefinitionList,
					// emoji.Emoji,
				),
				goldmark.WithParserOptions(
					parser.WithAutoHeadingID(),
				),
			)

			ar := NewRenderer(options)
			md.SetRenderer(
				renderer.NewRenderer(
					renderer.WithNodeRenderers(util.Prioritized(ar, 1000))))

			var buf bytes.Buffer
			err = md.Convert(in, &buf)
			if err != nil {
				t.Error(err)
			}

			// generate
			if generateIssues {
				err = os.WriteFile(tn, buf.Bytes(), 0o644)
				if err != nil {
					t.Fatal(err)
				}
				return
			}

			// verify
			td, err := os.ReadFile(tn)
			if err != nil {
				t.Fatal(err)
			}

			if !bytes.Equal(td, buf.Bytes()) {
				t.Errorf("Rendered output for %s doesn't match!\nExpected: `\n%s`\nGot: `\n%s`\n",
					bn, string(td), buf.String())
			}
		})
	}
}
