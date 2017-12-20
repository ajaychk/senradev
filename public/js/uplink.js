;
$(function(){
    console.log("uplink set up")
    rcvUplink()
});


//Receive uplink
function rcvUplink(){
    var ws = new WebSocket('ws://'+location.host+'/uplink')

    ws.onopen = function(e){
        console.log("connection opened")
        ws.send("connection opened")
    }
    ws.onclose = function(e){
        console.log("connection closed")
    }
    ws.onerror = function(e){
        console.log("connection error")
    }
    ws.onmessage = function(e){
       // var data = JSON.parse(e.data)
        console.log("message received ", e.data)
        updateUplink(e.data)
    }
}

function updateUplink(data){
    if (data == null) {
        console.log('null data')
        return
    }
    console.log('data is', data)
    $('#lastUpdate').text(Date())
    $('#uplink').text(data)
}