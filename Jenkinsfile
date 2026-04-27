// // pipeline {
// //     agent any

// //     environment {
// //         APP_NAME = "tubestahap2"
// //         DOCKER_USER = "umarx"
// //     }

// //     stages {

// //         stage('Unit Test') {
// //             steps {
// //                 bat 'cd user-service && go test -v'
// //                 bat 'cd order-service && go test -v'
// //             }
// //         }

// //         stage('Lint / Vet') {
// //             steps {
// //                 bat 'cd user-service && go vet ./...'
// //                 bat 'cd order-service && go vet ./...'
// //             }
// //         }

// //         stage('Build Docker Image') {
// //             steps {
// //                 bat 'docker compose build'
// //             }
// //         }
        
// //         stage('Functional Test') {
// //             steps {
// //                 bat 'docker compose up -d'
// //                 bat 'docker compose exec order-service go test -tags=functional -v'
// //             }
// //         }
        
// //         stage('Push Image') {
// //             steps {
// //                 bat 'docker tag tubestahap2-user-service %DOCKER_USER%/user-service:latest'
// //                 bat 'docker tag tubestahap2-order-service %DOCKER_USER%/order-service:latest'

// //                 bat 'docker push %DOCKER_USER%/user-service:latest'
// //                 bat 'docker push %DOCKER_USER%/order-service:latest'
// //             }
// //         }

// //         stage('Deploy Kubernetes') {
// //             steps {
// //                 bat 'kubectl apply -f k8s/'
// //             }
// //         }

// //         stage('Verify') {
// //             steps {
// //                 bat 'kubectl get pods'
// //                 bat 'kubectl get svc'
// //             }
// //         }
// //     }
// // }




















// pipeline {
//     agent any

//     environment {
//         APP_NAME = "tubestahap2"
//         DOCKER_USER = "umarx"
//         K8S_NAMESPACE = "default"
//     }

//     stages {

//         // =========================
//         // 1. CHECKOUT
//         // =========================
//         stage('1. Checkout Repo') {
//             steps {
//                 checkout scm
//             }
//         }

//         // =========================
//         // 2. UNIT TEST (NO DB / MOCK ONLY)
//         // =========================
//         stage('2. Unit Tests') {
//             steps {
//                 script {
        
//                     catchError(buildResult: 'SUCCESS', stageResult: 'UNSTABLE') {
//                         bat 'cd user-service && go test -v'
//                     }
        
//                     catchError(buildResult: 'SUCCESS', stageResult: 'UNSTABLE') {
//                         bat 'cd order-service && go test -v'
//                     }
        
//                     catchError(buildResult: 'SUCCESS', stageResult: 'UNSTABLE') {
//                         bat 'cd report-service && if exist go.mod (go test -v)'
//                     }
        
//                     catchError(buildResult: 'SUCCESS', stageResult: 'UNSTABLE') {
//                         bat 'cd courier-service && if exist go.mod (go test -v)'
//                     }
        
//                     catchError(buildResult: 'SUCCESS', stageResult: 'UNSTABLE') {
//                         bat 'cd gudang-service && if exist go.mod (go test -v)'
//                     }
//                 }
//             }
//         }

//         // =========================
//         // 3. LINT / VET
//         // =========================
//         stage('3. Lint / Vet') {
//             steps {
//                 bat 'cd user-service && go vet ./...'
//                 bat 'cd order-service && go vet ./...'
//                 bat 'cd report-service && if exist go.mod (go vet ./...) else echo skip'
//                 bat 'cd courier-service && if exist go.mod (go vet ./...) else echo skip'
//                 bat 'cd gudang-service && if exist go.mod (go vet ./...) else echo skip'
//             }
//         }

//         // =========================
//         // 4. BUILD DOCKER IMAGE (LOCAL)
//         // =========================
//         stage('4. Build Docker Images') {
//             steps {
//                 bat 'docker compose build'
//             }
//         }

//         // =========================
//         // 5. FUNCTIONAL TEST (BOLEH DB / SERVICE)
//         // =========================
//         stage('5. Functional Tests') {
//             steps {
//                 script {
        
//                     bat 'docker compose up -d'
//                     bat 'timeout /t 5'
        
//                     catchError(buildResult: 'SUCCESS', stageResult: 'UNSTABLE') {
//                         bat 'cd user-service && go test -tags=functional -v'
//                     }
        
