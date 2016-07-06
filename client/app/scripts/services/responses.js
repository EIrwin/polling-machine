'use strict';

angular.module('yapp')
    .factory('Responses', function($http,$q,$log,APIHelper,$localstorage) {

        var self = this;

        this.createResponse = function(item_id,poll_id){
            var key = 'pm-' + poll_id;
            var token = $localstorage.get(key)
            var d = $q.defer();
            var requestUrl = APIHelper.fillUrl(APIHelper.endpoints.responses, {id:poll_id});
            var data = {
                item_id:item_id,
                poll_id:poll_id,
                token:token
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

        this.getResponseCounts = function(poll_id){
            var d = $q.defer();
            var requestUrl = APIHelper.fillUrl(APIHelper.endpoints.responseCount, {id:poll_id});
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
        
        this.getResponseToken = function(poll_id){
            var d = $q.defer();
            var requestUrl = APIHelper.fillUrl(APIHelper.endpoints.responseToken, {id:poll_id});
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

        this.setResponseToken = function(poll_id,token){
            var key = 'pm-' + poll_id;
            $localstorage.set(key,token);
        }

        this.hasResponseToken = function(poll_id){
            var key = 'pm-' + poll_id;
            var token = $localstorage.get(key)
            if(token){
                return true;
            }else{
                return false;
            }
        }
        return self;
    });

