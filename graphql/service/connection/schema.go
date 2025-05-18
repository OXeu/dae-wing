/*
 * SPDX-License-Identifier: AGPL-3.0-only
 * Copyright (c) 2023, daeuniverse Organization <team@v2raya.org>
 */

package connection

func Schema() (string, error) {
    return `
type Connection {
    src: String!
    dst: String!
    sniffed: String!
    dscp: Int!
    l4proto: String!
    dialer: String!
    outbound: String!
    mac: String!
    pname: String!
}
`, nil
}
