
var N = function(){};

N.ws = new WebSocket("ws://192.168.1.20:1234/relays");
N.DataHandler = new Object();
N.Heat = new Object();
N.Pump = new Object();


N.DataHandler.handle = function (msg) {
  var o = JSON.parse(msg);
  switch (o.Type) {
    case "pumpStateChanged":
    N.Pump.callback(o);
    break;

    case "heatStateChanged":
    N.Heat.callback(o);
    break;

  }
};

N.ws.onmessage = function(e) {
    console.log("received:" + event.data);
    N.DataHandler.handle(event.data);
};


var ButtonMotor = React.createClass({
  getInitialState: function() {
    var self = this
    fetch('/control/pump')
      .then(function(response) {
        //alert(response.headers.get('Content-Type')); // application/json; charset=utf-8
        //alert(response.status); // 200
        if (response.status == 200) {
          //self.setState({message: "Pump: " + response.text});
          var txt = response.text();
          //alert (txt)
          return txt;
        } else {
          alert(response.statusText)
        }
       })
       .then(function(state) {
         self.props.val = state;
         self.setState({message: "Pump: " + state});
         //alert(state);
       })
    return {message: "checking..."};
  },

  handleSubmit: function() {
    /*
      N.ws.send(this.props.val);
      N.ws.send("{\"Command\":\"" + N.CommandOnPumpr1 + "\"}");
      */
      if (this.props.val === "Auto") {
        this.props.val = "On";
      } else if (this.props.val === "On") {
        this.props.val = "Off";
      } else {
        this.props.val = "Auto";
      }
      var v = this.props.val
      var self = this
      fetch('/control/pump?state=' + v)
        .then(function(response) {
          //alert(response.headers.get('Content-Type')); // application/json; charset=utf-8
          //alert(response.status); // 200
          if (response.status == 200) {
            //self.setState({message: "Pump: " + v});
            return response.text();
          } else {
            alert(response.statusText)
          }
         })
         .then(function(state) {
           self.props.val = state;
           self.setState({message: "Pump: " + state});
           alert(state);
         })
        .catch( alert );
  },

  handleChange(e) {
     this.setState({message: e.target.value})
   },

   componentWillMount() {
     N.Pump.callback = (data) => {
       switch (data.Key ) {
         case "state":
           this.setState({
             message: "Pump: " + data.Value
           });
           break;
       }
     };
   },


  render: function() {
    return (
      <div>
      <button className="mui-btn mui-btn--flat bdefault" onClick={this.handleSubmit} >
      {this.state.message}</button>
      </div>
    );
  }
});


var ButtonHeater = React.createClass({
  getInitialState: function() {
    var self = this
    fetch('/control/heat')
      .then(function(response) {

        if (response.status == 200) {
          return response.text();
        } else {
          alert(response.statusText)
        }
       })
       .then(function(state) {
         self.props.val = state;
         self.setState({message: "Heater: " + state});
         //alert(state);
       })
    return {message: "checking..."};
  },

  handleSubmit: function() {
      if (this.props.val === "Auto") {
        this.props.val = "On";
      } else if (this.props.val === "On") {
        this.props.val = "Off";
      } else {
        this.props.val = "Auto";
      }
      this.setState({message: "Heater: " + this.props.val})

      var v = this.props.val
      var self = this
      fetch('/control/heat?state=' + v)
        .then(function(response) {

          if (response.status == 200) {
            //self.setState({message: "Heater: " + v});
            return response;
          } else {
            alert(response.statusText)
          }
         })
          .catch( alert );
  },

  handleChange(e) {
     this.setState({message: e.target.value})
   },

   componentWillMount() {
     N.Heat.callback = (data) => {
       switch (data.Key ) {
         case "state":
           this.setState({
             message: "Heat: " + data.Value
           });
           break;
       }
     };
   },

   /*
  componentWillMount(){
    DataHandler.callback = (data) => {
       this.setState({message: data.message});
     };
  },
*/
  render: function() {
    return (
      <div>
      <button className="mui-btn mui-btn--flat bdefault" onClick={this.handleSubmit} >
      {this.state.message}</button>
      </div>
    );
  }
});



ReactDOM.render(<ButtonMotor msg="Pump: Auto" />, document.getElementById('relayMotor'));
ReactDOM.render(<ButtonHeater msg="Heater: Auto" />, document.getElementById('relayHeater'));
