function TasksCtrl($scope, $http) {
    $scope.tasks = [];

    $http.get('/tasks/').then(function(response) {
        console.log(response);
        $scope.tasks = response.data;
    });

}
