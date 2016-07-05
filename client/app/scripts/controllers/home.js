'use strict';

angular.module('yapp')
  .controller('HomeCtrl', function($scope, $state,$stateParams,Users,$log,Polls,Auth) {


      activate();
      
      function activate() {
          var model = {
              userId:$stateParams.id,
              user:null,
              polls:[]
          }

          $scope.model = model;
          $scope.edit = edit;
          $scope.disable = disable;
          $scope.logout = logout;

          Users.getUserById(model.userId)
              .then(function (user) {
                  Polls.getPollsByUserId(user.ID)
                      .then(function (polls) {
                          model.polls = polls;
                      },function(error){
                          $log.error(error);
                      })
                model.user = user;
              },function(error){
                  $log.error(error);
              })
      }

      function edit(id) {
        $state.go('poll-edit',{id:id});
      }

      function disable(id) {
        //TODO: Delete Polls
      }
      
      function logout() {
          Auth.logout();
          $state.go('login');
      }
  });
