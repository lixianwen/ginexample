pipeline {
    // Instructs Jenkins to allocate an executor and workspace for the Pipeline, 
    // ensures that the source repository is checked out and made available for steps in the subsequent stages.
    agent {
        kubernetes {
            yamlFile 'dind.yaml'
        }
    }

    stages {
        stage('Build Image') {
            steps {
                container('dind') {
                    script {
                        docker.withRegistry('http://192.168.31.162:5000') {
                            docker.build('lixianwen/ginexample').push(env.BUILD_TAG)
                        }
                    }
                }
            }
        }
        stage('Test') {
            agent {
                kubernetes {
                    yaml '''
                      apiVersion: v1
                      kind: Pod
                      spec:
                        containers:
                          - name: go
                            image: golang:1.24-alpine
                      '''
                }
            }
            steps {
                container('go') {
                    sh 'go test $(go list ./... | grep -v /vendor/)'
                }
            }
        }
    }
}
