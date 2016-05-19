var N = function(){};
N.HOST = "alprihodko.asuscomm.com:1234";

N.DataHandler = new Object();
N.TempDataHandler = new Object();

N.DataHandler.handle = function (msg) {
  //alert(msg)
  var o = JSON.parse(msg);

  if (o.TempInside) {
    N.TempDataHandler.callback(o);
    updateChart(o);
  } else {
    console.log("nothing to do: " + o.TempInside);
  }
};

//var N.ws = null;

function wsConnect(wsurl, handler) {
  var reopen = null;
  var ws = null;

  if (ws) {
    ws.close(3001);
  } else {
    ws = new WebSocket(wsurl);
    ws.onopen = function() {
      console.log("Connection: " + wsurl + " Opened");
      if (reopen) {
        clearInterval(reopen)
      }
    };

    ws.onmessage = function(msg) {
      console.log(msg);
      handler(msg.data);
    };

    ws.onclose = function(evt) {
      ws = null;
      console.log('ws error: ' + evt.type);
      reopen = setInterval(wsConnect, 3000, ws, wsurl, handler)
    };

    ws.onerror = function(evt) {
      if (_websocket.readyState == 1) {
        console.log('ws error: ' + evt.type);
      }
    };
  }
}

wsConnect("ws://" + N.HOST + "/echo", N.DataHandler.handle);

/*
N.openTempConn = function(){
  try {
    N.ws = new WebSocket("ws://" + N.HOST + "/echo");
  } catch (err) {
    console.log(err.message);
    setTimeout(N.openTempConn(), 3000);
  }
}

N.openTempConn();


N.ws.onopen = function(){
    console.log("Connection Opened");
}

N.ws.onmessage = function(e) {
    console.log("received:" + event.data);
    N.DataHandler.handle(event.data);
};

N.ws.onclose = function(){
    console.log("Sensor connection Closed");
    setTimeout(N.openTempConn(), 3000);
}

N.ws.onerror = function(evt){
    console.log("The following error occurred: " + evt.data);
    N.ws.close();
}
*/


var Ns = function(){};
Ns.ws = new WebSocket("ws://" + N.HOST + "/relays");
Ns.DataHandler = new Object();
Ns.Heat = new Object();
Ns.Pump = new Object();


Ns.DataHandler.handle = function (msg) {
  //alert(msg)
  var o = JSON.parse(msg);
  switch (o.Type) {
    case "pumpStateChanged":
    Ns.Pump.callback(o);
    break;

    case "heatStateChanged":
    Ns.Heat.callback(o);
    break;

  }
};

Ns.ws.onmessage = function(e) {
    console.log("received:" + event.data);
    Ns.DataHandler.handle(event.data);
};


jQuery(function($) {
  var $bodyEl = $('body'),
      $sidedrawerEl = $('#sidedrawer');


  function showSidedrawer() {
    // show overlay
    var options = {
      onclose: function() {
        $sidedrawerEl
          .removeClass('active')
          .appendTo(document.body);
      }
    };

    var $overlayEl = $(mui.overlay('on', options));

    // show element
    $sidedrawerEl.appendTo($overlayEl);
    setTimeout(function() {
      $sidedrawerEl.addClass('active');
    }, 20);
  }


  function hideSidedrawer() {
    $bodyEl.toggleClass('hide-sidedrawer');
  }


  $('.js-show-sidedrawer').on('click', showSidedrawer);
  $('.js-hide-sidedrawer').on('click', hideSidedrawer);

  var $titleEls = $('strong', $sidedrawerEl);

  $titleEls
    .next()
    .hide();

  $titleEls.on('click', function() {
    $(this).next().slideToggle(200);
  });
});


/*
fetch('/control/pump')
  .then(function(response) {
    alert(response.headers.get('Content-Type')); // application/json; charset=utf-8
    alert(response.status); // 200
    if (response.status == 200) {
      //self.setState({message: "Pump: " + response.text});
      var txt = response.text();
      return txt;
    } else {
      alert(response.statusText)
    }
   })
   .then(function(state) {
     self.props.val = state;
     self.setState({message: "Pump: " + state});
   })
return {message: "checking..."};
},
*/
