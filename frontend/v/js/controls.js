

var ButtonMotor = React.createClass({
  getInitialState: function() {
    var self = this
    fetch('/control/pump')
      .then(function(response) {
          if (response.status == 200) {
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

  handleSubmit: function() {
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
          if (response.status == 200) {
            return response.text();
          } else {
            alert(response.statusText)
          }
         })
         .then(function(state) {

         })
        .catch( alert );
  },



   componentWillMount() {
     window.Ns.Pump.callback = (data) => {
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
      //this.setState({message: "Heater: " + this.props.val})

      var v = this.props.val
      var self = this
      fetch('/control/heat?state=' + v)
        .then(function(response) {
          if (response.status == 200) {
            return response;
          } else {
            alert(response.statusText)
          }
         })
          .catch( alert );
  },

  /*
  handleChange(e) {
     this.setState({message: e.target.value})
   },
*/

   componentWillMount() {
     window.Ns.Heat.callback = (data) => {
       switch (data.Key ) {
         case "state":
           this.setState({
             message: "Heat: " + data.Value
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



ReactDOM.render(<ButtonMotor msg="Pump: Auto" />, document.getElementById('relayMotor'));
ReactDOM.render(<ButtonHeater msg="Heater: Auto" />, document.getElementById('relayHeater'));
