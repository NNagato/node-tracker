(function() {
  'use strict';

  angular
    .module('app')
    .controller('HomeController', HomeController);

  HomeController.$inject = ['$rootScope', '$scope', '$http', 'ServerService'];

  function HomeController($rootScope, $scope, $http, ServerService) {
    var vm = this;
    $scope.addChart = addChart;

    function addChart(id, rpc, data) {
      var trueId = "chart" + id;
      var trueData = makeChartData(data);
      // console.log("data make chart: ", trueId, rpc, trueData);
      var firstIndex;
      var lastIndex = trueData.length - 1;
      if (trueData.length > 50) {
        firstIndex = trueData.length - 50;
      } else {
        firstIndex = 0;
      }

      vm.chart = AmCharts.makeChart(trueId, {
        "type": "serial",
        "theme": "light",
        "marginTop": 20,
        "marginRight": 80,
        "marginLeft": 10,
        "marginBottom": 50,
        "dataProvider": trueData,
        "valueAxes": [{
          "axisAlpha": 0,
          "position": "left",
          "title": "time response"
        }],
        "graphs": [{
          "id":"g1",
          "balloonText": "[[category]]<br><b><span style='font-size:13px;'>[[value]] ms</span></b>",
          "bullet": "round",
          "bulletSize": 6,
          "lineColor": "#d1655d",
          "lineThickness": 2,
          "negativeLineColor": "#637bb6",
          "type": "smoothedLine",
          "valueField": "delay"
        }],
        "chartScrollbar": {
          "graph":"g1",
          "gridAlpha":0,
          "color":"#888888",
          "scrollbarHeight":55,
          "backgroundAlpha":0,
          "selectedBackgroundAlpha":0.1,
          "selectedBackgroundColor":"#888888",
          "graphFillAlpha":0,
          "autoGridCount":true,
          "selectedGraphFillAlpha":0,
          "graphLineAlpha":0.2,
          "graphLineColor":"#c2c2c2",
          "selectedGraphLineColor":"#888888",
          "selectedGraphLineAlpha":1
        },
        "chartCursor": {
          "categoryBalloonDateFormat": "JJ:NN:SS, DD MMMM",
          "cursorAlpha": 0,
          "valueLineEnabled":true,
          "valueLineBalloonEnabled":true,
          "valueLineAlpha":0.5,
          "fullWidth":true
        },
        // "dataDateFormat": "DD hh mm",
        "categoryField": "time",
        "categoryAxis": {
          "minPeriod": "ss",
          "parseDates": true,
          "minorGridAlpha": 0.1,
          "minorGridEnabled": true
        },
        "export": {
          "enabled": true,
          "dateFormat": "YYYY-MM-DD HH:NN:SS"
        },
        listeners: [{
          event: "init",
          method: function(e) {
            e.chart.zoomToIndexes(firstIndex, lastIndex); //set default zoom
         }
        }]
      });
      vm.chart.addListener("rendered", zoomChart);
      if(vm.chart.zoomChart){
        vm.chart.zoomChart();
      }
    }
    function zoomChart(){
      chart.zoomToIndexes(Math.round(chart.dataProvider.length * 0.4), Math.round(chart.dataProvider.length * 0.55));
    }

    (function initController() {
      // console.log("run init home controller");
      ServerService.GetDataTimeResponse().then(function(response){
        if (response.success) {
          vm.timeResponseData = response.data;
          // console.log(vm.timeResponseData);
        }
      })
    })();

    function makeChartData(data) {
      // console.log(data)
      var chartData = [];
      for(var key in data){
        var keyNum = parseInt(key);
        var date = getDateFormat(keyNum);
        // console.log("parse string: ", keyNum.toString(), data[key].toString());
        chartData.push({
          time: date,
          delay: Math.round(parseFloat(data[key]) * 1000 * 10) / 10
        });
      }
      return chartData;
    }

    function getDateFormat(keyNum) {
      var date = new Date(keyNum * 1000);
      // console.log("date: ", date.toISOString());      
      return date.toString();
    }
  }
})();
