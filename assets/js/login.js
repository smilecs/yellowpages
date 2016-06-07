var app = angular.module('yellowpages', ['ngCookies']);

app.controller('LoginCtrl',LoginCtrl);
LoginCtrl.$inject = ['$scope', '$location', 'AuthenticationService'];
function LoginCtrl($scope, $location, AuthenticationService){
  console.log("called");

  var vm = this;
  vm.login = login;
  $scope.hide = "true";
  (function initController(){
    //AuthenticationService.ClearCredentials();
  })();
function login(){
    console.log("called");
    vm.dataLoading = true;
   AuthenticationService.Login(vm.Username, vm.Password, function(response){
      if(response.id){
        console.log("true");
        AuthenticationService.SetCredent(vm.Username, vm.Password, response.id);
          $scope.hide = "false";
        $window.location = '/admin';
      } else{
        console.log("false");
        vm.dataLoading = false;
      }
    });
  }
}

app.controller('ClientListing', function($scope, $http){
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
      $scope.image = '';
      //$location.path('/');

    }, function(){
        Notification.error("Error Adding Data");
    });
  };
});
