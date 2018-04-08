(function() {
    'use strict';

    angular
      .module('app')
      .controller('HardWareController', HardWareController);

    HardWareController.$inject = ['$rootScope', '$http'];

    function HardWareController($rootScope,$http) {
      // var vm = this;

      // vm.user = null;
      // vm.register = register;
      // //  vm.allUsers = [];
      // // vm.deleteUser = deleteUser;
      // vm.auth = AuthenticationService.IsAuthenticated();
      // initController();

      // function initController() {
      //     delete $rootScope.flash;
      //     if (vm.auth) {
      //         loadCurrentUser();
      //     }

          

      //     //loadAllUsers();
      // }
    }

})();

