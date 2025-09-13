pipeline {
    // Instructs Jenkins to allocate an executor and workspace for the Pipeline, 
    // ensures that the source repository is checked out and made available for steps in the subsequent stages.
    agent {
        kubernetes {
            yaml '''
              apiVersion: v1
              kind: Pod
              spec:
                containers:
                - name: dind
                  image: docker:dind
                  imagePullPolicy: Always
                  args:
                    - "--insecure-registry=192.168.31.162:5000"
                  securityContext:
                    privileged: true
              '''
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
    }
}
