

var Temperature = React.createClass({
  getInitialState: function() {
    var self = this
    fetch('/control/currentState')
      .then(function(response) {
          if (response.status == 200) {
          var txt = response.text();
          return txt;
        } else {
          alert(response.statusText);
        }
       })
       .then(function(msg) {
         var o = JSON.parse(msg);
         //self.props.temp1 = o.TempInternal;
         //self.props.temp2 = o.TempInternal;
         self.set(o);
       })

    return {temp1: this.props.temp1, temp2: this.props.temp2, temp3: this.props.temp3, temp4: this.props.temp4};
  },

  set(data) {
    this.setState({
      temp1: data.TempInside
    });

    this.setState({
      temp2: data.TempReverse
    });

    this.setState({
      temp3: data.TempEntryRoom
    });


    this.setState({
      temp4: data.TempHeater
    });

    if (data.HeaterState) {
      this.setState({
        heater: "ON"
      });
    } else {
      this.setState({
        heater: "OFF"
      });
    }

    if (data.PumpState) {
      this.setState({
        heaterPump: "ON"
      });
    } else {
      this.setState({
        heaterPump: "OFF"
      });
    }

  },

  componentWillMount() {
    window.N.TempDataHandler.callback = (data) => {
      this.set(data);
    };
  },

  render: function() {

    return (
      <div>
        <div className="mui-container-fluid" >
          <span className="temp">Kitchen:</span> {this.state.temp1}
          <span className="temp">Entry room:</span> {this.state.temp3}
          <span className="temp">H Rev: </span>{this.state.temp2}
        </div>

        <div className="mui-container-fluid" >
          <span className="temp">Heater: </span> {this.state.temp4}
          , Heater: {this.state.heater}, HeaterPump: {this.state.heaterPump}
        </div>


      </div>
    );
  }
});

ReactDOM.render(<Temperature />, document.getElementById('sensors'));
