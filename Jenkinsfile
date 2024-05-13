pipeline {
    agent {
        docker {
            image 'golang:1.22.1'
            args '-v /usr/local/go/pkg:/usr/local/go/pkg'
        }
    }

    stages {
        stage('Operator Lint') {
            steps {
                checkout scm
                sh 'go get -u github.com/golangci/golangci-lint/cmd/golangci-lint'
                dir('components/operator') {
                    sh 'golangci-lint run ./...'
                }
            }
        }

        stage('Operator Unit Tests') {
            steps {
                checkout scm
                dir('components/operator') {
                    sh 'make test'
                }
            }
        }
    }
}
