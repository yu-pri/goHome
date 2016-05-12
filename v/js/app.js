

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
        <div className="mui-container-fluid, flexrow" >
          <span className="item">Kitchen: {this.state.temp1}</span>
          <span className="item">Entry room: {this.state.temp3}</span>
          <span className="item">H Rev: {this.state.temp2} </span>
        </div>

        <div className="mui-container-fluid, flexrow" >
          <span className="item">Heater: {this.state.temp4} </span>
          <span className="item"> Heater: {this.state.heater} </span>
          <span className="item"> HeaterPump: {this.state.heaterPump} </span>
        </div>


      </div>
    );
  }
});

ReactDOM.render(<Temperature />, document.getElementById('sensors'));
