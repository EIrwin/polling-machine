'use strict';

/**
 * @ngdoc function
 * @name yapp.controller:MainCtrl
 * @description
 * # MainCtrl
 * Controller of yapp
 */
angular.module('yapp')
  .controller('CreateCtrl', function($scope,Meetings,$log,APIHelper,$state) {

    var model = {
      meeting:null,
      meetingUrl:null
    };

    $scope.model = model;

    $scope.createMeeting = function(){
      Meetings.createMeeting()
        .then(function(meeting){
          $scope.model.meeting = meeting;
          $scope.model.meetingUrl = APIHelper.endpoints.meetings + '/' + meeting.id;
        },function(error){
          $log.error(error);
        })
    };

    $scope.joinMeeting = function(meetingId){
      $state.go('join',{meetingId:meetingId})
    }

  });
