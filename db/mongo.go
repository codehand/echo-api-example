package db

import (
	"crypto/tls"
	"net"
	"strings"

	"github.com/echo-restful-crud-api-example/config"
	"github.com/labstack/gommon/log"
	mgo "gopkg.in/mgo.v2"
)

var (
	session      *mgo.Session
	databaseName = "example"
)

func init() {
	stgConnection()
}

func stgConnection() {
	hosts := strings.Split(config.Config.Database.Address, ",")
	tlsConfig := &tls.Config{}
	dialInfo := &mgo.DialInfo{
		Addrs:          hosts,
		Database:       "admin",
		Username:       config.Config.Database.Username,
		Password:       config.Config.Database.Password,
		ReplicaSetName: "ClusterSVC-shard-0",
	}
	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		return conn, err
	}
	mgoSession, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Fatal(err)
	}
	session = mgoSession
}

func pullSession() *mgo.Session {
	return session.Copy()
}

// Ping connection
func Ping() error {
	sessionCopy := pullSession()
	defer sessionCopy.Close()
	return sessionCopy.Ping()
}
