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
            user_id:User.getCurrent().ID,
            date:new Date(),
            time:''
        };

        $scope.model = model;

        $scope.logout = logout;

        $scope.createPoll = function(title,end){
            var userId = User.getCurrent().ID;


            var exp = ISODateString($scope.model.date) + ISOTimeString($scope.model.time);

            Polls.createPoll(title,exp,userId)
                .then(function (poll) {
                    $state.go('user-home',{id:userId});
                },function (error) {
                    $log.error(error);
                })
        }

        function logout(){
            Auth.logout();
        }

        $scope.inlineOptions = {
            minDate: new Date(),
            showWeeks: true
        };

        $scope.dateOptions = {
            formatYear: 'yy',
            maxDate: new Date(2020, 5, 22),
            minDate: new Date(),
            startingDay: 1
        };

        $scope.open1 = function() {
            $scope.popup1.opened = true;
        };

        $scope.format = 'dd-MMMM-yyyy';

        $scope.popup1 = {
            opened: false
        };

        function ISODateString(d){
            function pad(n){return n<10 ? '0'+n : n}
            return d.getFullYear()+'-'
                + pad(d.getMonth()+1)+'-'
                + pad(d.getDate())+'T';
        }

        function ISOTimeString(d) {
            function pad(n){return n<10 ? '0'+n : n}
            return pad(d.getHours())+':'
            + pad(d.getMinutes())+':'
            + pad(d.getSeconds())+'Z';
        }
    });
