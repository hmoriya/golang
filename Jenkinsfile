pipeline {
  agent any
  stages {
    stage('Test') {
      steps {
        echo 'Hello World'
      }
    }
    stage('docker build') {
      steps {
        sh 'ls -la'
        sh 'pwd'
        sh 'who'
        sh 'docker build -t jenkins-master .'
      }
    }
  }
}