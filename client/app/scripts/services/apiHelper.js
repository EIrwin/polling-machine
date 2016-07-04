'use strict';

angular.module('yapp')
  .factory('APIHelper', function(ENV) {

    var self = this;
    var baseApiUrl = ENV.apiEndpoint;

    this.endpoints = {
        items:baseApiUrl + '/polls/:poll_id/items',
        item:baseApiUrl + '/polls/:poll_id/items/:item_id',
        meetings:baseApiUrl + '/meetings',
        meeting:baseApiUrl + '/meetings/:id',
        users:baseApiUrl + '/users',
        user:baseApiUrl + '/users/:id',
        polls:baseApiUrl + '/polls',
        poll:baseApiUrl + '/polls/:id',
        login:baseApiUrl + '/login'
    }

    this.fillUrl = function(urlFormat, pathParams, queryParams) {
       //i.e. APIHelper.endpoints.state
       var url = urlFormat;

       angular.forEach(pathParams, function (val, name) {
           if (typeof(val) === 'undefined' || val === null || val === '') {
               url = url.replace(new RegExp('/:' + name, 'g'), '');
           } else {
               url = url.replace(new RegExp(':' + name, 'g'), val);
           }

       });

       if(queryParams){
           var formattedQueryParams = [];
           angular.forEach(queryParams, function(val, key){
               //TODO format values (i.e.: arrays)
               formattedQueryParams.push(''+key+'='+val);
           });
           url += '?' + formattedQueryParams.join('&');
       }

       return url;
   }

    return self;

  });
