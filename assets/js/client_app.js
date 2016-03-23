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
	}).when('/result/:query', {
		templateUrl: '/result',
		controller:'SearchCtrl'
	});
}]);

app.controller('HomeCtrl', ['$scope', '$http', function($scope, $http){
$scope.result = {};
$http.get('/api/getcat').success(function(data, status){
  console.log(data);
  $scope.result = data;
});
$scope.send = function(data){
	$location.path('/result/data');
};

}]);

app.controller('ListingCtrl', function($scope, $http,  $routeParams){
$scope.result = {};
$scope.category = {};
$http.get('/api/getsingle?q='+$routeParams.id).success(function(data, status){
  $scope.result = data;
});
$http.get('/api/newview?q='+$routeParams.id).success(function(data,status){
	$scope.category = data;
	console.log(data);
});
$scope.send = function(data){
	$location.path('/result/data');
};

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
$scope.send = function(data){
	$location.path('/result/data');
};

});

app.controller('SearchCtrl', function($scope, $http, $routeParams){
	$scope.category = {};
	$scope.pages = {};

	$scope.send = function(data){
		$location.path('/result/data');
	};

	$http.get('/api/result?page=1&q='+$routeParams.query).success(function(data, status){
		console.log(data);
		$scope.category = data[0];
		$scope.pages = data[1];
	});
});

app.controller('Searchbtn', function($scope){
	$scope.send = function(data){
		$location.path('/result/data');
	};
});
