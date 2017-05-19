angular.module('yellowpages').controller('adCtrl', ['$scope', '$http', 'Notification', function($scope, $http, Notification){
$scope.result = {};
$scope.data = {};
$scope.show = "show";

/*$http.get('/api/getcat').success(function(data, status){
  $scope.result = data;
});
*/


$http.get('/api/adverts?p=1').success(function(data,status){
	console.log(data)
	$scope.adverts= data.Posts;
  $scope.show = "hide";
  
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


$scope.add = function(data){
data.image = $scope.f;
  $http.post('/api/newAd', data).success(function(data, status){
    Notification({message: 'Success', title: 'Advert Management'});
    $scope.f = {};
    $scope.result = data;
    console.log(data);
    if(status === 200){
      //$location.path('/');
    }
  });
};
$scope.newfile1 = function(file){


  var reader = new FileReader();
  reader.onload = function(u){
        //$scope.files.push(u.target.result);
        $scope.$apply(function($scope) {
          $scope.f = u.target.result;
          //console.log(u.target.result);
        });
  };
  reader.readAsDataURL(file);

};

}]);
