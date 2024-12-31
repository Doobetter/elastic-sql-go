/*
 *  Copyright 2020-present Doobetter. All rights reserved.
 *  Use of this source code is governed by a MIT-license.
 *
 */

package elasticsql

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Doobetter/elastic-sql-go/src/client"
	"github.com/Doobetter/elastic-sql-go/src/conf"
	"github.com/Doobetter/elastic-sql-go/src/grammer"
	"testing"
)

func TestConf(t *testing.T) {
	// default conf
	cj, _ := json.Marshal(conf.DefaultElasticSQLConf)
	fmt.Println(string(cj))
}
func PrintObject(out interface{}) {
	cj, _ := json.Marshal(out)
	fmt.Println(string(cj))
}

func TestESClient(t *testing.T) {
	ctx := context.Background()
	conn := client.GetDefaultClientConnection()
	res, _ := conn.Client.CatHealth().Do(ctx)
	PrintObject(res)
}

func TestPath(t *testing.T) {
	fmt.Println(grammer.GetPathArray("a$b$c"))
}
func TestLang(t *testing.T) {
	fmt.Println(fmt.Sprintf("%v", float64(234.0)))
}
