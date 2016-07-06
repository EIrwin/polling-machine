'use strict';

angular
  .module('yapp', [
    'ui.router',
    'chart.js',
      'ui.bootstrap',
      'ui.bootstrap.datepicker'
  ])
    .run(function($rootScope,User,$state){
        $rootScope.$on("$stateChangeStart", function (event, toState, toParams, fromState, fromParams) {
            if (toState.data.authenticate && !User.isAuthenticated()) {
                // User isn’t authenticated
                $state.transitionTo("login");
                event.preventDefault();
            }
        });
    })
  .config(function($stateProvider, $urlRouterProvider,$httpProvider) {

      $httpProvider.defaults.useXDomain = true;
      $httpProvider.defaults.headers.common["Accept"] = "application/json";
      $httpProvider.defaults.headers.common["Content-Type"] = "application/json";
      $httpProvider.interceptors.push('APIInterceptor');

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
            controller:'LoginCtrl',
            data: {
                authenticate: false
            }
        })
        .state('user-create',{
            url:'/users/create',
            parent:'base',
            templateUrl:'views/create_user.html',
            controller:'CreateUserCtrl',
            data: {
                authenticate: false
            }
        })
        .state('user-home',{
            url:'/users/:id',
            parent:'base',
            templateUrl:'views/home.html',
            controller:'HomeCtrl',
            data: {
                authenticate: true
            }
        })
        .state('polls-create',{
            url:'/polls/create',
            parent:'base',
            templateUrl:'views/create_poll.html',
            controller:'CreatePollCtrl',
            data: {
                authenticate: true
            }
        })
        .state('results',{
            url:'/polls/:id/results',
            parent:'base',
            templateUrl:'views/results.html',
            controller:'ResultsCtrl',
            data: {
                authenticate:true
            }
        })
        .state('poll-edit',{
            url:'/polls/:id',
            parent:'base',
            templateUrl:'views/edit_poll.html',
            controller:'EditPollCtrl',
            data: {
                authenticate: true
            }
        })
        .state('item-create',{
            url:'/items/create/:poll_id',
            parent:'base',
            templateUrl:'views/create_item.html',
            controller:'CreateItemCtrl',
            data: {
                authenticate: true
            }
        })
        .state('item-edit',{
            url:'/items/:item_id/edit/:poll_id',
            parent:'base',
            templateUrl:'views/edit_item.html',
            controller:'EditItemCtrl',
            data: {
                authenticate: true
            }
        })
        .state('share',{
            url:'/share/:id',
            parent:'base',
            templateUrl:'views/share.html',
            controller:'ShareCtrl',
            data: {
                authenticate: false
            }
        });

  });
