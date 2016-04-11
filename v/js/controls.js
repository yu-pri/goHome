var ButtonMotor = React.createClass({
  getInitialState: function() {
    this.props.val = "Auto"
    return {message: this.props.msg};
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

      this.setState({message: "Pump: " + this.props.val})
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
    this.props.val = "Auto";
    return {message: this.props.msg};
  },

  handleSubmit: function() {
      //N.ws.send(this.props.val);
      //this.setState({message: "clicked: " + this.props.val++})
      if (this.props.val === "Auto") {
        this.props.val = "On";
      } else if (this.props.val === "On") {
        this.props.val = "Off";
      } else {
        this.props.val = "Auto";
      }
      this.setState({message: "Heater: " + this.props.val})
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
