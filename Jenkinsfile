pipeline {
    agent any
    stages {
        stage('Source Control') {
            steps {
                checkout scm
            }
        }
        stage('Build Image') {
            steps {
                // Prüfen, ob das Image überhaupt baubar ist (Sicherheits-Check für PRs)
                sh 'podman-compose build --no-cache'
            }
        }
        stage('Deploy to VPS') {
            // Nur ausführen, wenn auf dem Hauptzweig main!
            when { branch 'main' }
            steps {
                // Den alten Port-Besetzer wegräumen 
                sh 'podman stop go-app-final || true'
                sh 'podman rm go-app-final || true'
                // Finales Hochfahren
                sh 'podman-compose up -d --no-build'
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
