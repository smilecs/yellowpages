angular.module('yellowpages').controller('unviewCtrl', ['$scope', '$http', function($scope, $http){
$scope.result = {};
  $http.get('/api/unapproved').success(function(data, status){
    $scope.result = data;
  });
  $scope.approve = function(data){
    $http.post('/api/approve?q='+data).success(function(data,status){
    $scope.result = data;
    });
  };
}]);
