var Knightspider = angular.module('phonecatControllers', []);

Knightspider.controller('PhoneListCtrl', ['$scope', '$http','$sce',
    function ($scope, $http, $sce) {
        $http.get('http://141.252.16.105/develop/api/v1/blogs').success(function(data) {
            $scope.posts = data;
        }).
        error(function() {
            $scope.error = true;
        });
        $scope.toTrustedHTML = function( html ){
            return $sce.trustAsHtml( html );
        }
        $scope.orderProp = 'titel';
    }]);




Knightspider.controller('PhoneDetailCtrl', ['$scope', '$http', '$routeParams','$sce',
    function($scope, $http, $routeParams, $sce) {
         $scope.toTrustedHTML = function( html ){
            return $sce.trustAsHtml( html );
        }
    $http.get('http://141.252.16.105/develop/api/v1/blogs/post/'+ $routeParams.postId).success(function(data)
        {
            //console.log(data);
            $scope.post = data[0];
            $scope.postId = $routeParams.postId;
        });
    }]);
    
    
    
Knightspider.controller('MyCtrl', function($scope, $location) {
    $scope.isActive = function(route) {
        return route === $location.path();
    }
});






Knightspider.controller('PageCtrl', ['$scope', '$routeParams', '$http',
    function($scope, $routeParams, $http) {
        $scope.data = $routeParams.pageName;
        $scope.isActive = function(route) {
        return route === $location.path();
       }
    }]);
    
    
 Knightspider.controller('DropdownCtrl', function ($scope, $log) {
  $scope.status = {
    isopen: false
  };
  $scope.toggleDropdown = function($event) {
    $event.preventDefault();
    $event.stopPropagation();
    $scope.status.isopen = !$scope.status.isopen;
  };
});

Knightspider.controller('CarouselDemoCtrl', function ($scope) {
  $scope.myInterval = 5000;
  var slides = $scope.slides = [];
  slides.push({
      image: 'http://141.252.16.105/develop/img/projectgroep.jpg',

    })
      slides.push({
      image: 'http://141.252.16.105/develop/img/web.jpg',
        text: 'Volledig online via Website en WebApp'
    })
});