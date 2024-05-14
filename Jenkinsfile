pipeline {
    agent any

    environment {
        GOPATH = "${WORKSPACE}"
        GOROOT = "/opt/homebrew/opt/go/"  // or wherever you have Go installed
    }

    stages {
        stage('Lint') {
            steps {
                sh '''
                    mkdir $WORKSPACE/temp
                    cd $WORKSPACE/temp
                    go get -u golang.org/x/lint/golint
                    cd $WORKSPACE
                    ${GOPATH}/temp/bin/golint -set_exit_status ${WORKSPACE}/components/operator/...
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
}
