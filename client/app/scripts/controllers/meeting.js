'use strict';

/**
 * @ngdoc function
 * @name yapp.controller:MainCtrl
 * @description
 * # MainCtrl
 * Controller of yapp
 */
angular.module('yapp')
  .controller('MeetingCtrl', function($scope,$stateParams,Meetings,$log,$interval,$localstorage,$state) {

    var model = {
      meetingId:$stateParams.meetingId,
      meeting:null
    }

    $scope.model = model;

    var getMeetingById = function(meetingId){
      Meetings.getById(model.meetingId)
        .then(function(meeting){
          $scope.model.meeting = meeting;
        },function(error){
          $log.error(error);
        })
    };

    getMeetingById(model.meetingId);

    $interval(function(){
      getMeetingById(model.meetingId);
    },5000);

    $scope.startMeeting = function(meetingId){
      Meetings.startMeeting(meetingId)
        .then(function(meeting){
          $scope.model.meeting = meeting;
        },function(error){
          $log.error(error);
        })
    }

    $scope.endMeeting = function(meetingId){
      Meetings.endMeeting(meetingId)
        .then(function(meeting){
          $scope.model.meeting = meeting;
        },function(error){
          $log.error(error);
        })
    }

    $scope.leaveMeeting = function(meetingId){
      var userId = $localstorage.get('userId');
      if(userId != undefined){
        Meetings.leaveMeeting(userId,meetingId)
          .then(function(meeting){
            $state.go('create');
          },function(error){
            $log.error(error);
          })
      }
    }

    $scope.showStartMeetingButton = function(){
      return $scope.model.meeting.startUtc == "0001-01-01T00:00:00.0000000Z";
    }

    $scope.showEndMeetingButton = function(){
      return $scope.model.meeting.endUtc == "0001-01-01T00:00:00.0000000Z";
    }

  });
