pipeline {
    agent any
    environment {
        GITHUB_URL = "http://github.com"
        GITHUB_ORG = "Am2901"
        GITHUB_REPO = "httpserver"
    }
    stages {
        stage('Checkout') {
            steps {
               echo "build"
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
                                                    credentialsId: 'github-password', 
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
               echo "build"
            }
        }

        stage('Deploy') {
            steps {
                echo 'Deployment in progress'
            }
        }
    }
}
