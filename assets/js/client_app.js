var app = angular.module('yellowclient', ['ngRoute', 'ngCookies']);
app.config(['$routeProvider', function($routeProvider){
	$routeProvider.when('/', {
		templateUrl: '/home',
    controller: 'HomeCtrl'
	}).when('/category/:id', {
		templateUrl:'/category',
		controller:'ListingCtrl'
	}).when('/search', {
		templateUrl:'/result'
	}).when('/listing/:cat/:id', {
		templateUrl: '/listing',
		controller:'PlusListingCtrl'
	});
}]);

app.controller('HomeCtrl', ['$scope', '$http', function($scope, $http){
$scope.result = {};
$http.get('/api/getcat').success(function(data, status){
  console.log(data);
  $scope.result = data;
});
}]);

app.controller('ListingCtrl', function($scope, $http,  $routeParams){
$scope.result = {};
$scope.category = {};
$http.get('/api/getcatList?q='+$routeParams.id).success(function(data, status){
	console.log($routeParams.id);
	console.log(data);
  $scope.result = data;
});
$http.get('/api/getsingle?q='+$routeParams.id).success(function(data,status){
	$scope.category = data;
});
});

app.controller('PlusListingCtrl', function($scope, $http,  $routeParams){
$scope.result = {};
$scope.category = {};
$http.get('/api/getsinglelist?q='+$routeParams.id).success(function(data, status){
  $scope.result = data;
});
$http.get('/api/getsingle?q='+$routeParams.cat).success(function(data,status){
	$scope.category = data;
});
});
