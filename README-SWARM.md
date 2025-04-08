# 📦 Phone Accessories Microservices Platform

Microservices-based e-commerce platform for selling phone accessories. This setup includes services for authentication, product management, ordering, and administration — all deployed using Docker Swarm.

---

## 🚀 3. Commandes de Déploiement avec Docker Swarm

### 🐳 Initialisation du cluster Swarm

```bash
# Rendre le script exécutable
chmod +x init-swarm.sh

# Lancer l'initialisation du cluster
./init-swarm.sh
```

Ce script :
- Initialise un cluster Docker Swarm.
- Crée un réseau overlay `microservices_net`.
- Affiche la commande `docker swarm join` pour connecter les nœuds secondaires.

---

### 📦 Déploiement de la stack

```bash
# Déployer l'application avec Swarm
docker stack deploy -c swarm-deploy.yml phone-accessories-app
```

---

### 🔍 Vérifier l’état de la stack

```bash
# Voir les services déployés
docker stack services phone-accessories-app

# Voir les tâches en cours (containers)
docker stack ps phone-accessories-app
```

---

## 🧼 Nettoyage (optionnel)

```bash
# Supprimer la stack
docker stack rm phone-accessories-app

# Supprimer le réseau Swarm (si plus utilisé)
docker network rm microservices_net
```

---

## 📁 Structure des Fichiers Importants

```
.
├── init-swarm.sh              # Script pour initier Docker Swarm
├── swarm-deploy.yml           # Fichier de déploiement Swarm
├── nginx/
│   └── nginx.conf             # Configuration du Gateway
├── README.md                  # Ce fichier
```

---

## 📝 Notes

- Le volume de chaque base de données est persisté via `docker volume`.
- Les services critiques sont répliqués pour assurer la haute disponibilité.
- Le réseau Swarm permet la communication sécurisée entre services.

---
