(function () {
	'use strict';

	angular
	.module('app')
	.factory('ServerService', ServerService);

	ServerService.$inject = ['$http'];
	function ServerService($http) {
		var service = {};
		
		service.GetDataTimeResponse = GetDataTimeResponse;

		return service;

		function GetDataTimeResponse() {
			return $http.get('/api/data').then(handleSuccess, handleError('Cannot get time-response data'));
		}

    // function SubmitMessage(form) {
    //     return $http.post('/api/blockchain/submitmessage',form).then(handleSuccess, handleError('Error submit message'));
    // }
    
    // function ProcessSubmit(form) {
    //     return $http.post('/api/blockchain/confirmsubmit',form).then(handleSuccess, handleError('Error submit message'));
    // }

    // function CheckMessage(form) {
    //     return $http.post('/api/blockchain/checkmessage',form).then(handleSuccess, handleError('Error check message'));
    // }

    // function QueryLengthTransaction() {
    //     return $http.post('/api/blockchain/lengthtransaction').then(handleSuccess, handleError('Error query length trasactions'));   
    // }

    // function QueryTransaction(startIndex,endIndex) {
    //     return $http.post('/api/blockchain/querytransaction',{"start":startIndex,"end":endIndex}).then(handleSuccess, handleError('Error query trasactions'));   
    // }

    // function QueryLengthTransactionAdmin() {
    //     return $http.post('/api/blockchain/admin/lengthtransaction').then(handleSuccess, handleError('Error query length trasactions'));   
    // }

    // function QueryTransactionAdmin(startIndex,endIndex) {
    //     return $http.post('/api/blockchain/admin/querytransaction',{"start":startIndex,"end":endIndex}).then(handleSuccess, handleError('Error query trasactions'));   
    // }

    // function LeaveDeposit(form) {
    //     return $http.post('/api/blockchain/leavedeposit',form).then(handleSuccess, handleError('Error deposit, please retry!'));   
    // }

    // // function ConfirmDeposit() {
    // //     return $http.post('/api/blockchain/confirmposit').then(handleSuccess, handleError('Error deposit, please retry!'));   
    // // }

    // function Gettimestamp(proof) {
    //     return $http.post('/api/blockchain/gettimestamp',proof).then(handleSuccess, handleError('Error get timestamp!'));   
    // }

    // function SatisticInfo() {
    //     return $http.post('/api/blockchain/statisticInfo').then(handleSuccess, handleError('Error get statistic informtaion!'));   
    // }
    // private functions

    function handleSuccess(res) {
    	return res.data;
    }

    function handleError(error) {
    	return function () {
    		return { success: false, message: error };
    	};
    }
  }

})();
