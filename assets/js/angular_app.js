var app = angular.module('yellowpages', ['ngRoute', 'ngCookies', 'ui-notification', 'ngMaterial']);
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
	}).when('/newad', {
		controller:'adCtrl',
		templateUrl:'/newad'
	})
	.when('/', {
		controller: 'unviewCtrl',
		templateUrl:'/viewlistingtemp'
	});
}]);
app.run(run);

run.$inject = ['$window','$rootScope', '$location', '$cookieStore', '$http'];
function run ($window, $rootScope, $location, $cookieStore, $http){
	$rootScope.globals = $cookieStore.get('globals') || {};
	if($rootScope.globals.currentUse){
		$http.defaults.headers.common['Authorization'] = 'Basic' + $rootScope.globals.currentUse.authdata;
	}
	$rootScope.$on('$locationChangeStart', function(event, next, current){
		var restrictedPage = $.inArray($location.path(), ['/login']) === -1;
		var loggedin = $rootScope.globals.currentUse;
		if(!loggedin){
			console.log("cam");
			/*var landingUrl = "http://localhost:8080/admin"; //URL complete
			window.location.href = landingUrl;*/
			$window.location = '/login';
		}
	});
}
