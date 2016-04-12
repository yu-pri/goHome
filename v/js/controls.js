var ButtonMotor = React.createClass({
  getInitialState: function() {
    var self = this
    fetch('/control/pump')
      .then(function(response) {
        //alert(response.headers.get('Content-Type')); // application/json; charset=utf-8
        //alert(response.status); // 200
        if (response.status == 200) {
          //self.setState({message: "Pump: " + response.text});
          //alert (response.text())
          return response.text();
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
            self.setState({message: "Pump: " + v});
            return response;
          } else {
            alert(response.statusText)
          }
         })
        //.then(function(user) {
          //alert(user.name); // iliakan
        //})
        .catch( alert );
  },

  handleChange(e) {
     this.setState({message: e.target.value})
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
            self.setState({message: "Heater: " + v});
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
