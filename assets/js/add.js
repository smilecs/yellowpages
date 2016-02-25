angular.module('yellowpages').controller('AddCtlr', ['$scope', '$http', function($scope, $http){
$scope.data = {};
$scope.cats = {};
$scope.show = [];
$http.get('/api/getcat').success(function(data, status){
  $scope.cats = data;
});
$scope.add = function(data){
  $http.post('/api/addlisting', data).success(function(data, status){
    $scope.data = {};
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
}]);
