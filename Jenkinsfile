pipeline {
    agent {
        docker {
            image 'golang:1.22.1'
            args '-v /usr/local/go/pkg:/usr/local/go/pkg'
        }
    }

    environment {
        GOROOT = '/usr/local/go' // Set GOROOT here
        GOPATH = "${WORKSPACE}/go" // go get installs to this directory
        GOCACHE = '/tmp/go-cache' // Set GOCACHE here
        PATH = "${WORKSPACE}/go/bin:${env.GOROOT}/bin:${env.PATH}" // Add to PATH
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
