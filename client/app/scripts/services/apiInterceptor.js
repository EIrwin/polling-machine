'use strict';

angular.module('yapp')
  .factory('APIInterceptor', function($rootScope,$log,User) {

      var self = this;

      var includeCommonAuthHeader = function(config){

          //right now, the only path we need to check
          //for is login, but this might include other
          //paths later which is why want to keep this
          //in its own function
          var loginPath = "/login";
          var index = config.url.indexOf(loginPath);
          return index == -1;
      }
      self.request = function(config) {
          //we don't want residual tokens to be passed in
          //to certain requests such as login so we need to
          //check to see if this is a request that we need to
          //ommit the common authentication header
          if(includeCommonAuthHeader(config)){
              if(User.isAuthenticated()){
                  var token = User.getToken();
                  if(token != undefined){
                      // config.headers.Authorization = 'Bearer ' + token;
                  }
              }
          }
          return config;
      };
      self.responseError = function(response) {
          if (response.status === 401) {
              User.clearToken();
          }
          return response;
      };
      return self;

  });
