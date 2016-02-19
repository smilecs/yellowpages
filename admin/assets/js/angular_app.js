var app = angular.module('yellowpages', ['ngRoute', 'ngCookies']);
app.config(['$routeProvider', function($routeProvider){
	$routeProvider.when('/', {
		controller:'MainCtlr',
        template:'<div>TestCase</div>'
	}).when('/addlisting', {
		controller:'AddCtlr',
		templateUrl:'/addlistingtemp'
	});
}]);
