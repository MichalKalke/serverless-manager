pipeline {
    agent {
        docker {
            image 'golang:1.22.1'
            args '-v /usr/local/go/pkg:/usr/local/go/pkg'
        }
    }

    environment {
        GOROOT = '/usr/local/go' 
        GOPATH = "${WORKSPACE}/go" 
        GOCACHE = '/tmp/go-cache'
    }
    
    stages {
        stage('Operator Lint') {
            steps {
                checkout scm
                sh 'echo ${WORKSPACE}'
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
