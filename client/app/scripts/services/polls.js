'use strict';

angular.module('yapp')
    .factory('Polls', function($http,$q,$log,APIHelper) {

        var self = this;

        this.createPoll = function(title,end,user_id){
            var d = $q.defer();
            var requestUrl = APIHelper.endpoints.polls;
            var data = {
                title:title,
                end:end,
                user_id:user_id
            };
            $http({
                url:requestUrl,
                dataType: 'json',
                method: 'POST',
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
        }

        this.getPollById = function(id){
          var d = $q.defer();
            var requestUrl = APIHelper.fillUrl(APIHelper.endpoints.poll, {id:id});
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
        
        this.getPollsByUserId = function (userId) {
            var d = $q.defer();
            var requestUrl = APIHelper.endpoints.polls + "?user_id=" + userId;
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

        this.updatePoll = function(id,user_id,start,end,title){
            var d = $q.defer();
            var requestUrl = APIHelper.fillUrl(APIHelper.endpoints.poll, {id:id});
            var data = {
                id:id,
                user_id:user_id,
                start:start,
                end:end,
                title:title
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
        }

        return self;
    });

