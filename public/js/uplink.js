$(function(){
    console.log("uplink set up")
    table = $('#tUplink').DataTable()

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
        var data = JSON.parse(e.data)
        console.log("message received ", e.data)
        updateUplink(data)
    }
}

function updateUplink(data){
    if (data == null) {
        console.log('null data')
        return
    }
    console.log('data is', data)
    table.row.add([
        data.seqno,
        data.txtime,
        data.devEui,
        data.gwEui,
        data.pdu,
        data.port
    ]).draw(false)
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