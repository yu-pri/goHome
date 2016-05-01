
var chart;
var data;
var options;

google.charts.load('current', {'packages':['corechart']});
      google.charts.setOnLoadCallback(drawChart);

      function drawChart() {
        data = new google.visualization.DataTable();

        data.addColumn('datetime', 'XAxis');
        data.addColumn('number', 'Internal temp');

        options = {
          title: 'Temperature chart',
          curveType: 'function',
          legend: { position: 'bottom' }
        };

        chart = new google.visualization.LineChart(document.getElementById('curve_chart'));

        fetch('/control/hdata')
          .then(function(response) {
            if (response.status == 200) {
              // Examine the text in the response

              response.json().then(function(dt) {
                console.log(dt);
                for (var i=0; i < dt.length; i++) {
                  var r = dt[i];
                  data.addRows(dt.Timestamp*1000, dt.TempInside);
                }
                chart.draw(data, options);
              });

            } else {
              alert(response.statusText)
            }
          })
          /*
          .then(function(resp) {
            console.log(resp);
            var o = JSON.parse(resp);
            for (var i=0; i < o.length; i++) {
              var r = o[i];
              data.addRows(r.Timestamp*1000, r.TempInside);
            }
            chart.draw(data, options);
          })
          */

        //chart.draw(data, options);
      };

      function updateChart() {
            //var dt = new google.visualization.DataTable();
            //d.addColumn('datetime', 'XAxis');
            //dt.addColumn('number', 'Internal temp');
            data.addRows([
              [new Date(1462035153*1000), 27.5],
              [new Date(1462036153*1000), 22.5],
              [new Date(1462047153*1000), 21.5]
            ]);
            chart.draw(data, options);
      }
