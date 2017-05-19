var app = angular.module('yellowclient', ['ngRoute']);
app.config(['$routeProvider', '$locationProvider', function($routeProvider, $locationProvider){
	$routeProvider.when('/', {
		templateUrl: '/home',
    controller: 'HomeCtrl'
	}).when('/category/:id', {
		templateUrl:'/category',
		controller:'ListingCtrl'
	}).when('/plus/:id', {
		templateUrl:'/category',
		controller:'FalseCtrl'
	}).when('/ads/:id', {
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
	}).when('/result/:query/:index', {
		templateUrl: '/result',
		controller: 'PageCtrl'
	}).when('/advert/:id', {
		templateUrl:'/category',
		controller:'AdvCtrl'
	});
	$locationProvider.html5Mode(true);
}]);

app.controller('HomeCtrl', ['$scope', '$http','$location', '$anchorScroll', function($scope, $http, $location, $anchorScroll){
$scope.result = {};
$scope.show = "show";
$scope.slides = [];
$http.get('/api/getcat').success(function(data, status){
	$scope.show = "hide";
  $scope.result = data;
});

$http.get('/slider?page=1').success(function(data,status){
	$scope.category = data.Data;
	$scope.pages = data;
	$scope.slides = data.Data;
	$scope.myInterval = 3000;

});


$scope.send = function(data){
	$location.path('/result/'+data);
};

$scope.scrollTo = function (id) {
  $anchorScroll(id);
}

}]);

app.controller('ListingCtrl', function($scope, $http, $location, $routeParams){
$scope.result = {};
$scope.category = {};
$scope.pages = {};
$scope.show = "show";
$scope.newerScope = [];
$http.get('/api/getsingle?q='+$routeParams.id).success(function(data, status){
  $scope.result = data;
});
$http.get('/api/newview?page=1&q='+$routeParams.id).success(function(data,status){
	$scope.category = data.Data;
	$scope.pages = data;
	$scope.show = "hide";
	$scope.newScope = data.Pag.Pages;
	for(var i =0; i < $scope.newScope.length; i++){
		var tmp = {"data": i+1};
		$scope.newerScope.push(tmp);
	}



});
$scope.send = function(data){
	$location.path('/result/'+data);
};

$scope.sends = function(data){
$scope.pages = {};
$scope.newScope = {};
$scope.newerScope = [];

	$http.get('/api/newview?page='+data+"&q="+$routeParams.id).success(function(data, status){
		$scope.category = data.Data;
		$scope.pages = data;
		$scope.newScope = data.Pag.Pages;
		for(var i =0; i < $scope.newScope.length; i++){
			var tmp = {"data": i+1};
			$scope.newerScope.push(tmp);
		}
//$scope.$apply();
	});
};


});

app.controller('FalseCtrl', function($scope, $http, $location, $routeParams){
$scope.result = {};
$scope.category = {};
$scope.pages = {};
$scope.newerScope = [];
$scope.show = "show";
$http.get('/api/getsingle?q='+$routeParams.id).success(function(data, status){
  $scope.result = data;
});
$http.get('/api/falseview?page=1').success(function(data,status){
	$scope.category = data.Data;
	$scope.pages = data;
	$scope.show = "hide";
	$scope.newScope = data.Pag.Pages;
	for(var i =0; i < $scope.newScope.length; i++){
		var tmp = {"data": i+1};
		$scope.newerScope.push(tmp);
	}


});

$scope.send = function(data){
	$location.path('/result/'+data);
};

$scope.sends = function(data){
$scope.pages = {};
$scope.newScope = {};
$scope.newerScope = [];

	$http.get('/api/falseview?page='+data).success(function(data, status){
		$scope.category = data.Data;
		$scope.pages = data;
		$scope.newScope = data.Pag.Pages;
		for(var i =0; i < $scope.newScope.length; i++){
			var tmp = {"data": i+1};
			$scope.newerScope.push(tmp);
		}
$scope.$apply();
	});
};

});

app.controller('AdvCtrl', function($scope, $http, $location, $routeParams){
$scope.result = {};
$scope.category = {};
$scope.pages = {};
$scope.newerScope = [];
$http.get('/api/getsingle?q='+$routeParams.id).success(function(data, status){
  $scope.result = data;
});
$http.get('/advert?page=1').success(function(data,status){
	$scope.category = data.Data;
	$scope.pages = data;
	$scope.newScope = data.Pag.Pages;
	for(var i =0; i < $scope.newScope.length; i++){
		var tmp = {"data": i+1};
		$scope.newerScope.push(tmp);
	}


});

$scope.send = function(data){
	$location.path('/result/'+data);
};

$scope.sends = function(data){
$scope.pages = {};
$scope.newScope = {};
$scope.newerScope = [];
	$http.get('/api/falseview?page='+data).success(function(data, status){
	$scope.category = data.Data;
	$scope.pages = data;
	$scope.newScope = data.Pag.Pages;
	for(var i =0; i < $scope.newScope.length; i++){
		var tmp = {"data": i+1};
		$scope.newerScope.push(tmp);
	}
$scope.$apply();

	});
};

});


app.controller('PlusListingCtrl', function($scope, $http,  $location, $routeParams){
$scope.result = {};
$scope.category = {};
$http.get('/api/getsinglelist?q='+$routeParams.id).success(function(data, status){
  $scope.result = data;
});
$http.get('/api/getsingle?q='+$routeParams.cat).success(function(data,status){
	$scope.category = data;
});
$scope.send = function(data){
$location.path('/result/'+data);
};

});

app.controller('SearchCtrl', function($scope, $http, $routeParams, $location){
	$scope.category = {};
	$scope.pages = {};
	$scope.newerScope = [];
	$scope.show = "show";

	$scope.send = function(data){
$location.path('/result/'+data);
	};

	$http.post('/api/result?page=1&q='+$routeParams.query).success(function(data, status){
		$scope.category = data.Data;
		$scope.show = "hide";
		$scope.pages = data;
		$scope.newScope = data.Pag.Pages;
		for(var i =0; i < $scope.newScope.length; i++){
			var tmp = {"data": i+1};
			$scope.newerScope.push(tmp);
		}
	//	$scope.$apply();

	});
});

app.controller('Searchbtn', function($scope){
	$scope.send = function(data){
		$location.path('/result/data');
	};
});

app.controller('PageCtrl', function($scope, $http, $routeParams){
	$http.post('/api/result?page='+$routeParams.index+'&q='+$routeParams.query).success(function(data, status){
		$scope.category = data.Data;
		$scope.pages = data;
	});
});
