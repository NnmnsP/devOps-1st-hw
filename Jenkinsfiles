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
                // Set the Go version to 1.20.4
                GOLANG_VERSION = '1.20.4'
            }
            steps {
                sh 'go version' // Verify the Go version
                sh 'go build -o myapp' // Build the binary artifact
            }
        }

        stage('Test') {
            steps {
                sh 'go test -v ./...' // Run tests
            }
        }
    }

    post {
        success {
            // Archive the binary artifact
            archiveArtifacts artifacts: 'myapp'
        }
    }
}
