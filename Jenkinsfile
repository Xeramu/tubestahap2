pipeline {
    agent any

    environment {
        APP_NAME = "tubestahap2"
        DOCKER_USER = "umarx"
    }

    stages {

        stage('Unit Test') {
            steps {
                bat 'cd user-service && go test -v'
                bat 'cd order-service && go test -v'
            }
        }

        stage('Lint / Vet') {
            steps {
                bat 'cd user-service && go vet ./...'
                bat 'cd order-service && go vet ./...'
            }
        }

        stage('Build Docker Image') {
            steps {
                bat 'docker compose build'
            }
        }
        
        stage('Functional Test') {
            steps {
                bat 'docker compose up -d'
                bat 'docker compose exec order-service go test -tags=functional -v'
            }
        }
        
        stage('Push Image') {
            steps {
                bat 'docker tag tubestahap2-user-service %DOCKER_USER%/user-service:latest'
                bat 'docker tag tubestahap2-order-service %DOCKER_USER%/order-service:latest'

                bat 'docker push %DOCKER_USER%/user-service:latest'
                bat 'docker push %DOCKER_USER%/order-service:latest'
            }
        }

        stage('Deploy Kubernetes') {
            steps {
                bat 'kubectl apply -f k8s/'
            }
        }

        stage('Verify') {
            steps {
                bat 'kubectl get pods'
                bat 'kubectl get svc'
            }
        }
    }
}
