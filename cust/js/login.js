angular.module('yellow').controller('LoginCtrl',LoginCtrl);
LoginCtrl.$inject = ['$scope', '$location', 'AuthenticationService'];
function LoginCtrl($scope, $location, AuthenticationService){
  console.log("called");
  $scope.vm = {};
  //var vm = this;
  //vm.login = login;

  $scope.hide = "true";
  (function initController(){
    //AuthenticationService.ClearCredentials();
  })();
$scope.login = function login(){
    console.log("called");
    vm.dataLoading = true;
    AuthenticationService.Login(vm.Username, vm.Password, function(response){
      if(response.id){
        console.log("true");
        AuthenticationService.SetCredent(vm.Username, vm.Password, response.id);
          $scope.hide = "false";
        $location.path('/');
      } else{
        console.log("false");
        vm.dataLoading = false;
      }
    });
  }
}
