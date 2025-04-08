#!/bin/bash

# Initie le Swarm sur le manager
docker swarm init

# Affiche la commande à exécuter sur les workers
echo " Pour ajouter des workers, exécutez cette commande sur les nœuds :"
docker swarm join-token worker

# Crée le réseau overlay pour la communication inter-service
docker network create \
  --driver overlay \
  --attachable \
  microservices_net
