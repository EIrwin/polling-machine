    'use strict';

    angular.module('yapp')
      .controller('LoginCtrl', function($scope, $location,Auth,$log,Users,User,$state) {

        $scope.submit = function(email,password) {

            var model = {
                email:'',
                password:''
            };

            $scope.model = model;

            Auth.login(email,password)
                .then(function(token){
                    var claims = User.getTokenClaims(token);
                    Users.getUserById(claims.user_id)
                        .then(function(user){
                            User.setCurrent(user);
                            $state.go('user-home',{id:claims.user_id});
                        },function(error){
                            $log.error(error);
                        })

                },function(error){
                    User.clearToken();
                });
        }

      });
