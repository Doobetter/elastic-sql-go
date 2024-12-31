/*
 *  Copyright 2020-present Doobetter. All rights reserved.
 *  Use of this source code is governed by a MIT-license.
 *
 */

package grammer

import (
	"errors"
	"github.com/olivere/elastic/v7"
)

type ScriptAdapter struct {
	source string
	lang   string
	param  map[string]interface{}
}

func (s *ScriptAdapter) ToScript() (*elastic.Script, error) {
	if s.source == "" {
		return nil, errors.New("script [source] must not be null")
	}

	if s.lang == "" {
		s.lang = "painless"
	}
	script := elastic.NewScript(s.source)
	script.Lang(s.lang)
	script.Params(s.param)
	return script, nil
}
