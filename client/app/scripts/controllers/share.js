'use strict';

angular.module('yapp')
  .controller('ShareCtrl', function($scope,Meetings,$log,APIHelper,$state,$stateParams,Polls,Items,Responses,$localstorage) {

      activate();

      function activate() {

          var model = {
              poll_id:$stateParams.id,
              poll:null,
              items:[],
              item_id:null,
              submitted:false,
              token:null,
              hasToken:false
          }

          $scope.model = model;

          $scope.submitAnswer = submitAnswer;

          var hasToken = Responses.hasResponseToken(model.poll_id);

          if(hasToken){
              $scope.model.hasToken = true;

          }else{
              Responses.getResponseToken(model.poll_id)
                  .then(function(token){
                      $scope.model.token = token;
                      loadPoll(model.poll_id);
                      loadItems(model.poll_id);
                  },function(error){
                      $log.error(error);
                  })
          }
      }
      


      function loadPoll(id) {
        Polls.getPollById(id)
            .then(function(poll){
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

                  //store token to prevent more submisions
                  Responses.setResponseToken(poll_id,$scope.model.token);
              },function (error) {
                  $log.error(error);
              })
      }

  });
