$(function(){
    console.log("uplink set up")
    table = $('#tUplink').DataTable()
    table.row.add([
        2,
        '3:6',
        'eeeeee',
        'www',
        'ddddddddddd',
        1
    ]).draw(false)

    enableSearchAll()

    rcvUplink()
});

var table;
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
    var $tr = `<tr><td>`+data.seqno+`</td>
                    <td>`+data.txtime+`</td>
                    <td>`+data.devEui+`</td>
                    <td>`+data.gwEui+`</td>
                    <td>`+data.pdu+`</td>
                    <td>`+data.port+`</td>
                </tr>`
    $('#tUplink tbody').prepend($tr)
}

function enableSearchAll(){
    table.columns( '.select-filter' ).every( function () {
        var that = this;
     
        // Create the select list and search operation
        var select = $('<select />')
            .appendTo(
                this.footer()
            )
            .on( 'change', function () {
                that
                    .search( $(this).val() )
                    .draw();
            } );
     
        // Get the search data for the first column and add to the select list
        this
            .cache( 'search' )
            .sort()
            .unique()
            .each( function ( d ) {
                select.append( $('<option value="'+d+'">'+d+'</option>') );
            } );
    } );
}