package handler

import (
	"context"
	"testing"
)

func TestHandleRequest(t *testing.T) {

	ctx := context.Background()

	resp, err := HandleRequest(ctx)

	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("wrong status code: %d", resp.StatusCode)
	}
}
