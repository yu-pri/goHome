
var ButtonSwitchChart = React.createClass({
  getInitialState: function() {
    this.props.val = "1 Hour";
    //this.setState({message: "1 Hour"});
    return {message: this.props.val};
  },

  handleSubmit: function() {
      if (this.props.val === "1 Hour") {
        this.props.val = "1 Day";
        window.ChartTimeLimit = 60*24*60;
      } else if (this.props.val === "1 Day") {
        this.props.val = "1 Month";
        window.ChartTimeLimit = 60*24*30*60;
      } else {
        this.props.val = "1 Hour";
        window.ChartTimeLimit = 60*60;
      }
      this.setState({message: this.props.val});
      window.setChartTimeLimit(this.props.val);
      window.loadChart();
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

ReactDOM.render(<ButtonSwitchChart msg="1 Hour" />, document.getElementById('switchChart'));
