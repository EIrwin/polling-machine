'use strict';

angular.module('yapp')
    .controller('EditItemCtrl', function($scope,$log,APIHelper,$state,User,Items,Auth) {

        activate();

        function activate() {

            var model = {
                poll_id:$state.params.poll_id,
                item_id:$state.params.item_id,
                item:null,
                user_id:User.getCurrent().ID
            }

            $scope.model = model;

            $scope.save = save;

            $scope.logout = logout;

            loadItem(model.poll_id,model.item_id);
        }
        
        function save(value,display) {
            var pollId = $scope.model.poll_id;
            var itemId = $scope.model.item_id;
            Items.updateItem(pollId,itemId,value,display)
                .then(function(item){
                    $state.go('poll-edit',{id:pollId});
                },function(error){
                    $log.error(error);
                })
        }
        
        function loadItem(poll_id,item_id) {
            Items.getItemById(poll_id,item_id)
                .then(function(item){
                    $scope.model.item = item;
                },function(error){
                    $log.error(error);
                })
        }
        
        function logout() {
            Auth.logout();
        }
    });
