#!/bin/bash
set -e

# Usage: ./deploy-image.sh <USER> <HOST> <SSH_KEY_PATH> <TARGET_ENV>

SSH_USER=$1
SSH_HOST=$2
SSH_KEY_PATH=$3
TARGET_ENV=$4 # "develop" or "production"

if [ -z "$SSH_USER" ] || [ -z "$SSH_HOST" ] || [ -z "$SSH_KEY_PATH" ] || [ -z "$TARGET_ENV" ]; then
  echo "Usage: $0 <USER> <HOST> <SSH_KEY_PATH> <TARGET_ENV>"
  exit 1
fi

echo "--------------------------------------------------"
echo "🚀 Starting Deployment to $TARGET_ENV"
echo "Target: $SSH_USER@$SSH_HOST"
echo "--------------------------------------------------"

# Ensure strict host key checking is off for CI/CD automation
# We use a custom ssh command wrapper for convenience
SSH_CMD="ssh -o StrictHostKeyChecking=no -i $SSH_KEY_PATH"
SCP_CMD="scp -o StrictHostKeyChecking=no -i $SSH_KEY_PATH"

# 1. Save and compress the Docker image
echo "📦 Saving and compressing Docker image (my_website_backend:latest)..."
# We pipe directly to gzip to save disk I/O
docker save my_website_backend:latest | gzip > backend.tar.gz
echo "✅ Image compressed to backend.tar.gz"

# 2. Transfer files
echo "📤 Transferring files to server..."
$SCP_CMD backend.tar.gz $SSH_USER@$SSH_HOST:/tmp/backend.tar.gz
$SCP_CMD docker-compose.prod.yml $SSH_USER@$SSH_HOST:~/docker-compose.yml

echo "   Securing .env directory..."
$SSH_CMD $SSH_USER@$SSH_HOST "if [ -f ~/.env ]; then rm ~/.env; fi && mkdir -p ~/.env && chmod 700 ~/.env"
if [ -f .env/.env.production ]; then
  $SCP_CMD .env/.env.production $SSH_USER@$SSH_HOST:~/.env/.env.production
  $SSH_CMD $SSH_USER@$SSH_HOST "chmod 600 ~/.env/.env.production"
else
  echo "❌ .env/.env.production not found! Aborting."
  exit 1
fi

# Transfer the database init script (mounted by compose)
echo "   Ensuring database directory exists..."
$SSH_CMD $SSH_USER@$SSH_HOST "mkdir -p ~/database"
$SCP_CMD database/init.sql $SSH_USER@$SSH_HOST:~/database/init.sql

# Transfer monitoring configs
echo "   Transferring monitoring configs..."
$SSH_CMD $SSH_USER@$SSH_HOST "mkdir -p ~/docker/monitoring/prometheus ~/docker/monitoring/loki ~/docker/monitoring/promtail ~/docker/monitoring/grafana/provisioning/datasources ~/docker/monitoring/grafana/provisioning/dashboards ~/docker/monitoring/grafana/provisioning/alerting"
$SCP_CMD docker/monitoring/prometheus/prometheus.yml $SSH_USER@$SSH_HOST:~/docker/monitoring/prometheus/prometheus.yml
$SCP_CMD docker/monitoring/loki/loki-config.yml $SSH_USER@$SSH_HOST:~/docker/monitoring/loki/loki-config.yml
$SCP_CMD docker/monitoring/promtail/promtail-config.yml $SSH_USER@$SSH_HOST:~/docker/monitoring/promtail/promtail-config.yml
$SCP_CMD docker/monitoring/grafana/grafana.ini $SSH_USER@$SSH_HOST:~/docker/monitoring/grafana/grafana.ini
$SCP_CMD -r docker/monitoring/grafana/provisioning/ $SSH_USER@$SSH_HOST:~/docker/monitoring/grafana/provisioning/
$SSH_CMD $SSH_USER@$SSH_HOST "chmod -R a+r ~/docker/monitoring/grafana/"

echo "   Transferring lifecycle scripts..."
$SSH_CMD $SSH_USER@$SSH_HOST "mkdir -p ~/lifecycle"
$SCP_CMD nixos/lifecycle/*.sh $SSH_USER@$SSH_HOST:~/lifecycle/
$SSH_CMD $SSH_USER@$SSH_HOST "chmod +x ~/lifecycle/*.sh"
echo "✅ Files transferred."

# 3. Execute deployment on server
echo "⚡ Executing remote deployment commands..."
$SSH_CMD $SSH_USER@$SSH_HOST << 'REMOTE_EOF'
  set -e

  # Load new image
  echo "   [Remote] Loading Docker image..."
  gunzip -c /tmp/backend.tar.gz | docker load
  rm /tmp/backend.tar.gz

  # Teardown old containers (graceful on first deploy)
  echo "   [Remote] Tearing down old containers..."
  cd ~
  ./lifecycle/kill.sh -f || echo "   [Remote] No existing stack to teardown (first deploy?)"

  # Start services using serve.sh with --detach (no log tailing in CI)
  echo "   [Remote] Starting services..."
  ./lifecycle/serve.sh --detach

  # Cleanup old images
  echo "   [Remote] Pruning old images..."
  docker image prune -f
REMOTE_EOF

echo "--------------------------------------------------"
echo "✅ Deployment to $TARGET_ENV complete!"
echo "--------------------------------------------------"

# Cleanup local artifact
rm backend.tar.gz
