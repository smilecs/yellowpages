var app = angular.module('yellow', ['ngRoute', 'ngCookies']).config(config).run(run);
config.$inject = ['$routeProvider', '$locationProvider'];
function config($routeProvider, $locationProvider){
	$routeProvider.when('/', {
		controller: 'UpdateCtrl',
		 templateUrl:'/cust/partials/main.html'
	}).when('/course', {
		controller: 'CourseCtrl',
		templateUrl:'/adm/course.html'
	}).when('/login', {
		controller: 'LoginCtrl',
		controllerAs: 'vm',
		templateUrl:'/client/assets/partials/login.html'
	}).when('/logout', {
		controller:'logoutCtrl'
	}).when('/profile', {
		templateUrl:'client/assets/partials/profile.html'
	});
}

run.$inject = ['$rootScope', '$location', '$cookieStore', '$http'];
function run ($rootScope, $location, $cookieStore, $http){
	$rootScope.globals = $cookieStore.get('globals') || {};
	if($rootScope.globals.currentUse){
		$http.defaults.headers.common['Authorization'] = 'Basic' + $rootScope.globals.currentUse.authdata;
	}
	$rootScope.$on('$locationChangeStart', function(event, next, current){
		var restrictedPage = $.inArray($location.path(), ['/login']) === -1;
		var loggedin = $rootScope.globals.currentUse;
		//if(restrictedPage && !loggedin){
		if(!loggedin){
			console.log("cam");
			/*var landingUrl = "http://localhost:8080/admin"; //URL complete
			window.location.href = landingUrl;*/
			//$location.path('/login');
		}
	});
}

app.controller('logoutCtrl',['$http', '$scope', '$location', '$cookieStore', 'AuthenticationService', function($http, $scope, $location, $cookieStore, AuthenticationService){
	//AuthenticationService.ClearCredentials();
	$cookieStore.remove('globals');
	$http.defaults.headers.common.Authorization = 'Basic';

}]);
