package jsonl

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"github.com/whosonfirst/go-writer"
	"strings"
	"testing"
)

func TestFeatureCollectionWriter(t *testing.T) {

	features := []string{
		`{"geometry":{"coordinates":[0,0],"type":"Point"},"properties":{"id":1},"type":"Feature"}`,
		`{"geometry":{"coordinates":[0,0],"type":"Point"},"properties":{"id":2},"type":"Feature"}`,
		`{"geometry":{"coordinates":[0,0],"type":"Point"},"properties":{"id":3},"type":"Feature"}`,
	}

	enc_features := strings.Join(features, "\n")
	enc_features = fmt.Sprintf("%s\n", enc_features)

	ctx := context.Background()

	var buf bytes.Buffer
	buf_wr := bufio.NewWriter(&buf)

	ctx, err := writer.SetIOWriterWithContext(ctx, buf_wr)

	if err != nil {
		t.Fatalf("Failed to set IOWriter context, %v", err)
	}

	writer_uri := "jsonl://?writer=io://"

	wr, err := writer.NewWriter(ctx, writer_uri)

	if err != nil {
		t.Fatalf("Failed to create new writer for '%s', %v", writer_uri, err)
	}

	for _, f := range features {

		sr := strings.NewReader(f)

		_, err := wr.Write(ctx, "", sr)

		if err != nil {
			t.Fatalf("Failed to write feature, %v", err)
		}
	}

	err = wr.Close(ctx)

	if err != nil {
		t.Fatalf("Failed to close feature collection writer, %v", err)
	}

	buf_wr.Flush()

	if !bytes.Equal([]byte(enc_features), buf.Bytes()) {
		t.Fatalf("Invalid output, got '%s', expected '%s'", string(buf.Bytes()), enc_features)
	}
}
