package globals

import (
	"context"

	pb "github.com/PretendoNetwork/grpc/go/friends"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	"google.golang.org/grpc/metadata"
)

func GetUserFriendPids(pid uint32) []uint32 {
	ctx := metadata.NewOutgoingContext(context.Background(), GRPCFriendsCommonMetadata)

	response, err := GRPCFriendsClient.GetUserFriendPIDs(ctx, &pb.GetUserFriendPIDsRequest{Pid: pid})
	if err != nil {
		globals.Logger.Error(err.Error())
		return []uint32{}
	}

	return response.Pids
}
