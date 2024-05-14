pipeline {
    agent {
        docker {
            image 'golang:latest'
        }
    }

    environment {
        GOPATH = "/go"
        PATH = "/go/bin:${env.PATH}"
    }

    stages {
        stage('Lint') {
            steps {
                sh '''
                    ls
                    go get -u golang.org/x/lint/golint
                    ls
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
