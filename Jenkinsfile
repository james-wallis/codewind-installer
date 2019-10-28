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
    }

    stages {

        stage ('Build') {
            steps {
                sh '''
                    echo "Script success"

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