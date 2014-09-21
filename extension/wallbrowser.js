console.log("HELLO")

window.innerWidth = 7168;
window.innerHeight = 3072;
document.body.style.width = '7168px';
document.body.style.height = '3072px';

var x; 
var y;

$.get("http://10.1.1.60:9191/location",
        function(d, txtstatus, jqxhdr) {
            x = parseInt(d.split(",")[0])
            y=  parseInt(d.split(",")[1])
            console.log("This window should scroll to "+x+","+y) 
            window.scrollTo(x,y) 
            setTimeout(function(){
                r()
            }, 1000);
        }
) 

function r() {
    window.innerWidth = 7168;
    window.innerHeight = 3072;

    document.body.style.width = '7168px';
    document.body.style.height = '3072px';
        
    window.scrollTo(x,y);
    setTimeout(r,1000)
}

// e = jQuery.Event("keydown"); 


var ws = new WebSocket("ws://10.1.1.60:9192") 
ws.onmessage = function(event) {

    //$("#ifr").focus() 
    
    console.log("Message from server:", event.data)
    switch (parseInt(event.data)) {
    case 37:
        console.log('left arrow key pressed!');
        break;
    case 38:
        console.log('up arrow key pressed!');
        break;
    case 39:
        console.log('right arrow key pressed!');
        break;
    case 40:
        console.log('down arrow key pressed!');
        break;
    case 82: 
        console.log('r key pressed!');
        break 
    }

    $('body')[0].click();
    
    /*
    e.which = parseInt(event.data)
    e.view = window
    e.keyCode = parseInt(event.data) 
    console.log(e) 
    $("body").trigger(e) 

    */
    /*
    var e = new KeyboardEvent('keydown');
    e.which = parseInt(event.data)
    e.view = window
    e.keyCode = parseInt(event.data) 
    console.log(e)
    document.querySelector('body').dispatchEvent(e);
    */
    simulateKeyEvent("a")
} 


function simulateKeyEvent(character) {
  var evt = document.createEvent("KeyboardEvent");
  evt.initKeyboardEvent("keypress", true, true, window, 0, 0, 0, 0, 0,
          character.charCodeAt(0)) 
  var canceled = !document.body.dispatchEvent(evt);
}
