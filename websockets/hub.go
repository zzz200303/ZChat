package websockets

type Hub struct {
	groups map[string]map[*Client]bool
}

var H = Hub{
	groups: make(map[string]map[*Client]bool),
}

func (h *Hub) BroadcastToGroup(groupID string, message []byte) {
	clients := h.groups[groupID]
	for client := range clients {
		select {
		case client.send <- message:
		default:
			close(client.send)
			delete(clients, client)
		}
	}
}

func (h *Hub) AddClientToGroup(groupID string, client *Client) {
	if _, ok := h.groups[groupID]; !ok {
		h.groups[groupID] = make(map[*Client]bool)
	}
	h.groups[groupID][client] = true
}
