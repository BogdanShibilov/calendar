package grpc

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
	"strings"
)

const localhost = "127.0.0.1"

func onlyLocalhost(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	p, ok := peer.FromContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.PermissionDenied, "invalid peer")
	}
	peerIpAddr, _, _ := strings.Cut(p.Addr.String(), ":")
	if peerIpAddr != localhost {
		return nil, status.Errorf(codes.PermissionDenied, "invalid peer")
	}

	return handler(ctx, req)
}
