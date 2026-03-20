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
            sh "echo 'DEPLOY BRANCH: ${env.GIT_BRANCH}'"
            // Nur ausführen, wenn auf dem Hauptzweig main!
            when { expression { env.GIT_BRANCH == 'origin/main' || env.GIT_BRANCH == 'main' } }

            steps {
                // Den alten Port-Besetzer wegräumen 
                sh 'podman-compose -f docker-compose.yml down || true'
                sh 'podman-compose -f docker-compose.yml build --no-cache'
                // Finales Hochfahren und --force-recreate
                sh 'podman-compose -f docker-compose.yml up -d --force-recreate'
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
