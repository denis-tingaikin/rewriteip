// Copyright (c) 2022 Cisco and/or its affiliates.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rewriteip

import (
	"context"
	"net"

	"github.com/coredns/coredns/plugin"

	"github.com/miekg/dns"
)

// RewriteIP is a plugin that rewrites IP response for A, AAA, SRV responses
type RewriteIP struct {
	Next plugin.Handler
	To   net.IP
}

// ServeDNS implements the plugin.Handler interface.
func (i *RewriteIP) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	return plugin.NextOrFailure(i.Name(), i.Next, ctx, &responseWriter{ip: i.To, ResponseWriter: w}, r)
}

// Name implements the Handler interface.
func (i *RewriteIP) Name() string { return "rewriteip" }
