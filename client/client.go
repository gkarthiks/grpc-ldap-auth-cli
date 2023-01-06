package client

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-ldap-auth-cli/proto"
	"os"
)

func EstablishServerConnection() (ldapServiceClient proto.SimpleLDAPServiceClient) {
	logrus.Infoln("establishing a connection to the ldap server over gRPC...")
	conn, err := grpc.DialContext(context.Background(), fmt.Sprintf("%s:%s", "localhost", "6000"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Errorf("error occurred while establishing the gRPC connection: %v", err)
		os.Exit(1)
	}
	ldapServiceClient = proto.NewSimpleLDAPServiceClient(conn)
	logrus.Infoln("connection established")
	return
}
