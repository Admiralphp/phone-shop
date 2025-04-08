# Initialiser Swarm
chmod +x init-swarm.sh
./init-swarm.sh

# Déployer la stack
docker stack deploy -c swarm-deploy.yml phone-accessories-app

# Voir l'état
docker stack services phone-accessories-app
docker stack ps phone-accessories-app