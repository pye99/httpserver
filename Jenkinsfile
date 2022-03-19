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
    HELM_VERSION = "${BUILD_NUMBER}"
  }
  stages {
    stage('Configure') {
      steps {
        echo "hello, starting"
      }
    }
    stage('image build and push') {
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
    stage('Helm Package and Push') {
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
  - name: helm
    image: alpine/helm:3.8.1
    command:
    - sleep
    args:
    - 99999
    tty: true
    volumeMounts:
      - name: docker-secret
        mountPath: /root/.config/helm/registry
        readOnly: true
  volumes:
  - name: docker-secret
    secret:
      secretName: regcred
"""
    }
  }
      steps {
        container('helm'){
            sh "helm package `pwd`/deploy/helm/httpserver --app-version ${DATED_GIT_HASH} --version 0.1.${HELM_VERSION}"
            sh "helm push `pwd`/httpserver-0.1.${HELM_VERSION}.tgz oci://cloudnative.azurecr.io/repo/stable"
        }
      }
    }
  }
}