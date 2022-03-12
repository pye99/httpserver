pipeline {
  agent {
    kubernetes {
      yaml """
apiVersion: v1
kind: Pod
metadata:
  labels:
    jenkins: worker
spec:
  containers:
  - name: kaniko
    image: gcr.io/kaniko-project/executor:debug
    command:
    - sleep
    args:
    - 99999
    tty: true
    volumeMounts:
      - name: dockercred
        mountPath: /root/.docker/
  volumes:
  - name: dockercred
    secret:
      secretName: regcred
"""
    }
  }
  environment {
      GITHUB_URL = "http://github.com"
      GITHUB_ORG = "Am2901"
      GITHUB_REPO = "httpserver"
  }
  stages {
    stage('Stage 1: Build with Kaniko') {
      steps {
        container('kaniko') {
          sh '/kaniko/executor -f `pwd`/Dockerfile -c `pwd`/src --cache=true \
                  --destination=cloudnative.azurecr.io/httpserver:$CI_COMMIT_TAG \
                  --insecure \
                  --skip-tls-verify  \
                  -v=debug'
        }
      }
    }  
  }
}
