/*
 * SPDX-License-Identifier: AGPL-3.0-only
 * Copyright (c) 2023, daeuniverse Organization <team@v2raya.org>
 */

package connection

import (
	"github.com/daeuniverse/dae-wing/dae"
	"github.com/daeuniverse/dae/control"
)

type Resolver struct{}

func (r *Resolver) Connection() []control.ConnectionInfo {
	connections := dae.InConnections()
	list := make([]control.ConnectionInfo, 0)
	if connections != nil {
		connections.Range(func(key, value interface{}) bool {
			list = append(list, value.(control.ConnectionInfo))
			return true
		})
	}
	return list
}
