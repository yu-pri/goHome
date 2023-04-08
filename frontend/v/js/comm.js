var N = function(){};
N.HOST = "sasha123.ddns.ukrtel.net:1234";

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




var Ns = function(){};
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

/*
Ns.ws = new WebSocket("ws://" + N.HOST + "/relays");

Ns.ws.onmessage = function(e) {
    console.log("received:" + event.data);
    Ns.DataHandler.handle(event.data);
};
*/

wsConnect("ws://" + N.HOST + "/echo", N.DataHandler.handle);
wsConnect("ws://" + N.HOST + "/relays", Ns.DataHandler.handle);

function wsConnect(wsurl, handler) {
  var ws = null;

  if (ws) {
    ws.close(3001);
  } else {
    ws = new WebSocket(wsurl);
    ws.onopen = function() {
      console.log("Connection: " + wsurl + " Opened");
    };

    ws.onmessage = function(msg) {
      console.log(msg);
      handler(msg.data);
    };

    ws.onclose = function(evt) {
      ws = null;

      console.log('ws error: ' + evt.type);
      setTimeout(function(){wsConnect(wsurl, handler)}, 30000);
    };

    ws.onerror = function(evt) {
      if (ws.readyState == 1) {
        console.log('ws error: ' + evt.type);
      }
    };
  }
}

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
