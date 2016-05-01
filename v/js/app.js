

var Temperature = React.createClass({
  getInitialState: function() {

    return {temp1: this.props.temp1, temp2: this.props.temp2, temp3: this.props.temp3, temp4: this.props.temp4};
  },

  componentWillMount() {
    window.N.TempDataHandler.callback = (data) => {
      switch (data.Name) {
        case "temp1":
          this.setState({
            temp1: data.Val
          });
          window.updateChart(data.Val);
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
    };
  },

  render: function() {

    return (
      <div>
        <div className="mui-container-fluid" >
          <span className="temp"> t (int):</span> {this.state.temp1}
        </div>

        <div className="mui-container-fluid" >
          <span className="temp">t (out): </span>{this.state.temp2}
        </div>

        <div className="mui-container-fluid" >
          <span className="temp"> Heater (top):</span> {this.state.temp3}
        </div>

        <div className="mui-container-fluid" >
          <span className="temp">Heater (rev): </span> {this.state.temp4}
        </div>


      </div>
    );
  }
});

ReactDOM.render(<Temperature />, document.getElementById('sensors'));
