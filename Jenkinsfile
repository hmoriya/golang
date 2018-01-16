pipeline {
  agent any
  stages {
    stage('Test') {
      steps {
        echo 'Hello World'
      }
    }
    stage('ls') {
      steps {
        echo '$USER'
        sh 'docker build -t jenkins-master .'
      }
    }
  }
}