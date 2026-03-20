pipeline {
    agent any
    stages {
        stage('Source & Build Test') {
            steps {
                checkout scm
                // Prüfen, ob das Image überhaupt baubar ist (Sicherheits-Check für PRs)
                sh 'podman-compose build'
            }
        }
        stage('Deploy to VPS') {
            // Nur ausführen, wenn auf dem Hauptzweig main!
            when { branch 'main' }
            steps {
                // Aufräum-Schritt falls Container vorhanden oder läuft 
                sh 'podman stop go-app-final || true'
                sh 'podman rm go-app-final || true'
                // Finales Hochfahren
                sh 'podman-compose up -d'
            }
        }
    }
    post {
        always {
            // Hält den VPS sauber
            sh 'podman image prune -f'
        }
    }
}
