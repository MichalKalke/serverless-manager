pipeline {
    agent {
        docker {
            image 'golang:latest'
            args '-v $HOME/go/pkg:/usr/local/go/pkg'  // Align home dir
        }
    }

     environment {
        CGO_ENABLED = '0'
        GO111MODULE = 'on'
        GOPATH = '/go' 
        GOROOT = '/usr/local/go' // Added this line
        GOLANGCI_LINT_CACHE = '/go/go-cache'
        PATH = "/go/bin:/usr/local/go/bin:${env.PATH}"
    }

    stages {
        stage('Operator Lint') {
            steps {
                checkout scm
                script {  
                   sh 'go version'
                   sh '''
                      curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin latest
                      golangci-lint --version
                      golangci-lint run components/operator/
                  '''
                }
            }
        }

        stage('Operator Unit Tests') {
            steps {
                script {
                   sh '''
                      cd components/operator
                      go test ./...
                   '''
                }
            }
        }
    }
}
