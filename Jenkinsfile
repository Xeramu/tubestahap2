pipeline {
    agent any

    environment {
        APP_NAME = "tubestahap2"
        DOCKER_USER = "umarx"
    }

    stages {

        stage('Checkout Repo') {
            steps {
                git 'https://github.com/Xeramu/tubestahap2.git'
            }
        }

        stage('Unit Test') {
            steps {
                dir('user-service') {
                    sh 'go test -v'
                }

                dir('order-service') {
                    sh 'go test -v'
                }
            }
        }

        stage('Lint / Vet') {
            steps {
                dir('user-service') {
                    sh 'go vet ./...'
                }

                dir('order-service') {
                    sh 'go vet ./...'
                }
            }
        }

        stage('Build Docker Image') {
            steps {
                sh 'docker compose build'
            }
        }

        stage('Functional Test') {
            steps {
                sh 'docker compose up -d'
                sh 'docker compose exec order-service go test -v'
            }
        }

        stage('Push Image') {
            steps {
                sh 'docker tag tubestahap2-user-service ${DOCKER_USER}/user-service:latest'
                sh 'docker tag tubestahap2-order-service ${DOCKER_USER}/order-service:latest'

                sh 'docker push ${DOCKER_USER}/user-service:latest'
                sh 'docker push ${DOCKER_USER}/order-service:latest'
            }
        }

        stage('Deploy Kubernetes') {
            steps {
                sh 'kubectl apply -f k8s/'
            }
        }

        stage('Verify') {
            steps {
                sh 'kubectl get pods'
                sh 'kubectl get svc'
            }
        }
    }
}