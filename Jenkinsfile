pipeline {
    agent any

    stages {

        stage('1. Checkout Repo') {
            steps {
                checkout scm
            }
        }

        stage('2. Build Docker Images') {
            steps {
                bat 'docker compose build'
            }
        }

        stage('3. Start Services') {
            steps {
                bat 'docker compose up -d'

                // debug
                bat 'docker compose ps'
                bat 'docker compose logs mysql'
            }
        }

        stage('4. Wait Until Ready') {
            steps {
                powershell '''
                Start-Sleep -Seconds 20
                '''
            }
        }

        stage('5. Run Tests Inside Containers') {
            steps {

                // user-service
                catchError(buildResult: 'UNSTABLE', stageResult: 'FAILURE') {
                    bat 'docker compose exec user-service go test -v ./...'
                }

                // order-service
                catchError(buildResult: 'UNSTABLE', stageResult: 'FAILURE') {
                    bat 'docker compose exec order-service go test -v ./...'
                }
            }
        }

        stage('6. Show Logs If Crash') {
            steps {
                bat 'docker compose logs user-service'
                bat 'docker compose logs order-service'
            }
        }

        stage('7. Push Images') {
            steps {
                bat 'docker tag tubestahap2-pipeline-user-service umarx/user-service:latest'
                bat 'docker tag tubestahap2-pipeline-order-service umarx/order-service:latest'

                bat 'docker push umarx/user-service:latest'
                bat 'docker push umarx/order-service:latest'
            }
        }

        stage('8. Deploy Kubernetes') {
            steps {
                bat 'kubectl apply -f k8s/'
            }
        }
    }

    post {
        always {
            bat 'docker compose down'
        }

        success {
            echo 'PIPELINE SUCCESS'
        }

        failure {
            echo 'PIPELINE FAILED'
        }
    }
}
