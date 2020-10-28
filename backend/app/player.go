import (
    "sync"
)

type playerId int

type player struct {
    id playerId 
    displayName string
    conn *websocket.Conn
    doneChannel chan *sync.WaitGroup 
}

// todo define message type
// todo do I need two goroutines per player connection? how many goroutines is too many
func (p *player) receiveThread(toManagingThread chan wsMessage, fromManagingThread chan wsMessage) {
    for {
        // Read from connection with timeout
        // on timeout, poll doneChannel and exit if there's anything there
        // otherwise, process message from websocket
        // then still poll the doneChannel so that we don't get stuck
        // if the client keeps sending messages but the backend wants to stop this thread

        // Outline below - feels very strange...
        // Wait I need to handle messages to/from the managing thread too..
        // maybe I do need another thread just to read and forward messages from conn to here?
        
        /*

        go func() {
            for {
                // poll connection and send on channel
                // but this thread needs to be closed too!!!
                // I must be structuring this wrong...
            }
        }()

        msg, err := p.conn.ReadJSON(timeout=0.1)
        if err == nil {
            messageChannel <- msg
        } else if err != timeout {
            // handle actual error
        }

        // poll doneChannel
        select {
        case wg <- p.doneChannel:
            defer wg.Done()
            return
        default:
            break
        }
        */
    }

}

