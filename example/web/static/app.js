setTimeout(() => {
    let msgId = 0
    let channel
    const source = new EventSource("http://localhost:3000/open")
    source.onmessage = (event) => {
        if (msgId == 0) {
            channel = event.data
            msgId++
            return
        }

        let p = document.createElement("pre");
        p.textContent = event.data
        document.getElementById("messages").append(p)

        msgId++
    }

    setInterval(() => {
        if (channel) {
            fetch("http://localhost:3000/notify?channel=" + channel, {
                method: 'POST',
                body: JSON.stringify({ message: "test", id: msgId })
            });
        }
    }, 5000)
}, 2000)