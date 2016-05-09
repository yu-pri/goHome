

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
  },

  componentWillMount() {
    window.N.TempDataHandler.callback = (data) => {
      this.set(data);
      /*
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
      */
    };
  },

  render: function() {

    return (
      <div>
        <div className="mui-container-fluid" >
          <span className="temp">Kitchen:</span> {this.state.temp1}
        </div>

        <div className="mui-container-fluid" >
          <span className="temp">Entry room:</span> {this.state.temp3}
        </div>

        <div className="mui-container-fluid" >
          <span className="temp">H Rev: </span>{this.state.temp2}
        </div>


        <div className="mui-container-fluid" >
          <span className="temp">Heater: </span> {this.state.temp4}
        </div>


      </div>
    );
  }
});

ReactDOM.render(<Temperature />, document.getElementById('sensors'));
