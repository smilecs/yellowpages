angular.module('yellowpages').controller('unviewCtrl', ['$scope', '$http', 'Notification', function($scope, $http, Notification){
$scope.result = {};
  $http.get('/api/unapproved').success(function(data, status){
    $scope.result = data;
  });
  $scope.approve = function(pos, data){
    console.log("clicked");
    $scope.result = $scope.result.splice(pos, 1);
    $http.post('/api/approve?q='+data).success(function(data,status){
    //$scope.result = data;
      Notification({message: 'Listing Approved', title: 'Listing Management'});
    /*$http.get('/api/unapproved').success(function(data, status){
      $scope.result = data;
      console.log("gotten");
    });*/
    });
  };
}]);
