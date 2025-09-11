pipeline {
    // Instructs Jenkins to allocate an executor and workspace for the Pipeline, 
    // ensures that the source repository is checked out and made available for steps in the subsequent stages.
    agent any

    stages {
        stage('Test') {
            steps {
                echo 'Testing..'
            }
        }
        stage('Build') {
            steps {
                sh 'go build -o myapp'
                archiveArtifacts artifacts: 'myapp', fingerprint: true
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying....'
            }
        }
    }
}
