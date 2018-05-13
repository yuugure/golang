package main

type room struct {
	forward chan []byte
	join    chan *client
	leave   chan *client
	clients map[*client]bool
}

func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			//join
			r.clients[client] = true
		case client := <-r.leave:
			//leave
			delete(r.clients, client)
			close(client.send)
		case msg := <-r.forward:
			//send msg all client
			for client := range r.clients {
				select {
				case client.send <- msg:
					//message send
				default:
					//send fail
					delete(r.clients, client)
					close(client.send)
				}
			}
		}
	}
}
