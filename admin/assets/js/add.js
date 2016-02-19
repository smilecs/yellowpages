angular.module('yellowpages').controller('AddCtlr', ['$scope', '$http', function($scope, $http){
$scope.data = {};
$scope.add = function(data){
  $http.post('/api/addlisting', data).success(function(data, status){
    $scope.data = {};
    if(status === 200){
      $location.path('/');
    }
  });
};
}]);
