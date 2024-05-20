pipeline {
    agent {
        docker {
            image 'golang:1.12'
        }
    }

    environment {
        GOPATH = "/go"
        GOCACHE = "${WORKSPACE}/.gocache"
        PATH = "/go/bin:${env.PATH}"
    }

    stages {
        stage('Lint') {
            steps {
                sh '''
                    go get -u golang.org/x/lint/golint
                    golint -set_exit_status ${WORKSPACE}/components/operator/...
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
