pipeline {
    agent any

    stages {
        stage('Source Control') {
            steps {
                // Holt den frischen Code vom GitHub-Branch
                checkout scm
            }
        }

        stage('Podman Build & Deploy') {
            steps {
                // 1. Das neue, kleine Multi-Stage Image bauen
                sh 'podman-compose build'
                
                // 2. Den alten Container stoppen und den neuen im Hintergrund starten
                // up -d kümmert sich automatisch um das Ersetzen des alten Containers
                sh 'podman-compose up -d'
            }
        }

        stage('Cleanup') {
            steps {
                // Löscht alte, ungenutzte Image-Reste (Dangling Images)
                sh 'podman image prune -f'
            }
        }
    }
}
