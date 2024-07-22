// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package chatconn

import (
	"ZChat/apps/group-chat/internal/svc"
	"ZChat/apps/group-chat/internal/types"
	"ZChat/apps/user/rpc/pb/user"
	"context"
	"fmt"
	"net/http"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) connAllUser(svc *svc.ServiceContext, w http.ResponseWriter, r *http.Request) {
	userList, err := svc.UserRpcService.AllUser(context.Background(), &user.AllUserReq{})
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, userEntity := range userList.User {
		var u = types.User{}
		u.Id = userEntity.Id
		u.Name = userEntity.Name
		fmt.Println(u)
		serveWs(h, w, r, u)
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
