var socket = new WebSocket("ws://localhost:8080/ws");

let connect = cb => {
    
}

let sendMsg = msg => {
    console.log("sending msg: ", msg);
    socket.send(msg);
}

let googleLoginPath = "http://localhost:8080/auth/google/login"


export {connect, sendMsg, googleLoginPath};