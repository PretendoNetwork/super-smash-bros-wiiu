package globals

import (
	"database/sql"

	pb "github.com/PretendoNetwork/grpc/go/account"
	"github.com/PretendoNetwork/nex-go/v2"
	common_globals "github.com/PretendoNetwork/nex-protocols-common-go/v2/globals"
	mmextension "github.com/PretendoNetwork/nex-protocols-common-go/v2/matchmake-extension"
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
var HPPServer *nex.HPPServer
var GRPCAccountClientConnection *grpc.ClientConn
var GRPCAccountClient pb.AccountClient
var GRPCAccountCommonMetadata metadata.MD
var S3Bucket string
var S3Key string
var S3Secret string
var S3Url string
var MinIOClient *minio.Client
var Presigner *S3Presigner
var Postgres *sql.DB
var MatchmakeExtensionCommonProtocol *mmextension.CommonProtocol
var MatchmakingManager *common_globals.MatchmakingManager
