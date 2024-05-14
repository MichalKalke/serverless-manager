pipeline {
    agent any

    environment {
        GOPATH = "${WORKSPACE}"
    }

    stages {
        stage('Lint') {
            steps {
                sh ''' 
                    go get -u golang.org/x/lint/golint
                    ${GOPATH}/bin/golint -set_exit_status ${WORKSPACE}/components/operator/...
                ''' 
            }
        }
    }
    
    post {
        failure {
            echo 'Linting process failed. Fix above issues.'
        }
        success {
            echo 'Linting process passed.'
        }
    }
