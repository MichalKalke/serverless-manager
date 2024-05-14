pipeline {
    agent {
        docker {
            image 'golang:latest'
        }
    }
    environment {
        GOLANGCI_LINT_CACHE = '/tmp'
    }
    stages {
        stage('Operator Lint') {
            steps {
                checkout scm
                script {
                    sh '''
                    curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin latest
                    golangci-lint run components/operator/
                    '''
                }
            }
        }
    }
}
