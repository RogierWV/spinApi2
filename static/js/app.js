var phonecatApp = angular.module('phonecatApp', [
    'ngRoute',
    'ui.bootstrap',
    'phonecatControllers'
]);

phonecatApp.config(['$routeProvider',
    function($routeProvider, $localStorage) {
        $routeProvider.
        when('/blog',{
            templateUrl: 'partials/blog.html',
                      controller: 'PhoneListCtrl'
        }).
        when('/blog/post/:postId', {
            templateUrl: 'partials/post.html',
            controller: 'PhoneDetailCtrl'
        }).
        when('/analytics',{
             templateUrl: 'partials/analytics.html',
        }).
        when('/contact',{
             templateUrl: 'partials/contact.html',
        }).
        when('/',{
             templateUrl: 'partials/home.html',
        }).
        when('/termsOfService',{
            templateUrl: 'partials/termsOfService.html',
        }).
        when('/privacy',{
            templateUrl: 'partials/privacy.html',
        }).
        otherwise({
             templateUrl: 'partials/404.html',
        });
    }]);
    
    
