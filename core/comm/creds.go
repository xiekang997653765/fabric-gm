/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package comm

import (
	"context"
	"errors"
	"net"

	"github.com/xiekang997653765/CryptoGM/tls"
	"github.com/hyperledger/fabric/common/flogging"
	"google.golang.org/grpc/credentials"
)

var (
	ErrClientHandshakeNotImplemented = errors.New("core/comm: client handshakes are not implemented with serverCreds")
	ErrServerHandshakeNotImplemented = errors.New("core/comm: server handshakes are not implemented with clientCreds")
	ErrOverrideHostnameNotSupported  = errors.New("core/comm: OverrideServerName is not supported")

	// alpnProtoStr are the specified application level protocols for gRPC.
	alpnProtoStr = []string{"h2"}
)

// NewServerTransportCredentials returns a new initialized
// grpc/credentials.TransportCredentials
func NewServerTransportCredentials(
	serverConfig *tls.Config,
	logger *flogging.FabricLogger) credentials.TransportCredentials {

	// NOTE: unlike the default grpc/credentials implementation, we do not
	// clone the tls.Config which allows us to update it dynamically
	serverConfig.NextProtos = alpnProtoStr

	// GMTLS support
	if serverConfig.GMSupport != nil {
		serverConfig.MinVersion = tls.VersionGMSSL
	} else {
		serverConfig.MinVersion = tls.VersionTLS12
	}
	serverConfig.MaxVersion = tls.VersionTLS12
	return &serverCreds{
		serverConfig: serverConfig,
		logger:       logger}
}

// serverCreds is an implementation of grpc/credentials.TransportCredentials.
type serverCreds struct {
	serverConfig *tls.Config
	logger       *flogging.FabricLogger
}

// ClientHandShake is not implemented for `serverCreds`.
func (sc *serverCreds) ClientHandshake(context.Context,
	string, net.Conn) (net.Conn, credentials.AuthInfo, error) {
	return nil, nil, ErrClientHandshakeNotImplemented
}

// ServerHandshake does the authentication handshake for servers.
func (sc *serverCreds) ServerHandshake(rawConn net.Conn) (net.Conn, credentials.AuthInfo, error) {
	conn := tls.Server(rawConn, sc.serverConfig)
	if err := conn.Handshake(); err != nil {
		if sc.logger != nil {
			sc.logger.With("remote address",
				conn.RemoteAddr().String()).Errorf("TLS handshake failed with error %s", err)
		}
		return nil, nil, err
	}
	return conn, credentials.TLSInfo{State: conn.ConnectionState()}, nil
}

// Info provides the ProtocolInfo of this TransportCredentials.
func (sc *serverCreds) Info() credentials.ProtocolInfo {
	return credentials.ProtocolInfo{
		SecurityProtocol: "tls",
		SecurityVersion:  "1.2",
	}
}

// Clone makes a copy of this TransportCredentials.
func (sc *serverCreds) Clone() credentials.TransportCredentials {
	creds := NewServerTransportCredentials(sc.serverConfig, sc.logger)
	return creds
}

// OverrideServerName overrides the server name used to verify the hostname
// on the returned certificates from the server.
func (sc *serverCreds) OverrideServerName(string) error {
	return ErrOverrideHostnameNotSupported
}

type DynamicClientCredentials struct {
	TLSConfig  *tls.Config
	TLSOptions []TLSOption
}

func (dtc *DynamicClientCredentials) latestConfig() *tls.Config {
	tlsConfigCopy := dtc.TLSConfig.Clone()
	for _, tlsOption := range dtc.TLSOptions {
		tlsOption(tlsConfigCopy)
	}
	return tlsConfigCopy
}

func (dtc *DynamicClientCredentials) ClientHandshake(ctx context.Context, authority string, rawConn net.Conn) (net.Conn, credentials.AuthInfo, error) {
	return credentials.NewTLS(dtc.latestConfig()).ClientHandshake(ctx, authority, rawConn)
}

func (dtc *DynamicClientCredentials) ServerHandshake(rawConn net.Conn) (net.Conn, credentials.AuthInfo, error) {
	return nil, nil, ErrServerHandshakeNotImplemented
}

func (dtc *DynamicClientCredentials) Info() credentials.ProtocolInfo {
	return credentials.NewTLS(dtc.latestConfig()).Info()
}

func (dtc *DynamicClientCredentials) Clone() credentials.TransportCredentials {
	return credentials.NewTLS(dtc.latestConfig())
}

func (dtc *DynamicClientCredentials) OverrideServerName(name string) error {
	dtc.TLSConfig.ServerName = name
	return nil
}
