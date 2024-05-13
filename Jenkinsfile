pipeline {
    agent {
        docker {
            image 'golang:1.22.1'
            args '-v /usr/local/go/pkg:/usr/local/go/pkg'
        }
    }

    environment {
        GOROOT = '/usr/local/go' 
        GOPATH = '/go'
        GOCACHE = '/go/go-cache'
        PATH = "${env.PATH}:${env.GOROOT}/bin:${env.GOPATH}/bin"
    }

    stages {
        stage('Operator Lint') {
            steps {
                checkout scm
                sh 'which go'
                sh 'go version'
                sh 'go env'
                sh '''
            curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b /go/bin v1.58.1
            golangci-lint --version
        '''
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
