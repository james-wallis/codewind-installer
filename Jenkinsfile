#!groovy​

pipeline {

    agent {
        kubernetes {
              label 'go-pod'
            yaml """
apiVersion: v1
kind: Pod
spec:
  containers:
  - name: go
    image: golang:1.11-stretch
    tty: true
    command:
    - cat
    resources:
      limits:
        memory: "2Gi"
        cpu: "1"
      requests:
        memory: "2Gi"
        cpu: "1"
"""
        }
    }

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
                container('go') {
                    sh '''
                        echo "Script success"

                    '''
                }
            }
        }

    }

    post {
        success {
            echo 'Build SUCCESS'
        }
    }
}