package main

import (
	"crypto/rand"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"

	pb_account "github.com/PretendoNetwork/grpc/go/account"
	pb_friends "github.com/PretendoNetwork/grpc/go/friends"
	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/plogger-go"
	"github.com/PretendoNetwork/super-smash-bros-wiiu/globals"
	"github.com/joho/godotenv"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func init() {
	globals.Logger = plogger.NewLogger()

	var err error

	err = godotenv.Load()
	if err != nil {
		globals.Logger.Warning("Error loading .env file")
	}

	authenticationServerPort := os.Getenv("PN_SSBWIIU_AUTHENTICATION_SERVER_PORT")
	secureServerHost := os.Getenv("PN_SSBWIIU_SECURE_SERVER_HOST")
	secureServerPort := os.Getenv("PN_SSBWIIU_SECURE_SERVER_PORT")
	accountGRPCHost := os.Getenv("PN_SSBWIIU_ACCOUNT_GRPC_HOST")
	accountGRPCPort := os.Getenv("PN_SSBWIIU_ACCOUNT_GRPC_PORT")
	accountGRPCAPIKey := os.Getenv("PN_SSBWIIU_ACCOUNT_GRPC_API_KEY")
	globals.S3Bucket = os.Getenv("PN_SSBWIIU_DATASTORE_S3BUCKET")
	globals.S3Key = os.Getenv("PN_SSBWIIU_DATASTORE_S3KEY")
	globals.S3Secret = os.Getenv("PN_SSBWIIU_DATASTORE_S3SECRET")
	globals.S3Url = os.Getenv("PN_SSBWIIU_DATASTORE_S3URL")
	friendsGRPCHost := os.Getenv("PN_SSBWIIU_FRIENDS_GRPC_HOST")
	friendsGRPCPort := os.Getenv("PN_SSBWIIU_FRIENDS_GRPC_PORT")
	friendsGRPCAPIKey := os.Getenv("PN_SSBWIIU_FRIENDS_GRPC_API_KEY")

	kerberosPassword := make([]byte, 0x10)
	_, err = rand.Read(kerberosPassword)
	if err != nil {
		globals.Logger.Error("Error generating Kerberos password")
		os.Exit(0)
	}

	globals.KerberosPassword = string(kerberosPassword)

	globals.AuthenticationServerAccount = nex.NewAccount(types.NewPID(1), "Quazal Authentication", globals.KerberosPassword)
	globals.SecureServerAccount = nex.NewAccount(types.NewPID(2), "Quazal Rendez-Vous", globals.KerberosPassword)

	if strings.TrimSpace(authenticationServerPort) == "" {
		globals.Logger.Error("PN_SSBWIIU_AUTHENTICATION_SERVER_PORT environment variable not set")
		os.Exit(0)
	}

	if port, err := strconv.Atoi(authenticationServerPort); err != nil {
		globals.Logger.Errorf("PN_SSBWIIU_AUTHENTICATION_SERVER_PORT is not a valid port. Expected 0-65535, got %s", authenticationServerPort)
		os.Exit(0)
	} else if port < 0 || port > 65535 {
		globals.Logger.Errorf("PN_SSBWIIU_AUTHENTICATION_SERVER_PORT is not a valid port. Expected 0-65535, got %s", authenticationServerPort)
		os.Exit(0)
	}

	if strings.TrimSpace(secureServerHost) == "" {
		globals.Logger.Error("PN_SSBWIIU_SECURE_SERVER_HOST environment variable not set")
		os.Exit(0)
	}

	if strings.TrimSpace(secureServerPort) == "" {
		globals.Logger.Error("PN_SSBWIIU_SECURE_SERVER_PORT environment variable not set")
		os.Exit(0)
	}

	if port, err := strconv.Atoi(secureServerPort); err != nil {
		globals.Logger.Errorf("PN_SSBWIIU_SECURE_SERVER_PORT is not a valid port. Expected 0-65535, got %s", secureServerPort)
		os.Exit(0)
	} else if port < 0 || port > 65535 {
		globals.Logger.Errorf("PN_SSBWIIU_SECURE_SERVER_PORT is not a valid port. Expected 0-65535, got %s", secureServerPort)
		os.Exit(0)
	}

	if strings.TrimSpace(accountGRPCHost) == "" {
		globals.Logger.Error("PN_SSBWIIU_ACCOUNT_GRPC_HOST environment variable not set")
		os.Exit(0)
	}

	if strings.TrimSpace(accountGRPCPort) == "" {
		globals.Logger.Error("PN_SSBWIIU_ACCOUNT_GRPC_PORT environment variable not set")
		os.Exit(0)
	}

	if port, err := strconv.Atoi(accountGRPCPort); err != nil {
		globals.Logger.Errorf("PN_SSBWIIU_ACCOUNT_GRPC_PORT is not a valid port. Expected 0-65535, got %s", accountGRPCPort)
		os.Exit(0)
	} else if port < 0 || port > 65535 {
		globals.Logger.Errorf("PN_SSBWIIU_ACCOUNT_GRPC_PORT is not a valid port. Expected 0-65535, got %s", accountGRPCPort)
		os.Exit(0)
	}

	if strings.TrimSpace(accountGRPCAPIKey) == "" {
		globals.Logger.Warning("Insecure gRPC server detected. PN_SSBWIIU_ACCOUNT_GRPC_API_KEY environment variable not set")
	}

	if strings.TrimSpace(friendsGRPCHost) == "" {
		globals.Logger.Error("PN_SSBWIIU_FRIENDS_GRPC_HOST environment variable not set")
		os.Exit(0)
	}

	if strings.TrimSpace(friendsGRPCPort) == "" {
		globals.Logger.Error("PN_SSBWIIU_FRIENDS_GRPC_PORT environment variable not set")
		os.Exit(0)
	}

	if port, err := strconv.Atoi(friendsGRPCPort); err != nil {
		globals.Logger.Errorf("PN_SSBWIIU_FRIENDS_GRPC_PORT is not a valid port. Expected 0-65535, got %s", friendsGRPCPort)
		os.Exit(0)
	} else if port < 0 || port > 65535 {
		globals.Logger.Errorf("PN_SSBWIIU_FRIENDS_GRPC_PORT is not a valid port. Expected 0-65535, got %s", friendsGRPCPort)
		os.Exit(0)
	}

	if strings.TrimSpace(friendsGRPCAPIKey) == "" {
		globals.Logger.Warning("Insecure gRPC server detected. PN_SSBWIIU_FRIENDS_GRPC_API_KEY environment variable not set")
	}

	if strings.TrimSpace(globals.S3Bucket) == "" {
		globals.Logger.Error("PN_SSBWIIU_DATASTORE_S3BUCKET environment variable not set")
		os.Exit(0)
	}

	if strings.TrimSpace(globals.S3Key) == "" {
		globals.Logger.Error("PN_SSBWIIU_DATASTORE_S3KEY environment variable not set")
		os.Exit(0)
	}

	if strings.TrimSpace(globals.S3Secret) == "" {
		globals.Logger.Error("PN_SSBWIIU_DATASTORE_S3SECRET environment variable not set")
		os.Exit(0)
	}

	if strings.TrimSpace(globals.S3Url) == "" {
		globals.Logger.Error("PN_SSBWIIU_DATASTORE_S3URL environment variable not set")
		os.Exit(0)
	}

	globals.GRPCAccountClientConnection, err = grpc.NewClient(fmt.Sprintf("%s:%s", accountGRPCHost, accountGRPCPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		globals.Logger.Criticalf("Failed to connect to account gRPC server: %v", err)
		os.Exit(0)
	}

	globals.GRPCAccountClient = pb_account.NewAccountClient(globals.GRPCAccountClientConnection)
	globals.GRPCAccountCommonMetadata = metadata.Pairs(
		"X-API-Key", accountGRPCAPIKey,
	)

	globals.GRPCFriendsClientConnection, err = grpc.Dial(fmt.Sprintf("%s:%s", friendsGRPCHost, friendsGRPCPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		globals.Logger.Criticalf("Failed to connect to friends gRPC server: %v", err)
		os.Exit(0)
	}

	globals.GRPCFriendsClient = pb_friends.NewFriendsClient(globals.GRPCFriendsClientConnection)
	globals.GRPCFriendsCommonMetadata = metadata.Pairs(
		"X-API-Key", friendsGRPCAPIKey,
	)

	globals.Postgres, err = sql.Open("postgres", os.Getenv("PN_SSBWIIU_POSTGRES_URI"))
	if err != nil {
		globals.Logger.Critical(err.Error())
	}

	staticCredentials := credentials.NewStaticV4(globals.S3Key, globals.S3Secret, "")

	minIOClient, err := minio.New(globals.S3Url, &minio.Options{
		Creds:  staticCredentials,
		Secure: true,
	})
	if err != nil {
		globals.Logger.Criticalf("Error occured during minio connection: %v", err)
	}

	globals.MinIOClient = minIOClient
	globals.Presigner = globals.NewS3Presigner(globals.MinIOClient)
}
