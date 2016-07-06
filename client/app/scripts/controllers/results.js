'use strict';

angular.module('yapp')
    .controller('ResultsCtrl', function($scope,$log,APIHelper,$state,User,Polls,Responses,Auth) {

        activate();

        function activate() {

            var model = {
                pollId:$state.params.id,
                poll:null,
                counts:[],
                user_id:User.getCurrent().ID,
                labels:[],
                series:[],
                data:[[]]
            };

            $scope.model = model;

            $scope.logout = logout;

            loadPoll(model.pollId);

            loadResponseCounts(model.pollId);
        }

        function loadPoll(poll_id) {
            Polls.getPollById(poll_id)
                .then(function(poll){
                    $scope.model.series.push(poll.title);
                    $scope.model.poll = poll;
                },function(error){
                    $log.error(error);
                })
        }

        function loadResponseCounts(poll_id) {
            Responses.getResponseCounts(poll_id)
                .then(function(counts){
                    for(var i in counts){
                        var count = counts[i]
                        $scope.model.labels.push(count.display)
                        $scope.model.data[0].push(count.count)
                    }
                    $scope.model.counts = counts;
                },function(error){
                    $log.error(error);
                })
        }

        function logout(){
            Auth.logout();
            $state.go('login');
        }
    });
