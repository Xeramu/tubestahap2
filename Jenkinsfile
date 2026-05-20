pipeline {
    agent any

    environment {
        USER_IMAGE  = "umarx/user-service:latest"
        ORDER_IMAGE = "umarx/order-service:latest"
    }

    stages {

        stage('1. Checkout Repo') {
            steps {
                checkout scm
            }
        }

        stage('2. Unit Tests') {
            steps {

                catchError(buildResult: 'UNSTABLE', stageResult: 'FAILURE') {
                    bat '''
                    cd user-service
                    go test -v
                    '''
                }

                catchError(buildResult: 'UNSTABLE', stageResult: 'FAILURE') {
                    bat '''
                    cd order-service
                    go test -v
                    '''
                }

            }
        }

        stage('3. Lint / Vet') {
            steps {

                catchError(buildResult: 'UNSTABLE', stageResult: 'FAILURE') {
                    bat '''
                    cd user-service
                    go vet ./...
                    '''
                }

                catchError(buildResult: 'UNSTABLE', stageResult: 'FAILURE') {
                    bat '''
                    cd order-service
                    go vet ./...
                    '''
                }

            }
        }

        stage('4. Build Docker Images') {
            steps {

                bat '''
                docker compose build
                '''

            }
        }

        stage('5. Functional Tests') {
            steps {

                script {

                    bat '''
                    docker compose up -d
                    '''

                    powershell '''
                    Start-Sleep -Seconds 60
                    '''

                    bat '''
                    docker compose ps
                    '''

                    catchError(buildResult: 'UNSTABLE', stageResult: 'FAILURE') {
                        bat '''
                        docker compose exec user-service go test -tags=functional -v
                        '''
                    }

                    catchError(buildResult: 'UNSTABLE', stageResult: 'FAILURE') {
                        bat '''
                        docker compose exec order-service go test -tags=functional -v
                        '''
                    }

                }

            }
        }

        stage('6. Push Images') {
            when {
                expression { currentBuild.currentResult != 'FAILURE' }
            }

            steps {

                bat '''
                docker tag tubestahap2-pipeline-user-service %USER_IMAGE%
                '''

                bat '''
                docker tag tubestahap2-pipeline-order-service %ORDER_IMAGE%
                '''

                bat '''
                docker push %USER_IMAGE%
                '''

                bat '''
                docker push %ORDER_IMAGE%
                '''

            }
        }

        stage('7. Deploy Kubernetes') {
            when {
                expression { currentBuild.currentResult != 'FAILURE' }
            }

            steps {

                bat '''
                kubectl apply -f k8s/ --validate=false
                '''

            }
        }

        stage('8. Verify') {
            when {
                expression { currentBuild.currentResult != 'FAILURE' }
            }

            steps {

                bat '''
                kubectl get pods
                '''

                bat '''
                kubectl get svc
                '''

            }
        }

    }

    post {

        always {

            bat '''
            docker compose down
            '''

        }

        success {
            echo 'PIPELINE SUCCESS'
        }

        unstable {
            echo 'PIPELINE UNSTABLE'
        }

        failure {
            echo 'PIPELINE FAILED'
        }

    }
}
