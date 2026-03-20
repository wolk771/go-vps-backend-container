# Go VPS Backend Container

Dieses Projekt dient als **Pilot-System** für eine hochverfügbare und abgesicherte **CI/CD-Infrastruktur** auf einem Debian 13 (Trixie) VPS. Es demonstriert den modernen Workflow von der lokalen Entwicklung bis zum automatisierten Deployment mittels **Jenkins** und **Podman-Compose**.

## 🚀 Key Features

*   **Automatisierter Workflow:** GitHub Webhooks triggern bei jedem Push auf den `main`-Branch eine Jenkins-Pipeline.
*   **Container Orchestrierung:** Migration von manuellen `podman run`-Befehlen zu einer deklarativen Steuerung via `podman-compose`.
*   **Security First:**
    *   **Nginx Reverse Proxy:** Absicherung über SSL (Let's Encrypt) mit automatischer Background-Erneuerung.
    *   **Webhook-Härtung:** Mehrstufige Filterung (HTTP-Methode + GitHub-spezifische Header) direkt im Nginx-Layer und Jenkins-Trigger.
    *   **Netzwerk-Isolation:** Betrieb in einem dedizierten Backend-Netzwerk (`dev_edu-network-dev`) ohne direkte Port-Exponierung nach außen.
*   **Resource Optimized:** Einsatz eines **Multi-Stage Dockerfiles** (Build auf Golang-Alpine, Deployment auf minimalem Alpine-Linux), um die Image-Größe und Angriffsfläche zu minimieren.

## 🛠 Technologie-Stack

*   **Backend:** Go (Golang) 1.21
*   **Infrastruktur:** Podman & Podman-Compose
*   **CI/CD:** Jenkins (Pipeline-as-Code via `Jenkinsfile`)
*   **Webserver:** Nginx (als Reverse Proxy & Security Shield)
*   **OS:** Debian GNU/Linux 13 (trixie)

## 🏗 Architektur-Übersicht

1.  **Local Dev:** Code-Änderungen in VS Code -> Git Push.
2.  **Webhook:** GitHub sendet POST-Request mit Signatur an den VPS.
3.  **Nginx:** Validiert den Request (Method/Header) und leitet ihn intern an Jenkins weiter.
4.  **Jenkins:** Checkt Code aus, baut das Docker-Image und rollt es via `podman-compose` neu aus.
5.  **Runtime:** Die Go-App läuft isoliert und ist nur über den Nginx-Proxy erreichbar.

---
*Dieses Projekt ist Teil eines größeren Ökosystems (Trainer-Buchungssystem) und dient als Blaupause für die Microservice-Infrastruktur.*
