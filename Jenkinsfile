pipeline {
    agent any

    environment {
        DOCKER_USER = "umarx"
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

                catchError(buildResult: 'SUCCESS', stageResult: 'UNSTABLE') {
                    bat 'cd user-service && go test -v'
                }

                catchError(buildResult: 'SUCCESS', stageResult: 'UNSTABLE') {
                    bat 'cd order-service && go test -v'
                }
            }
        }

        // =========================
        // 3. LINT / VET
        // =========================
        stage('3. Lint / Vet') {
            steps {

                catchError(buildResult: 'SUCCESS', stageResult: 'UNSTABLE') {
                    bat 'cd user-service && go vet ./...'
                }

                catchError(buildResult: 'SUCCESS', stageResult: 'UNSTABLE') {
                    bat 'cd order-service && go vet ./...'
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

                    // start semua container
                    bat 'docker compose up -d'

                    // tunggu mysql healthy
                    powershell 'Start-Sleep -Seconds 20'

                    // cek container
                    bat 'docker compose ps'

                    // =========================
                    // USER SERVICE FUNCTIONAL TEST
                    // =========================
                    catchError(buildResult: 'SUCCESS', stageResult: 'UNSTABLE') {
                        bat 'docker compose exec user-service go test -tags=functional -v'
                    }

                    // =========================
                    // ORDER SERVICE FUNCTIONAL TEST
                    // =========================
                    catchError(buildResult: 'SUCCESS', stageResult: 'UNSTABLE') {
                        bat 'docker compose exec order-service go test -tags=functional -v'
                    }
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

                bat 'docker push %DOCKER_USER%/user-service:latest'
                bat 'docker push %DOCKER_USER%/order-service:latest'
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
                bat 'kubectl get pods'
                bat 'kubectl get svc'
                bat 'kubectl get deployments'
            }
        }
    }

    post {

        always {
            bat 'docker compose down'
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
