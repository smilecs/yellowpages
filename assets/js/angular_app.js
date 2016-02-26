var app = angular.module('yellowpages', ['ngRoute', 'ngCookies']);
app.config(['$routeProvider', function($routeProvider){
	$routeProvider.when('/viewlisting', {
		controller: 'unviewCtrl',
		templateUrl:'/viewlistingtemp'
	}).when('/addlisting', {
		controller:'AddCtlr',
		templateUrl:'/addlistingtemp'
	}).when('/unapprovedView', {
		controller: 'unviewCtrl',
		templateUrl:'viewlistingtemp'
	}).when('/addcat', {
		controller:'catCtrl',
		templateUrl:'/addcattemp'
	})
	.when('/', {
		controller: 'unviewCtrl',
		templateUrl:'/viewlistingtemp'
	});
}]);
