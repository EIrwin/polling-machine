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
          $log.info($scope.model.token);
          //do local check for slug for dupe check
          $scope.model.loading = true;
          $scope.model.submitted = true;
          Responses.createResponse(item_id,poll_id,$scope.model.token)
              .then(function(resp){
                  //show success message
                  $scope.model.loading = false;
                  //store token to prevent more submisions
                  Responses.setResponseToken(poll_id,$scope.model.token);
              },function (error) {
                  $scope.model.loading = false;
                  $log.error(error);
              })
      }

  });
