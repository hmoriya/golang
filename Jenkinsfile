pipeline {
    agent any
    stages {
        stage('Test') {
            steps {
                echo 'Hello World'
            }
        }
	stage('GoVersion'){
	    steps {
		sh 'go --version'
		}
	}
    }
}
