'use strict';

/**
 * @ngdoc function
 * @name yapp.controller:MainCtrl
 * @description
 * # MainCtrl
 * Controller of yapp
 */
angular.module('yapp')
    .controller('CreatePollCtrl', function($scope,Users,$log,APIHelper,$state,User,Polls,Auth) {

        var model = {
            title:'',
            end:'2014-04-15T18:00:15-07:00',
            user_id:User.getCurrent().ID
        };

        $scope.model = model;

        $scope.logout = logout;

        $scope.createPoll = function(title,end){
            var userId = User.getCurrent().ID;
            console.log(userId);
            Polls.createPoll(title,end,userId)
                .then(function (poll) {
                    $state.go('user-home',{id:userId});
                },function (error) {
                    $log.error(error);
                })
        }

        function logout(){
            Auth.logout();
        }
    });
