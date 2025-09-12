pipeline {
    // Instructs Jenkins to allocate an executor and workspace for the Pipeline, 
    // ensures that the source repository is checked out and made available for steps in the subsequent stages.
    agent any

    environment {
        GITHUB_PERSIONAL_ACCESS_TOKEN = credentials('GitHub-PAT')        
    }

    parameters {
       string(name: 'DEPLOY_ENV', defaultValue: 'staging', description: 'whatever')
       text(name: 'DEPLOY_INFO', defaultValue: 'One\nTwo\nThree\n', description: 'What can I say?')
       booleanParam(name: 'IS_ADMIN', defaultValue: false, description: '')
       choice(name: 'PLATFORM', choices: ['Linux', 'Windows', 'MacOS'], description: '')
       password(name: 'PASSWORD', defaultValue: 'SECRET', description: 'A secret password')
    }

    stages {
        stage('Test') {
            steps {
                echo 'Testing..'
            }
        }
        stage('Build') {
            environment {
                CGO_ENABLED = 0
            }
            steps {
                // sh 'go build -o myapp'
                // archiveArtifacts artifacts: 'myapp', fingerprint: true
                echo env.BRANCH_NAME
                echo "$CGO_ENABLED"
                sh('echo $GITHUB_PERSIONAL_ACCESS_TOKEN')
                sh('echo $GITHUB_PERSIONAL_ACCESS_TOKEN_USR')
                sh('echo $GITHUB_PERSIONAL_ACCESS_TOKEN_PSW')
            }
        }
        stage('Deploy') {
            steps {
                echo params.DEPLOY_ENV
                echo "DEPLOY_INFO: $params.DEPLOY_INFO"
                echo "$params.IS_ADMIN"
                echo "Which platform did you chose: ${params.PLATFORM}"
                echo "What is your password: ${params.PASSWORD}"
            }
        }
    }
}

