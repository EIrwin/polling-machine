'use strict';

angular.module('yapp')
  .controller('ShareCtrl', function($scope,Meetings,$log,APIHelper,$state,$stateParams,Polls,Items,Responses) {

      activate();

      function activate() {
          var model = {
              poll_id:$stateParams.id,
              poll:null,
              items:[],
              item_id:null,
              submitted:false
          }

          $scope.model = model;

          $scope.submitAnswer = submitAnswer;

          loadPoll(model.poll_id);

          loadItems(model.poll_id);
      }

      function loadPoll(id) {
        Polls.getPollById(id)
            .then(function(poll){
                console.log(poll);
                $scope.model.poll = poll;
            },function (error) {
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
      
      function submitAnswer(item_id,poll_id) {
          //do local check for slug for dupe check
          
          Responses.createResponse(item_id,poll_id)
              .then(function(resp){
                  //show success message
                  $scope.model.submitted = true;
                  //store confirmation token for or something locally
              },function (error) {
                  $log.error(error);
              })
      }

  });
