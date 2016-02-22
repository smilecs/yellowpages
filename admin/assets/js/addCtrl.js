angular.module('yellowpages').controller('catCtrl', ['$scope', '$http', function($scope, $http){
$scope.result = {};
$scope.datas = {};
$http.get('/api/getcat').success(function(data, status){
  $scope.result = data;
});
$scope.add = function(data){
  $http.post('/api/addcat', data).success(function(data, status){
    $scope.result = data;
    console.log(data);
    if(status === 200){
      //$location.path('/');
    }
  });
};
}]);
