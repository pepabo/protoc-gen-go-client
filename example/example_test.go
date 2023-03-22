package example

import (
	"context"
	"testing"

	"github.com/k1LoW/grpcstub"
	"github.com/pepabo/protoc-gen-go-client/example/gen/go/myapp"
)

func TestClient(t *testing.T) {
	ctx := context.Background()
	ts := grpcstub.NewServer(t, "proto/myapp/*.proto")
	t.Cleanup(func() {
		ts.Close()
	})
	ts.ResponseDynamic()
	client := myapp.New(ts.Conn())
	if _, err := client.UserService().CreateUser(ctx, &myapp.DummyRequest{}); err != nil {
		t.Fatal(err)
	}
	{
		got := len(ts.Requests())
		if want := 1; got != want {
			t.Errorf("got %v\nwant %v", got, want)
		}
	}
}
