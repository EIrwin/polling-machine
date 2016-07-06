'use strict';

angular.module('yapp')
  .controller('HomeCtrl', function($scope, $state,$stateParams,Users,$log,Polls,Auth,$location,ENV) {

      activate();
      
      function activate() {
          var model = {
              userId:$stateParams.id,
              user:null,
              polls:[]
          }

          $scope.model = model;
          $scope.logout = logout;
          $scope.getShareUrl = getShareUrl;

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
      
      function logout() {
          Auth.logout();
          $state.go('login');
      }



      function getShareUrl(pollId) {
          var host = $location.host();
          var port = $location.port();
          var url = "http://" + host + (ENV.name == "local" ? port : "") + "/#/share/" + pollId;
          return url;
      }
  });
