pipeline {
    agent {
        label 'common'
    }
    environment {
        GITHUB_URL = "http://github.com"
        GITHUB_ORG = "Am2901"
        GITHUB_REPO = "httpserver"
    }
    options {
        skipDefaultCheckout()
        ansiColor('xterm')
    }
    stages {
        stage('Checkout') {
            steps {
                script{
                    def scmVars =   checkout(
                                        [$class: 'GitSCM', branches: [[name: "${ghprbActualCommit}"]], 
                                        doGenerateSubmoduleConfigurations: false,
                                        submoduleCfg: [], 
                                        extensions: [
                                            [$class: 'RelativeTargetDirectory', relativeTargetDir: 'codes'],
                                            [$class: 'CleanBeforeCheckout']
                                        ],
                                        userRemoteConfigs: [
                                                [
                                                    credentialsId: 'ghe_account', 
                                                    name: 'origin', 
                                                    refspec: '+refs/pull/*:refs/remotes/origin/pr/*', 
                                                    url: "${GITHUB_URL}/${GITHUB_ORG}/${GITHUB_REPO}.git"
                                                ]
                                            ]
                                        ]
                                    )
                    env.GIT_BRANCH = "${scmVars.GIT_BRANCH}"
                    env.GIT_COMMIT = "${scmVars.GIT_COMMIT}"
                }
            }
        }
        stage('Build') {
            steps {
                dir('codes') {
                    sh '''#!/bin/bash -l
                        echo "Start building!"
                        cat README.md
                    '''
                }
            }
        }


    }
    post {
        always {
            junit allowEmptyResults: true, testResults: 'codes/build/test-results/test/**/*.xml'
            // send to email
            emailext(
                subject: '$DEFAULT_SUBJECT',
                body: '$DEFAULT_CONTENT',
                recipientProviders: [
                [$class: 'CulpritsRecipientProvider'],
                [$class: 'RequesterRecipientProvider'],
                [$class : 'DevelopersRecipientProvider']
                ],
                replyTo: '$DEFAULT_REPLYTO',
                to: '$DEFAULT_RECIPIENTS'
            )
        }
    }
}
