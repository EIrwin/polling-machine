'use strict';

angular.module('yapp')
    .factory('Responses', function($http,$q,$log,APIHelper) {

        var self = this;

        this.createResponse = function(item_id,poll_id){
            var d = $q.defer();
            var requestUrl = APIHelper.fillUrl(APIHelper.endpoints.responses, {id:poll_id});;
            var data = {
                item_id:item_id,
                poll_id:poll_id
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

        // this.getResponsesByPollId = function (pollId) {
        //     var d = $q.defer();
        //     var requestUrl = APIHelper.endpoints.polls + "?user_id=" + userId;
        //     $http
        //         .get(requestUrl)
        //         .success(function(data,status,headers,config){
        //             d.resolve(data);
        //         })
        //         .error(function(data,status,headers,config){
        //             d.reject(data);
        //         });
        //     return d.promise;
        // }

        return self;
    });

