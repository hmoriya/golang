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
        sh 'echo $USER'
      }
    }
  }
}