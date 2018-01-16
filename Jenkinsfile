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
        echo $USER
        sh 'docker build -t jenkins-master .'
      }
    }
  }
}
