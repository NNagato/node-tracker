(function () {
	angular
    .module('app', ['ngRoute', 'ngCookies', 'ui.bootstrap', 'ngAnimate', 'ngSanitize'])
    .config(config)
    // .run(run)
    .directive('myChart', ['$parse', '$timeout', function($parse, $timeout) {
      return {
        restrict: 'A',
        link: function(scope, element, attrs){
          $timeout(function(){
            var handle = $parse(attrs.myChart);
            handle(scope);
          })
        }
      };
    }]);

    config.$inject = ['$routeProvider', '$locationProvider','$qProvider']
    function config($routeProvider, $locationProvider, $qProvider) {
    	$qProvider.errorOnUnhandledRejections(false);
      $routeProvider
      .when('/', {
        controller:   'HomeController',
        templateUrl:  'app-pages/home/home.view.html',
        controllerAs: 'vm' 
      })

      .when('/hardware', {
        controller:   'HardWareController',
        templateUrl:  'app-pages/hardware/hardware.view.html',
        controllerAs: 'vm'
      })

      .otherwise({ redirectTo: '/' });
      console.log("this is config")
    }
})()
