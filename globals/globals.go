package globals

import (
	pb "github.com/PretendoNetwork/grpc-go/account"
	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/plogger-go"
	"github.com/minio/minio-go/v7"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var Logger *plogger.Logger
var KerberosPassword = "password" // * Default password
var AuthenticationServer *nex.PRUDPServer
var AuthenticationEndpoint *nex.PRUDPEndPoint
var SecureServer *nex.PRUDPServer
var SecureEndpoint *nex.PRUDPEndPoint
var GRPCAccountClientConnection *grpc.ClientConn
var GRPCAccountClient pb.AccountClient
var GRPCAccountCommonMetadata metadata.MD
var S3Bucket string
var S3Key string
var S3Secret string
var S3Url string
var MinIOClient *minio.Client
var Presigner *S3Presigner
