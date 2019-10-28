#!groovy​

pipeline {

    agent none

    options {
        timestamps()
        skipStagesAfterUnstable()
    }

    environment {
        CODE_DIRECTORY_FOR_GO = 'src/github.com/eclipse/codewind-installer'
        DEFAULT_WORKSPACE_DIR_FILE = 'temp_default_dir'
        JENKINS_HOME = "/var/lib/jenkins/workspace/codewind-cli-testing/"
    }

    stages {

        stage('Build') {
            agent any
            steps {
                sh '''#!/usr/bin/env bash
                            
                            echo success
                            



                        '''

        
            }
        }

    }

    post {
        success {
            echo 'Build SUCCESS'
        }
    }
}

















