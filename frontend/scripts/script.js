if (window["WebSocket"]) {
    conn = new WebSocket("ws://" + document.location.host + "/ws")

}