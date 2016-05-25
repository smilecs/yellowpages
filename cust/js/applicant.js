angular.module('yellow').controller('UpdateCtrl',['$scope', '$http', '$cookieStore', function($scope, $http, $cookieStore){
$scope.employment = [];
$scope.education = [];
$scope.dat = {};
$scope.data = {};
$scope.courses = {};
$scope.finish = function(data){
$http.post("/api/update", data).success(function(data, status){

});
};
console.log($cookieStore.get('globals').currentUse.id);

$scope.emp = function(data){
  $scope.dat = {};
  var user = $cookieStore.get('globals');
  $scope.education.push(data);
  $http.post("/api/education?id="+$cookieStore.get('globals').currentUse.id, data).success(function(res){
    console.log(res);
  });
};

$scope.emp1 = function(data){
  $scope.employment.push(data);
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
$http.get('/api/course').success(function(data){
console.log(data);
  $scope.courses = data;
});

}]);
