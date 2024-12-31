/*
 *  Copyright 2020-present Doobetter. All rights reserved.
 *  Use of this source code is governed by a MIT-license.
 *
 */

package elasticsql

import "testing"

func TestMetricDistinct(t *testing.T) {
	sql := "select * from simba_online where ctime>'now-1d' | distinct() map r1"
	RunAndPrintResult(sql, "r1")
}
