angular.module('yellowpages').controller('AddCtlr', ['$scope', '$http', function($scope, $http){
$scope.data = {};
$scope.cats = {};
$scope.show = [];
$scope.files = [];
$http.get('/api/getcat').success(function(data, status){
  $scope.cats = data;
});
$scope.add = function(data){
  data.images = $scope.files;
  data.image = $scope.f;
  console.log(data);

  $http.post('/api/addlisting', data).success(function(data, status){
    $scope.data = {};
    $scope.files = [];
    $scope.image = '';
    if(status === 200){
      $location.path('/');
    }
  });
};
$scope.change = function(data){
  console.log("start");
  if(data.Plus === true){
    $scope.show = true;
  }
};

$scope.newfile = function(file){

  console.log(file);
  var reader = new FileReader();
  reader.onload = function(u){
        //$scope.files.push(u.target.result);
        $scope.$apply(function($scope) {
          $scope.files.push(u.target.result);
          //console.log(u.target.result);
        });
  }
  reader.readAsDataURL(file);

}

$scope.newfile1 = function(file){


  var reader = new FileReader();
  reader.onload = function(u){
        //$scope.files.push(u.target.result);
        $scope.$apply(function($scope) {
          $scope.f = u.target.result;
          //console.log(u.target.result);
        });
  }
  reader.readAsDataURL(file);

}
}]);
