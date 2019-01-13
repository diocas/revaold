package main

import (
	"context"
	"fmt"

	rpcpb "github.com/cernbox/go-cs3apis/cs3/rpc"
	storageproviderv0alphapb "github.com/cernbox/go-cs3apis/cs3/storageprovider/v0alpha"
)

func statCommand() *command {
	cmd := newCommand("stat")
	cmd.Description = func() string { return "get the metadata for a file or folder" }
	cmd.Action = func() error {
		fn := "/"
		if cmd.NArg() >= 1 {
			fn = cmd.Args()[0]
		}

		ctx := context.Background()
		client, err := getStorageProviderClient()
		if err != nil {
			return err
		}

		req := &storageproviderv0alphapb.StatRequest{Filename: fn}
		res, err := client.Stat(ctx, req)
		if err != nil {
			return err
		}

		if res.Status.Code != rpcpb.Code_CODE_OK {
			return formatError(res.Status)
		}

		fmt.Println(res.Metadata)
		return nil
	}
	return cmd
}
