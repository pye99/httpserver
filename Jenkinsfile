podTemplate(containers: [
    containerTemplate(name: 'golang', image: 'golang:1.17', ttyEnabled: true, command: 'cat'),
  ],
  yaml: """\
apiVersion: v1
kind: Pod
metadata:
  name: kaniko
spec:
  containers:
  - name: kaniko
    image: gcr.io/kaniko-project/executor
    command:
    - /busybox/cat
    tty: true
    volumeMounts:
      - name: kaniko-secret
        mountPath: /kaniko/.docker
  restartPolicy: Never
  volumes:
    - name: kaniko-secret
      secret:
        secretName: regcred
    """.stripIndent()
  ) {
    node(POD_LABEL) {
        stage('Clone') {
            git url: '{{REPO}}'
        }
        stage('Compile') {
            container('golang') {
                    sh """
                    make  
                    """
            }
        }
        stage('Build Image')
            container('kaniko') {
                sh """
                /kaniko/executor -c `pwd`/ -f `pwd`Dockerfile -d {{IMAGE}}
                """
            }
    }
}
