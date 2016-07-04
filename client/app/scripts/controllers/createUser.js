'use strict';

/**
 * @ngdoc function
 * @name yapp.controller:MainCtrl
 * @description
 * # MainCtrl
 * Controller of yapp
 */
angular.module('yapp')
    .controller('CreateUserCtrl', function($scope,Users,$log,APIHelper,$state) {

        var model = {
            email:null,
            password:''
        };

        $scope.model = model;

        $scope.createUser = function (email,password) {
            Users.createUser(email,password)
                .then(function(user){
                    $state.go('login');
                },function(error){
                    $log.error(error);
                })
        }

    });
