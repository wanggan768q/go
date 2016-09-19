var socket;

console.log("111")
$(document).ready(function () {
	console.log("Create a socket")
    // Create a socket
    socket = new WebSocket('ws://' + window.location.host + '/Connect');
console.log(socket.readyState)
    // Message received on the socket
    socket.onmessage = function (event) {
        var data = JSON.parse(event.data);
        console.log(data);
        switch (data.Type) {
        case 0: // JOIN
            if (data.User == $('#uname').text()) {
                $("#chatbox li").first().before("<li>You joined the chat room.</li>");
            } else {
                $("#chatbox li").first().before("<li>" + data.User + " joined the chat room.</li>");
            }
            break;
        case 1: // LEAVE
            $("#chatbox li").first().before("<li>" + data.User + " left the chat room.</li>");
            break;
        case 2: // MESSAGE
            $("#chatbox li").first().before("<li><b>" + data.User + "</b>: " + data.Content + "</li>");
            break;
        }
    };

	$('#Connect').click(function () {
        postConecnt();
    });
});