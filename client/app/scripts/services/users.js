'use strict';

angular.module('yapp')
  .factory('Users', function($http,$q,$log,APIHelper) {

    var self = this;

    this.createUser = function(email,password){
      var d = $q.defer();
      var requestUrl = APIHelper.endpoints.users;
      var data = {
          email:email,
          password:password
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
    };


      this.getUserById = function(id){
          var d = $q.defer();
          var requestUrl = APIHelper.fillUrl(APIHelper.endpoints.user, {id:id}, {});
          $http({
              url:requestUrl,
              dataType: 'json',
              method: 'GET',
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
