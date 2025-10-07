pipeline {
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
                        serviceAccountName: jenkins-agent
                        containers:
                          - name: helm
                            image: alpine/helm:3
                            command:
                              - sleep
                            args:
                              - 9999999
                      '''
                }
            }
            steps {
                container('helm') {
                    sh '''
                      helm upgrade ginexample ./mychart/ \
                      --install \
                      --wait \
                      --timeout=5m0s \
                      --set image.tag=${BUILD_TAG}
                      '''
                }
            }
        }
    }
}
