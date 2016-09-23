angular.module('yellowpages').controller('AddCtlr', ['$scope', '$http', 'Notification', function($scope, $http, Notification){
$scope.data = {};
$scope.cats = {};
$scope.show = [];
$scope.files = [];
$scope.show = "hide";
$http.get('/api/getcat').success(function(data, status){
  $scope.cats = data;
});
$scope.add = function(data){
  data.images = $scope.files;
  data.image = $scope.f;
  console.log(data);
  $scope.show = "show";

  $http.post('/api/addlisting', data).then(function(){
    $scope.data = {};
    Notification({message: 'Success', title: 'Listing Management'});
    $scope.show = "hide";
    $scope.files = [];
    $scope.f = {};
    $scope.image = '';
    //$location.path('/');

  }, function(){
      Notification.error("Error Adding Data");
  });
};
$scope.change = function(data){
  console.log("start");
  if(data.Plus === true){
    $scope.show = true;
  }
};

$scope.newfile = function(file){

  var reader = new FileReader();
  reader.onload = function(u){
        //$scope.files.push(u.target.result);
        $scope.$apply(function($scope) {
          $scope.files.push(u.target.result);
          //console.log(u.target.result);
        });
  };
  reader.readAsDataURL(file);

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
