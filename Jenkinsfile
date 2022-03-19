import java.text.SimpleDateFormat;

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
      - name: docker-secret
        mountPath: /kaniko/.docker
        readOnly: true
  volumes:
  - name: docker-secret
    secret:
      secretName: regcred
"""
    }
  }
  environment {
    DATED_GIT_HASH = "${new SimpleDateFormat("yyMMddHHmmss").format(new Date())}${GIT_COMMIT.take(6)}"
  }
  stages {
    stage('Configure') {
      steps {
        echo "hello, starting"
      }
    }
    stage('Build with Kaniko and Push') {
      steps {
        container('kaniko') {
          sh '/kaniko/executor -f `pwd`/Dockerfile -c `pwd`/src --cache=true \
          --destination=cloudnative.azurecr.io/httpserver:${DATED_GIT_HASH} \
                  --insecure \
                  --skip-tls-verify  \
                  -v=debug'
        }
      }
    }
    stage('Helm') {
     agent {
       kubernetes {
             containerTemplate {
               name 'helm'
               image 'alpine/helm:3.1.1'
               ttyEnabled true
               command 'cat'
          }
        }
     }
      steps {
        container('helm'){
            sh "ls -l `pwd`/*/*/*/*/*"
            sh "helm package `pwd`/deploy/helm/httpserver"
        }
      }
    }
  }
}