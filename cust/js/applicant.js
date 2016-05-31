angular.module('yellow').controller('UpdateCtrl',['$scope', '$http', '$cookieStore', function($scope, $http, $cookieStore){
$scope.result = {};

console.log($cookieStore.get('globals').currentUse.id);

  
  var user = $cookieStore.get('globals');
  $http.get("/api/plus?id="+$cookieStore.get('globals').currentUse.id).success(function(res){
    $scope.result = res;
    console.log(res);
  });




}]);
