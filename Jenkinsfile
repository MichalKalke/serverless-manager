pipeline {
    agent {
        docker {
            image 'golang:1.22.1'
            args '-v /usr/local/go/pkg:/usr/local/go/pkg'
        }
    }

    environment {
        GOROOT = '/usr/local/go' 
        GOPATH = '/Users/I571889/.jenkins/workspace/serverless-cicd@2/go' 
        GOCACHE = '/tmp/go-cache'
        PATH = "${env.PATH}:${env.GOROOT}/bin:${env.GOPATH}/bin"
    }
    
    stages {
        stage('Operator Lint') {
            steps {
                checkout scm
                sh 'which go'
                sh 'go version'
                sh 'go env'
                sh 'go get -v -u github.com/golangci/golangci-lint/cmd/golangci-lint'
                sh 'ls -al "${GOPATH}/bin"'
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