//                     catchError(buildResult: 'SUCCESS', stageResult: 'UNSTABLE') {
//                         bat 'cd order-service && go test -tags=functional -v'
//                     }
        
//                     catchError(buildResult: 'SUCCESS', stageResult: 'UNSTABLE') {
//                         bat 'cd courier-service && if exist go.mod go test -tags=functional -v'
//                     }
        
//                     catchError(buildResult: 'SUCCESS', stageResult: 'UNSTABLE') {
//                         bat 'cd gudang-service && if exist go.mod go test -tags=functional -v'
//                     }
        
//                     catchError(buildResult: 'SUCCESS', stageResult: 'UNSTABLE') {
//                         bat 'cd report-service && if exist go.mod go test -tags=functional -v'
//                     }
        
//                     bat 'docker compose down'
//                 }
//             }
//         }

//         // =========================
//         // 6. PUSH IMAGE
//         // =========================
//         stage('6. Push Images') {
//             steps {
//                 bat 'docker tag tubestahap2-user-service %DOCKER_USER%/user-service:latest'
//                 bat 'docker tag tubestahap2-order-service %DOCKER_USER%/order-service:latest'
//                 bat 'docker tag tubestahap2-report-service %DOCKER_USER%/report-service:latest'
//                 bat 'docker tag tubestahap2-courier-service %DOCKER_USER%/courier-service:latest'
//                 bat 'docker tag tubestahap2-gudang-service %DOCKER_USER%/gudang-service:latest'

//                 bat 'docker push %DOCKER_USER%/user-service:latest'
//                 bat 'docker push %DOCKER_USER%/order-service:latest'
//                 bat 'docker push %DOCKER_USER%/report-service:latest'
//                 bat 'docker push %DOCKER_USER%/courier-service:latest'
//                 bat 'docker push %DOCKER_USER%/gudang-service:latest'
//             }
//         }

//         // =========================
//         // 7. DEPLOY KUBERNETES
//         // =========================
//         stage('7. Deploy Kubernetes') {
//             steps {
//                 bat 'kubectl apply -f k8s/'
//             }
//         }

//         // =========================
//         // 8. VERIFY
//         // =========================
//         stage('8. Verify') {
//             steps {
//                 bat 'kubectl get pods'
//                 bat 'kubectl get svc'
//                 bat 'kubectl get deployments'
//             }
//         }
//     }

//     post {
//         always {
//             bat 'docker compose down || exit 0'
//         }
//         success {
//             echo "PIPELINE SUCCESS"
//         }
//         failure {
//             echo "PIPELINE FAILED"
//         }
//     }
// }





