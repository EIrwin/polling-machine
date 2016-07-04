'use strict';

angular.module('yapp')
    .factory('Items', function($http,$q,$log,APIHelper) {

        var self = this;

        this.createItem = function(poll_id,display,value){
            var d = $q.defer();
            var requestUrl = APIHelper.fillUrl(APIHelper.endpoints.items, {poll_id:poll_id});
            var data = {
                poll_id:poll_id,
                display:display,
                value:value
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

        this.getItemById = function(poll_id,item_id){
            var d = $q.defer();
            var requestUrl = APIHelper.fillUrl(APIHelper.endpoints.item, {item_id:item_id,poll_id:poll_id});
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

        this.getItemsByPollId = function(id){
          var d = $q.defer();
            var requestUrl = APIHelper.fillUrl(APIHelper.endpoints.items, {poll_id:id});
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

        this.deleteItem = function(poll_id,item_id){
            var d = $q.defer();
            var requestUrl = APIHelper.fillUrl(APIHelper.endpoints.item, {item_id:item_id,poll_id:poll_id});
            $http
                .post(requestUrl)
                .success(function(data,status,headers,config){
                    d.resolve(data);
                })
                .error(function(data,status,headers,config){
                    d.reject(data);
                });
            return d.promise;
        }
        
        this.updateItem = function (poll_id,item_id,value,display) {
            var d = $q.defer();
            var requestUrl = APIHelper.fillUrl(APIHelper.endpoints.item, {item_id:item_id,poll_id:poll_id});
            var data = {
                poll_id:poll_id,
                display:display,
                value:value
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

