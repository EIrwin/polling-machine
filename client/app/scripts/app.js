'use strict';

angular
  .module('yapp', [
    'ui.router',
    'ngAnimate'
  ])
  .config(function($stateProvider, $urlRouterProvider,$httpProvider) {

    $httpProvider.defaults.useXDomain = true;
    $httpProvider.defaults.headers.common["Accept"] = "application/json";
    $httpProvider.defaults.headers.common["Content-Type"] = "application/json";

    $urlRouterProvider.otherwise('/login');

    $stateProvider
        .state('base', {
          abstract: true,
          url: '',
          templateUrl: 'views/base.html'
        })
        .state('login',{
            url:'/login',
            parent:'base',
            templateUrl:'views/login.html',
            controller:'LoginCtrl'
        })
        .state('user-create',{
            url:'/users/create',
            parent:'base',
            templateUrl:'views/create_user.html',
            controller:'CreateUserCtrl'
        })
        .state('user-home',{
            url:'/users/:id',
            parent:'base',
            templateUrl:'views/home.html',
            controller:'HomeCtrl'
        })
        .state('polls-create',{
            url:'/polls/create',
            parent:'base',
            templateUrl:'views/create_poll.html',
            controller:'CreatePollCtrl'
        })
        .state('poll-edit',{
            url:'/polls/:id',
            parent:'base',
            templateUrl:'views/edit_poll.html',
            controller:'EditPollCtrl'
        })
        .state('item-create',{
            url:'/items/create/:poll_id',
            parent:'base',
            templateUrl:'views/create_item.html',
            controller:'CreateItemCtrl'
        })
        .state('item-edit',{
            url:'/items/:item_id/edit/:poll_id',
            parent:'base',
            templateUrl:'views/edit_item.html',
            controller:'EditItemCtrl'
        })
        .state('share',{
            url:'/share/:id',
            parent:'base',
            templateUrl:'views/share.html',
            controller:'ShareCtrl'
        })
        .state('polls',{
            url: '/polls',
            parent: 'dashboard',
            templateUrl:'views/'
        });

  });