pipeline {
    agent any

    environment {
        APP_NAME      = "tubestahap2"
        DOCKER_USER   = "umarx"
        K8S_NAMESPACE = "default"
    }

    stages {

        // =========================
        // 1. CHECKOUT
        // =========================
        stage('1. Checkout Repo') {
            steps {
                checkout scm
            }
        }

        // =========================
        // 2. UNIT TEST
        // =========================
        stage('2. Unit Tests') {
            steps {
                script {

                    catchError(buildResult: 'SUCCESS', stageResult: 'UNSTABLE') {
                        bat 'cd user-service && go test -v'
                    }

                    catchError(buildResult: 'SUCCESS', stageResult: 'UNSTABLE') {
                        bat 'cd order-service && go test -v'
                    }

                    catchError(buildResult: 'SUCCESS', stageResult: 'UNSTABLE') {
                        bat 'cd report-service && if exist go.mod (go test -v)'
                    }

                    catchError(buildResult: 'SUCCESS', stageResult: 'UNSTABLE') {
                        bat 'cd courier-service && if exist go.mod (go test -v)'
                    }

                    catchError(buildResult: 'SUCCESS', stageResult: 'UNSTABLE') {
                        bat 'cd gudang-service && if exist go.mod (go test -v)'
                    }

                    catchError(buildResult: 'SUCCESS', stageResult: 'UNSTABLE') {
                        bat 'cd tracking-service && if exist go.mod (go test -v)'
                    }
                }
            }
        }

        // =========================
        // 3. LINT / VET
        // =========================
        stage('3. Lint / Vet') {
            steps {
                script {

                    catchError(buildResult: 'SUCCESS', stageResult: 'UNSTABLE') {
                        bat 'cd user-service && go vet ./...'
                    }

                    catchError(buildResult: 'SUCCESS', stageResult: 'UNSTABLE') {
                        bat 'cd order-service && go vet ./...'
                    }

                    catchError(buildResult: 'SUCCESS', stageResult: 'UNSTABLE') {
                        bat 'cd report-service && if exist go.mod (go vet ./...)'
                    }

                    catchError(buildResult: 'SUCCESS', stageResult: 'UNSTABLE') {
                        bat 'cd courier-service && if exist go.mod (go vet ./...)'
                    }

                    catchError(buildResult: 'SUCCESS', stageResult: 'UNSTABLE') {
                        bat 'cd gudang-service && if exist go.mod (go vet ./...)'
                    }

                    catchError(buildResult: 'SUCCESS', stageResult: 'UNSTABLE') {
                        bat 'cd tracking-service && if exist go.mod (go vet ./...)'
                    }
                }
            }
        }

        // =========================
        // 4. BUILD IMAGE
        // =========================
        stage('4. Build Docker Images') {
            steps {
                bat 'docker compose build'
            }
        }

        // =========================
        // 5. FUNCTIONAL TEST
        // =========================
        stage('5. Functional Tests') {
            steps {
                script {

                    bat 'docker compose up -d'
                    bat 'timeout /t 10'

                    catchError(buildResult: 'SUCCESS', stageResult: 'UNSTABLE') {
                        bat 'cd user-service && go test -tags=functional -v'
                    }

                    catchError(buildResult: 'SUCCESS', stageResult: 'UNSTABLE') {
                        bat 'cd order-service && go test -tags=functional -v'
                    }

                    catchError(buildResult: 'SUCCESS', stageResult: 'UNSTABLE') {
                        bat 'cd report-service && if exist go.mod (go test -tags=functional -v)'
                    }

                    catchError(buildResult: 'SUCCESS', stageResult: 'UNSTABLE') {
                        bat 'cd courier-service && if exist go.mod (go test -tags=functional -v)'
                    }

                    catchError(buildResult: 'SUCCESS', stageResult: 'UNSTABLE') {
                        bat 'cd gudang-service && if exist go.mod (go test -tags=functional -v)'
                    }

                    catchError(buildResult: 'SUCCESS', stageResult: 'UNSTABLE') {
                        bat 'cd tracking-service && if exist go.mod (go test -tags=functional -v)'
                    }

                    bat 'docker compose down'
                }
            }
        }

        // =========================
        // 6. PUSH IMAGE
        // =========================
        stage('6. Push Images') {
            steps {

                bat 'docker tag tubestahap2-user-service %DOCKER_USER%/user-service:latest'
                bat 'docker tag tubestahap2-order-service %DOCKER_USER%/order-service:latest'
                bat 'docker tag tubestahap2-report-service %DOCKER_USER%/report-service:latest'
                bat 'docker tag tubestahap2-courier-service %DOCKER_USER%/courier-service:latest'
                bat 'docker tag tubestahap2-gudang-service %DOCKER_USER%/gudang-service:latest'
                bat 'docker tag tubestahap2-tracking-service %DOCKER_USER%/tracking-service:latest'

                bat 'docker push %DOCKER_USER%/user-service:latest'
                bat 'docker push %DOCKER_USER%/order-service:latest'
                bat 'docker push %DOCKER_USER%/report-service:latest'
                bat 'docker push %DOCKER_USER%/courier-service:latest'
                bat 'docker push %DOCKER_USER%/gudang-service:latest'
                bat 'docker push %DOCKER_USER%/tracking-service:latest'
            }
        }

        // =========================
        // 7. DEPLOY KUBERNETES
        // =========================
        stage('7. Deploy Kubernetes') {
            steps {
                bat 'kubectl apply -f k8s/'
            }
        }

        // =========================
        // 8. VERIFY
        // =========================
        stage('8. Verify') {
            steps {
                bat 'kubectl get pods -n %K8S_NAMESPACE%'
                bat 'kubectl get svc -n %K8S_NAMESPACE%'
                bat 'kubectl get deployments -n %K8S_NAMESPACE%'
            }
        }
    }

    post {
        always {
            bat 'docker compose down || exit 0'
        }

        success {
            echo "PIPELINE SUCCESS"
        }

        unstable {
            echo "PIPELINE UNSTABLE (ada test gagal)"
        }

        failure {
            echo "PIPELINE FAILED"
        }
    }
}
