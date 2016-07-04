'use strict';

angular.module('yapp')
    .factory('User', function($localstorage,$window) {

        var self = this;

        this.setCurrent = function(user){
            $localstorage.setObject('user',user);
        };

        this.getCurrent = function(){
            return $localstorage.getObject('user');
        };

        this.clearCurrent = function(){
            $localstorage.setObject('user',undefined);
        };

        this.isAuthenticated = function(){
            var isAuthenticated = $localstorage.get('jwt-token') != 'undefined';
            return isAuthenticated;
        }

        this.getToken = function(){
            return $localstorage.get('jwt-token');
        };

        this.setToken = function(token){
            $localstorage.set('jwt-token',token)
        };

        this.clearToken = function(){
            $localstorage.set('jwt-token',undefined);
        }

        this.getTokenClaims = function(token){
            var base64Url = token.split('.')[1];
            var base64 = base64Url.replace('-', '+').replace('_', '/');
            return JSON.parse($window.atob(base64));
        };

        return self;


    });

