
var ws = new WebSocket("ws://localhost:1234/echo");
var DataHandler = new Object();

ws.onmessage = function(e) {
    console.log("received:" + event.data);
    DataHandler.callback({message: "incoming " + event.data});
};

var Button = React.createClass({
  getInitialState: function() {
    this.props.val = 1
    return {message: this.props.msg};
  },

  handleSubmit: function() {
      ws.send(this.props.val);
      this.setState({message: "clicked: " + this.props.val++})
  },

  handleChange(e) {
     this.setState({email: e.target.value})
   },

  componentWillMount(){
    DataHandler.callback = (data) => {
       this.setState({message: data.message});
     };
  },

  render: function() {
    return (
      <div>
      <button className="mui-btn mui-btn--flat" onClick={this.handleSubmit} >
      Button: {this.state.message}</button>
      </div>
    );
  }
});

var Temperature = React.createClass({
  getInitialState: function() {
    return {message: this.props.msg};
  },

  componentWillMount(){
    DataHandler.callback = (data) => {
       this.setState({message: data.message});
     };
  },

  render: function() {
    return (
      <div className="mui-container">
      Temperature sensor: {this.state.message}
      </div>
    );
  }
});


ReactDOM.render(<Button msg="John" />, document.getElementById('relayMotor'));
ReactDOM.render(<Temperature />, document.getElementById('sensorInside'));
