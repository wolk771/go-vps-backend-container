pipeline {
    agent any
    stages {
        stage('Source Control') {
            steps {
                checkout scm
            }
        }
        stage('Build & Validation') {
            steps {
                sh "echo 'PROZESS STARTET AUF BRANCH: ${env.GIT_BRANCH}'"
                // Sicherheit: ohne Cache, um Code-Inkonsistenzen (main.go) zu vermeiden
                sh 'podman-compose -f docker-compose.yml build --no-cache'
            }
        }
        stage('Safe Deployment') {
            // Nur auf main ausführen
            when { expression { env.GIT_BRANCH == 'origin/main' || env.GIT_BRANCH == 'main' } }
            steps {
                sh "echo 'STARTE PROD-DEPLOYMENT...'"
                
                // 1. Absolute Port-Garantie: Den alten Container hart beenden
                sh 'podman-compose -f docker-compose.yml down || true'
                
                // 2. Verwendet bereits gebauten Image
                sh 'podman-compose -f docker-compose.yml up -d --no-build'
            }
        }
    }
    post {
        always {
            // Befreit VPS von übrigen Images
            sh 'podman image prune -f'
        }
    }
}
