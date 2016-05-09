

var Temperature = React.createClass({
  getInitialState: function() {

    return {temp1: this.props.temp1, temp2: this.props.temp2, temp3: this.props.temp3, temp4: this.props.temp4};
  },

  componentWillMount() {
    window.N.TempDataHandler.callback = (data) => {
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
      /*
      if (data.TempInside) {
        this.setState({
          temp1: data.TempInside
        });
      }


      /*
      switch (data.Name) {
        case "temp1":
          this.setState({
            temp1: data.Val
          });
          //window.updateChart(data.Val);
          break;

        case "temp2":
          this.setState({
            temp2: data.Val
          });
          break;

        case "temp3":
          this.setState({
            temp3: data.Val
          });
          break;
        case "temp4":
          this.setState({
            temp4: data.Val
          });
          break;
      }
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
