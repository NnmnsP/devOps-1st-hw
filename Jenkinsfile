pipeline {
    agent any

    stages {
        stage('Fetch Source Code') {
            steps {
                git 'https://github.com/NnmnsP/devOps-1st-hw'
            }
        }

        stage('Build') {
            environment {
                GO_VERSION = '1.20.4'
            }
            steps {
                sh '''
                    # Install Go
                    curl -O https://dl.google.com/go/go$GO_VERSION.linux-amd64.tar.gz
                    tar -xvf go$GO_VERSION.linux-amd64.tar.gz
                    export PATH=$PATH:/path/to/go/bin

                    # Build the Go application
                    go build -o app
                '''
            }
        }

        stage('Test') {
            steps {
                sh '''
                    # Run the tests
                    go test -v ./...
                '''
            }
        }

        stage('Produce Artefact') {
            steps {
                sh 'cp app /path/to/artefacts'
            }
        }
    }
}
