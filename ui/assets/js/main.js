var app = angular.module('yellowpages', ['ngRoute', 'ngCookies']).config(config);
config.$inject = ['$routeProvider', '$locationProvider'];
function config($routeProvider, $locationProvider){
	$routeProvider.when('/', {
		controller: 'UpdateCtrl',
		 templateUrl:'/login'
	});
}
