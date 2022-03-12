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
    command: ["/busybox/cat"]
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
  stages {
    stage('Stage 1: Build with Kaniko') {
      steps {
        container('kaniko') {
          sh '/kaniko/executor --context=https://github.com/Am2901/httpserver.git \
                  --destination=cloudnative.azurecr.io/httpserver \
                  --insecure \
                  --skip-tls-verify  \
                  -v=debug'
        }
      }
    }  
  }
}
