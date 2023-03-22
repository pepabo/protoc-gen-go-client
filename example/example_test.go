package example

import (
	"context"
	"errors"
	"io"
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
	if _, err := client.UserService().CreateUser(ctx, &myapp.CreateUserRequest{
		Name:  "alice",
		Email: "alice@example.com",
	}); err != nil {
		t.Fatal(err)
	}
	res, err := client.ProjectService().ListProjects(ctx, &myapp.ListProjectsRequest{
		Page: 1,
	})
	if err != nil {
		t.Fatal(err)
	}
	for {
		_, err := res.Recv()
		if err != nil {
			if errors.Is(io.EOF, err) {
				break
			}
			t.Fatal(err)
		}
	}
	{
		got := len(ts.Requests())
		if want := 2; got != want {
			t.Errorf("got %v\nwant %v", got, want)
		}
	}
}
