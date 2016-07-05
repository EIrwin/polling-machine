'use strict';

angular.module('yapp')
    .controller('CreateItemCtrl', function($scope,$log,APIHelper,$state,User,Items) {

        activate();

        function activate() {

            var model = {
                value:'',
                display:'',
                poll_id:$state.params.poll_id
            }

            $scope.model = model;

            $scope.create = create;
        }
        
        
        function create(value,display) {
            var pollId = $scope.model.poll_id;
            Items.createItem(pollId,value,display)
                .then(function(item){
                    $state.go('poll-edit',{id:pollId});
                },function(error){
                    $log.error(error);
                })
        }
    });