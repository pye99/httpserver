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
    stage('Configure') {
      steps {
        script {    
          GIT_COMMIT_REV = sh(returnStdout: true, script: "git log -n 1 --pretty=format:'%h'").trim()
        }
      }
    }    
    stage('Build with Kaniko') {
      steps {
        container('kaniko') {
          sh '/kaniko/executor -f `pwd`/Dockerfile -c `pwd`/src --cache=true \
                  --destination=cloudnative.azurecr.io/httpserver:$GIT_COMMIT_REV \
                  --insecure \
                  --skip-tls-verify  \
                  -v=debug'
        }
      }
    }  
  }
}
