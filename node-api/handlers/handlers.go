// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2025, Berachain Foundation. All rights reserved.
// Use of this software is governed by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package handlers

import (
	"fmt"
	"net/http"

	"github.com/berachain/beacon-kit/errors"
	"github.com/berachain/beacon-kit/log"
	"github.com/labstack/echo/v4"
)

type Context = echo.Context

// handlerFn enforces a signature for all handler functions.
type handlerFn func(c Context) (any, error)

// Handlers is an interface that all handlers must implement.
type Handlers interface {
	// RegisterRoutes is a method that registers the routes for the handler.
	RegisterRoutes(logger log.Logger)
	RouteSet() *RouteSet
}

// BaseHandler is a base handler for all handlers. It abstracts the route set
// and logger from the handler.
type BaseHandler struct {
	routes *RouteSet
	logger log.Logger
}

// NewBaseHandler initializes a new base handler with the given routes and
// logger.
func NewBaseHandler(routes *RouteSet) *BaseHandler {
	return &BaseHandler{
		routes: routes,
	}
}

// NotImplemented is the handler for API endpoints that are defined in the Ethereum Beacon Node API spec,
// but not yet implemented.
func (b *BaseHandler) NotImplemented(Context) (any, error) {
	return nil, errors.New("not implemented")
}

// Deprecated handles deprecated API endpoints that are no longer supported according to the Ethereum Beacon Node API spec.
func (b *BaseHandler) Deprecated(c Context) (any, error) {
	endpoint := c.Request().URL.Path

	return nil, NewHTTPError(
		http.StatusInternalServerError,
		fmt.Sprintf("The endpoint %s is deprecated.", endpoint),
	)
}

// RouteSet returns the route set for the base handler.
func (b *BaseHandler) RouteSet() *RouteSet {
	return b.routes
}

// Logger is used to access the logger for the base handler.
func (b *BaseHandler) Logger() log.Logger {
	return b.logger
}

func (b *BaseHandler) SetLogger(logger log.Logger) {
	b.logger = logger
}

// AddRoutes adds the given slice of routes to the base handler.
func (b *BaseHandler) AddRoutes(routes []*Route) {
	b.routes.Routes = append(b.routes.Routes, routes...)
}
