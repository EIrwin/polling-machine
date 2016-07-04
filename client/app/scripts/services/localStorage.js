'use strict';

/**
 * @ngdoc function
 * @name yapp.controller:MainCtrl
 * @description
 * # MainCtrl
 * Controller of yapp
 */
angular.module('yapp')
  .factory('$localstorage', function($window) {
    return {
        set: function(key, value) {
          $window.localStorage[key] = value;
        },
        get: function(key, defaultValue) {
          return $window.localStorage[key] || defaultValue;
        },
        setObject: function(key, value) {
          $window.localStorage[key] = JSON.stringify(value);
        },
        getObject: function(key) {
          var value = $window.localStorage[key];
          if(value != undefined || value != null){
            value = JSON.parse(value);
          }
          return value;
        },
        clear:function(){
          $window.localStorage.clear();
        }
      }
  });
