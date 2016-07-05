'use strict';

angular.module('yapp')
    .controller('ResultsCtrl', function($scope,$log,APIHelper,$state,User,Polls,Responses) {


        activate();

        function activate() {

            var model = {
                pollId:$state.params.id,
                poll:null,
                counts:[]
            };

            $scope.model = model;

            loadPoll(model.pollId);

            loadResponseCounts(model.pollId);
        }

        function loadPoll(poll_id) {
            Polls.getPollById(poll_id)
                .then(function(poll){
                    $scope.model.poll = poll;
                },function(error){
                    $log.error(error);
                })
        }

        function loadResponseCounts(poll_id) {
            Responses.getResponseCounts(poll_id)
                .then(function(counts){
                    $scope.model.counts = counts;
                },function(error){
                    $log.error(error);
                })
        }
    });
