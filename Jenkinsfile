pipeline {
    // Instructs Jenkins to allocate an executor and workspace for the Pipeline, 
    // ensures that the source repository is checked out and made available for steps in the subsequent stages.
    agent none

    environment {
        REGISTRY = '192.168.31.162:5000'
        REPO = 'lixianwen/ginexample'
    }

    stages {
        stage('Build Image') {
            agent {
                kubernetes {
                    yamlFile 'dind.yaml'
                }
            }
            steps {
                container('dind') {
                    script {
                        docker.withRegistry("http://${env.REGISTRY}") {
                            docker.build(env.REPO).push(env.BUILD_TAG)
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
                            command:
                              - sleep
                            args:
                              - 9999999
                      '''
                }
            }
            steps {
                container('go') {
                    sh 'go test $(go list ./... | grep -v /vendor/)'
                }
            }
        }
        stage('Deploy') {
            agent {
                kubernetes {
                    yaml '''
                      apiVersion: v1
                      kind: Pod
                      spec:
                        containers:
                          - name: kubectl
                            image: lachlanevenson/k8s-kubectl
                            command:
                              - sleep
                            args:
                              - 9999999
                      '''
                }
            }
            steps {
                container('kubectl') {
                    sh '''
                      sed -i "s|${REGISTRY}/${REPO}.*|${REGISTRY}/${REPO}:${BUILD_TAG}|" deployment.yaml
                      kubectl apply -f deployment.yaml
                      '''
                }
            }
        }
    }
}
