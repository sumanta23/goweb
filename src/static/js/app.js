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