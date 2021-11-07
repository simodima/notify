setTimeout(() => {
    let msgId = 0
    const channel = "test-channel-id"
    const source = new EventSource("http://localhost:3000/open?channel=" + channel)

    source.addEventListener('test', (event) => {
        let p = document.createElement("pre");
        p.textContent = event.data
        document.getElementById("messages").append(p)
    })

    setInterval(() => {
        if (channel) {
            fetch("http://localhost:3000/notify?channel=" + channel, {
                method: 'POST',
                body: JSON.stringify({ message: "test", id: msgId })
            }).then(() => {
                msgId++
            });
        }
    }, 5000)
}, 2000)