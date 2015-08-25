// Copyright © 2014, 2015 Maxim Tishchenko.
// All Rights Reserved.

'use strict';

/* App Module */

var iamApp = angular.module('pt', [
  'ngRoute'
]);

var iamServices = angular.module('iamServices', []);

iamServices.service('AuthService', ['$cookieStore', '$cookies', AuthService]);
iamServices.service('LoginService', ['$http', '$location', '$cookies', '$rootScope', '$window', 'Settings', 'AuthService', LoginService]);
iamServices.service('User', ['LoginService', '$location', '$q', '$rootScope', User]);


var iamControllers = angular.module('iamControllers', []);

iamControllers.controller('MainCtrl', ['$scope', 'User', MainCtrl]);
iamControllers.controller('LoginCtrl', ['$scope', '$auth', 'User', LoginCtrl]);
iamControllers.controller('EditProfileCtrl', ['$scope', '$auth', 'User', EditProfileCtrl]);
iamControllers.controller('EditProfilePasswordCtrl', ['$scope', '$auth', 'User', '$rootScope', EditProfilePasswordCtrl]);
iamControllers.controller('NotificationsCtrl', ['$scope', '$rootScope', NotificationsCtrl]);

iamControllers.controller('DateCtrl', ['$scope', DateCtrl]);

iamApp.config(['$routeProvider', '$locationProvider', function($routeProvider, $locationProvider){
  $locationProvider.html5Mode(true);
  $routeProvider
    .when('/', {
      templateUrl: 'partials/login.html',
      controller: 'LoginCtrl',
      auth: false
    })
    .when('/register', {
      templateUrl: 'partials/register.html',
      controller: 'LoginCtrl',
      auth: false
    })
    .when('/password-renew', {
      templateUrl: 'partials/reset.html',
      controller: 'LoginCtrl',
      auth: false
    })
    .when('/me', {
      templateUrl: 'partials/main.html',
      controller: 'MainCtrl',
      auth: true
    })
    .when('/me/edit-profile', {
      templateUrl: 'partials/edit.html',
      controller: 'MainCtrl',
      auth: true
    })
    .when('/me/edit-profile/password', {
      templateUrl: 'partials/password.html',
      controller: 'MainCtrl',
      auth: true
    })
    // .when('/oauth2callback', {
    //   template: 'FUCK',
    //   controller: 'MainCtrl',
    //   auth: false
    // })

    .otherwise({redirectTo: '/'});
}]);

iamApp.run(function($location, $rootScope, $route, AuthService) {
  $rootScope.$on('$locationChangeStart', function(evt, next, current) {
    var nextPath = $location.path(),
      nextRoute = $route.routes[nextPath];

    if (nextRoute && nextRoute.auth && !AuthService.isAuthenticated()) {
      $location.path("/");
    }

    if (nextRoute && !nextRoute.auth && AuthService.isAuthenticated()) {
      $location.path("/me");
    }
  });
});

iamApp.config(['$httpProvider', function ($httpProvider) {
  $httpProvider.defaults.withCredentials = true;
  //Reset headers to avoid OPTIONS request (aka preflight)
  $httpProvider.defaults.headers.common = {};
  $httpProvider.defaults.headers.post = {};
  $httpProvider.defaults.headers.put = {};
  $httpProvider.defaults.headers.patch = {};
  $httpProvider.defaults.useXDomain = true;
  delete $httpProvider.defaults.headers.common['X-Requested-With'];
}]);

iamApp.factory('authHttpResponseInterceptor',['$q', 'AuthService', function($q, AuthService){
    return {
        response: function(response){
            if (response.status === 401) {
                console.log("Response 401");
            }
            return response || $q.when(response);
        },
        responseError: function(rejection) {
            if (rejection.status === 401) {
                console.log("Response Error 401", rejection);
                AuthService.logout();
                // $location.path('/');
            }
            return $q.reject(rejection);
        }
    }
}]);

iamApp.config(['$httpProvider',function($httpProvider) {
    //Http Intercpetor to check auth failures for xhr requests
    $httpProvider.interceptors.push('authHttpResponseInterceptor');
}]);