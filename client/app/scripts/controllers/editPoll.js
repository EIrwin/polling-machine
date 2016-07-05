'use strict';

angular.module('yapp')
    .controller('EditPollCtrl', function($scope,$log,APIHelper,$state,User,Polls,Items) {


        activate();

        function activate() {

            var model = {
                pollId:$state.params.id,
                poll:null,
                items:[],
                user_id:User.getCurrent().ID
            };

            $scope.model = model;

            loadPoll(model.pollId);

            loadItems(model.pollId);

            $scope.save = save;
            $scope.deleteItem  = deleteItem;
        }
        
        
        function save(id,user_id,start,end,title) {
            var userId = User.getCurrent().ID;
            Polls.updatePoll(id,user_id,start,end,title)
                .then(function (poll) {
                    $state.go('user-home',{id:userId});
                },function (error) {
                    $log.error(error);
                })
        }

        function loadPoll(poll_id) {
            Polls.getPollById(poll_id)
                .then(function(poll){
                    $scope.model.poll = poll;
                },function(error){
                    $log.error(error);
                })
        }
        
        function loadItems(poll_id) {
            Items.getItemsByPollId(poll_id)
                .then(function(items){
                    $scope.model.items = items;
                },function (error){
                    $log.error(error);
                })
        }

        function deleteItem(poll_id,item_id) {
            Items.deleteItem(poll_id,item_id)
                .then(function(item){
                    loadItems(poll_id);
                })
        }
    });
