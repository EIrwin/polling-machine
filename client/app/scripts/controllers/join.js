'use strict';

/**
 * @ngdoc function
 * @name yapp.controller:MainCtrl
 * @description
 * # MainCtrl
 * Controller of yapp
 */
angular.module('yapp')
  .controller('JoinCtrl', function($scope, $location,$stateParams,Meetings,Users,$log,$state,$localstorage) {

    var model = {
      salary:null,
      hourlyRate:null,
      meetingId:$stateParams.meetingId
    }

    $scope.model = model;

    $scope.joinMeeting = function(salary,hourlyRate){
      Users.createUser(salary,hourlyRate)
        .then(function(user){
          if(user != null){
            $localstorage.set('userId',user.id);
            Meetings.joinMeeting(user.id,model.meetingId)
              .then(function(meeting){
                $state.go('meeting',{meetingId:meeting.id})
              },function(error){
                $log.error(error);
              })
          }
        },function(error){
          $log.error(error);
        })
    }

  });
