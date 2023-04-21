package handlers

import (
	"github.com/CloudyKit/jet/v6"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sort"
)

// wsChan is a channel for websocket payloads
var wsChan = make(chan WsPayload)

// clients is a map of websocket connections
var clients = make(map[WebSocketConnection]string)

var response WsJsonResponse

// views is a set of jet templates
var views = jet.NewSet(
	jet.NewOSFileSystemLoader("./html"),
	jet.InDevelopmentMode(),
)

// upgradeConnection is a websocket upgrader
var upgradeConnection = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Home renders the home page
func Home(w http.ResponseWriter, r *http.Request) {
	log.Println("Home handler called")

	err := renderPage(w, "home.jet", nil)
	if err != nil {
		log.Println(err)
	}
}

// renderPage renders a page using the jet template engine
func renderPage(w http.ResponseWriter, tmpl string, data jet.VarMap) error {
	view, err := views.GetTemplate(tmpl)
	if err != nil {
		log.Println(err)
		return err
	}

	err = view.Execute(w, data, nil)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// WebSocketConnection defines a websocket connection
type WebSocketConnection struct {
	*websocket.Conn
}

// WsJsonResponse defines the response sent back from websocket
type WsJsonResponse struct {
	Action         string   `json:"action"`
	Message        string   `json:"message"`
	MessageType    string   `json:"message_type"`
	ConnectedUsers []string `json:"connected_users"`
}

// WsPayload defines the payload sent to the websocket
type WsPayload struct {
	Action   string              `json:"action"`
	Username string              `json:"username"`
	Message  string              `json:"message"`
	Conn     WebSocketConnection `json:"conn"`
}

// WsEndpoint handles the websocket connection
func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgradeConnection.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client socket successfully connected")

	conn := WebSocketConnection{Conn: ws}
	//clients[conn] = ""

	// inform actual users
	users := getUserList()
	if len(users) > 0 {
		response.ConnectedUsers = users
		response.Action = "users_list"

		err = ws.WriteJSON(response)
		if err != nil {
			log.Println(err)
		}
	}

	go ListenForWs(&conn)
}

// ListenForWs listens for websocket messages
func ListenForWs(conn *WebSocketConnection) {
	defer func() {
		if err := conn.Close(); err != nil {
			log.Println("Websocket connection closed with error: ", err)
		}
	}()

	for {
		var payload WsPayload
		err := conn.ReadJSON(&payload)
		if err != nil {
			log.Println(err)
			return
		}

		payload.Conn = *conn
		wsChan <- payload
	}
}

// ListenToWsChannel listens to the websocket channel
func ListenToWsChannel() {
	var response WsJsonResponse

	for {
		e := <-wsChan
		switch e.Action {
		case "set_username":
			log.Printf("New user from %s connected %s", e.Conn, e.Username)
			clients[e.Conn] = e.Username

			users := getUserList()
			response.Action = "users_list"
			response.ConnectedUsers = users

			broadcastToAll(response)
		}

	}
}

func getUserList() []string {
	var userList []string

	for _, user := range clients {
		userList = append(userList, user)
	}
	sort.Strings(userList)
	return userList
}

// broadcastToAll broadcasts a message to all connected clients
func broadcastToAll(response WsJsonResponse) {
	for client := range clients {
		err := client.WriteJSON(response)

		if err != nil {
			log.Println(err)
			_ = client.Close()
			delete(clients, client)
			return
		}
	}
}

// WsPingEndpoint handles the websocket ping connection for testing connectivity
func WsPingEndpoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgradeConnection.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Client successfully connected")

	ws.SetPingHandler(func(appData string) error {
		log.Println("Ping received")
		return nil
	})
	ws.PingHandler()
}
