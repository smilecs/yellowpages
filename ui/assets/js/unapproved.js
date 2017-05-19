angular.module('yellowpages').controller('unviewCtrl', ['$scope', '$http', 'Notification', function($scope, $http, Notification){
$scope.result = {};
$scope.show = "hide";
$scope.but = false;
  $http.get('/api/unapproved').success(function(data, status){
    console.info(data)
    $scope.result = data.Data;
  });
  $scope.approve = function(pos, data){
    $scope.but = true;
    $scope.show = "show";
    console.log("clicked");
    console.log(data);
    //$scope.result = $scope.result.splice(pos, 1);

    //console.log($scope.result);
    $http.post('/api/approve?q='+data).success(function(data,status){
    //$scope.result = data;
    //console.log(data);
    $scope.show = "hide";
    $scope.but = false;
    console.log($scope.result.splice(pos, 1));
      Notification({message: 'Listing Approved', title: 'Listing Management'});
    /*$http.get('/api/unapproved').success(function(data, status){
      $scope.result = data;
      console.log("gotten");
    });*/
    });
  };
}]);
