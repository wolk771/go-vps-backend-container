pipeline {
    agent any

    stages {
        stage('Source Control') {
            steps {
                // Jenkins holt sich hier den neuesten Code
                checkout scm
            }
        }

        stage('Podman Build') {
            steps {
                // das Image neu bauen
                sh 'podman build -t go-app-v2 .'
            }
        }

        stage('Deploy App') {
            steps {
                // Zuerst alten Container wegräumen
                // Das "|| true" verhindert, dass die Pipeline stoppt, falls kein Container da ist
                sh 'podman stop go-app-final || true'
                sh 'podman rm go-app-final || true'
                
                // Jetzt den neuen Container auf deinem Port 9001 starten
                sh 'podman run -d --name go-app-final -p 9001:9001 go-app-v2'
            }
        }
    }
}
