'use strict';

/**
 * @ngdoc function
 * @name yapp.controller:MainCtrl
 * @description
 * # MainCtrl
 * Controller of yapp
 */
angular.module('yapp')
  .factory('Meetings', function($q,$http,$log,APIHelper) {

    var self = this;

    this.getById = function(meetingId){
      var d = $q.defer();
      var requestUrl = APIHelper.fillUrl(APIHelper.endpoints.meeting,{id:meetingId});
      $http
        .get(requestUrl)
        .success(function(data,status,headers,config){
          d.resolve(data);
        })
        .error(function(data,status,headers,config){
          d.reject(data);
        });
      return d.promise;
    }

    this.createMeeting = function(){
      var d = $q.defer();
      var requestUrl = APIHelper.endpoints.meetings;
      $http
        .post(requestUrl)
        .success(function(data,status,headers,config){
          d.resolve(data);
        })
        .error(function(data,status,headers,config){
          d.reject(data);
        });
      return d.promise;
    };



    this.joinMeeting = function(userId,meetingId){
      var d = $q.defer();
      var requestUrl = APIHelper.endpoints.meetings + '/_join';
      var data = {
        userId:userId,
        meetingId:meetingId
      };
      $http({
           url:requestUrl,
           dataType: 'json',
           method: 'PUT',
           data:data,
           headers: {
               "Content-Type": "application/json"
           }
       })
        .success(function(data,status,headers,config){
          d.resolve(data);
        })
        .error(function(data,status,headers,config){
          d.reject(data);
        });
      return d.promise;
    };

    this.leaveMeeting = function(userId,meetingId){
      var d = $q.defer();
      var requestUrl = APIHelper.endpoints.meetings + '/_leave';
      var data = {
        userId:userId,
        meetingId:meetingId
      };
      $http({
           url:requestUrl,
           dataType: 'json',
           method: 'PUT',
           data:data,
           headers: {
               "Content-Type": "application/json"
           }
       })
        .success(function(data,status,headers,config){
          d.resolve(data);
        })
        .error(function(data,status,headers,config){
          d.reject(data);
        });
      return d.promise;
    };

    this.startMeeting = function(meetingId){
      var d = $q.defer();
      var requestUrl = APIHelper.endpoints.meetings + '/_start';
      var data = {
        meetingId:meetingId
      };
      $http({
           url:requestUrl,
           dataType: 'json',
           method: 'PUT',
           data:data,
           headers: {
               "Content-Type": "application/json"
           }
       })
        .success(function(data,status,headers,config){
          d.resolve(data);
        })
        .error(function(data,status,headers,config){
          d.reject(data);
        });
      return d.promise;
    };

    this.endMeeting = function(meetingId){
      var d = $q.defer();
      var requestUrl = APIHelper.endpoints.meetings + '/_end';
      var data = {
        meetingId:meetingId
      };
      $http({
           url:requestUrl,
           dataType: 'json',
           method: 'PUT',
           data:data,
           headers: {
               "Content-Type": "application/json"
           }
       })
        .success(function(data,status,headers,config){
          d.resolve(data);
        })
        .error(function(data,status,headers,config){
          d.reject(data);
        });
      return d.promise;
    };

    return self;


  });
