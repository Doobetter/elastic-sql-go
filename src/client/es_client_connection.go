/*
 *  Copyright 2020-present Doobetter. All rights reserved.
 *  Use of this source code is governed by a MIT-license.
 *
 */

package client

import (
	"github.com/Doobetter/elastic-sql-go/src/conf"
	"log"
	"os"
	"time"

	"github.com/olivere/elastic/v7"
)

var DefaultClientConnection *ESClientConnection

func init() {
	DefaultClientConnection, _ = NewESClientConnection(conf.DefaultElasticSQLConf)
}

type ESClientConnection struct {
	Conf   *conf.ElasticSQLConfiguration
	Client *elastic.Client
}

func NewESClientConnection(conf *conf.ElasticSQLConfiguration) (*ESClientConnection, error) {
	conn := new(ESClientConnection)
	client, err := elastic.NewClient(
		elastic.SetURL(conf.ESClient.HttpHosts...),
		elastic.SetBasicAuth(conf.ESClient.Username, conf.ESClient.Password),
		elastic.SetSniff(false),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetGzip(true),
		elastic.SetErrorLog(log.New(os.Stdout, "", log.LstdFlags)),
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
	)
	if err != nil {
		return nil, err
	}
	conn.Conf = conf
	conn.Client = client
	return conn, nil
}

func GetDefaultClientConnection() *ESClientConnection {
	return DefaultClientConnection
}
