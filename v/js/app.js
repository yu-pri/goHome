
var N = function(){};

N.ws = new WebSocket("ws://192.168.1.20:1234/echo");
N.DataHandler = new Object();
N.TempDataHandler = new Object();

N.CommandOnPumpr1 = "onPumpr1";
N.CommandOffPumpr1 = "offPumpr1";

N.DataHandler.handle = function (msg) {
  var o = JSON.parse(msg);
  switch (o.Key) {
    case "temperature":
    N.TempDataHandler.callback(o);
    break;
  }
};

N.ws.onmessage = function(e) {
    console.log("received:" + event.data);
    N.DataHandler.handle(event.data);
};




jQuery(function($) {
  var $bodyEl = $('body'),
      $sidedrawerEl = $('#sidedrawer');


  function showSidedrawer() {
    // show overlay
    var options = {
      onclose: function() {
        $sidedrawerEl
          .removeClass('active')
          .appendTo(document.body);
      }
    };

    var $overlayEl = $(mui.overlay('on', options));

    // show element
    $sidedrawerEl.appendTo($overlayEl);
    setTimeout(function() {
      $sidedrawerEl.addClass('active');
    }, 20);
  }


  function hideSidedrawer() {
    $bodyEl.toggleClass('hide-sidedrawer');
  }


  $('.js-show-sidedrawer').on('click', showSidedrawer);
  $('.js-hide-sidedrawer').on('click', hideSidedrawer);

  var $titleEls = $('strong', $sidedrawerEl);

  $titleEls
    .next()
    .hide();

  $titleEls.on('click', function() {
    $(this).next().slideToggle(200);
  });
});


var Temperature = React.createClass({
  getInitialState: function() {

    return {temp1: this.props.temp1, temp2: this.props.temp2, temp3: this.props.temp3, temp4: this.props.temp4};
  },

  componentWillMount() {
    N.TempDataHandler.callback = (data) => {
      switch (data.Name) {
        case "temp1":
          this.setState({
            temp1: data.Val
          });
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
