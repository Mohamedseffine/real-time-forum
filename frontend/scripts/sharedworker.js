

let windows = []


onmessage((evt)=>{
 let port = evt.ports[0]
 windows.push(port)
 port.postMessage("the shared worker is connected")
 port.start()
});
