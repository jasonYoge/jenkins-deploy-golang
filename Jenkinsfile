pipeline {
    agent any

    stages {
        stage('build') {
            steps {
                sh 'docker build -t golang-staging .'
                sh 'docker image ls'
            }
        }

        stage('Deploy') {
            steps {
                echo 'deploy...'
            }
        }
    }
}