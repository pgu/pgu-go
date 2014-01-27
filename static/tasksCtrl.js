'use strict';

angular.module('PguGo', [])

    .run(function ($http) {
        $http.defaults.headers.common.Authorization = 'Basic cGd1Omdv';
    })

    .controller('TasksCtrl', function ($scope, $http) {

        $scope.displayUI = false;

        $scope.tasks = [];
        $scope.newTask = {
            Title: ''
        };

        $scope.isWorkingCreation = false;
        $scope.isWorkingEdition = false;

        $scope.editionTask = null;

        function fetchTasks() {
            return $http.get('/tasks/').then(function (response) {

                $scope.tasks = response.data;
                $scope.displayUI = true;

            }).catch(function (response) {
                    $scope.displayUI = response.status !== 401;
                });
        }

        fetchTasks();

        $scope.tasksSort = function (task) {

            var prefix = '';

            if (task.Done) {
                prefix = '4';
            } else {
                if (task.Priority === 'normal') {
                    prefix = '3';

                } else if (task.Priority === 'warning') {
                    prefix = '2';

                } else if (task.Priority === 'asap') {
                    prefix = '1';

                }
            }

            return prefix + task.Title;
        };

        $scope.addTask = function () {

            $scope.isWorkingCreation = true;

            $http.post('/tasks/', $scope.newTask)

                .then(function (response) {

                    var pathToNewTask = response.headers().location;
                    return $http.get(pathToNewTask).then(function (response) {

                        $scope.tasks.push(response.data);
                    });

                })
                .finally(function () {
                    $scope.newTask.Title = '';
                    $scope.isWorkingCreation = false;
                });
        };

        var theTask = null;

        $scope.editTask = function (task) {
            theTask = task;
            $scope.editionTask = angular.copy(task);

            $('html, body').animate({scrollTop: 1000}, 'slow');
        };

        function endsEdition() {
            $scope.editionTask = null;
            theTask = null;
            $scope.isWorkingEdition = false;

            $('html, body').animate({scrollTop: 0}, 'fast');
        }

        $scope.updateTask = function () {
            $scope.isWorkingEdition = true;

            $http.put('/tasks/' + $scope.editionTask.ID, $scope.editionTask)

                .then(function () {
                    theTask.Title = $scope.editionTask.Title;
                    theTask.Done = $scope.editionTask.Done;
                    theTask.Priority = $scope.editionTask.Priority;
                })

                .finally(function () {
                    endsEdition();
                })
            ;
        };

        $scope.deleteTask = function () {
            $scope.isWorkingEdition = true;

            $http.delete('/tasks/' + $scope.editionTask.ID)

                .then(function () {
                    $scope.tasks = _.reject($scope.tasks, function (task) {
                        return task.ID === $scope.editionTask.ID;
                    });
                })

                .finally(function () {
                    endsEdition();
                });

        };

        $scope.cancelEdition = function () {
            endsEdition();
        };

    });
