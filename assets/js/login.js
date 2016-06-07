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
