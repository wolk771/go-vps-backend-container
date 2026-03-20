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
                sh 'podman rm -f go-app-final || true'
                // Finales Hochfahren und --force-recreate
                sh 'podman-compose up -d --force-recreate'
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
