async function send(context) {
    const action = context.action;
    delete context.action;
    return fetch(action, context);
}

async function increment(){
    send({action: "/count", method: "POST"})
}

async function getCounter(){
    const response = await send({action: "/count", method: "GET"});
    const data = await response.json()
    document.getElementById("counter").innerHTML = data.count;
}

async function resetCounter(){
    const response = await send({action: "/count", method: "DELETE"});
    const data = await response.json()
    return true;
}

async function getwsurl(){
    const response = await send({action: "/getwsurl", method: "GET"});
    const data = await response.json()
    return data.ws;
}

async function connectwithWS(){
    const conn = new WebSocket("ws://" + document.location.host + (await getwsurl()));
    conn.onclose = function (evt) {
        var item = document.createElement("div");
        item.innerHTML = "<b>Connection closed.</b>";
        appendLog(item);
    };
    conn.onmessage = function (msg) {
        document.getElementById("counter").innerHTML = msg.data;

        let root = document.documentElement
        let spinspeed = (msg.data % 10)+1;
        root.style.setProperty('--spinspeed', spinspeed+ "s");
    };
}
