angular.module('yellowpages').controller('adCtrl', ['$scope', '$http', function($scope, $http){
$scope.result = {};
$scope.data = {};
/*$http.get('/api/getcat').success(function(data, status){
  $scope.result = data;
});
*/
$scope.add = function(data){
data.image = $scope.f;
  $http.post('/api/newAd', data).success(function(data, status){
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
