'use strict';

angular.module('yapp')
  .factory('Auth', function($http,$q,$log,APIHelper,User) {

    var self = this;

      this.login = function(email,password){
          var d = $q.defer();
          var requestUrl = APIHelper.endpoints.login;
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
                  console.log('logged in:' + data);
                  User.setToken(data);
                  
                  d.resolve(data);
              })
              .error(function(data,status,headers,config){
                  d.reject(data);
              });
          return d.promise;
      }

    return self;

  });
